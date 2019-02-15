package eobehttp

import (
	"io/ioutil"
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
}

func (h *httpGetHandler) doMethod(param methodParam) (*RequestData, int, error) {
	rs := h.rs
	r := param.req
	w := param.resp
	logger := param.logger

	//get requested Target
	//	target is the HTML template full path or resources full path or Action
	target, urlResult := h.parseURLPath(r, rs.module, logger)
	if urlResult == cIntURLNil {
		h.httpSendError(w, cStrNULLURL, r.URL.String(), http.StatusInternalServerError, logger)
		return nil, http.StatusInternalServerError, nil
	}

	if urlResult == cIntURLRootDomain || urlResult == cIntURLRootModule {
		target = rs.indexActionName
	}

	if urlResult == cIntURLBadModule {
		h.send404NotFoundHTML(w, r.URL.String())
		return nil, http.StatusNotFound, nil //404 sent, finish this request
	}

	//Try to send raw static files
	staticPath := combineFullPath(rs.path+rs.module, target)
	if h.handleFileDownloads(w, staticPath) {
		return nil, http.StatusOK, nil
	}
	/* All static works done above, handle dynamic things below: */

	queryValues, pqErr := url.ParseQuery(r.URL.RawQuery)
	if pqErr != nil {
		errFmtStr := cStrParseQueryFailedError + pqErr.Error()
		h.httpSendError(w, errFmtStr, r.URL.String(), http.StatusInternalServerError, logger)
		return nil, http.StatusInternalServerError, nil
	}

	qv := h.parseQryValues(queryValues)
	reqData := h.getReqData(qv, target, logger, rs)
	return &reqData, http.StatusContinue, nil
}

//Handling the File Downloads for: static HTML, js, png, jpg, svg, css and so on..
func (h *httpGetHandler) handleFileDownloads(w http.ResponseWriter, path string) bool {
	if len(path) == 0 {
		return false
	}

	data, err := ioutil.ReadFile(string(path))
	if err != nil {
		return false
	}

	var contentType string
	if strings.HasSuffix(path, cStrSuffixCSS) {
		contentType = cStrTextCSS
	} else if strings.HasSuffix(path, cStrSuffixJS) {
		contentType = cStrAppJS
	} else if strings.HasSuffix(path, cStrSuffixPNG) {
		contentType = cStrImgPNG
	} else if strings.HasSuffix(path, cStrSuffixJPG) {
		contentType = cStrImgJPG
	} else if strings.HasSuffix(path, cStrSuffixSVG) {
		contentType = cStrImgSVG
	} else if strings.HasSuffix(path, cStrSuffixICO) {
		contentType = cStrImgICON
	} else {
		contentType = cStrTextHTML
	}

	w.Header().Add(cStrContentType, contentType)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return true
}
