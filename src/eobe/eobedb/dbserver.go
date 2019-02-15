package eobedb

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

//DBServer DB Server will start the main DB thread to accept query/exec requests.
//	DB Server acts like an co-ordinator/scheduler to manager below issues:
//	1. DB validation and connection test
type DBServer struct {
	//Logger inside DB Server, recommand a global logger, not for anyother threads.
	//For each request, a log could use the Query/Execute method to provides.
	logger *DBLogger
	//Concurrency objects
	chanQuery chan exchgnData
	chanQuit  chan bool
	quitWg    *sync.WaitGroup
	//DB connection info
	dbtype     string
	connectStr string
	//DB instance, sql.DB related operations should only involed in run() and init()
	//	init() - get a instinct errror response when try to open in the caller's goroutine.
	//	run()  - make sure sql.DB.close() is called, query is scoped.
	//	Do not involve this member in any other methods.
	dbInstance *sql.DB
	//DB options, for supported options and its types, see OptionDB* constants
	dbOptions map[int]interface{}
}

//RunNewDBServer Initiate a DB server instance, run it and get an interface to use.
//Don't give any logger here!!! it will be deleted
func RunNewDBServer(dbtype, connectStr string) (DBQueryInf, DBServerInf) {
	var dbSrv DBServer

	dbSrv.dbtype = dbtype
	dbSrv.connectStr = connectStr
	//dbSrv.logger = logger

	dbSrv.chanQuery = make(chan exchgnData)
	dbSrv.chanQuit = make(chan bool)

	return dbSrv, &dbSrv
}

func (dbSrv *DBServer) Init() error {
	return dbSrv.init()
}

//StopDBServer Stop the DB server
//	Exit DB running go routine, all chanels closed, DB logger interface set to null.
func (dbSrv *DBServer) StopDBServer() {
	dbSrv.stop()
}

//SetDBOptions Set DB Server options, for supported options and its types, see eobedb.OptionDB* constants
//	DBOptions is not protected by any lock,
//  so call this method right AFTER Get DB server and BEFORE Init.
//  And Don't play concurrency with this method.
//	It's not just risky, but also meanless to do so.
func (dbSrv *DBServer) SetDBOptions(option int, value interface{}) error {
	if dbSrv.dbOptions == nil {
		dbSrv.dbOptions = make(map[int]interface{})
	}
	switch v := value.(type) {
	case int:
		switch option {
		case OptionDBMaxConcurrencyInt:
			dbSrv.dbOptions[OptionDBMaxConcurrencyInt] = value.(int)
		case OptionEnableDEVLog:
			if dbSrv.logger == nil {
				dbSrv.logger = &DBLogger{bDevLogEnable: true}
			} else {
				dbSrv.logger.bDevLogEnable = true
			}
		}
	case string:
		if option == OptionDBLogRoot {
			if dbSrv.logger == nil {
				dbSrv.logger = &DBLogger{logPath: v}
			} else {
				dbSrv.logger.reInitLogger(v)
			}
			dbSrv.logger.TraceInfo(cStrSetDBLogPath, v)
		} else {
			fmt.Printf("%s\n", value.(string))
		}
	default:
		if dbSrv.logger != nil {
			return dbSrv.logger.TraceError(cstrTypeCastError, v)
		}
		return fmt.Errorf(cstrTypeCastError, v)
	}
	return nil
}

//ExecDBAction Interface implementation
//	Param: logger - recommand to use a client specific logger, not a global one.
func (dbSrv DBServer) ExecDBAction(qryData QueryData, callerLogger DBLoggerInf) (QueryResult, error) {
	return dbSrv.execDBAction(qryData, callerLogger)
}

func (dbSrv *DBServer) init() (err error) {
	if dbSrv.logger == nil {
		//if no logger, then consider it as dev debug.
		dbSrv.logger = &DBLogger{logPath: DefaultLogRoot, bDevLogEnable: true}
	}

	//initiate *sql.DB
	dbSrv.dbInstance, err = sql.Open(dbSrv.dbtype, dbSrv.connectStr)
	if err != nil {
		return dbSrv.logger.TraceError(cstrDBOpenError, dbSrv.dbtype, dbSrv.connectStr, err.Error())
	}

	//verify the connection is fine.
	err = dbSrv.dbInstance.Ping()
	if err != nil {
		return dbSrv.logger.TraceError(cstrDBConnectionError, err.Error())
	}

	dbSrv.quitWg = &sync.WaitGroup{}

	dbSrv.logger.TraceInfo(cstrDBSuccessfullyInited)
	var wg sync.WaitGroup
	wg.Add(1)
	go dbSrv.run(&wg)
	wg.Wait() //Wait for the DB server successfully started and running without errors.

	return nil
}

func (dbSrv *DBServer) run(wg *sync.WaitGroup) {
	defer dbSrv.dbInstance.Close()

	dbSrv.logger.TraceInfo(cstrDBSuccessfullyOpened, dbSrv.dbtype)

	dbWorker := databaseWorker{dbInstance: dbSrv.dbInstance}

	//release dbSrv.init()
	wg.Done()

	//read Options
	optDBMaxConcurrencyInt, found := dbSrv.dbOptions[OptionDBMaxConcurrencyInt].(int)
	if !found {
		optDBMaxConcurrencyInt = DefaultOptionDBMaxConcurrencyInt
	}

	//prepare some local vars
	exitReason := cstrUnknown //Desc of any possible reason for DB server Stop

	//Max go routine management
	chanDBActionDone := make(chan bool)
	defer close(chanDBActionDone)
	actionQue := dbActionQueue{chanDBActionDone: chanDBActionDone, itemQueue: []queItem{}, dbWroker: dbWorker}
	usedGoRoutineCount := 0

	//main loop
	quit := false
	for !quit {
		select {
		case exchData := <-dbSrv.chanQuery:
			dbSrv.logger.TraceDev(cstrExchangeData, exchData)
			qryDfn := exchData.qryValue.QryActionDfn
			if usedGoRoutineCount < optDBMaxConcurrencyInt {
				go dbWorker.doDBAction(qryDfn, exchData, chanDBActionDone)
				usedGoRoutineCount++
			} else {
				actionQue.addToQueue(qryDfn, exchData)
			}

		case quit = <-dbSrv.chanQuit:
			if quit {
				dbSrv.logger.TraceInfo(cstrDBServerStopGood)
				dbSrv.quitWg.Done()
				return
			}

		case done := <-chanDBActionDone:
			if done && usedGoRoutineCount > 0 {
				usedGoRoutineCount--
			}

			hasNewRoutine := actionQue.processNextOne()
			if hasNewRoutine {
				usedGoRoutineCount++
			}

			//default:
		}
	}

	dbSrv.logger.TraceInfo(cstrDBServerStopBad, exitReason)
	return
}

//Do not involve any actual sql.DB related operations here.
func (dbSrv *DBServer) stop() {
	dbSrv.quitWg.Add(1)
	dbSrv.chanQuit <- true
	dbSrv.quitWg.Wait()

	close(dbSrv.chanQuery)
	close(dbSrv.chanQuit)

	dbSrv.logger = nil
}

//Structure to pass both result value and result recieving chanel to DB Server --> DB Worker
//	Data flow:
//		1. External caller of ExecDBAction() ==chanQuery(type exchgnData)==> DB server run() call DB Worker
//		2. DB Worker ==exchgnData.QueryResult(type QueryResult)==> ExecDBAction()
type exchgnData struct {
	rsltChan     chan QueryResult
	qryValue     QueryData
	callerLogger DBLoggerInf
}

func (dbSrv DBServer) execDBAction(qryData QueryData, callerLogger DBLoggerInf) (QueryResult, error) {
	rsltChan := make(chan QueryResult)
	exData := exchgnData{rsltChan: rsltChan, qryValue: qryData, callerLogger: callerLogger}
	dbSrv.chanQuery <- exData
	qryRslt := <-rsltChan

	//callerLogger.TraceDev("Result: %v. \n", qryRslt) //TODO: DELETE THIS
	return qryRslt, nil
}
