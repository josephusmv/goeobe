package eobeapiimpl

import (
	"fmt"
	"strconv"
)

//filterMultiRows
//	Define: filterMultiRows(startIndex, count)
//	Return:
//		1. Map: map[name]values	temporary variables
//		2. Errors
type filterMultiRows struct {
	apiBase
	//Specific data
	paramCount int
	startIndex int
	count      int

	//filterSource
	msrcNames []string
	msrcRows  [][]string
	mNames    []string
	mrows     [][]string
}

//API filterMultiRows
const CAPIFilterMultiRowss = "FilterMultiRows"

//newFilterMultiRows API filterMultiRows(paramSrcStr, paramRegexpStr) retIsValid
func newFilterMultiRows(apiParamInput string) (*ApiInf, error) {
	var api filterMultiRows
	api.paramCount = 2

	//api.apiParamNames = []string{cStrValidateStrRexParamSrc, cStrValidateStrRexParamReg}
	api.apiRetrnnNames = []string{CAPIFilterMultiRowss}
	api.parseParameter(apiParamInput)

	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, CAPIFilterMultiRowss, api.paramCount)
	}

	var retIf ApiInf
	retIf = &api

	return &retIf, nil
}

func (api *filterMultiRows) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	//filter multi rows just save the start index and counts, all real works don in GetResultRows().
	values, err := api.getInputVarValues(qryKVMap, preCallRslts)
	if err != nil {
		return nil, err
	}

	if len(values) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, CAPIFilterMultiRowss, api.paramCount)
	}

	api.startIndex, err = strconv.Atoi(values[0])
	if err != nil {
		return nil, fmt.Errorf(cStrInvalidParameterTypeMore, CAPIFilterMultiRowss, values[0])
	}

	api.count, err = strconv.Atoi(values[1])
	if err != nil {
		return nil, fmt.Errorf(cStrInvalidParameterTypeMore, CAPIFilterMultiRowss, values[1])
	}

	//Filter
	start := api.startIndex
	end := api.startIndex + api.count

	if start < 0 {
		start = 0
	}

	api.mNames = api.msrcNames

	if end > len(api.msrcNames) {
		api.mrows = api.msrcRows[start:]
	} else {
		api.mrows = api.msrcRows[start:end]
	}

	return nil, nil
}

func (api *filterMultiRows) SetFilterSource(mNames []string, mrows [][]string) {
	api.msrcNames = mNames
	api.msrcRows = mrows
}

func (api *filterMultiRows) GetResultRows(srcMNames []string, srcMultiRow [][]string) ([]string, [][]string) {
	return api.mNames, api.mrows
}
