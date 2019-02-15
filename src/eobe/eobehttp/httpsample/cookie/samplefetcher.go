package httpsample3

import (
	"eobe/eobehttp"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//RespFetcherImpl ...
type RespFetcherImpl struct {
	cookieCount int
}

const cSampleTempIdxActionName = "GetIndex"

//FetchResponse As an sample, we only return the index.html
func (rf *RespFetcherImpl) FetchResponse(req eobehttp.RequestData) (rsp eobehttp.ResponseData, err error) {
	logger := req.Logger
	logger.TraceDev("Fetch Response for request: %s.%s", req.Module, req.QueryTarget)

	switch req.QueryTarget {
	case cSampleTempIdxActionName:
		return rf.indexAction(req)
	case "AddCookie":
		return rf.addCookie(req)
	default:
		return rsp, logger.TraceError("Cannot recognize the action: %s", req.QueryTarget)
	}
}

func (rf *RespFetcherImpl) addCookie(req eobehttp.RequestData) (rsp eobehttp.ResponseData, err error) {
	for k, v := range req.UploadedFiles {
		ioutil.WriteFile(k, v, 0666)
		break
	}
	rsp.ContentType = "application/json"
	rsp.HTMLTmpltName = ""
	rsp.Body = []byte("{Success:true}")

	//make a new cookie
	newName := fmt.Sprintf("CookieName_%d", rf.cookieCount)
	newValu := fmt.Sprintf("Cookie Value : %d", rf.cookieCount)
	expire := time.Now().Add(time.Hour * 72)
	newCookie := http.Cookie{Name: newName, Value: newValu, Path: "/", Expires: expire, MaxAge: 86400}

	rsp.CookieList = make([]*http.Cookie, 1)
	rsp.CookieList[0] = &newCookie
	rf.cookieCount++
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
