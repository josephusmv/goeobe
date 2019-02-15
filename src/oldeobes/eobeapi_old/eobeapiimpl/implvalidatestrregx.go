package eobeapiimpl

import "regexp"
import "fmt"
import "strconv"

//ValidateStrRegX use regexp to match a string
//	Define: ValidateStrRex(src, regex string, FAIL_DESC_STR)
//	true: regexp matches
//	nil map
//	nil error
type ValidateStrRegX struct {
	apiBase
	paramCount int
}

//API ValidateStrRex
const cStrValidateStrRexName = "ValidateStrRegX"

//const cStrValidateStrRexParamSrc = "paramSrcStr"
//const cStrValidateStrRexParamReg = "paramRegexpStr"
const cStrValidateStrRexReturnValid = "retIsValid"

//newAPIValidateStrRegX API ValidateStrRegX(paramSrcStr, paramRegexpStr) retIsValid
func newAPIValidateStrRegX(apiParamInput string) (*ApiInf, error) {
	var api ValidateStrRegX
	api.paramCount = 3

	//api.apiParamNames = []string{cStrValidateStrRexParamSrc, cStrValidateStrRexParamReg}
	api.apiRetrnnNames = []string{cStrValidateStrRexReturnValid}
	api.parseParameter(apiParamInput)

	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrValidateStrRexName, api.paramCount)
	}

	var retIf ApiInf
	retIf = &api

	return &retIf, nil
}

func (api ValidateStrRegX) RunAPI(qryKVMap map[string]string, callsResults map[string]string) (map[string]string, error) {
	var found bool
	value := make([]string, len(api.apiValueVarInput))
	for i, varName := range api.apiValueVarInput {
		value[i], found = api.getParamValue(varName, qryKVMap, callsResults)
		if !found {
			return nil, fmt.Errorf(cStrParameterNotFound, varName)
		}
	}

	ok := api.validateRegex(value[0], value[1])

	result := make(map[string]string)
	for _, v := range api.apiRetrnnNames {
		result[v] = strconv.FormatBool(ok)
	}

	if !ok {
		result[cStrRetFailDescStr] = value[2]
		return result, fmt.Errorf(cStrEmptyErrorStr, value[2])
	}

	return result, nil
}

func (api ValidateStrRegX) validateRegex(src, regex string) bool {
	regExp := regexp.MustCompile(regex)
	return regExp.MatchString(src)
}
