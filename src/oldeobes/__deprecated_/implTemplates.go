package eobeapiimpl

import "fmt"

//SPECIFICAPIIMPL
//	Define: SPECIFICAPIIMPL()
//	Return:
//		1. Map: map[name]values	temporary variables
//		2. Errors
type SPECIFICAPIIMPL struct {
	apiBase
	//Specific data
	paramCount int
}

//API SPECIFICAPIIMPL
const cStrSPECIFICAPIIMPL = "SPECIFICAPIIMPL"
const cStrSPECIFICAPIIMPLRetAffectedRows = "retAffectedRows"

//newAPISPECIFICAPIIMPL API SPECIFICAPIIMPL(paramSrcStr, paramRegexpStr) retIsValid
func newSPECIFICAPIIMPL(apiParamInput string) (ApiInf, error) {
	var api SPECIFICAPIIMPL
	api.paramCount = 4

	//api.apiParamNames = []string{cStrValidateStrRexParamSrc, cStrValidateStrRexParamReg}
	api.apiRetrnnNames = []string{cStrSPECIFICAPIIMPL}
	api.parseParameter(apiParamInput)

	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrSPECIFICAPIIMPL, api.paramCount)
	}

	return &api, nil
}

func (api *SPECIFICAPIIMPL) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	return nil, nil
}
