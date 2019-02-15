package eobeapiimpl

import "fmt"
import "strconv"

//ValidateStrEqual use regexp to match a string
//	Define: ValidateStrEqual(src, regex string, FAIL_DESC_STR)
//	true: regexp matches
//	nil map
//	nil error
type ValidateStrEqual struct {
	apiBase
	paramCount int
}

//API ValidateStrEqual
const cStrValidateStrEqualName = "ValidateStrEqual"
const cStrValidateStrEqualReturnValid = "retIsValid"

//newAPIValidateStrEqual API ValidateStrEqual(paramSrcStr, paramDstStr) retIsValid
func newAPIValidateStrEqual(apiParamInput string) (ApiInf, error) {
	var api ValidateStrEqual
	api.paramCount = 3

	api.apiRetrnnNames = []string{cStrValidateStrEqualReturnValid}
	api.parseParameter(apiParamInput)

	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrValidateStrEqualName, api.paramCount)
	}

	return &api, nil
}

func (api ValidateStrEqual) RunAPI(qryKVMap map[string]string, callsResults map[string]string) (map[string]string, error) {
	var found bool
	value := make([]string, len(api.apiValueVarInput))
	for i, varName := range api.apiValueVarInput {
		value[i], found = api.getParamValue(varName, qryKVMap, callsResults)
		if !found {
			return nil, fmt.Errorf(cStrParameterNotFound, varName)
		}
	}

	ok := false
	if value[0] == value[1] {
		ok = true
	}

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
