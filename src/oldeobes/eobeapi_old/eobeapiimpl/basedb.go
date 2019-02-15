package eobeapiimpl

import (
	"eobe/eobedb"
	"fmt"
	"strings"
)

type apiDBBase struct {
	apiBase
	//Specific data
	//init during construct
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
	if len(qryPart) <= 0 {
		return fmt.Errorf(cStrErrorParsingDBAction, apiParamInput)
	}

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

	if len(api.qryWhereVarNames) == 1 && api.qryWhereVarNames[0] == "" {
		api.qryWhereVarNames = nil
	}

	return nil

}

func (api *apiDBBase) runDBActions(qryKVMap map[string]string, preCallRslts map[string]string) (qryRes eobedb.QueryResult, err error) {
	qryData := &eobedb.QueryData{}

	qryData.QryActionDfn = api.qryDef

	switch qryData.QryActionDfn.QueryType {
	case eobedb.CStrDBINSERT: // INSERT don't need where
		qryData, err = api.getExpectedParamValues(qryData, qryKVMap, preCallRslts)
	case eobedb.CStrDBDELETE: // DELECT don't need Expected
		//api.qryWhereVarNames = strings.Split(qryPart[0], cStrComma)
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
		return qryRes, err
	}

	qryRes, err = api.dbQryInf.ExecDBAction(*qryData, api.logger)
	if qryRes.QueryErr != nil {
		return qryRes, qryRes.QueryErr
	}

	return
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

func (api *apiDBBase) initParams(apiParamInput string, dbQry eobedb.DBQueryInf, qryDfn eobedb.QueryDefn) error {
	api.dbQryInf = dbQry
	err := api.prepareQueryDefine(apiParamInput, qryDfn)
	if err != nil {
		return err
	}

	//Return name should be: {"retAffectedRows", "retLastIndex", DB action ExpectedColNames}
	retVarCount := len(qryDfn.ExpectedColNames) + 2
	api.apiRetrnnNames = make([]string, retVarCount)
	api.apiRetrnnNames[0] = cStrDBActionExecRetAffectedRows
	api.apiRetrnnNames[1] = cStrDBActionExecRetLastIndex
	for i, v := range qryDfn.ExpectedColNames {
		api.apiRetrnnNames[i+2] = v
	}

	return nil
}
