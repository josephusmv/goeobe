package eobehttp

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime"
)

/* templtHandler
 *	This struct responsible for handling the HTTP template rendering.
 *  As HTTP templates have data binding, which uses a map to do that.
 *  When Done the respFetcher.FetchResponse() with response bytes.
 *  The bytes will convert into a map of response map as defined by struct htmlBindData{} above
 */
type templtHandler struct {
	templateName  string
	htmlTemplates *template.Template
	hData         TemplateData
}

//renderTemplates will send Template HTML, or error and end this request.
// need caller to help log some fatal errors
func (t *templtHandler) renderTemplates(w http.ResponseWriter, logger HttpLoggerInf) {
	tmplFullName := t.templateName
	/* I don;t knwon why I did this before....:-(
	if strings.HasSuffix(t.templateName, cStrSuffixTmplt) {
		tmplFullName = t.templateName
	} else {
		tmplFullName = t.templateName + cStrSuffixTmplt
	}
	*/

	if t.htmlTemplates == nil {
		_, file, line, _ := runtime.Caller(0)
		logger.TraceError(cStrProgramLogicError, file, line)
		http.Error(w, cStrProgramLogicError, http.StatusInternalServerError)
		return
	}

	err := t.htmlTemplates.ExecuteTemplate(w, tmplFullName, t.hData)
	if err != nil {
		errorStr := fmt.Sprintf(cStrExcuteTemplateError, tmplFullName, err.Error())
		logger.TraceDev(errorStr)
		http.Error(w, errorStr, http.StatusBadRequest)
		return
	}
}
