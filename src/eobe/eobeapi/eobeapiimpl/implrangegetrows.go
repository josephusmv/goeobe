package eobeapiimpl

//RangeGetRows run a DB query to get multiple row results
import (
	"eobe/eobedb"
	"strconv"
)

//	Define: RangeGetRows(DBActionName; expectedParams; whereParams)
//	true: regexp matches
//	nil map
//	nil error
type rangeGetRows struct {
	apiDBBase
	msrcNames []string
	msrcRows  [][]string
	mNames    []string
	mrows     [][]string
}

//API ValidateStrRex
const CAPIRangeGetRows = "RangeGetRows"

//newRangeGetRows apiParamInput should be processed by upper callers
func newRangeGetRows(dbActionName string, apiParamInput string) (ApiInf, error) {
	var api rangeGetRows
	api.apiName = CAPIGetMultiRows
	api.dbActionName = dbActionName
	api.apiParamInput = apiParamInput

	return &api, nil
}

func (api *rangeGetRows) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, *APIError) {
	//Modify preCallRslts to get put range values into it.
	rQryMap := make(map[string]string) //copy preCallRslts
	for k, v := range preCallRslts {
		rQryMap[k] = v
	}

	//range execute
	qryRes := make([]eobedb.QueryResult, len(api.msrcRows))
	var err *APIError
	rowAffected := int64(0)
	for i, row := range api.msrcRows {
		for i, name := range api.msrcNames {
			rQryMap[name] = row[i]
		}
		//fmt.Println(rQryMap)
		qryRes[i], err = api.runDBActions(qryKVMap, rQryMap)
		if err.HasError() {
			return nil, err
		}
		rowAffected += qryRes[i].AffectedRows
	}

	result := make(map[string]string)
	result[cStrDBActionExecRetAffectedRows] = strconv.FormatInt(rowAffected, 10)
	result[cStrDBActionExecRetLastIndex] = strconv.FormatInt(qryRes[len(api.msrcRows)-1].LastIndex, 10)
	for i := range api.msrcRows {
		for j, v := range api.apiRetrnnNames {
			if j < 2 {
				continue
			}
			if len(qryRes[i].QueryRows) == 0 {
				break
			}
			result[v] = qryRes[i].QueryRows[0][j-2]
			break
		}
	}

	//Save multi-rows results
	api.mNames = api.qryDef.ExpectedColNames
	for _, v := range qryRes {
		api.mrows = append(api.mrows, v.QueryRows...)
	}

	return result, ApiSuccess()
}

func (api *rangeGetRows) GetResultRows() ([]string, [][]string) {
	return api.mNames, api.mrows
}

func (api *rangeGetRows) SetRangeSource(mNames []string, mrows [][]string) {
	api.msrcNames = mNames
	api.msrcRows = mrows
}
