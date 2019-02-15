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
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//RequestServer Main factory for EOBEHTTP
type RequestServer struct {
	path             string //Module resources root path, includes the module name: root/modulename/. mandatory
	module           string //Module Name, mandatory
	htmlTemplates    *template.Template
	htmlTemplatesStr []string
	respFetcher      RespFetcherInf    //Resonpse formatter, mandatory
	loggerFactory    LoggerFactoryInf  //User defined logger, optional, will use stdout if not provided.
	resFileMap       map[string]string //map[fileInfo]content : all resources files under Path, avoid Disk IO access.
	//Var for special cases
	indexActionName string // IF no action, like root path, still need an action name.
	html404Content  string
}

//NewRequestServer Get a new Request Server Instance by Copy
//	@Params:
//		path: root path not include module name, but will store only the module root path
//		indexActionName string: specify the Action for Index.html
//		loggerFactory LoggerFactoryInf	: Log factory for get logger of each request, if nil will use the fmt pacakge's default output.
func NewRequestServer(module, path string, lgFact LoggerFactoryInf) *RequestServer {
	var rs RequestServer
	rs.module = strings.ToLower(module)
	//rs.path = combineFullPath(path, rs.module)
	rs.path = path

	if lgFact != nil {
		rs.loggerFactory = lgFact
	} else {
		rs.loggerFactory = DummyHTTPLogFactory{logger: DummyIHTTPLoggerImpl{}}
	}

	return &rs
}

//Init Init Request server with validation and testing
//	@Params:
//		htmlTemplates []string	: HTML template slice, includes all the HTML template FULL PATH, must have Index.html
//		respFetcher   RespFetcherInf	: Interface implementation, cannot be nil
func (rs *RequestServer) Init(indexActionName string, htmlTemplates []string, respFetcher RespFetcherInf) (err error) {
	if rs.loggerFactory == nil {
		return fmt.Errorf(cStrIllegalReqServer)
	}
	logger := rs.loggerFactory.GetLogger(cStrHTTPMainLoggerFileName, nil, nil)
	logger.TraceDev(cStrInitRequestServer, rs.module)

	err = rs.initHTMLTemplates(htmlTemplates, logger)
	if err != nil {
		return err
	}

	rs.resFileMap = make(map[string]string)
	modulePath := combineFullPath(rs.path, rs.module)
	err = rs.readResources(modulePath, logger)
	if err != nil {
		return err
	}

	if respFetcher == nil /*|| *respFetcher == nil*/ {
		return logger.TraceError(cStrRequireRespFetcherError)
	}
	rs.respFetcher = respFetcher

	rs.indexActionName = indexActionName

	return nil
}

//ServeHTTP Handler for HTTP requests.
func (rs RequestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Init log for this request client.
	ipAddr, portNum := parseIPPort(r)
	logger := rs.loggerFactory.GetLogger(ipAddr, r.Cookies(), r.Header)
	logger.TraceDev(cStrNewRequest, ipAddr, portNum, r.URL.String())

	//Get method implementation and DoMethod()
	mthInf, param := methodFactory(w, r, logger, &rs)
	req, httpStatues, unhandledErr := mthInf.doMethod(param)

	if httpStatues != http.StatusContinue { //Request already handled inside, no matter error or not, just return
		return
	}

	if unhandledErr != nil { //Error not handled inside, send response here.
		errorStr := fmt.Sprintf(cStrStatusInternalServerError, r.URL.String())
		logger.TraceDev(cStrDevErrorDebug2Param, errorStr, unhandledErr.Error())
		http.Error(w, errorStr, http.StatusInternalServerError)
		return
	}

	//other information
	req.IP = ipAddr
	req.Port = portNum
	req.CookieList = r.Cookies() //for necessary infors, like SID

	//Beyond this package could handle, pass to upper to decide.
	resp, respErr := rs.respFetcher.FetchResponse(*req)
	if respErr != nil { //Error not handled inside, send response here.
		errorStr := fmt.Sprintf(cStrBadRequest, r.URL.String())
		logger.TraceDev(cStrDevErrorDebug2Param, errorStr, respErr.Error())
		http.Error(w, errorStr, http.StatusBadRequest)
		return
	}

	//Set fetched cookie
	for _, v := range resp.CookieList {
		http.SetCookie(w, v)
	}

	//Treat as Template HTML if with type  "text/html", end this request
	if resp.ContentType == cStrTextHTML && resp.HTMLTmpltName != "" {
		th := templtHandler{templateName: resp.HTMLTmpltName,
			hData:         resp.HTMLTmpltData,
			htmlTemplates: rs.htmlTemplates}
		th.renderTemplates(w, logger)
		return
	}

	//If not HTML then resp.Body is storing all response data formated.
	w.Header().Add(cStrContentType, resp.ContentType)
	w.Header().Add(cStrCharSet, cStrDefaultCharSetUTF8)
	w.Write(resp.Body)

	return
}

//initHTMLTemplates Validate the given template is ok or not, to ensure there is no file missing during parse.
func (rs *RequestServer) initHTMLTemplates(templates []string, logger HttpLoggerInf) error {
	notFoundIndexHTML := true
	for _, v := range templates {
		if strings.HasSuffix(v, cStrIndexHTML) ||
			strings.HasSuffix(v, cStrIndexTMPLT) {
			notFoundIndexHTML = false
		}

		if _, err := os.Stat(v); err != nil {
			logger.TraceError(cStrInvalidHTMLTemplateFileError, v)
		} else {
			rs.htmlTemplatesStr = append(rs.htmlTemplatesStr, v)
		}
		logger.TraceDev(cStrDEVLoadTmplateFile, v)
	}

	if notFoundIndexHTML {
		return logger.TraceError(cStrIndexHTMLNotFound)
	}

	if len(rs.htmlTemplatesStr) <= 0 {
		return logger.TraceError(cStrErrorNoTemplates) //a Warning error
	}

	tmps, err := template.ParseFiles(rs.htmlTemplatesStr...)
	if tmps == nil || err != nil {
		return logger.TraceError(cStrErrorParseTemplates, err.Error())
	}

	rs.htmlTemplates = template.Must(tmps, err)
	if rs.htmlTemplates == nil {
		return logger.TraceError(cStrErrorMustTemplates)
	}
	return nil
}

func (rs *RequestServer) readResources(path string, logger HttpLoggerInf) error {
	files, derr := ioutil.ReadDir(path)
	if derr != nil {
		return logger.TraceError(cStrReadModuleDirError, path, derr)
	}

	for _, file := range files {
		//Skip files with suffix .tmplt, since they are already handled by initHTMLTemplates.
		if strings.HasSuffix(file.Name(), cStrSuffixTmplt) {
			continue
		}

		filePath := combineFullPath(path, file.Name())

		if file.IsDir() {
			rs.readResources(filePath, logger) //ignore errors
			continue
		}

		if file.Size() <= 0 {
			continue
		}

		if rs.readSpecialRes(filePath, logger) {
			continue
		}
		fdata, err := ioutil.ReadFile(filePath)
		if err != nil {
			logger.TraceDev(cStrReadModuleDirError, file.Name(), err.Error())
			continue
		}

		rs.resFileMap[file.Name()] = string(fdata)
		logger.TraceDev(cStrLoadFileSuccess, filePath, rs.module)
	}

	return nil
}

//readSpecialRes Return true means already handled.
// If special resources found will return true no matter already handled or not.
func (rs *RequestServer) readSpecialRes(filePath string, logger HttpLoggerInf) bool {
	var fdata []byte
	var err error

	switch {
	case strings.HasSuffix(filePath, VStr404PageFileName):
		fdata, err = ioutil.ReadFile(filePath)
		if err != nil {
			logger.TraceDev(cStrReadModuleDirError, filePath, err.Error())
		} else {
			rs.html404Content = string(fdata)
		}
		logger.TraceDev(cStrLoadFileSuccess, filePath, rs.module)
		return true
	default:
		return false
	}
}

//GetModuleName   ...
func (rs RequestServer) GetModuleName() string {
	return rs.module
}
