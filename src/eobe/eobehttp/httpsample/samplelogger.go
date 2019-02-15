package main

import (
	"eobe/eobehttp"
	"fmt"
	"log"
	"net/http"
	"os"
)

type httpLoggerImpl struct {
	logger *log.Logger
}

const cStrFilePath = "./httplogs/"

func newhttpLoggerImpl(fileName string) *httpLoggerImpl {
	os.MkdirAll(cStrFilePath, os.ModePerm)
	file, err := os.OpenFile(cStrFilePath+fileName+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
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
	return newhttpLoggerImpl(ipAddr)
}
