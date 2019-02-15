package eobeapiimpl

import (
	"fmt"
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
func newLogInUser(apiParamInput string) (*ApiInf, error) {
	var api logInUser
	api.apiRetrnnNames = []string{cStrRetSuccess}
	api.parseParameter(apiParamInput)

	api.paramCount = 2
	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrLogInUser, api.paramCount)
	}

	var retIf ApiInf
	retIf = &api
	return &retIf, nil
}

func (api *logInUser) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	values, err := api.getInputVarValues(qryKVMap, preCallRslts)
	if err != nil {
		return nil, err
	}

	if len(values) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrLogInUser, api.paramCount)
	}

	var expDays int
	expDays, err = strconv.Atoi(values[1])
	//fmt.Println("&&& user " + values[0] + "Expire days " + values[1])
	if err != nil {
		return nil, fmt.Errorf(cStrErrorParameterTypeError, cStrLogInUser, api.apiValueVarInput[1], values[1])
	}

	uname := api.getCurrentUser()

	retKVMap := make(map[string]string)

	//Already login
	if uname != "" {
		retKVMap[cStrRetSuccess] = cStrFalse
		retKVMap[cStrLogInUserRetUserName] = uname
		return retKVMap, nil
	}

	err = api.hr.SignInUser(values[0], expDays)
	if err != nil {
		return nil, fmt.Errorf(cStrCookieServerInternalError, err.Error())
	}

	retKVMap[cStrRetSuccess] = cStrTure
	retKVMap[cStrLogInUserRetUserName] = values[0]
	return retKVMap, nil
}
