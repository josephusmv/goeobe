package eobeapiimpl

import (
	"eobe/eobedb"
	"fmt"
	"strings"
)

type apiDBBase struct {
	apiBase
	//Specific data
	dbActionName string
	//init during construct
	apiParamInput    string
	dbQryInf         eobedb.DBQueryInf
	qryDef           eobedb.QueryDefn
	qryExptVarNames  []string
	qryWhereVarNames []string
}

const cStrDBActionExecRetAffectedRows = "retAffectedRows"
const cStrDBActionExecRetLastIndex = "retLastIndex"

func (api *apiDBBase) prepareQueryDefine(apiParamInput string, qryDefn eobedb.QueryDefn) error {
	//Prepare for api.qryDef    eobedb.QueryDefn
	api.qryDef = qryDefn

	//Prepare for api.qryExptVarNames and api.qryWhereVarNames
	qryPart := strings.Split(apiParamInput, cStrSemiColon)
	/*	This should be unreacheable codes, strings.Split will give a "" for "" and impossible to get a zero array.
		if len(qryPart) <= 0 {
			return fmt.Errorf(cStrErrorParsingDBAction, apiParamInput)
		}
	*/

	switch api.qryDef.QueryType {
	case eobedb.CStrDBINSERT: // INSERT don't need where
		api.qryExptVarNames = strings.Split(qryPart[0], cStrComma)
	case eobedb.CStrDBDELETE: // DELECT don't need Expected
		//api.qryWhereVarNames = strings.Split(qryPart[0], cStrComma)
		fallthrough
	case eobedb.CStrDBSELECT: // SELECT don't need Expected
		api.qryWhereVarNames = strings.Split(qryPart[0], cStrComma)
	case eobedb.CStrDBUPDATE:
		api.qryWhereVarNames = strings.Split(qryPart[1], cStrComma)
		api.qryExptVarNames = strings.Split(qryPart[0], cStrComma)
	default:
		return fmt.Errorf(cStrUnexpectedDBAction, api.qryDef.QueryType)
	}

	if len(api.qryExptVarNames) == 1 && api.qryExptVarNames[0] == "" {
		api.qryExptVarNames = nil
	}

	if len(api.qryWhereVarNames) == 1 && strings.Trim(api.qryWhereVarNames[0], cStrSpace) == "" {
		api.qryWhereVarNames = nil
	}

	return nil

}

func (api *apiDBBase) runDBActions(qryKVMap map[string]string, preCallRslts map[string]string) (qryRes eobedb.QueryResult, apiErr *APIError) {
	qryData := &eobedb.QueryData{}

	qryData.QryActionDfn = api.qryDef
	if api.qryDef.QueryActionName == "" || api.qryDef.TableName == "" {
		apiErr = NewAPIErrorf(CErrServerInternalError, cStrDBInvalidQueryDefn)
		return
	}

	var err error
	switch qryData.QryActionDfn.QueryType {
	case eobedb.CStrDBINSERT: // INSERT don't need where
		qryData, err = api.getExpectedParamValues(qryData, qryKVMap, preCallRslts)
	case eobedb.CStrDBDELETE: // DELECT don't need Expected
		fallthrough
	case eobedb.CStrDBSELECT: // SELECT don't need Expected
		qryData, err = api.getParameterParamValues(qryData, qryKVMap, preCallRslts)
	case eobedb.CStrDBUPDATE:
		qryData, err = api.getExpectedParamValues(qryData, qryKVMap, preCallRslts)
		qryData, err = api.getParameterParamValues(qryData, qryKVMap, preCallRslts)
	default:
		err = fmt.Errorf(cStrUnexpectedDBAction, qryData.QryActionDfn.QueryType)
	}

	if err != nil {
		apiErr = NewAPIErrorf(CErrBadCallError, cStrDBParseParameterErr, api.qryDef.QueryActionName, err.Error())
		return
	}

	qryRes, err = api.dbQryInf.ExecDBAction(*qryData, api.logger)
	if qryRes.QueryErr != nil {
		apiErr = NewAPIErrorf(CErrServerInternalError, cStrErrorFromEobeDB, qryRes.QueryErr.Error())
		return
	}

	return qryRes, ApiSuccess()
}

func (api *apiDBBase) getExpectedParamValues(qryData *eobedb.QueryData, qryKVMap map[string]string, preCallRslts map[string]string) (*eobedb.QueryData, error) {
	if len(api.qryDef.ExpectedColNames) != len(api.qryExptVarNames) {
		return qryData, fmt.Errorf(cStrDBAParamMismatch, api.qryDef.ExpectedColNames, api.qryExptVarNames)
	}

	var found bool
	qryData.ExpectedValues = make([]string, len(api.qryExptVarNames))
	for i, v := range api.qryExptVarNames {
		qryData.ExpectedValues[i], found = api.getParamValue(v, qryKVMap, preCallRslts)
		if !found {
			return qryData, fmt.Errorf(cStrParameterNotFound, v)
		}
	}

	return qryData, nil
}

func (api *apiDBBase) getParameterParamValues(qryData *eobedb.QueryData, qryKVMap map[string]string, preCallRslts map[string]string) (*eobedb.QueryData, error) {
	if len(api.qryDef.ParamterColNames) != len(api.qryWhereVarNames) && len(api.qryDef.WhereReadyStr) == 0 {
		//ParamterColNames is a not supported in this version....so ignore this line during corverage testing.
		return qryData, fmt.Errorf(cStrDBAParamMismatch, api.qryDef.ParamterColNames, api.qryWhereVarNames)
	}

	var found bool
	qryData.ParameterValues = make([]string, len(api.qryWhereVarNames))
	for i, v := range api.qryWhereVarNames {
		qryData.ParameterValues[i], found = api.getParamValue(v, qryKVMap, preCallRslts)
		if !found {
			return qryData, fmt.Errorf(cStrParameterNotFound, v)
		}
	}

	return qryData, nil
}

func (api *apiDBBase) IsDBAction() (string, bool) {
	return api.dbActionName, true
}

func (api *apiDBBase) SetDBInfo(dbQry eobedb.DBQueryInf, dbActDefn eobedb.QueryDefn) error {
	api.dbQryInf = dbQry
	api.qryDef = dbActDefn

	err := api.prepareQueryDefine(api.apiParamInput, api.qryDef)
	if err != nil {
		return err
	}

	retVarCount := len(api.qryDef.ExpectedColNames) + 2
	api.apiRetrnnNames = make([]string, retVarCount)
	api.apiRetrnnNames[0] = cStrDBActionExecRetAffectedRows
	api.apiRetrnnNames[1] = cStrDBActionExecRetLastIndex
	for i, v := range api.qryDef.ExpectedColNames {
		api.apiRetrnnNames[i+2] = v
	}

	return nil
}
