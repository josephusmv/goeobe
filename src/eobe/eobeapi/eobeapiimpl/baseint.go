package eobeapiimpl

import "fmt"

type apiIntBase struct {
	apiBase
}

func (api apiIntBase) doIntCompareAction(operator string, src, value int) bool {
	switch operator {
	case cStrEqualTo:
		fallthrough
	case cStrSymbolEqualTo:
		return src == value
	case cStrGreaterThan:
		fallthrough
	case cStrSymbolGreaterThan:
		return src > value
	case cStrSmallerThan:
		fallthrough
	case cStrSymbolSmallerThan:
		return src < value
	case cStrGreaterEqualThan:
		fallthrough
	case cStrSymbolGreaterEqualThan:
		return src >= value
	case cStrSmallerEqualThan:
		fallthrough
	case cStrSymbolSmallerEqualThan:
		return src <= value
	case cStrNotEqualTo:
		fallthrough
	case cStrSymbolNotEqualTo:
		return src != value
	default:
		return src == value
	}
}
func (api apiIntBase) validateSymbol(operator string) bool {
	switch operator {
	case cStrEqualTo:
		fallthrough
	case cStrSymbolEqualTo:
		fallthrough
	case cStrGreaterThan:
		fallthrough
	case cStrSymbolGreaterThan:
		fallthrough
	case cStrSmallerThan:
		fallthrough
	case cStrSymbolSmallerThan:
		fallthrough
	case cStrGreaterEqualThan:
		fallthrough
	case cStrSymbolGreaterEqualThan:
		fallthrough
	case cStrSmallerEqualThan:
		fallthrough
	case cStrSymbolSmallerEqualThan:
		fallthrough
	case cStrNotEqualTo:
		fallthrough
	case cStrSymbolNotEqualTo:
		return true
	default:
		return false
	}
}

//getInputVarValuesForInt Int operation follows a role: VAR, literal symbot(>, =, < ...), and literal number(0, 1, 2...), others var
func (api apiIntBase) getInputVarValuesForInt(qryKVMap map[string]string, callsResults map[string]string, expected int, apiName string) ([]string, error) {
	var found bool
	lenInut := len(api.apiValueVarInput)
	value := make([]string, lenInut)

	//first parameter must be variable:
	value[0], found = api.getParamValue(api.apiValueVarInput[0], qryKVMap, callsResults)
	if !found {
		return nil, fmt.Errorf(cStrParameterNotFound, api.apiValueVarInput[0])
	}

	//only accept literal for the second as symbol
	value[1] = api.apiValueVarInput[1]

	//the third one could be either literal or const user defined
	value[2], found = api.getParamValue(api.apiValueVarInput[2], qryKVMap, callsResults)
	if !found { //if not found, then we think it's literal
		value[2] = api.apiValueVarInput[2]
	}

	if lenInut == 4 {
		//the third one could be either literal or const user defined
		value[3], found = api.getParamValue(api.apiValueVarInput[3], qryKVMap, callsResults)
		if !found { //if not found, then we think it's literal
			value[3] = api.apiValueVarInput[3]
		}
	}

	if len(value) != expected {
		return nil, fmt.Errorf(cStrParameterCountError, apiName, expected)
	}

	return value, nil
}
