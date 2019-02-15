package eobedb

import "database/sql"

type dbActor interface {
	doDBAction(db *sql.DB, qryDefn QueryDefn, exData exchgnData) QueryResult
}

//DatabaseWorker worker for the DB...
type databaseWorker struct {
	dbInstance *sql.DB
}

func (wrk databaseWorker) doDBAction(qryDefn QueryDefn, exData exchgnData, chanDBActionDone chan bool) {
	if exData.callerLogger == nil {
		exData.callerLogger = &DBLogger{}
	}

	exData.callerLogger.TraceDev(cstrDBActorInitMsg, qryDefn, exData.qryValue)

	var actor dbActor
	var qryRslt QueryResult
	switch qryDefn.QueryType {
	case "SELECT":
		actor = queryActor{}
	case "INSERT":
		fallthrough
	case "UPDATE":
		fallthrough
	case "DELETE":
		actor = execActor{}
	default:
		qryRslt.QueryErr = exData.callerLogger.TraceError(cstrInvalidExecType, qryDefn.QueryType)
	}

	if qryRslt.QueryErr == nil {
		qryRslt = actor.doDBAction(wrk.dbInstance, qryDefn, exData)
	}
	exData.rsltChan <- qryRslt
	chanDBActionDone <- true
}
