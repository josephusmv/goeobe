package eobehttp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//No reentrances, only for one single request processing
type responseWritter struct {
	w            http.ResponseWriter
	rootRedirect bool
	statuesCode  int

	//Header Options
	headerWritter
}

func newResponseWritter(w http.ResponseWriter, rootRedirect bool) *responseWritter {
	rspW := &responseWritter{w: w, rootRedirect: rootRedirect}
	rspW.statuesCode = http.StatusOK

	//default header options
	rspW.setDefaultRespHeaders()

	return rspW
}

func (rspw *responseWritter) setStatuesCode(stcd int) {
	rspw.statuesCode = stcd
}

func (rspw *responseWritter) send200Response(cntType string, body []byte) {
	if rspw.statuesCode != http.StatusOK {
		rspw.w.WriteHeader(rspw.statuesCode)
	}

	rspw.w.Header().Add(cStrContentType, cntType)

	rspw.addDefaultRespHeaders(rspw.w)

	rspw.w.Write(body)
}

func (rspw *responseWritter) sendBadRequestError(url *url.URL, err error, httpCode int, rr resourceReaderInf, logger HttpLoggerInf) {
	if err == nil {
		err = fmt.Errorf(cStrGenError)
	}

	errorStr := fmt.Sprintf(cStrBadRequestWithErr, url.String(), err.Error(), httpCode)
	logger.TraceDev(errorStr) //cannot use TraceError here..

	rspw.w.Header().Add(cStrContentType, cStrTextHTML)
	rspw.addDefaultRespHeaders(rspw.w)

	pkgcfg := getPckConfig()
	if pkgcfg.Page404FilePath != "" {
		var valid bool
		var fbytes []byte
		if getPckConfig().PreloadStaticFiles {
			valid, fbytes = rr.findFile(pkgcfg.Page404FilePath)
		} else {
			valid, fbytes = rr.openDiskFile(pkgcfg.Page404FilePath)
		}

		if valid && fbytes != nil {
			rspw.w.WriteHeader(httpCode)
			fmt.Fprintf(rspw.w, string(fbytes), errorStr)
			return
		}
	}

	http.Error(rspw.w, errorStr, httpCode)
}

func (rspw *responseWritter) sendServerInternalError(url *url.URL, err error, logger HttpLoggerInf) {
	if err == nil {
		err = fmt.Errorf(cStrGenError)
	}
	errorStr := fmt.Sprintf(cStrStatusInternalServerErrorWithErr, url.String(), err.Error())
	logger.TraceDev(errorStr) //cannot use TraceError here..
	http.Error(rspw.w, errorStr, http.StatusInternalServerError)
}

func (rspw *responseWritter) sendRedirectIndexURL(r *http.Request, workPath string, indexHTMLName string, logger HttpLoggerInf) {
	logger.TraceDev(cStrDEVRedirectURL, indexHTMLName) //cannot use TraceError here..

	if rspw.rootRedirect {
		r.URL.Path = cStrSlash
		http.Redirect(rspw.w, r, indexHTMLName, http.StatusSeeOther)
		return
	}

	//try to found work path from both old and input PATH. add work path if required.
	wpIndxOld := strings.Index(r.URL.Path, workPath)
	wpIndxRedi := strings.Index(indexHTMLName, workPath)
	if wpIndxOld < 0 && wpIndxRedi < 0 {
		indexHTMLName = combineFullPath(workPath, indexHTMLName)
	}

	http.Redirect(rspw.w, r, indexHTMLName, http.StatusSeeOther)
}

//Send 404 Not found page, must return success and terminate the current request.
func (rspw *responseWritter) send404NotFoundHTML(url string, rscReader resourceReaderInf) {
	var htmlContent string

	found, fBytes := rscReader.findFile(cStrPageNotFound404)

	if !found {
		htmlContent = fmt.Sprintf(cStrDefaultHTML404, url)
	} else {
		htmlContent = fmt.Sprintf(string(fBytes), url)
	}

	rspw.w.Header().Add(cStrContentType, cStrTextHTML)
	rspw.w.WriteHeader(http.StatusNotFound)
	rspw.w.Write([]byte(htmlContent))
}
