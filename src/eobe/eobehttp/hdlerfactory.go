package eobehttp

import (
	"net/http"
)

type httpMethodInf interface {
	//!!Important**: If HTTPStatus is not http.StatusContinue, then means internally, the http Response has been sent
	doMethod(r *http.Request) (HTTPStatus int, err error)
	initURLParser(string, string)
	getRequestData() *RequestData
}

type methodFactory struct {
	//Var for special cases
	indxAction   string // IF no action, like root path, still need an action name.
	page404fName string
}

//This function should have no Error reported
func (mf methodFactory) newMethodHandler(method string, logger HttpLoggerInf, rr resourceReaderInf, rw *responseWritter) httpMethodInf {
	var mthInf httpMethodInf

	//Route Method first
	switch method {
	case "GET":
		mthInf = newGetHandler(logger, rr, rw)
	case "POST":
		mthInf = newPostHandler(logger, rr, rw)
	default:
	}

	mthInf.initURLParser(mf.indxAction, mf.page404fName)

	return mthInf
}
