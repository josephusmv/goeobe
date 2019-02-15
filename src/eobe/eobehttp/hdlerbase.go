package eobehttp

import (
	"net/url"
)

type httpBaseHandler struct {
	rscReader resourceReaderInf
	rspWrtter *responseWritter
	logger    HttpLoggerInf
	up        *urlParser
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

func (h *httpBaseHandler) initURLParser(indexAction string, page404FName string) {
	h.up = &urlParser{rr: h.rscReader}
	h.up.indexAction = indexAction
	h.up.page404FName = page404FName
}
