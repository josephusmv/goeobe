package httpsample3

import (
	"eobe/eobehttp"
	"io/ioutil"
	"strings"
)

//RespFetcherImpl ...
type RespFetcherImpl struct {
}

const cSampleTempIdxActionName = "GetIndex"

//FetchResponse As an sample, we only return the index.html
func (rf *RespFetcherImpl) FetchResponse(req eobehttp.RequestData) (rsp eobehttp.ResponseData, err error) {
	logger := req.Logger
	logger.TraceDev("Fetch Response for request: %s.%s", req.Module, req.QueryTarget)

	var target string
	tIdx := strings.LastIndex(req.QueryTarget, "/")
	if tIdx > 0 && tIdx < len(req.QueryTarget) {
		target = req.QueryTarget[tIdx+1:]
	} else {
		target = req.QueryTarget //this is already best efforts... will not works for most times.
	}
	switch target {
	case cSampleTempIdxActionName:
		return rf.indexAction(req)
	case "UploadFile":
		return rf.saveFile(req)
	default:
		return rsp, logger.TraceError("Cannot recognize the action: %s", req.QueryTarget)
	}
}

func (rf *RespFetcherImpl) saveFile(req eobehttp.RequestData) (rsp eobehttp.ResponseData, err error) {
	for k, v := range req.UploadedFiles {
		ioutil.WriteFile(k, v, 0666)
		break
	}
	rsp.ContentType = "application/json"
	rsp.HTMLTmpltName = ""
	rsp.Body = []byte("{Success:true}")

	return
}

func (rf *RespFetcherImpl) indexAction(req eobehttp.RequestData) (rsp eobehttp.ResponseData, err error) {
	var tp eobehttp.TemplateData
	tp.Rows = [][]string{
		[]string{"Value1111", "Value1112", "Vao & bar斯卡拉lue1113"},
		[]string{"V卡alue1222", "Value2222", "Value22卡拉23"},
		[]string{"Value3卡拉331", "Value3卡332", "Value卡拉3333"},
	}

	rsp.ContentType = "text/html"
	rsp.HTMLTmpltName = "index.tmplt"
	rsp.HTMLTmpltData = tp
	rsp.Body = nil
	rsp.CookieList = req.CookieList //just copy paste for index.
	return
}
