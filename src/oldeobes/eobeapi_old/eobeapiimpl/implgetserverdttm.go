package eobeapiimpl

import (
	"fmt"
	"time"
)

//getServerDTTM
//	Define: getServerDTTM()
//	Return:
//		1. Map: map[name]values	temporary variables
//		2. Errors
type getServerDTTM struct {
	apiBase
	//Specific data
	paramCount int
}

//API getServerDTTM
const cStrGetServerDTTM = "GetServerDTTM"
const cStrGetServerDTTMRetServerDTTM = "retServerDTTM"

//newAPIgetServerDTTM API getServerDTTM(paramSrcStr, paramRegexpStr) retIsValid
func newgetServerDTTM(apiParamInput string) (*ApiInf, error) {
	var api getServerDTTM
	api.paramCount = 1

	//api.apiParamNames = []string{cStrValidateStrRexParamSrc, cStrValidateStrRexParamReg}
	api.apiRetrnnNames = []string{cStrGetServerDTTM}
	api.parseParameter(apiParamInput)

	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrGetServerDTTM, api.paramCount)
	}

	var retIf ApiInf
	retIf = &api

	return &retIf, nil
}

func (api *getServerDTTM) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	// port should support both literal and variable
	format, found := api.getParamValue(api.apiValueVarInput[0], qryKVMap, preCallRslts)
	if !found {
		format = api.apiValueVarInput[0]
	}

	t := time.Now()
	nowStr := fmt.Sprintf(format,
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	rsltMap := make(map[string]string)
	rsltMap[cStrGetServerDTTMRetServerDTTM] = nowStr
	return rsltMap, nil
}
