package eobeapiimpl

import (
	"strconv"
)

//dbActionExec Execute the DB actions, Internally treated as API.
//	This Action only support Single row select, if need multiple rows return, should use: API: GetRowsFromDB
//	Define: DB_ACTION_NAME(expectedParams; whereParams)
//	true if row affected or returned values > 0
//  retAffectedRows, retLastIndex, select results.
//	nil error
type dbActionExec struct {
	apiDBBase
	mNames []string
	mrows  [][]string
}

const cStrDBActionExec = "dbActionExec"

//newAPIValidateStrRegX API ValidateStrRegX(paramSrcStr, paramRegexpStr) retIsValid
func newdbActionExec(dbActionName string, apiParamInput string) (ApiInf, error) {
	var api dbActionExec
	api.apiName = cStrDBActionExec
	api.dbActionName = dbActionName
	api.apiParamInput = apiParamInput

	return &api, nil
}

func (api *dbActionExec) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, *APIError) {
	qryRes, err := api.runDBActions(qryKVMap, preCallRslts)
	if err.HasError() {
		return nil, err
	}

	result := make(map[string]string)
	result[cStrDBActionExecRetAffectedRows] = strconv.FormatInt(qryRes.AffectedRows, 10)
	result[cStrDBActionExecRetLastIndex] = strconv.FormatInt(qryRes.LastIndex, 10)
	for i, v := range api.apiRetrnnNames {
		if i < 2 {
			continue
		}
		if len(qryRes.QueryRows) == 0 {
			break
		}
		result[v] = qryRes.QueryRows[0][i-2]
	}

	//Save multi-rows results
	api.mNames = api.qryDef.ExpectedColNames
	api.mrows = qryRes.QueryRows

	return result, ApiSuccess()
}

func (api *dbActionExec) GetResultRows() ([]string, [][]string) {
	return api.mNames, api.mrows
}
