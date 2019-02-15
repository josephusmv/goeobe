package eobesession

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const cStrDummyLogPrefix = "eobesession(dummy):"
const cStrLogFileName = "eobesession_dummy.log"

//Dummy interface implementation
//LogInf
type dummyLogger struct {
	logPath       string
	logger        *log.Logger
	bDevLogEnable bool
}

func NewDummyLogger() LoggerInf {
	var dLogger dummyLogger
	dLogger.logPath = cStrCurrentPathLocal
	dLogger.bDevLogEnable = FlagEnableDevLog
	dLogger.initLogger()

	return &dLogger
}

func (logger *dummyLogger) initLogger() {
	if logger.logger != nil {
		return
	}

	path := strings.TrimRight(logger.logPath, cStrSlash) + cStrSingleSlash
	filePath := path + cStrLogFileName

	os.MkdirAll(path, os.ModePerm)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger.logger = log.New(file, cStrDummyLogPrefix, log.Ldate|log.Ltime)
}

func (logger *dummyLogger) TraceError(format string, a ...interface{}) error {
	logger.initLogger()
	logger.logger.Printf("Error: "+format, a...)
	return fmt.Errorf("Error"+format, a...)
}

func (logger *dummyLogger) TraceDev(format string, a ...interface{}) {
	logger.initLogger()
	if logger.bDevLogEnable {
		logger.logger.Printf("DEV: "+format, a...)
	}
}

func (logger *dummyLogger) TraceInfo(format string, a ...interface{}) {
	logger.initLogger()
	logger.logger.Printf("Info: "+format, a...)
}

//Dummy interface implementation
//UserDataInf
type DummyBindDataImpl struct {
}

func (ud *DummyBindDataImpl) Validate(name string, sid string, pCList map[string]*http.Cookie) bool {
	return true
}
func (ud *DummyBindDataImpl) Clear() {

}
func (ud *DummyBindDataImpl) Login(name string) {

}
func (ud *DummyBindDataImpl) Logout() {

}
