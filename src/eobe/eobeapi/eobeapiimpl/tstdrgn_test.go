package eobeapiimpl

import (
	"eobe/eobedb"
	"fmt"
	"testing"
)

func testGetDBActionsForRange() []*testDBQueryParam {
	qryDfn := make([]*testDBQueryParam, 4)

	//SELECT
	qryRangeTble := eobedb.QueryDefn{
		QueryActionName:  "GetRangeValues",
		TableName:        "DBRANGETEST_TBL",
		QueryType:        "SELECT",
		WhereReadyStr:    "srchkey=?",
		ExpectedColNames: []string{"testfield1"}}
	qryRangeTbleParam := make(map[string]string)
	qryRangeTbleParam["srchkey"] = "2014"
	qryDfn[0] = &testDBQueryParam{
		dbActName:  "GetRangeValues",
		ParamInput: "?srchkey",
		qryDefn:    qryRangeTble,
		qryKVMap:   qryRangeTbleParam}

	//SELECT
	qryValueTble := eobedb.QueryDefn{
		QueryActionName:  "GetValuesByRange",
		TableName:        cTstStrTableName,
		QueryType:        "SELECT",
		WhereReadyStr:    "testfield1=?",
		ExpectedColNames: []string{"id", "username", "testfield1"}}
	qryDfn[1] = &testDBQueryParam{
		dbActName:  "GetValuesByRange",
		ParamInput: "^testfield1",
		qryDefn:    qryValueTble,
		qryKVMap:   nil}

	//Get All for filterrows api
	qrySelect := eobedb.QueryDefn{
		QueryActionName:  "GetByUName",
		TableName:        cTstStrTableName,
		QueryType:        "SELECT",
		WhereReadyStr:    "username=?",
		ExpectedColNames: []string{"id", "username", "testfield1", "testfield3"}}
	qrySelectParam := make(map[string]string)
	qrySelectParam["username"] = "sadmin"
	qryDfn[2] = &testDBQueryParam{
		dbActName:  "GetByUName",
		ParamInput: "?username",
		qryDefn:    qrySelect,
		qryKVMap:   qrySelectParam}

	return qryDfn
}

func testDoGetMultirows(apiParamInput string, param *testDBQueryParam, dbQry eobedb.DBQueryInf) (ApiInf, map[string]string) {
	const apiName = CAPIGetMultiRows

	api, err := GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")

	if _, dbAction := api.IsDBAction(); dbAction {
		api.SetDBInfo(dbQry, param.qryDefn)
	}

	result, apiErr := api.RunAPI(param.qryKVMap, nil)
	checkAPIError(apiErr, "RunAPI() "+apiName+" error.")

	testDumpResultMap(result, apiName+": "+apiParamInput)

	return api, result
}

func testRunRangeGetRows(param *testDBQueryParam, dbQry eobedb.DBQueryInf, preAPI ApiInf, preResult map[string]string) {
	const apiName = CAPIRangeGetRows
	const apiParamInput = "GetValuesByRange$$ ^testfield1"

	api, err := GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")

	if _, dbAction := api.IsDBAction(); dbAction {
		api.SetDBInfo(dbQry, param.qryDefn)
	}

	api.SetRangeSource(preAPI.GetResultRows())

	result, apiErr := api.RunAPI(param.qryKVMap, preResult)
	checkAPIError(apiErr, "RunAPI() "+apiName+" error.")

	testDumpResultMap(result, apiName+": "+apiParamInput)
	testDumpMultiRows(api.GetResultRows())
}

//TestDBRangeRows Smoke test for DB API RangeGetRows
//	go test -v  -run DBRangeRows
func TestDBRangeRows(t *testing.T) {
	dbQry, dbSrv := testConnectDB()
	defer dbSrv.StopDBServer()

	dbParaList := testGetDBActionsForRange()
	const apiParamInput = "GetRangeValues$$ ?srchkey"

	api, result := testDoGetMultirows(apiParamInput, dbParaList[0], dbQry)
	testDumpMultiRows(api.GetResultRows())

	testRunRangeGetRows(dbParaList[1], dbQry, api, result)
}

func testDoFilterRows(apiParamInput string, startIndx, count string, dbQry eobedb.DBQueryInf, preAPI ApiInf, preResult map[string]string) (ApiInf, *APIError) {
	const apiName = CAPIFilterMultiRowss

	api, err := GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")

	api.SetFilterSource(preAPI.GetResultRows())
	qryKVMap := make(map[string]string)
	qryKVMap["startIndx"] = startIndx
	qryKVMap["count"] = count
	_, apiErr := api.RunAPI(qryKVMap, preResult)

	return api, apiErr
}

func testDoFilterRowsBadParamInput(apiParamInput string, dbQry eobedb.DBQueryInf, preAPI ApiInf, preResult map[string]string) *APIError {
	const apiName = CAPIFilterMultiRowss

	api, err := GetAPIImplementation(apiName, apiParamInput)
	if err != nil {
		return NewAPIErrorf(CErrGenericError, err.Error())
	}

	api.SetFilterSource(preAPI.GetResultRows())
	qryKVMap := make(map[string]string)
	qryKVMap["startIndx"] = "0"
	qryKVMap["count"] = "10"
	_, apiErr := api.RunAPI(qryKVMap, preResult)

	return apiErr
}

//TestDBFilterRows Smoke test for DB API FilterRows
//	go test -v  -run DBFilterRows
func TestDBFilterRows(t *testing.T) {
	dbQry, dbSrv := testConnectDB()
	defer dbSrv.StopDBServer()

	dbParaList := testGetDBActionsForRange()
	apiParamInput := "GetByUName$$ ?username"
	api, result := testDoGetMultirows(apiParamInput, dbParaList[2], dbQry)
	//testDumpMultiRows(api.GetResultRows())

	apiParamInput = "?startIndx, ?count"
	apiFilter, apiErr := testDoFilterRows(apiParamInput, "0", "10", dbQry, api, result)
	checkAPIError(apiErr, "RunAPI() for filter rows error.")
	_, mrows := apiFilter.GetResultRows()
	fmt.Printf("Get total %d rows.\n", len(mrows))
	if len(mrows) != 10 { //Check DB and Do inerst enough values!!!!
		panic("result row count errors")
	}

	//Bad scenarios start < 0
	apiFilter, apiErr = testDoFilterRows(apiParamInput, "-1", "10", dbQry, api, result)
	checkAPIError(apiErr, "RunAPI() for filter rows error.")
	_, mrows = apiFilter.GetResultRows()
	fmt.Printf("Get total %d rows.\n", len(mrows))
	if len(mrows) != 10 { //Check DB and Do inerst enough values!!!!
		panic("result row count errors")
	}

	//Bad scenarios huge end, expect total counts
	apiFilter, apiErr = testDoFilterRows(apiParamInput, "-1", "10000", dbQry, api, result)
	checkAPIError(apiErr, "RunAPI() for filter rows error.")
	_, mrows = apiFilter.GetResultRows()
	fmt.Printf("Get total %d rows.\n", len(mrows))
	if len(mrows) != 36 { //Check DB and Do inerst enough values!!!!
		panic("result row count errors, expect total all rows for sadmin")
	}

	//Bad scenarios start > end - get none
	apiFilter, apiErr = testDoFilterRows(apiParamInput, "99", "10000", dbQry, api, result)
	checkAPIError(apiErr, "RunAPI() for filter rows error.")
	_, mrows = apiFilter.GetResultRows()
	fmt.Printf("Get total %d rows.\n", len(mrows))
	if len(mrows) != 0 { //Check DB and Do inerst enough values!!!!
		panic("result row count errors, expect total all rows for sadmin")
	}

	//Bad scenarios invalid input -- nondigit
	_, apiErr = testDoFilterRows(apiParamInput, "XXNOTDITG START", "10000", dbQry, api, result)
	checkAPIErrorType(apiErr, CErrBadCallError)

	//Bad scenarios invalid input -- nondigit
	_, apiErr = testDoFilterRows(apiParamInput, "0", "XXNOTDITG END", dbQry, api, result)
	checkAPIErrorType(apiErr, CErrBadCallError)

	//Bad scenarios invalid param input: not existed precall
	apiParamInput = "?startIndx, ^NotExisted"
	apiErr = testDoFilterRowsBadParamInput(apiParamInput, dbQry, api, result)
	checkAPIErrorType(apiErr, CErrBadCallError)

	//Bad scenarios invalid param input: not existed qryparam
	apiParamInput = "?NotExisted, ?count"
	apiErr = testDoFilterRowsBadParamInput(apiParamInput, dbQry, api, result)
	checkAPIErrorType(apiErr, CErrBadCallError)

	//Bad scenarios invalid param input: not enough params.
	apiParamInput = "?onlyone"
	apiErr = testDoFilterRowsBadParamInput(apiParamInput, dbQry, api, result)
	checkAPIErrorType(apiErr, CErrGenericError) //specified in testDoFilterRowsBadParamInput() for this test
}
