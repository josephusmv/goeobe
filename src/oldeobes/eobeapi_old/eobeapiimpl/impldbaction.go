package eobeapiimpl

import (
	"eobe/eobedb"
	"fmt"
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
}

const cStrDBActionExec = "dbActionExec"

//newAPIValidateStrRegX API ValidateStrRegX(paramSrcStr, paramRegexpStr) retIsValid
func newdbActionExec(apiParamInput string, dbQry eobedb.DBQueryInf, qryDfn eobedb.QueryDefn) (*ApiInf, error) {
	if dbQry == nil {
		return nil, fmt.Errorf(cStrDBInterfaceImplError)
	}

	var api dbActionExec
	api.apiName = cStrDBActionExec
	api.initParams(apiParamInput, dbQry, qryDfn)

	var retIf ApiInf
	retIf = &api

	return &retIf, nil
}

func (api dbActionExec) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	qryRes, err := api.runDBActions(qryKVMap, preCallRslts)
	if err != nil {
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

	return result, nil
}
