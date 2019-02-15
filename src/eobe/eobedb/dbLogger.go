package eobedb

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const cStrLogFileName = "dbtrace.log"

//Dummy interface implementation
//LogInf
type DBLogger struct {
	logPath       string
	logger        *log.Logger
	bDevLogEnable bool
}

const cStrSlash = "\\/"
const cStrSingleSlash = "/"
const cStrSpace = " "

func (dblogger *DBLogger) initLogger() {
	if dblogger.logger != nil {
		return
	}

	dblogger.logPath = strings.TrimRight(dblogger.logPath, cStrSlash+cStrSpace)
	if dblogger.logPath == "" {
		dblogger.logPath = "."
	}

	filePath := dblogger.logPath + cStrSingleSlash + cStrLogFileName

	os.MkdirAll(dblogger.logPath+cStrSingleSlash, os.ModePerm)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	dblogger.logger = log.New(file, "eobe DB worker trace : ", log.Ldate|log.Ltime)
}

func (dblogger *DBLogger) reInitLogger(logPath string) {
	dblogger.logger = nil
	dblogger.logPath = logPath
	dblogger.initLogger()
}

func (dblogger *DBLogger) TraceError(format string, a ...interface{}) error {
	dblogger.initLogger()
	dblogger.logger.Printf("Error: "+format, a...)
	return fmt.Errorf("Error: "+format, a...)
}

func (dblogger *DBLogger) TraceDev(format string, a ...interface{}) {
	dblogger.initLogger()
	if dblogger.bDevLogEnable {
		dblogger.logger.Printf("DEV: "+format, a...)
	}
}

func (dblogger *DBLogger) TraceInfo(format string, a ...interface{}) {
	dblogger.initLogger()
	dblogger.logger.Printf("Info: "+format, a...)
}
