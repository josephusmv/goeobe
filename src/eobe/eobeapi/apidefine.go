package eobeapi

import (
	impl "eobe/eobeapi/eobeapiimpl"
	"eobe/eobehttp"
	"strings"
)

//function names need preaction, import from implementation package
const cStrAPIRangeGetRows = impl.CAPIRangeGetRows
const cStrAPISaveFile = impl.CAPISaveFile
const CStrFilterMultiRowss = impl.CAPIFilterMultiRowss //---consider deprecate!

//Pre-Action
const cStrPreActSetRange = "setrange"
const cStrPreActSetFile = "setfile"
const cStrPreSetMultiRows = "filterrows"

//POST Action
const cStrSaveRslt = "(SaveRslt)"
const cStrSaveRows = "(SaveRows)"
const cStrSaveAll = "(SaveAll)"

type apiDefine struct {
	apiName        string
	preAction      string
	postAction     string
	implementation impl.ApiInf
	logger         eobehttp.HttpLoggerInf
}

//apiDefine should decide pre and post action by apiName and actionStr
func newApiDefine(callStr, action string, logger eobehttp.HttpLoggerInf) *apiDefine {
	//Decide post action by action string
	var postAction string
	switch action {
	case "":
		fallthrough
	case cStrSaveRslt:
		fallthrough
	case cStrSaveAll:
		fallthrough
	case cStrSaveRows:
		postAction = action
	default:
		return nil //unexpected action
	}

	//decide preaction by api name
	var preAction string
	switch {
	case strings.HasPrefix(callStr, CStrFilterMultiRowss):
		preAction = cStrPreSetMultiRows
	case strings.HasPrefix(callStr, cStrAPIRangeGetRows):
		preAction = cStrPreActSetRange
	case strings.HasPrefix(callStr, cStrAPISaveFile):
		preAction = cStrPreActSetFile
	}

	return &apiDefine{preAction: preAction, postAction: postAction, logger: logger}
}

func (apiDefn *apiDefine) doStepPreAction(ce callExec, preRslts CallSeqResult) {
	switch apiDefn.preAction {
	case cStrPreActSetRange:
		apiDefn.implementation.SetRangeSource(preRslts.MNames, preRslts.MultiRow)
	case cStrPreActSetFile:
		apiDefn.implementation.SetFileBytes(ce.fileMap)
	case cStrPreSetMultiRows:
		apiDefn.implementation.SetFilterSource(preRslts.MNames, preRslts.MultiRow)
	}
	return
}

func (apiDefn *apiDefine) doStepPostAction(ce callExec, callRslt *CallSeqResult, stepRslt map[string]string) CallSeqResult {

	if apiDefn.postAction == cStrSaveAll || apiDefn.postAction == cStrSaveRslt {
		if callRslt.SingleRow == nil {
			callRslt.SingleRow = make(map[string]string)
		}
		for k, v := range stepRslt {
			callRslt.SingleRow[k] = v
		}
	}

	if apiDefn.postAction == cStrSaveAll || apiDefn.postAction == cStrSaveRows {
		callRslt.MNames, callRslt.MultiRow = apiDefn.implementation.GetResultRows()
	}

	callRslt.ApiErr = impl.ApiSuccess()

	return *callRslt
}
