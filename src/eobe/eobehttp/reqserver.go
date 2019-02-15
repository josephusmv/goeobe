//Package eobehttp Responsible for Handling all HTTP requestes, including:
//	1. Push requested HTML to the client browser, combine templates
//	2. Push resources(image, js, css...) to the client browser
//	3. Extract the HTTP request Parameters, unify the different HTTP methods
//	(TODO)
//	4. Providing flexible template loading
//	5. Providing flexible template data binding
//	6. (TODO in next version) JS API??
package eobehttp

import (
	"fmt"
	"net/http"
	"strings"
)

//RequestServer Main factory for EOBEHTTP
type RequestServer struct {
	methodFactory
	*resourcesManager
	isRootModule  bool
	module        string //Module is used for notify outside, should not be used inside!!
	rootPath      string
	respFetcher   RespFetcherInf   //Resonpse formatter, mandatory
	loggerFactory LoggerFactoryInf //User defined logger, optional, will use stdout if not provided.
}

//NewRequestServer
func NewRequestServer(module, path string, lgFact LoggerFactoryInf) *RequestServer {

	var rs RequestServer
	rs.isRootModule = false
	rs.module = strings.ToLower(module)
	rs.rootPath = path

	if lgFact != nil {
		rs.loggerFactory = lgFact
	} else {
		rs.loggerFactory = DummyHTTPLogFactory{logger: DummyIHTTPLoggerImpl{}}
	}

	rs.resourcesManager = &resourcesManager{}

	return &rs
}

//Init Init Request server with validation and testing
//	@Params:
//		htmlTemplates []string	: HTML template slice, includes all the HTML template FULL PATH, must have Index.html
//		respFetcher   RespFetcherInf	: Interface implementation, cannot be nil
func (rs *RequestServer) Init(respFetcher RespFetcherInf, indexActionName string) (err error) {
	if !glbPckgConfig.initiated {
		return fmt.Errorf(cStrPackageNotInit)
	}
	if rs.loggerFactory == nil {
		return fmt.Errorf(cStrIllegalReqServer)
	}
	logger := rs.loggerFactory.GetLogger(cStrHTTPMainLoggerFileName, nil, nil)

	if respFetcher == nil /*|| *respFetcher == nil*/ {
		return logger.TraceError(cStrRequireRespFetcherError)
	}
	rs.respFetcher = respFetcher

	rs.readResources(rs.rootPath, rs.module, glbPckgConfig.TemplateFileList, logger)

	rs.indxAction = indexActionName
	rs.page404fName = glbPckgConfig.Page404FilePath

	logger.TraceDev(cStrInitRequestServer, rs.module)

	return nil
}

//ServeHTTP Handler for HTTP requests.
func (rs RequestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ipAddr, portNum := parseIPPort(r)

	rspWrter := newResponseWritter(w, rs.isRootModule)
	logger := rs.loggerFactory.GetLogger(ipAddr, r.Cookies(), r.Header)
	logger.TraceDev(cStrNewRequest, ipAddr, portNum, r.URL.String())

	//Get method implementation and DoMethod()
	mthInf := rs.newMethodHandler(r.Method, logger, rs.resourcesManager, rspWrter)
	httpStatues, unhandledErr := mthInf.doMethod(r)
	if httpStatues != http.StatusContinue { //Request already handled inside, no matter error or not, just return
		return
	}

	if unhandledErr != nil { //Error not handled inside, send response here.
		rspWrter.sendServerInternalError(r.URL, unhandledErr, logger)
		return
	}

	//other information
	req := mthInf.getRequestData()
	req.HTTPMethod = r.Method
	req.Accept = r.Header.Get(cStrAccepct)
	req.Module = rs.module
	req.IP = ipAddr
	req.CookieList = r.Cookies() //for necessary infors, like SID

	//Beyond this package could handle, pass to upper to decide.
	resp, respErr := rs.respFetcher.FetchResponse(*req)
	if respErr != nil { //Error not handled inside, send response here.
		rspWrter.sendBadRequestError(r.URL, respErr, http.StatusInternalServerError, rs.resourcesManager, logger)
		return
	}
	if resp.APIErr != nil { //this is for all API returned errors
		rspWrter.sendBadRequestError(r.URL, resp.APIErr, http.StatusBadRequest, rs.resourcesManager, logger)
		return
	}

	//Set fetched cookie
	for _, v := range resp.CookieList {
		http.SetCookie(w, v)
	}

	//Treat as Template HTML if with type  "text/html", end this request
	if resp.ContentType == cStrTextHTML && resp.HTMLTmpltName != "" {
		if getPckConfig().DevMode {
			newtempltDevHandler(rs.resourcesManager, &resp).renderTemplates(w, logger)
		} else {
			newTempltHandler(rs.resourcesManager, &resp).renderTemplates(w, logger) //newTempltHandler should never return nil
		}
		return
	}

	rspWrter.send200Response(resp.ContentType, resp.Body)
	return
} //end of ServeHTTP

//GetModuleName   ...
func (rs RequestServer) GetModuleName() string {
	return rs.module
}

//SetRootModule This Request server should be accessed by URL root: /
func (rs *RequestServer) SetRootModule() {
	rs.isRootModule = true
}
