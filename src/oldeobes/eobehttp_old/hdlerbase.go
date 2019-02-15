package eobehttp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type httpBaseHandler struct {
	rs *RequestServer
}

func (h httpBaseHandler) parseQryValues(queryValues url.Values) map[string]string {
	result := make(map[string]string)
	for k, v := range queryValues {
		if v != nil && len(v) > 0 {
			result[k] = v[0]
		} else {
			result[k] = ""
		}
	}
	return result
}

func (h httpBaseHandler) parseURLPath(r *http.Request, module string, logger HttpLoggerInf) (target string, urlResult int) {
	if r.URL == nil {
		logger.TraceDev(cStrNULLURL) //this is a very frequently happens error, Cannot use TraceError
		return "", cIntURLNil        //response server internal error
	}

	path := r.URL.Path
	if len(path) == 0 || (len(path) == 1 && path == cStrSlash) {
		return "", cIntURLRootDomain //without any content URL, direct domain name access
	}

	//trim leading cStrSlash
	if path[0] == '/' {
		path = path[1:]
	}

	if len(path) == len(module) || (len(path) == len(module)+1 && path == module+cStrSlash) {
		return "", cIntURLRootModule
	}

	if len(path) < len(module) || !strings.HasPrefix(strings.ToLower(path), module) {
		return "", cIntURLBadModule
	}

	target = path[len(module+cStrSlash):]

	return target, cIntURLOK
}

func (h httpBaseHandler) httpSendError(w http.ResponseWriter, errFmtStr, url string, httpCode int, logger HttpLoggerInf) {
	errorStr := fmt.Sprintf(errFmtStr, url)
	logger.TraceDev(errorStr)
	http.Error(w, errorStr, httpCode)
}

//Send 404 Not found page, must return success and terminate the current request.
func (h httpBaseHandler) send404NotFoundHTML(w http.ResponseWriter, url string) {
	var htmlContent string

	if len(h.rs.html404Content) == 0 {
		htmlContent = fmt.Sprintf(cStrDefaultHTML404, url)
	} else {
		htmlContent = fmt.Sprintf(h.rs.html404Content, url)
	}

	w.Header().Add(cStrContentType, cStrTextHTML)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(htmlContent))
	//TODO: Is there any way I could send HTTP error codes with templates????
	//return templates.ExecuteTemplate(w, VStr404PageFileName, nil)

}

func (h httpBaseHandler) validateActionTarget(actionList []string, target string) bool {

	for i := range actionList {
		if actionList[i] == target {
			return true
		}
	}

	return false
}

func (h httpBaseHandler) getReqData(vls map[string]string, target string, logger HttpLoggerInf, rs *RequestServer) (reqData RequestData) {
	reqData.Module = rs.module
	reqData.QueryTarget = target
	reqData.QueryKeyValueMap = vls
	reqData.Logger = logger

	return
}
