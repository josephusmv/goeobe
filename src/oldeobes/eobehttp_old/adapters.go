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
	CookieList    []*http.Cookie //read cookie from cltmgmt module
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
