package eobereqhdler

import (
	"eobe/eobehttp"
)

//Error string to the front end
const cStrBadRequest = "Bad request, cannot found resouces: %s"
const cStrBadRequestMethod = "Bad request, no action defined for: %s:%s"

//DEV LOG string
const cStrDEVInvalidAction = "Response Fetcher reported error, Invalid Action: %s"
const cStrDEVDebugResultJSONString = "Make JSON response, length of singleRow: %d, length of names: %d, length of multirows: %d"

//Error LOG strings
const cStrClientManagmentErrorNewAccess = "Client Manament module error for register new access, request:%s, error: %s."
const cStrErrorFailedToGetClientInfo = "Failed to get client info from eobereqhdler.RequestHandler.handlingClientInfo, check server client managment logging"
const cStrErrorFailedToGetCallSeqInstance = "Get an Call sequence instance failed"

//keywords
const cStrExpectedXML = eobehttp.CStrExpectedXML
const cStrExpectedJSON = eobehttp.CStrExpectedJSON
const cStrExpectedHTML = eobehttp.CStrExpectedHTML
const cStrStringFmt = "%s"
const cStrSlash = "/"

/*
//type abbreviations
type rqstDataWrapper eobehttp.RequestData
type respDataWrapper eobehttp.ResponseData
type apiCallResults eobeapi.CallSeqResult
type httpActDefn eobeload.HTTPActionDefn
type loggerInf eobehttp.HttpLoggerInf
*/

//content types
const cStrTextHTML = "text/html"
const cStrAppJSON = "application/json"
