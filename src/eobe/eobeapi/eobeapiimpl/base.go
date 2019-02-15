package eobeapiimpl

import (
	cltmgrcmd "eobe/eobecltmgmt"
	"eobe/eobedb"
	"eobe/eobehttp"
	"fmt"
	"strings"
)

const cStrRetFailDescStr = "FAIL_DESC_STR"

type apiBase struct {
	//Values constant  --- part of API define.
	apiName string
	//apiParamNames  []string //Name in
	apiRetrnnNames []string //Name out

	//values parsed --- user input
	apiValueVarInput []string //Variable name for parameter

	//Usr constant strings
	ucStr map[string]string
	hr    *cltmgrcmd.Herald

	//logger from Request handler
	logger eobehttp.HttpLoggerInf
}

//apiParamInput Calls defined parameters, like: "?username, UNAMEREX" from "ValidateStrREX: ?username, UNAMEREX"
func (ab *apiBase) parseParameter(apiParamInput string, expectParmCount int, apiName string) error {
	params := strings.Split(apiParamInput, cStrComma)
	ab.apiValueVarInput = make([]string, len(params)) //even for empty string, len(params) at least = 1

	for i := range params {
		paramVar := strings.Trim(params[i], cStrSpace)
		if paramVar == cStrNotApplicable {
			ab.apiValueVarInput[i] = ""
		}

		ab.apiValueVarInput[i] = paramVar
	}

	if len(ab.apiValueVarInput) == 1 && ab.apiValueVarInput[0] == "" {
		ab.apiValueVarInput = nil
	}

	if ab.apiValueVarInput == nil && expectParmCount == 0 {
		return nil
	}

	if len(ab.apiValueVarInput) != expectParmCount {
		return fmt.Errorf(cStrParameterCountError, apiName, expectParmCount)
	}

	return nil
}

const cStrPreResultPrefix = '^'
const cStrQryValuesPrefix = '?'
const cStrUserConstPrefix = '$'

func (ab apiBase) getFromPreResult(name string, callsResults map[string]string) (value string, found bool) {
	name = strings.Trim(name, string(cStrPreResultPrefix))
	value, found = callsResults[name]
	return
}

func (ab apiBase) getFromQryValues(name string, qryKVMap map[string]string) (value string, found bool) {
	name = strings.Trim(name, string(cStrQryValuesPrefix))
	value, found = qryKVMap[name]
	return
}

func (ab *apiBase) SetDataSrc(ucKVMap map[string]string, hr *cltmgrcmd.Herald) {
	ab.ucStr = ucKVMap
	ab.hr = hr
}

func (ab apiBase) getFromUserConstValue(name string) (value string, found bool) {
	// UserConst should not be trimed.
	name = strings.Trim(name, string(cStrUserConstPrefix))
	value, found = ab.ucStr[name]
	return
}

func (ab apiBase) getParamValue(name string, qryKVMap map[string]string, callsResults map[string]string) (value string, found bool) {
	name = strings.Trim(name, cStrSpace)
	if len(name) == 0 {
		return "", false
	}
	switch name[0] {
	case cStrPreResultPrefix:
		return ab.getFromPreResult(name, callsResults)
	case cStrQryValuesPrefix:
		return ab.getFromQryValues(name, qryKVMap)
	case cStrUserConstPrefix:
		return ab.getFromUserConstValue(name)
	default:
		return "", false
	}
}

func (ab apiBase) getParamName(name string) string {
	name = strings.Trim(name, cStrSpace)
	switch name[0] {
	case cStrPreResultPrefix:
		return strings.Trim(name, string(cStrPreResultPrefix))
	case cStrQryValuesPrefix:
		return strings.Trim(name, string(cStrQryValuesPrefix))
	case cStrUserConstPrefix:
		return strings.Trim(name, string(cStrUserConstPrefix))
	default:
		return name
	}
}

func (ab apiBase) getInputVarValues(qryKVMap map[string]string, callsResults map[string]string, expected int, apiName string) ([]string, error) {
	var found bool
	values := make([]string, len(ab.apiValueVarInput))
	for i, varName := range ab.apiValueVarInput {
		values[i], found = ab.getParamValue(varName, qryKVMap, callsResults)
		if !found {
			return nil, fmt.Errorf(cStrParameterNotFound, varName)
		}
	}

	if len(values) < expected {
		return nil, fmt.Errorf(cStrParameterCountError, apiName, expected)
	}

	return values, nil
}

func (ab apiBase) getCurrentUser() string {
	//check currently logged in user info using IP-Port and SID
	user, _ := ab.hr.GetCurrentUser(ab.logger)
	return user
}

func (ab *apiBase) SetRangeSource(mNames []string, mrows [][]string) {
}

func (ab *apiBase) GetResultRows() ([]string, [][]string) {
	return nil, nil
}

func (ab *apiBase) SetFileBytes(map[string][]byte) {
	return
}

func (ab *apiBase) SetLogger(logger eobehttp.HttpLoggerInf) {
	ab.logger = logger
}

func (ab *apiBase) SetFilterSource(mNames []string, mrows [][]string) {

}

func (ab *apiBase) IsDBAction() (string, bool) {
	return "", false
}

func (ab *apiBase) SetDBInfo(dbQry eobedb.DBQueryInf, dbActDefn eobedb.QueryDefn) error {
	return nil
}