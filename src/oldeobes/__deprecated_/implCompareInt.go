package eobeapiimpl

import (
	"fmt"
	"strconv"
)

//compareInt
//	Define: compareInt(var, eq/gt/ge/lt/le int)
//	Return:
//		1. Bool: success of not.
//		2. retCmpIntResult: TRUE/ FALSE
//		3. Errors
type compareInt struct {
	apiIntBase
	//Specific data
	paramCount int
}

//API ValidateStrRex
const cStrCompareInt = "CompareInt"
const cStrCompareIntRetCmpIntResult = "retCmpIntResult"

//newAPIValidateStrRegX API ValidateStrRegX(paramSrcStr, paramRegexpStr) retIsValid
func newCompareInt(apiParamInput string) (ApiInf, error) {
	var api compareInt
	api.paramCount = 3

	api.apiRetrnnNames = []string{cStrCompareIntRetCmpIntResult}
	api.parseParameter(apiParamInput)

	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrCompareInt, api.paramCount)
	}

	return &api, nil
}

func (api *compareInt) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrCompareInt, 3)
	}

	//first parameter
	strSrc, found := api.getParamValue(api.apiValueVarInput[0], qryKVMap, preCallRslts)
	if !found {
		return nil, fmt.Errorf(cStrParameterNotFound, api.apiValueVarInput[0])
	}

	var src, cval int
	var err error

	src, err = strconv.Atoi(strSrc)
	if err != nil {
		return nil, fmt.Errorf(cStrInvalidParameterTypeMore, cStrCompareInt, strSrc)
	}
	cval, err = strconv.Atoi(api.apiValueVarInput[2])
	if err != nil {
		return nil, fmt.Errorf(cStrInvalidParameterTypeMore, cStrCompareInt, api.apiValueVarInput[2])
	}

	success := api.doIntCompareAction(api.apiValueVarInput[1], src, cval)

	result := make(map[string]string)
	for _, v := range api.apiRetrnnNames {
		result[v] = strconv.FormatBool(success)
	}

	//CompareInt will carry result by ret var map, this is different to ValidateInt
	return result, nil
}
