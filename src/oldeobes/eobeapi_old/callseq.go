package eobeapi

import (
	"eobe/eobeapi/eobeapiimpl"
	cltmgrcmd "eobe/eobecltmgmt"
	"eobe/eobedb"
	"eobe/eobehttp"
	"fmt"
)

//CallSeqResult struct holding Call Sequence Result
type CallSeqResult struct {
	SingleRow map[string]string
	MNames    []string
	MultiRow  [][]string
}

//CallSeq  CallSequences class
type CallSeq struct {
	af        eobeapiimpl.ApiFactory
	hr        *cltmgrcmd.Herald
	callsList []*apiDefine
	callRslt  CallSeqResult
	fileMap   map[string][]byte
	logger    loggerInf
}

//GetCallSequence Builder function
//	return nil, nil means no API calls, need not do anything
func GetCallSequence() *CallSeq {
	return &CallSeq{}
}

//SetLogger  Get HTTP logger inf from higher layer, convert to internal alias
func (cs *CallSeq) SetLogger(logger eobehttp.HttpLoggerInf) {
	cs.logger = loggerInf(logger)
	cs.af.SetLogger(logger)
}

//PrepareForAPICalls ...
func (cs *CallSeq) PrepareForAPICalls(userConsts map[string]string, hr *cltmgrcmd.Herald) error {
	cs.hr = hr
	return cs.af.PrepareForAPICalls(userConsts, cs.hr)
}

//PrepareForDBActions Load DB related interfaces and defines.
func (cs *CallSeq) PrepareForDBActions(dbQry eobedb.DBQueryInf, dbActMapmap map[string]*eobedb.QueryDefn) error {
	return cs.af.PrepareForDBActions(dbQry, dbActMapmap)
}

func (cs *CallSeq) SetFileMap(fileMap map[string][]byte) {
	cs.fileMap = fileMap
}

func (cs *CallSeq) copyResult(results map[string]string) {
	if cs.callRslt.SingleRow == nil {
		cs.callRslt.SingleRow = make(map[string]string)
	}

	for k, v := range results {
		cs.callRslt.SingleRow[k] = v
	}

}

//Execute execute the call sequences.
func (cs *CallSeq) Execute(qryKVMap map[string]string, calls []string) (rslts CallSeqResult, resErr error) {
	if qryKVMap == nil {
		resErr = fmt.Errorf(cStrErrorInput)
		return
	}

	resErr = cs.buildAPICallsList(calls)
	if resErr != nil {
		return
	}

	//Temporary variables
	tempResults := make(map[string]string)
	filterVar := CallSeqResult{}
	rangeVar := CallSeqResult{}

	//Execute call sequence
	for _, call := range cs.callsList {
		cs.logger.TraceDev(cStrRunPresteps, call.apiName)
		//presteps
		switch call.apiName {
		case eobeapiimpl.CAPIRangeGetRows: //Refactory not show any names in higher package, use call.options to do bit calc
			call.implementation.SetRangeSource(rangeVar.MNames, rangeVar.MultiRow)
		case eobeapiimpl.CAPISaveFile: //Refactory not show any names in higher package, use call.options to do bit calc
			call.implementation.SetFileBytes(cs.fileMap)
		case eobeapiimpl.CAPIFilterMultiRowss:
			call.implementation.SetFilterSource(filterVar.MNames, filterVar.MultiRow)
		default:
		}

		cs.logger.TraceDev(cStrRunsAPI, call.apiName)
		results, err := call.implementation.RunAPI(qryKVMap, tempResults)
		if err != nil {
			rslts.SingleRow = results //Error may need some error info stored in result returned
			resErr = fmt.Errorf(cStrFailedWhenExe, call.apiName, err.Error())
			return
		}

		//Save actions
		switch call.options {
		case cIntSingleResult:
			cs.logger.TraceDev(cStrRunPoststepsSnglSave, call.apiName)
			cs.copyResult(results)
		case cIntMultiRowResult:
			cs.logger.TraceDev(cStrRunPoststepsMultSave, call.apiName)
			if call.apiName == eobeapiimpl.CAPIGetMultiRows || call.apiName == eobeapiimpl.CAPIFilterMultiRowss || call.apiName == eobeapiimpl.CAPIRangeGetRows {
				mname, mrows := call.implementation.GetResultRows(cs.callRslt.MNames, cs.callRslt.MultiRow)
				if mname != nil && mrows != nil {
					cs.callRslt.MNames = mname
					cs.callRslt.MultiRow = mrows
				}
			}
		case cIntFilterRows:
			cs.logger.TraceDev(cStrRunPoststepFilter, call.apiName)
			if call.apiName == eobeapiimpl.CAPIGetMultiRows || call.apiName == eobeapiimpl.CAPIRangeGetRows {
				mname, mrows := call.implementation.GetResultRows(cs.callRslt.MNames, cs.callRslt.MultiRow)
				if mname != nil && mrows != nil {
					filterVar.MNames = mname
					filterVar.MultiRow = mrows
				}
			}
		case cIntAsRange:
			if call.apiName == eobeapiimpl.CAPIGetMultiRows || call.apiName == eobeapiimpl.CAPIFilterMultiRowss {
				mname, mrows := call.implementation.GetResultRows(cs.callRslt.MNames, cs.callRslt.MultiRow)
				if mname != nil && mrows != nil {
					rangeVar.MNames = mname
					rangeVar.MultiRow = mrows
				}
			}
		case cIntNotSave:
		default:
		}

		for k, v := range results {
			tempResults[k] = v
		}
	}

	return cs.callRslt, nil
}

func (cs *CallSeq) buildAPICallsList(calls []string) error {
	if calls == nil || len(calls) <= 0 {
		return nil
	}

	cs.callsList = make([]*apiDefine, len(calls))
	for i, callStr := range calls {
		apiName, api, saveTo, err := cs.af.MakeNewAPI(callStr)
		if err != nil {
			return err
		}

		var apiDfn apiDefine
		apiDfn.apiName = apiName
		apiDfn.implementation = *api
		apiDfn.options = saveTo

		cs.callsList[i] = &apiDfn
	}
	return nil
}
