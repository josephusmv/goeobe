package eobeapiimpl

import (
	"fmt"
	"strconv"
)

//validateInt
//	Define: validateInt(var, eq/gt/ge/lt/le int, FAIL_DESC_STR)
//	Return:
//		1. Bool: success of not.
//		2. retCmpIntResult: TRUE/ FALSE
//		3. Errors
type validateInt struct {
	apiIntBase
	//Specific data
	paramCount int
}

//API ValidateStrRex
const cStrValidateInt = "ValidateInt"
const cStrValidateIntRetCmpIntResult = "retVldIntResult"

//newAPIValidateStrRegX API ValidateStrRegX(paramSrcStr, paramRegexpStr) retIsValid
func newValidateInt(apiParamInput string) (*ApiInf, error) {
	var api validateInt
	api.paramCount = 4

	api.apiRetrnnNames = []string{cStrValidateIntRetCmpIntResult}
	api.parseParameter(apiParamInput)

	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrValidateInt, api.paramCount)
	}

	var retIf ApiInf
	retIf = &api

	return &retIf, nil
}

func (api *validateInt) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	values, err := api.getInputVarValuesForInt(qryKVMap, preCallRslts)
	if err != nil {
		return nil, err
	}

	if len(values) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrValidateInt, 3)
	}

	//first parameter
	var src, cval int64

	src, err = strconv.ParseInt(values[0], 0, 64)
	if err != nil {
		return nil, fmt.Errorf(cStrInvalidParameterTypeMore, cStrValidateInt, values[0])
	}
	cval, err = strconv.ParseInt(values[2], 0, 64)
	if err != nil {
		return nil, fmt.Errorf(cStrInvalidParameterTypeMore, cStrValidateInt, values[2])
	}

	success := api.doIntCompareAction(values[1], int(src), int(cval))

	result := make(map[string]string) //error or good result
	for _, v := range api.apiRetrnnNames {
		result[v] = strconv.FormatBool(success)
	}

	if !success {
		result[cStrRetFailDescStr] = values[3]
		return result, fmt.Errorf(cStrEmptyErrorStr, values[3])
	}

	//ValidateInt will carry result by ret var map, this is different to ValidateInt
	return result, nil
}
