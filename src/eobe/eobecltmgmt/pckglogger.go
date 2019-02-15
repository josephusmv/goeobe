package eobecltmgmt

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const cStrLogFileName = "clientmanagment.log"
const cStrCltMgrLogPrefix = "clientmanagment:"

type pckgLogger struct {
	logPath      string
	logger       *log.Logger
	enableDevLog bool
}

func (logger *pckgLogger) initLogger() {
	if logger.logger != nil {
		return
	}

	if logger.logPath == "" {
		logger.logPath = "./"
	}

	path := strings.TrimRight(logger.logPath, cStrSlash) + cStrSingleSlash
	filePath := path + cStrLogFileName

	os.MkdirAll(path, os.ModePerm)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger.logger = log.New(file, cStrCltMgrLogPrefix, log.Ldate|log.Ltime)
}

func (logger pckgLogger) TraceError(format string, a ...interface{}) error {
	logger.initLogger()
	logger.logger.Printf("Error: "+format, a...)
	return fmt.Errorf("Error"+format, a...)
}

func (logger pckgLogger) TraceDev(format string, a ...interface{}) {
	logger.initLogger()
	if logger.enableDevLog {
		logger.logger.Printf("DEV: "+format, a...)
	}
}

func (logger pckgLogger) TraceInfo(format string, a ...interface{}) {
	logger.initLogger()
	logger.logger.Printf("Info: "+format, a...)
}
