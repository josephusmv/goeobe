package eobehttp

import (
	"fmt"
	"html/template"
	"net/http"
)

/* templtHandler
 *	This struct responsible for handling the HTTP template rendering.
 *  As HTTP templates have data binding, which uses a map to do that.
 *  When Done the respFetcher.FetchResponse() with response bytes.
 *  The bytes will convert into a map of response map as defined by struct htmlBindData{} above
 */
type templtHandler struct {
	headerWritter
	templateName  string
	htmlTemplates *template.Template
	hData         TemplateData
}

const cStrHTMLSrcFileInvalid = `HTMLSrcFileInvalid: 
		Source HTML files are not properly loaded, 
		Server is down, 
		Administrator should check ` + cStrHTTPMainLoggerFileName

func newTempltHandler(rr resourceReaderInf, resp *ResponseData) *templtHandler {
	t := templtHandler{templateName: resp.HTMLTmpltName,
		hData:         resp.HTMLTmpltData,
		htmlTemplates: rr.getTemplates()}

	//default header options
	t.setDefaultRespHeaders()

	return &t
}

//renderTemplates will send Template HTML, or error and end this request.
// need caller to help log some fatal errors
func (t *templtHandler) renderTemplates(w http.ResponseWriter, logger HttpLoggerInf) {
	t.addDefaultRespHeaders(w)

	if t.htmlTemplates == nil {
		logger.TraceError(cStrHTMLSrcFileInvalid)
		http.Error(w, cStrHTMLSrcFileInvalid, http.StatusServiceUnavailable)
		return
	}

	err := t.htmlTemplates.ExecuteTemplate(w, t.templateName, t.hData)
	if err != nil {
		errorStr := fmt.Sprintf(cStrExcuteTemplateError, t.templateName, err.Error())
		logger.TraceDev(errorStr)
		http.Error(w, errorStr, http.StatusBadRequest)
		return
	}
}
