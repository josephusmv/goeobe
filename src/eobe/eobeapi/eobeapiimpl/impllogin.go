package eobeapiimpl

import (
	"strconv"
)

//logInUser
//	Define: logInUser(username, expireDays)
//		LogInUsr API only add cookie, not do permission check.
//		If already logged in a user, will return with success and the logged in user name
//	Return:
//		1. Bool: success of not.
//		2. Map:
//			[retSuccess]BOOL, [retLoginUser]string
//		3. Errors
type logInUser struct {
	apiBase
	//Specific data
	paramCount int
}

//API ValidateStrRex
const cStrLogInUser = "LogInUser"
const cStrLogInUserRetUserName = "retLoginUser"

//newAPIValidateStrRegX API ValidateStrRegX(paramSrcStr, paramRegexpStr) retIsValid
func newLogInUser(apiParamInput string) (ApiInf, error) {
	var api logInUser
	api.apiRetrnnNames = []string{cStrRetSuccess}
	api.paramCount = 2

	return &api, api.parseParameter(apiParamInput, api.paramCount, cStrLogInUser)
}

const cStrLoginAPIParameterErr = "LogIn API error: %s"

func (api *logInUser) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, *APIError) {
	values, err := api.getInputVarValues(qryKVMap, preCallRslts, api.paramCount, CAPIFilterMultiRowss)
	if err != nil {
		return nil, NewAPIErrorf(CErrBadCallError, cStrLoginAPIParameterErr, err.Error())
	}

	var expDays int
	expDays, err = strconv.Atoi(values[1])
	if err != nil {
		return nil, NewAPIErrorf(CErrBadCallError, cStrErrorParameterTypeError, cStrLogInUser, api.apiValueVarInput[1], values[1])
	}

	uname := api.getCurrentUser()

	retKVMap := make(map[string]string)

	//Already login
	if uname != "" {
		retKVMap[cStrRetSuccess] = cStrFalse
		retKVMap[cStrLogInUserRetUserName] = uname
		return retKVMap, ApiSuccess()
	}

	err = api.hr.SignInUser(values[0], expDays, api.logger)
	if err != nil {
		return nil, NewAPIErrorf(CErrServerInternalError, cStrCookieServerInternalError, err.Error())
	}

	retKVMap[cStrRetSuccess] = cStrTure
	retKVMap[cStrLogInUserRetUserName] = values[0]
	return retKVMap, ApiSuccess()
}
