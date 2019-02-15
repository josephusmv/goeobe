package eobeapi

import (
	impl "eobe/eobeapi/eobeapiimpl"
	clt "eobe/eobecltmgmt"
	edb "eobe/eobedb"
	eh "eobe/eobehttp"
)

const cStrErrBuildAPI = "Build API List Error: %s"

//CallSeq  CallSequences class
type CallSeq struct {
	callExec
	af apiFactory
	hr *clt.Herald
}

//GetCallSequence Builder function
//	return nil, nil means no API calls, need not do anything
func GetCallSequence(logger eh.HttpLoggerInf) *CallSeq {
	var cs CallSeq
	cs.logger = logger
	return &cs
}

//InitAPIFactory Init the API Factory
func (cs *CallSeq) InitAPIFactory(userConsts map[string]string,
	hr *clt.Herald,
	dbQry edb.DBQueryInf,
	dbActMapmap map[string]*edb.QueryDefn) {

	factory := apiFactory{
		dbQry:       dbQry,
		dbActDfnMap: dbActMapmap,
		userConsts:  userConsts,
		hr:          hr}
	cs.af = factory
	cs.hr = hr
}

//SetFileMap  ...
func (cs *CallSeq) SetFileMap(fileMap map[string][]byte) {
	cs.fileMap = fileMap
}

func (cs CallSeq) buildAPICallsList(calls []string) ([]*apiDefine, error) {

	callsList := make([]*apiDefine, len(calls))
	for i, callStr := range calls {
		apiDfn, err := cs.af.makeNewAPI(callStr, cs.logger)
		if err != nil {
			return nil, err
		}

		callsList[i] = apiDfn
	}

	return callsList, nil
}

//Execute execute the call sequences.
func (cs *CallSeq) Execute(qryKVMap map[string]string, calls []string) CallSeqResult {
	var callRslt CallSeqResult
	if calls == nil || len(calls) <= 0 { //Just ignore if this is an empty call
		callRslt.ApiErr = impl.ApiSuccess()
		return callRslt
	}

	if qryKVMap == nil { //means no query map, make an empty one
		qryKVMap = make(map[string]string)
	}

	//build []*apiDefine list
	callsList, err := cs.buildAPICallsList(calls)
	if err != nil {
		callRslt.ApiErr = impl.NewAPIErrorf(impl.CErrServerInternalError, cStrErrBuildAPI, err.Error())
		return callRslt
	}

	//Run API execution and return
	return cs.runExecute(callsList, qryKVMap)
}
