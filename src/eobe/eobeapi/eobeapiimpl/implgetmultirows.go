package eobeapiimpl

//getMultiRows run a DB query to get multiple row results
import (
	"strconv"
)

//	Define: getMultiRows(DBActionName; expectedParams; whereParams)
//	true: regexp matches
//	nil map
//	nil error
type getMultiRows struct {
	apiDBBase
	mNames []string
	mrows  [][]string
}

//API ValidateStrRex
const CAPIGetMultiRows = "GetMultiRows"

//newgetMultiRows apiParamInput should be processed by upper callers
func newgetMultiRows(dbActionName string, apiParamInput string) (ApiInf, error) {
	var api getMultiRows
	api.apiName = CAPIGetMultiRows
	api.dbActionName = dbActionName
	api.apiParamInput = apiParamInput

	return &api, nil
}

func (api *getMultiRows) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, *APIError) {
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

func (api *getMultiRows) GetResultRows() ([]string, [][]string) {
	return api.mNames, api.mrows
}
