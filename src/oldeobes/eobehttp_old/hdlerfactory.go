package eobehttp

import (
	"net/http"
)

type methodParam struct {
	resp   http.ResponseWriter
	req    *http.Request
	logger HttpLoggerInf
}

type httpMethodInf interface {
	//!!Important**: If HTTPStatus is not http.StatusContinue, then means internally, the http Response has been sent
	doMethod(param methodParam) (reqData *RequestData, HTTPStatus int, err error)
}

//This function should have no Error reported
func methodFactory(w http.ResponseWriter, r *http.Request, logger HttpLoggerInf, rs *RequestServer) (httpMethodInf, methodParam) {
	var param methodParam
	param.req = r
	param.resp = w
	param.logger = logger

	var mthInf httpMethodInf

	//Route Method first
	switch r.Method {
	case "GET":
		mthInf = &httpGetHandler{httpBaseHandler{rs: rs}}
	case "POST":
		mthInf = &httpPostHandler{httpBaseHandler{rs: rs}}
	default:
	}

	return mthInf, param
}
