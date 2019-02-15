package eobeapiimpl

import (
	"strconv"
)

//logOutUser
//	Define: logOutUser()
//		If already logged out, will return with success and do nothing
//	Return:
//		1. Map:
//			[retSuccess]always true
//		2. Errors
type logOutUser struct {
	apiBase
	//Specific data
	paramCount int
}

//API ValidateStrRex
const cStrLogOutUser = "LogOutUser"
const cStrLogOutUserRetSuccess = "retSuccess"

//newAPIValidateStrRegX API ValidateStrRegX(paramSrcStr, paramRegexpStr) retIsValid
func newLogOutUser(apiParamInput string) (ApiInf, error) {
	var api logOutUser
	api.apiRetrnnNames = []string{cStrLogOutUserRetSuccess}

	api.paramCount = 0

	return &api, api.parseParameter(apiParamInput, api.paramCount, cStrLogOutUser)
}

func (api *logOutUser) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, *APIError) {
	//constant result
	retKVMap := make(map[string]string)
	retKVMap[cStrLogOutUserRetSuccess] = strconv.FormatBool(true)

	err := api.hr.SignOutUser(api.logger)
	if err != nil {
		return nil, NewAPIErrorf(CErrServerInternalError, cStrCookieServerInternalError, err)
	}

	return retKVMap, ApiSuccess()
}
