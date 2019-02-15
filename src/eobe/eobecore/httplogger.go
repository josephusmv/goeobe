package eobecore

import (
	"eobe/eobecltmgmt"
	"eobe/eobehttp"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type httpLoggerImpl struct {
	logger *log.Logger
}

var vStrLogFilePath = "./httplogs/"
var vEnableDevLogs = false

func newhttpLoggerImpl(fileName string) *httpLoggerImpl {
	os.MkdirAll(vStrLogFilePath, os.ModePerm)
	file, err := os.OpenFile(vStrLogFilePath+fileName+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	var loggerImpl httpLoggerImpl
	loggerImpl.logger = log.New(file, "TRACE  : ", log.Ldate|log.Ltime)
	return &loggerImpl
}

func (li httpLoggerImpl) TraceError(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	errorLog := "Error: " + format
	li.logger.Printf(errorLog, a...)
	return err
}

func (li httpLoggerImpl) TraceDev(format string, a ...interface{}) {
	errorLog := "Dev: " + format
	li.logger.Printf(errorLog, a...)
}

func (li httpLoggerImpl) TraceInfo(format string, a ...interface{}) {
	errorLog := "Info: " + format
	li.logger.Printf(errorLog, a...)
}

//LoggerFactoryImpl ..
type loggerFactoryImpl struct {
}

//GetLogger ...
func (lf loggerFactoryImpl) GetLogger(ipAddr string, reqCookie []*http.Cookie, Header map[string][]string) eobehttp.HttpLoggerInf {
	fileName := ipAddr

	//Find Client ID from Cookie first
	for _, cky := range reqCookie {
		if cky.Name == eobecltmgmt.CStrKeyWordClientID {
			return newhttpLoggerImpl(ipAddr + cStrUnderscore + cky.Value)
		}
	}

	//Find Client agent infor "User-Agent" from header with IP as the file name
	uaStr, found := Header[cStrHeaderUA]
	if found {
		indxSlash := strings.Index(uaStr[0], cStrSingleSlash)
		if indxSlash > 0 && indxSlash < len(uaStr[0]) {
			fileName = ipAddr + cStrUnderscore + uaStr[0][:indxSlash]
		} else {
			fileName = ipAddr
		}
	}

	return newhttpLoggerImpl(fileName)
}
