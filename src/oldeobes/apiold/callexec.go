package eobeapi

import (
	"eobe/eobeapi/eobeapiimpl"
	"eobe/eobehttp"
	"fmt"
)

type callExec struct {
	fileMap map[string][]byte
	logger  eobehttp.HttpLoggerInf
}

func (ce callExec) executeCalls(callsList []*apiDefine, qryKVMap map[string]string) (rslts CallSeqResult, resErr error) {

	tempResults := make(map[string]string)
	filterVar := CallSeqResult{}
	rangeVar := CallSeqResult{}

	//Execute call sequence
	for _, call := range callsList {

		ce.logger.TraceDev(cStrRunsAPI, call.apiName)
		results, err := call.implementation.RunAPI(qryKVMap, tempResults)
		if err != nil {
			rslts.SingleRow = results //Error may need some error info stored in result returned
			resErr = fmt.Errorf(cStrFailedWhenExe, call.apiName, err.Error())
			return
		}
		for k, v := range results {
			tempResults[k] = v
		}
	}
}

func (ce callExec) stepPostAction(call apiDefine) {
	//Save actions
	switch call.options {
	case cIntSingleResult:
		ce.logger.TraceDev(cStrRunPoststepsSnglSave, call.apiName)
		ce.copyResult(results)
	case cIntMultiRowResult:
		ce.logger.TraceDev(cStrRunPoststepsMultSave, call.apiName)
		if call.apiName == eobeapiimpl.CAPIGetMultiRows || call.apiName == eobeapiimpl.CAPIFilterMultiRowss || call.apiName == eobeapiimpl.CAPIRangeGetRows {
			mname, mrows := call.implementation.GetResultRows(ce.callRslt.MNames, ce.callRslt.MultiRow)
			if mname != nil && mrows != nil {
				ce.callRslt.MNames = mname
				ce.callRslt.MultiRow = mrows
			}
		}
	case cIntFilterRows:
		ce.logger.TraceDev(cStrRunPoststepFilter, call.apiName)
		if call.apiName == eobeapiimpl.CAPIGetMultiRows || call.apiName == eobeapiimpl.CAPIRangeGetRows {
			mname, mrows := call.implementation.GetResultRows(ce.callRslt.MNames, ce.callRslt.MultiRow)
			if mname != nil && mrows != nil {
				filterVar.MNames = mname
				filterVar.MultiRow = mrows
			}
		}
	case cIntAsRange:
		if call.apiName == eobeapiimpl.CAPIGetMultiRows || call.apiName == eobeapiimpl.CAPIFilterMultiRowss {
			mname, mrows := call.implementation.GetResultRows(ce.callRslt.MNames, ce.callRslt.MultiRow)
			if mname != nil && mrows != nil {
				rangeVar.MNames = mname
				rangeVar.MultiRow = mrows
			}
		}
	case cIntNotSave:
	default:
	}

}

func (ce callExec) copyResult(src map[string]string) (dst map[string]string) {
	dst = make(map[string]string)

	for k, v := range results {
		dst[k] = v
	}
}
