package eobehttp

import (
	"io/ioutil"
	"net/http"
	"strings"
)

/* httpPostHandler
 *	POST method is not for static resources for most times,
 *  so the httpPostHandler will only responsible for parse the query forms.
 */
type httpPostHandler struct {
	httpBaseHandler
	target string
	qv     map[string]string
	fMap   map[string][]byte
}

func newPostHandler(logger HttpLoggerInf, rr resourceReaderInf, rw *responseWritter) httpMethodInf {
	var h httpPostHandler

	h.logger = logger
	h.rscReader = rr
	h.rspWrtter = rw

	return &h
}

func (h *httpPostHandler) doMethod(r *http.Request) (HTTPStatus int, err error) {
	//Paser
	err = r.ParseForm() //ignore errors here.
	if err != nil {
		h.logger.TraceDev(cStrParsePostFormError, err.Error())
	}
	reqCt := strings.ToLower(r.Header.Get(cStrContentType))
	if strings.HasPrefix(reqCt, cStrMultipart) {
		//For content type(r.Header.Get("content-type")) is multipart/form-data, we do ParseMultipartForm
		err = r.ParseMultipartForm(cIntDefaultMaxMemory)
		if err != nil {
			h.rspWrtter.sendServerInternalError(r.URL, err, h.logger)
			return
		}
	}
	h.qv = h.parseQryValues(r.Form)
	postFormValues := h.parseQryValues(r.PostForm) //also add form value not in post form to here...
	for k, v := range postFormValues {
		if _, found := h.qv[k]; !found {
			h.qv[k] = v
		}
	}

	//get File bytes
	h.fMap, err = h.readFileBytes(r, h.qv)
	if err != nil {
		h.rspWrtter.sendServerInternalError(r.URL, err, h.logger)
		return
	}

	//get requested Target
	//	target is the HTML template full path or resources full path or Action
	h.target, _ = h.up.parseURL(r.URL.Path, h.logger)

	return http.StatusContinue, nil
}

func (h httpPostHandler) readFileBytes(r *http.Request, qryKVMap map[string]string) (map[string][]byte, error) {
	fQryKey, found := qryKVMap[CStrUploadFileInKey]
	if !found {
		return nil, nil
	}

	//Ignore the file header, only accept the file bytes.
	file, _, err := r.FormFile(fQryKey)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []byte
	data, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fileMap := make(map[string][]byte)
	fileMap[fQryKey] = data

	return fileMap, nil
}

func (h *httpPostHandler) getRequestData() *RequestData {
	var reqData RequestData
	reqData.QueryTarget = h.target
	reqData.QueryKeyValueMap = h.qv
	reqData.Logger = h.logger
	reqData.UploadedFiles = h.fMap
	return &reqData
}
