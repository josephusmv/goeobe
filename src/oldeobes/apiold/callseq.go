package eobeapi

import (
	cltmgrcmd "eobe/eobecltmgmt"
	"eobe/eobedb"
	"eobe/eobehttp"
	"fmt"
)

//CallSeq  CallSequences class
type CallSeq struct {
	callExec
	af       apiFactory
	hr       *cltmgrcmd.Herald
	callRslt CallSeqResult
	logger   loggerInf
}

//GetCallSequence Builder function
//	return nil, nil means no API calls, need not do anything
func GetCallSequence() *CallSeq {
	return &CallSeq{}
}

//SetLogger  Get HTTP logger inf from higher layer
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

//Execute execute the call sequences.
func (cs *CallSeq) Execute(qryKVMap map[string]string, calls []string) (rslts CallSeqResult, resErr error) {
	if qryKVMap == nil {
		resErr = fmt.Errorf(cStrErrorInput)
		return
	}

	var callsList []*apiDefine
	callsList, resErr = cs.buildAPICallsList(calls)
	if resErr != nil {
		return
	}

	//Temporary variables
	return cs.callRslt, nil
}
