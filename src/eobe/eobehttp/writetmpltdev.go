package eobehttp

import (
	"fmt"
	"html/template"
	"net/http"
)

/* templtDevHandler
 *	This struct responsible for handling the HTTP template rendering.
 *  As HTTP templates have data binding, which uses a map to do that.
 *  When Done the respFetcher.FetchResponse() with response bytes.
 *  The bytes will convert into a map of response map as defined by struct htmlBindData{} above
 */
type templtDevHandler struct {
	headerWritter
	templateName string

	htmlTmpltStrs []string
	hData         TemplateData
}

func newtempltDevHandler(rr resourceReaderInf, resp *ResponseData) *templtDevHandler {
	t := templtDevHandler{templateName: resp.HTMLTmpltName,
		hData: resp.HTMLTmpltData}

	var i int
	tMap := rr.getTemplatesDev()
	t.htmlTmpltStrs = make([]string, len(tMap))
	for _, v := range tMap {
		t.htmlTmpltStrs[i] = v
		i++
	}

	//default header options
	t.setDefaultRespHeaders()

	return &t
}

//renderTemplates will send Template HTML, or error and end this request.
// need caller to help log some fatal errors
func (t *templtDevHandler) renderTemplates(w http.ResponseWriter, logger HttpLoggerInf) {

	t.addDefaultRespHeaders(w)

	if t.htmlTmpltStrs == nil || len(t.htmlTmpltStrs) <= 0 {
		logger.TraceError(cStrHTMLSrcFileInvalid)
		http.Error(w, cStrHTMLSrcFileInvalid, http.StatusServiceUnavailable)
		return
	}

	tmps, err := template.ParseFiles(t.htmlTmpltStrs...)
	if tmps == nil || err != nil {
		errStr := fmt.Sprintf(cStrErrorParseTemplates, err.Error())
		logger.TraceError(errStr)
		http.Error(w, errStr, http.StatusServiceUnavailable)
		return
	}

	err = tmps.ExecuteTemplate(w, t.templateName, t.hData)
	if err != nil {
		errorStr := fmt.Sprintf(cStrExcuteTemplateError, t.templateName, err.Error())
		logger.TraceDev(errorStr)
		http.Error(w, errorStr, http.StatusBadRequest)
		return
	}
}
