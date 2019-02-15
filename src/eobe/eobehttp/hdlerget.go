package eobehttp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

/* httpGetHandler take below responsibilities:
 *  0. Validate GET Method URL
 *  Static Resources:
 *	1. Response all kinds of static resources to clients, including static html files. (DO: Finish request)
 *  2. Send the user specified STATIC 404 page or default 404 page.(DO: Finish request)
 *  Dynamic Resources:
 *	1. If not specified in the request.
 *  3. Parse Get Method Query Values and send to next part to handle.
 */
type httpGetHandler struct {
	httpBaseHandler
	qv     map[string]string
	target string
}

func newGetHandler(logger HttpLoggerInf, rr resourceReaderInf, rw *responseWritter) httpMethodInf {
	var h httpGetHandler

	h.logger = logger
	h.rscReader = rr
	h.rspWrtter = rw

	return &h
}

func (h *httpGetHandler) doMethod(r *http.Request) (int, error) {
	//Why this happens????
	if r.URL == nil {
		h.rspWrtter.sendServerInternalError(r.URL, fmt.Errorf(cStrNULLURL, r.RequestURI), h.logger)
		return http.StatusInternalServerError, nil
	}

	//Get and save query values first
	queryValues, pqErr := url.ParseQuery(r.URL.RawQuery)
	if pqErr != nil {
		h.rspWrtter.sendServerInternalError(r.URL, pqErr, h.logger)
		return http.StatusInternalServerError, nil
	}
	h.qv = h.parseQryValues(queryValues)

	//get requested Target
	//	target is the HTML template full path or resources full path or Action
	var urlResult int
	h.target, urlResult = h.up.parseURL(r.URL.Path, h.logger)
	switch urlResult {
	case cIntFileFound:
		if h.staticFileDownload(h.target) {
			return http.StatusOK, nil
		}
	case cIntRedirectIndex:
		_, wrkPath := h.rscReader.getPathes()
		h.rspWrtter.sendRedirectIndexURL(r, wrkPath, h.target, h.logger)
		return http.StatusSeeOther, nil
	}

	return http.StatusContinue, nil
}

//Handling the File Downloads for: static HTML, js, png, jpg, svg, css and so on..
func (h *httpGetHandler) staticFileDownload(path string) bool {
	if len(path) == 0 {
		return false
	}

	var fname string
	fIndex := strings.LastIndex(path, cStrSlash)
	if fIndex > 0 && fIndex < len(path) {
		fname = path[fIndex:]
		fname = strings.Trim(fname, cStrSlash)
	} else {
		fname = strings.Trim(path, cStrSlash)
	}

	var found bool
	var fbytes []byte
	if getPckConfig().PreloadStaticFiles {
		found, fbytes = h.rscReader.findFile(fname)
	} else {
		found, fbytes = h.rscReader.openDiskFile(path)
	}

	if !found {
		return false
	}

	var contentType string
	if strings.HasSuffix(fname, cStrSuffixCSS) {
		contentType = cStrTextCSS
	} else if strings.HasSuffix(fname, cStrSuffixJS) {
		contentType = cStrAppJS
	} else if strings.HasSuffix(fname, cStrSuffixPNG) {
		contentType = cStrImgPNG
	} else if strings.HasSuffix(fname, cStrSuffixJPG) {
		contentType = cStrImgJPG
	} else if strings.HasSuffix(fname, cStrSuffixSVG) {
		contentType = cStrImgSVG
	} else if strings.HasSuffix(fname, cStrSuffixICO) {
		contentType = cStrImgICON
	} else {
		contentType = cStrTextHTML
	}

	h.rspWrtter.setCacheControl(cStrDefaultValCacheCtrlSTATIC, "", "")
	h.rspWrtter.send200Response(contentType, fbytes)
	return true
}

func (h *httpGetHandler) getRequestData() *RequestData {
	var reqData RequestData
	reqData.QueryTarget = h.target
	reqData.QueryKeyValueMap = h.qv
	reqData.Logger = h.logger
	return &reqData
}
