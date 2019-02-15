package eobeapiimpl

import (
	"strconv"
)

//filterMultiRows
//	Define: FilterMultiRows(startIndex, count)
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

//API filterMultiRows	_deprecate, dbActionExec could done this.
const CAPIFilterMultiRowss = "FilterMultiRows"

//newFilterMultiRows API filterMultiRows(paramSrcStr, paramRegexpStr) retIsValid
func newFilterMultiRows(apiParamInput string) (ApiInf, error) {
	var api filterMultiRows
	api.paramCount = 2

	//api.apiParamNames = []string{cStrValidateStrRexParamSrc, cStrValidateStrRexParamReg}
	api.apiRetrnnNames = []string{CAPIFilterMultiRowss}

	return &api, api.parseParameter(apiParamInput, api.paramCount, CAPIFilterMultiRowss)
}

const cStrFilterMultiRowsAPIParameterErr = "FilterMultiRows API error: %s"

func (api *filterMultiRows) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, *APIError) {
	values, err := api.getInputVarValues(qryKVMap, preCallRslts, api.paramCount, CAPIFilterMultiRowss)
	if err != nil {
		return nil, NewAPIErrorf(CErrBadCallError, cStrFilterMultiRowsAPIParameterErr, err.Error())
	}

	api.startIndex, err = strconv.Atoi(values[0])
	if err != nil {
		return nil, NewAPIErrorf(CErrBadCallError, cStrInvalidParameterTypeMore, CAPIFilterMultiRowss, values[0])
	}

	api.count, err = strconv.Atoi(values[1])
	if err != nil {
		return nil, NewAPIErrorf(CErrBadCallError, cStrInvalidParameterTypeMore, CAPIFilterMultiRowss, values[1])
	}

	start := api.startIndex
	if start < 0 {
		start = 0
	}

	end := start + api.count
	if end > len(api.msrcRows) {
		end = len(api.msrcRows) //use as count, because do slicing will includes this index.
	}
	if start >= end {
		start = end
	}

	result := make(map[string]string)
	result[cStrDBActionExecRetAffectedRows] = values[1]
	result[cStrDBActionExecRetLastIndex] = strconv.Itoa(end)

	api.mrows = api.msrcRows[start:end]
	api.mNames = api.msrcNames
	return result, ApiSuccess()
}

func (api *filterMultiRows) SetFilterSource(mNames []string, mrows [][]string) {
	api.msrcNames = mNames
	api.msrcRows = mrows
}

func (api *filterMultiRows) GetResultRows() ([]string, [][]string) {
	return api.mNames, api.mrows
}
