package eobeapiimpl

import "fmt"

//getCurrentUser
//	Define: GetCurrentUser()
//		will get an empty string if not login
//	Return:
//		1. Map: map[retCurrentUserName]UserName
//		2. Errors
type getCurrentUser struct {
	apiBase
	//Specific data
	paramCount int
}

//API GetCurrentUser
const cStrGetCurrentUser = "GetCurrentUser"
const cStrGetCurrentUserRetUserName = "retCurrentUserName"

//newGetCurrentUser API getCurrentUser(paramSrcStr, paramRegexpStr) retIsValid
func newGetCurrentUser(apiParamInput string) (*ApiInf, error) {
	var api getCurrentUser
	api.paramCount = 0

	//api.apiParamNames = []string{cStrValidateStrRexParamSrc, cStrValidateStrRexParamReg}
	api.apiRetrnnNames = []string{cStrGetCurrentUserRetUserName}
	api.parseParameter(apiParamInput)

	if len(api.apiValueVarInput) < api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrGetCurrentUser, api.paramCount)
	}

	var retIf ApiInf
	retIf = &api

	return &retIf, nil
}

func (api *getCurrentUser) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	uName := api.getCurrentUser()

	rsltMap := make(map[string]string)
	rsltMap[cStrGetCurrentUserRetUserName] = uName

	return rsltMap, nil
}
