package eobeapi

import (
	impl "eobe/eobeapi/eobeapiimpl"
	"eobe/eobehttp"
	"strings"
)

//function names need preaction, import from implementation package
const cStrAPIRangeGetRows = impl.CAPIRangeGetRows

//Pre-Action
const cStrPreActSetRange = "setrange"

//POST Action
const cStrSaveRslt = "(SaveRslt)"
const cStrSaveRows = "(SaveRows)"
const cStrFilterRows = "(FilterRows)"
const cStrAsRange = "(AsRange)"
const cStrSaveFile = "(SaveFile)"

type stepAction struct {
	logger     eobehttp.HttpLoggerInf
	preAction  string
	postAction string
}

//stepAction should decide pre and post action by apiName and actionStr
func newStepAction(fullCallStr string, logger eobehttp.HttpLoggerInf) (*stepAction, string) {
	var callStr string

	//Decide post action by action string
	var postAction string
	switch {
	case strings.HasPrefix(fullCallStr, cStrSaveRslt):
		postAction = cStrSaveRslt
		callStr = string(callStr[len(cStrSaveRslt):])

	case strings.HasPrefix(fullCallStr, cStrSaveRows):
		postAction = cStrSaveRows
		callStr = string(callStr[len(cStrSaveRows):])

	case strings.HasPrefix(fullCallStr, cStrFilterRows):
		postAction = cStrFilterRows
		callStr = string(callStr[len(cStrFilterRows):])

	case strings.HasPrefix(fullCallStr, cStrAsRange):
		postAction = cStrAsRange
		callStr = string(callStr[len(cStrAsRange):])

	case strings.HasPrefix(fullCallStr, cStrSaveFile):
		postAction = cStrSaveFile
		callStr = string(callStr[len(cStrSaveFile):])
	}

	//decide preaction by api name
	var preAction string
	switch {
	case strings.HasPrefix(callStr, cStrAPIRangeGetRows):
		preAction = cStrPreActSetRange
	}

	return &stepAction{preAction: preAction, postAction: postAction, logger: logger}, callStr
}

func (sact stepAction) doStepPreAction(call apiDefine) {
	ce.logger.TraceDev(cStrRunPresteps, call.apiName)
	//presteps
	switch sact.preAction {
	case impl.CAPIRangeGetRows: //Refactory not show any names in higher package, use call.options to do bit calc
		call.implementation.SetRangeSource(rangeVar.MNames, rangeVar.MultiRow)
	case impl.CAPISaveFile: //Refactory not show any names in higher package, use call.options to do bit calc
		call.implementation.SetFileBytes(ce.fileMap)
	case impl.CAPIFilterMultiRowss:
		call.implementation.SetFilterSource(filterVar.MNames, filterVar.MultiRow)
	default:
	}
}

func (sact stepAction) doStepPostAction(call apiDefine) {
	//Save actions
	switch call.options {
	case cIntSingleResult:
		ce.logger.TraceDev(cStrRunPoststepsSnglSave, call.apiName)
		ce.copyResult(results)
	case cIntMultiRowResult:
		ce.logger.TraceDev(cStrRunPoststepsMultSave, call.apiName)
		if call.apiName == impl.CAPIGetMultiRows || call.apiName == impl.CAPIFilterMultiRowss || call.apiName == impl.CAPIRangeGetRows {
			mname, mrows := call.implementation.GetResultRows(ce.callRslt.MNames, ce.callRslt.MultiRow)
			if mname != nil && mrows != nil {
				ce.callRslt.MNames = mname
				ce.callRslt.MultiRow = mrows
			}
		}
	case cIntFilterRows:
		ce.logger.TraceDev(cStrRunPoststepFilter, call.apiName)
		if call.apiName == impl.CAPIGetMultiRows || call.apiName == impl.CAPIRangeGetRows {
			mname, mrows := call.implementation.GetResultRows(ce.callRslt.MNames, ce.callRslt.MultiRow)
			if mname != nil && mrows != nil {
				filterVar.MNames = mname
				filterVar.MultiRow = mrows
			}
		}
	case cIntAsRange:
		if call.apiName == impl.CAPIGetMultiRows || call.apiName == impl.CAPIFilterMultiRowss {
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
