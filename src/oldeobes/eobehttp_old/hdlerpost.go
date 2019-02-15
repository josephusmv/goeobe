package eobehttp

import (
	"fmt"
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
}

func (h *httpPostHandler) doMethod(param methodParam) (reqData *RequestData, HTTPStatus int, err error) {
	rs := h.rs
	qryKeyValueMap := make(map[string]string)
	logger := param.logger
	r := param.req
	w := param.resp

	//get requested Target
	qryTarget, urlResult := h.parseURLPath(r, rs.module, logger)
	if urlResult != cIntURLOK || qryTarget == "" {
		errorStr := fmt.Sprintf(cStrInvalidPostRequestError, r.URL.String())
		logger.TraceDev(errorStr)
		http.Error(w, errorStr, http.StatusBadRequest)
		//return http.StatusBadRequest to indicates that request has been handled,
		//	tmplErr for further debugging
		return nil, http.StatusBadRequest, nil
	}

	//Parse Form
	err = r.ParseForm() //ignore errors here.
	if err != nil {
		logger.TraceError(cStrParsePostFormError, err.Error())
	}
	reqCt := strings.ToLower(r.Header.Get(cStrContentType))
	if strings.HasPrefix(reqCt, cStrMultipart) {
		//For content type(r.Header.Get("content-type")) is multipart/form-data, we do ParseMultipartForm
		err = r.ParseMultipartForm(cIntDefaultMaxMemory)
		if err != nil {
			logger.TraceError(cStrParseMultipartFormError, err.Error())
			return nil, http.StatusInternalServerError, nil
		}
	}

	qryKeyValueMap = h.parseQryValues(r.Form)
	postFormValues := h.parseQryValues(r.PostForm) //also add form value not in post form to here...
	for k, v := range postFormValues {
		if _, found := qryKeyValueMap[k]; !found {
			qryKeyValueMap[k] = v
		}
	}

	//get File bytes
	fMap, err := h.readFileBytes(r, qryKeyValueMap)
	if err != nil {
		logger.TraceError(err.Error())
	}

	req := h.getReqData(qryKeyValueMap, qryTarget, logger, rs)
	req.UploadedFiles = fMap

	return &req, http.StatusContinue, nil
}

func (h *httpPostHandler) readFileBytes(r *http.Request, qryKVMap map[string]string) (map[string][]byte, error) {
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
