package eobehttp

import (
	"fmt"
	"net/http"
)

const cDummyIntLogLevel = 0

const cDummyIntLogLevelDev = 0
const cDummyIntLogLevelInfo = 10
const cDummyIntLogLevelError = 100

type DummyIHTTPLoggerImpl struct {
	prefix string
}

func (logger DummyIHTTPLoggerImpl) TraceError(format string, a ...interface{}) error {

	fmt.Printf(logger.prefix+"Error: "+format, a...)
	fmt.Println()
	return fmt.Errorf("Error"+format, a...)
}

func (logger DummyIHTTPLoggerImpl) TraceDev(format string, a ...interface{}) {
	if cDummyIntLogLevel > cDummyIntLogLevelDev {
		return
	}
	fmt.Printf(logger.prefix+"DEV: "+format, a...)
	fmt.Println()
}

func (logger DummyIHTTPLoggerImpl) TraceInfo(format string, a ...interface{}) {
	if cDummyIntLogLevel > cDummyIntLogLevelInfo {
		return
	}
	fmt.Printf(logger.prefix+"Info: "+format, a...)
	fmt.Println()
}

type DummyHTTPLogFactory struct {
	logger DummyIHTTPLoggerImpl
}

func (logfact DummyHTTPLogFactory) GetLogger(ipAddr string, reqCookie []*http.Cookie, Header map[string][]string) HttpLoggerInf {
	//DELETE me: Keep IP and Port Debug codes before it verified works using browser and remote debug...
	//logFileName := strings.Replace(ip, ".", "_", -1)
	logfact.logger.prefix = ipAddr + ":"
	return logfact.logger
}
