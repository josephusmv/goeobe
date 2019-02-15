package eobehttp

import "net/http"

//LoggerFactoryInf Log Factory for get log interface object
//	If no concrete implementation provided, will use the internal httpLoggerDummyImpl factory
type LoggerFactoryInf interface {
	GetLogger(ipAddr string, reqCookie []*http.Cookie, Header map[string][]string) HttpLoggerInf
}

//HttpLoggerInf HTTP Log Interface Adapter.
// If no concrete implementation provided, will use stdout(fmt.Print) as default.
type HttpLoggerInf interface {
	TraceError(format string, a ...interface{}) error
	TraceDev(format string, a ...interface{})
	TraceInfo(format string, a ...interface{})
}

//RespFetcherInf Response Fetcher Interface Adapter
// This is a mandatory interface need to be implemented.
type RespFetcherInf interface {
	FetchResponse(req RequestData) (rsp ResponseData, err error)
}

//RequestData The request from client
//	e.g.: Get /Modl/DoSomething?KEY1=VALUE1&KEY2=VALUE2
//		Module = Modl
//		QueryTarget = DoSomething
//		QueryKeyValueMap={KEY1:VALUE1, KEY2:VALUE2}
type RequestData struct {
	IP               string
	Port             string
	Module           string
	HTTPMethod       string
	Accept           string //support only: Accept: text/html, application/json, application/xml for now
	QueryTarget      string
	QueryKeyValueMap map[string]string
	UploadedFiles    map[string][]byte //[Filename]bytes
	CookieList       []*http.Cookie    //cookie from front end
	Logger           HttpLoggerInf
}

//ResponseData
type ResponseData struct {
	ContentType   string //String format content type, standard HTTP, like: application/xml, text/html...
	Body          []byte
	HTMLTmpltName string
	HTMLTmpltData TemplateData
	CookieList    []*http.Cookie //read cookie from cltmgmt package
	APIErr        error
}

//TemplateData wrap the data to render the HTML, user MUST use names here to write HTML template
type TemplateData struct {
	KVMap map[string]string
	Rows  [][]string
}

//GetContentType return predefined content type to HTML content type string
func GetContentType(respType string) string {
	switch respType {
	case CStrExpectedXML:
		return cStrAppXML
	case CStrExpectedJSON:
		return cStrAppJSON
	case CStrExpectedHTML:
		fallthrough
	default:
		return "text/html"
	}
}

//some of the exposed const strings
const CHTTPTypeStrHTML = cStrTextHTML
const CHTTPTypeStrJSON = cStrAppJSON
const CHTTPTypeStrXML = cStrAppXML

//CStrExpectedXML EXPECTEDRESP = RESP_XML
const CStrExpectedXML = "RESP_XML"

//CStrExpectedJSON EXPECTEDRESP = RESP_JSON
const CStrExpectedJSON = "RESP_JSON"

//CStrExpectedHTML EXPECTEDRESP = RESP_HTML
const CStrExpectedHTML = "RESP_HTML"

//CStrUploadFileInKey w
const CStrUploadFileInKey = "HasUploadFileInKey"

//VStr404PageFileName file name for 404Page, user could modify this file but no security provided about this file
// If specified 404 file is not found, will jump to index.html page, with random Module(depends on route).
var VStr404PageFileName = "page404.html"
