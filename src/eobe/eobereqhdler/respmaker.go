package eobereqhdler

import (
	"eobe/eobeapi"
	"eobe/eobehttp"
)

type respMakerInf interface {
	makeResponse(eobeapi.CallSeqResult, string, eobehttp.HttpLoggerInf) (eobehttp.ResponseData, error)
}

//doMakeResponse   a factory method to use differnt RespMakerInf depending on ExpectedFMT defined in resources Define.
func doMakeResponse(apiRslt eobeapi.CallSeqResult, actualResName string, respType string, logger eobehttp.HttpLoggerInf) (eobehttp.ResponseData, error) {
	var rspMaker respMakerInf
	switch respType {
	case cStrExpectedXML:
	case cStrExpectedJSON:
		rspMaker = respJSONMaker{}
	case cStrExpectedHTML:
		rspMaker = respHTMLMaker{}
	default:
	}

	return rspMaker.makeResponse(apiRslt, actualResName, logger)
}
