package eobeapiimpl

import (
	"fmt"
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
func newLogOutUser(apiParamInput string) (*ApiInf, error) {
	var api logOutUser
	api.apiRetrnnNames = []string{cStrLogOutUserRetSuccess}
	api.parseParameter(apiParamInput)

	api.paramCount = 0
	if len(api.apiValueVarInput) < api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrLogOutUser, api.paramCount)
	}

	var retIf ApiInf
	retIf = &api
	return &retIf, nil
}

func (api *logOutUser) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	//constant result
	retKVMap := make(map[string]string)
	retKVMap[cStrLogOutUserRetSuccess] = strconv.FormatBool(true)

	err := api.hr.SignOutUser()
	if err != nil {
		return nil, fmt.Errorf(cStrCookieServerInternalError, err)
	}

	return retKVMap, nil
}
