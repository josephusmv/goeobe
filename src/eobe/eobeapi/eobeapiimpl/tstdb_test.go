package eobeapiimpl

import (
	"eobe/eobedb"
	"fmt"
	"strconv"
	"testing"
)

const cTstStrTableName = "DBTEST_TBL"

type testDBQueryParam struct {
	dbActName  string
	ParamInput string
	qryDefn    eobedb.QueryDefn
	qryKVMap   map[string]string
}

func checkAPIError(err *APIError, contextStr string) {
	if err != nil && err.HasError() {
		fmt.Println(contextStr, "Error: ")
		fmt.Println(err.Error())
		panic(err)
	}
}

func testConnectDB() (eobedb.DBQueryInf, eobedb.DBServerInf) {
	const cTstStrDBType = "mysql"
	const cTstStrDBConnStr = "root:123456@/THDATABASE"
	dbQry, dbSrv := eobedb.RunNewDBServer(cTstStrDBType, cTstStrDBConnStr)
	if dbQry == nil || dbSrv == nil {
		checkError(fmt.Errorf("Nil DB Server"), "testConnectDB() create DB server failed.")
	}

	dbSrv.SetDBOptions(eobedb.OptionDBMaxConcurrencyInt, 80)
	dbSrv.SetDBOptions(eobedb.OptionDBLogRoot, "./LOGS")
	err := dbSrv.Init()
	checkError(err, "testConnectDB() DB init failed.")

	return dbQry, dbSrv
}

func testGetDBActions() []*testDBQueryParam {
	qryDfn := make([]*testDBQueryParam, 5)

	//SELECT
	qrySelect := eobedb.QueryDefn{
		QueryActionName:  "GetByUName",
		TableName:        cTstStrTableName,
		QueryType:        "SELECT",
		WhereReadyStr:    "username=?",
		ExpectedColNames: []string{"id", "username", "testfield1", "testfield3"}}
	qrySelectParam := make(map[string]string)
	qrySelectParam["username"] = "sadmin"
	qryDfn[0] = &testDBQueryParam{
		dbActName:  "GetByUName",
		ParamInput: "?username",
		qryDefn:    qrySelect,
		qryKVMap:   qrySelectParam}

	//INSERT
	qryInsert := eobedb.QueryDefn{
		QueryActionName:  "InsertNewItem",
		TableName:        cTstStrTableName,
		QueryType:        "INSERT",
		ExpectedColNames: []string{"testfield1", "testfield2", "testfield3", "username"}}
	qryInsertParam := make(map[string]string)
	qryInsertParam["username"] = "sadmin"
	qryInsertParam["testfield1"] = "testfield111111111111111"
	qryInsertParam["testfield2"] = "testfield2222222222222222"
	qryInsertParam["testfield3"] = "testfield33333333"
	qryDfn[1] = &testDBQueryParam{
		dbActName:  "InsertNewItem",
		ParamInput: "?testfield1, ?testfield2, ?testfield3, ?username",
		qryDefn:    qryInsert,
		qryKVMap:   qryInsertParam}

	//UPDATE
	qryUpdate := eobedb.QueryDefn{
		QueryActionName:  "UpdateNewItem",
		TableName:        cTstStrTableName,
		QueryType:        "UPDATE",
		WhereReadyStr:    "testfield1 LIKE ?",
		ExpectedColNames: []string{"username", "testfield1", "testfield2", "testfield3"}}
	qryUpdateParam := make(map[string]string)
	qryUpdateParam["testfield1old"] = "test%"
	qryUpdateParam["username"] = "sadmin"
	qryUpdateParam["testfield1"] = "test New Change"
	qryUpdateParam["testfield2"] = "updated testfield2222222"
	qryUpdateParam["testfield3"] = "updated testfield33333333"
	qryDfn[2] = &testDBQueryParam{
		dbActName:  "UpdateNewItem",
		ParamInput: "?username, ?testfield1, ?testfield2, ?testfield3; ?testfield1old",
		qryDefn:    qryUpdate,
		qryKVMap:   qryUpdateParam}

	//Insert Before Delete
	qryInsertPreDelete := eobedb.QueryDefn{
		QueryActionName:  "InsertOneForDelete",
		TableName:        cTstStrTableName,
		QueryType:        "INSERT",
		ExpectedColNames: []string{"testfield1", "testfield2", "testfield3", "username"}}
	qryInsertPreDeleteParam := make(map[string]string)
	qryInsertPreDeleteParam["username"] = "sadmin"
	qryInsertPreDeleteParam["testfield1"] = "testchange"
	qryInsertPreDeleteParam["testfield2"] = "testfield2222222222222222"
	qryInsertPreDeleteParam["testfield3"] = "testfield33333333"
	qryDfn[3] = &testDBQueryParam{
		dbActName:  "InsertOneForDelete",
		ParamInput: "?testfield1, ?testfield2, ?testfield3, ?username",
		qryDefn:    qryInsertPreDelete,
		qryKVMap:   qryInsertPreDeleteParam}

	//Delete
	qryDelete := eobedb.QueryDefn{
		QueryActionName: "DeleteByUName",
		TableName:       cTstStrTableName,
		QueryType:       "DELETE",
		WhereReadyStr:   "testfield1 LIKE ?"}
	qryDeleteParam := make(map[string]string)
	qryDeleteParam["testfield1"] = "test%"
	qryDfn[4] = &testDBQueryParam{
		dbActName:  "DeleteByUName",
		ParamInput: "?testfield1",
		qryDefn:    qryDelete,
		qryKVMap:   qryDeleteParam}

	return qryDfn
}

func testRunAPIFor(apiName, apiParamInput string, param *testDBQueryParam, dbQry eobedb.DBQueryInf) (ApiInf, map[string]string) {
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

func testDumpResultMap(result map[string]string, stepName string) {
	fmt.Println("\n -------- Result for: " + stepName)
	for k, v := range result {
		fmt.Printf("\t---->%s: %s\n", k, v)
	}
	fmt.Println("\n-------- Done ------\n")
}

func testDumpMultiRows(cnames []string, rows [][]string) {
	fmt.Printf("\n----> Row names:\n\t")
	for _, v := range cnames {
		fmt.Printf("%s,  ", v)
	}

	fmt.Println("\n----> Rows:")
	for _, v := range rows {
		fmt.Printf("\t %v\n", v)
	}
	fmt.Println("\n-------- Done ------\n")
}

func testSmokeDBActionRun(paramIndx int) map[string]string {
	dbQry, dbSrv := testConnectDB()
	defer dbSrv.StopDBServer()

	dbParaList := testGetDBActions()

	_, result := testRunAPIFor(dbParaList[paramIndx].dbActName, dbParaList[paramIndx].ParamInput, dbParaList[paramIndx], dbQry)
	return result
}

//TestDBActionAPISelect test connect to remote Server
//	go test -v  -run DBActionAPISelect
//	go tool cover -html=cover.out -o cover.html
func TestDBActionAPISelect(t *testing.T) {
	testSmokeDBActionRun(0)
}

//TestDBActionAPIInsert test connect to remote Server
//	go test -v  -run DBActionAPIInsert
//	go tool cover -html=cover.out -o cover.html
func TestDBActionAPIInsert(t *testing.T) {
	testSmokeDBActionRun(1)
}

//TestDBActionAPIUpdate test connect to remote Server
//	go test -v  -run DBActionAPIUpdate
//	go tool cover -html=cover.out -o cover.html
func TestDBActionAPIUpdate(t *testing.T) {
	testSmokeDBActionRun(2)
}

//TestDBActionAPIDelete test connect to remote Server
//	go test -v  -run DBActionAPIDelete
func TestDBActionAPIDelete(t *testing.T) {
	testSmokeDBActionRun(3) //add a new row for delete
	result := testSmokeDBActionRun(4)
	affectedRows, err := strconv.Atoi(result["retAffectedRows"])
	if err != nil || affectedRows <= 0 {
		panic("TestDBActionAPIDelete failed for not affected the newly addd rows")
	}
}

//TestDBMultiRows test connect to remote Server
//	go test -v  -run DBMultiRows
func TestDBMultiRows(t *testing.T) {
	const apiName = CAPIGetMultiRows
	const apiParamInput = "GetByUName$$ ?username"

	dbQry, dbSrv := testConnectDB()
	defer dbSrv.StopDBServer()

	dbParaList := testGetDBActions()

	api, _ := testRunAPIFor(apiName, apiParamInput, dbParaList[0], dbQry)
	testDumpMultiRows(api.GetResultRows())
}

/* **********************************
 * Error runs
 * **********************************/
func testGetBadDBActions() []*testDBQueryParam {
	qryDfn := make([]*testDBQueryParam, 9)

	//Invalid DB Query Type
	qrySelect := eobedb.QueryDefn{
		QueryActionName:  "GetByUName",
		TableName:        cTstStrTableName,
		QueryType:        "BADOP",
		WhereReadyStr:    "username=?",
		ExpectedColNames: []string{"id", "username", "testfield1", "testfield3"}}
	qrySelectParam := make(map[string]string)
	qrySelectParam["username"] = "sadmin"
	qryDfn[0] = &testDBQueryParam{
		dbActName:  "GetByUName",
		ParamInput: "?username",
		qryDefn:    qrySelect,
		qryKVMap:   qrySelectParam}

	//Empty DB ParamInput
	qryEmpty := eobedb.QueryDefn{
		QueryActionName:  "GetByUName",
		TableName:        cTstStrTableName,
		QueryType:        "INSERT",
		WhereReadyStr:    "username=?",
		ExpectedColNames: []string{"id", "username", "testfield1", "testfield3"}}
	qryEmptyParam := make(map[string]string)
	qryEmptyParam["username"] = "sadmin"
	qryDfn[1] = &testDBQueryParam{
		dbActName:  "GetByUName",
		ParamInput: "",
		qryDefn:    qryEmpty,
		qryKVMap:   qryEmptyParam}

	//UPDATE wrong sequence
	qryUpdate := eobedb.QueryDefn{
		QueryActionName:  "UpdateNewItem",
		TableName:        cTstStrTableName,
		QueryType:        "UPDATE",
		WhereReadyStr:    "testfield1 LIKE ?",
		ExpectedColNames: []string{"username", "testfield1", "testfield2", "testfield3"}}
	qryUpdateParam := make(map[string]string)
	qryUpdateParam["testfield1old"] = "test%"
	qryUpdateParam["username"] = "sadmin"
	qryUpdateParam["testfield1"] = "test New Change"
	qryUpdateParam["testfield2"] = "updated testfield2222222"
	qryUpdateParam["testfield3"] = "updated testfield33333333"
	qryDfn[2] = &testDBQueryParam{
		dbActName:  "UpdateNewItem",
		ParamInput: "?testfield1old; ?username, ?testfield1, ?testfield2, ?testfield3",
		qryDefn:    qryUpdate,
		qryKVMap:   qryUpdateParam}

	//INSERT missing new value
	qryInsert := eobedb.QueryDefn{
		QueryActionName:  "InsertNewItem",
		TableName:        cTstStrTableName,
		QueryType:        "INSERT",
		ExpectedColNames: []string{"testfield1", "testfield2", "testfield3", "username"}}
	qryInsertParam := make(map[string]string)
	qryInsertParam["username"] = "sadmin"
	qryInsertParam["testfield1"] = "testfield111111111111111"
	qryInsertParam["testfield3"] = "testfield33333333"
	qryDfn[3] = &testDBQueryParam{
		dbActName:  "InsertNewItem",
		ParamInput: "?testfield1, ?testfield2, ?testfield3, ?username; ,",
		qryDefn:    qryInsert,
		qryKVMap:   qryInsertParam}

	//Parameter nil
	qrySelectPValMissing := eobedb.QueryDefn{
		QueryActionName:  "GetByUName",
		TableName:        cTstStrTableName,
		QueryType:        "SELECT",
		WhereReadyStr:    "username=?",
		ExpectedColNames: []string{"id", "username", "testfield1", "testfield3"}}
	qryDfn[4] = &testDBQueryParam{
		dbActName:  "GetByUName",
		ParamInput: "?username",
		qryDefn:    qrySelectPValMissing,
		qryKVMap:   nil}

	//Parameter nil
	qryPValEmpty := eobedb.QueryDefn{
		QueryActionName:  "GetByUName",
		TableName:        cTstStrTableName,
		QueryType:        "SELECT",
		WhereReadyStr:    "username=?",
		ExpectedColNames: []string{"id", "username", "testfield1", "testfield3"}}
	qryDfn[5] = &testDBQueryParam{
		dbActName:  "GetByUName",
		ParamInput: " ;  ,",
		qryDefn:    qryPValEmpty,
		qryKVMap:   nil}

	//Parameter nil
	qryTbleEmpty := eobedb.QueryDefn{
		QueryActionName:  "GetByUName",
		TableName:        "",
		QueryType:        "SELECT",
		WhereReadyStr:    "username=?",
		ExpectedColNames: []string{"id", "username", "testfield1", "testfield3"}}
	qryDfn[6] = &testDBQueryParam{
		dbActName:  "GetByUName",
		ParamInput: " ;  ,",
		qryDefn:    qryTbleEmpty,
		qryKVMap:   nil}

	//SELECT with zero result
	qrySelectNone := eobedb.QueryDefn{
		QueryActionName:  "GetByUName",
		TableName:        cTstStrTableName,
		QueryType:        "SELECT",
		WhereReadyStr:    "username=?",
		ExpectedColNames: []string{"id", "username", "testfield1", "testfield3"}}
	qrySelectNoneParam := make(map[string]string)
	qrySelectNoneParam["username"] = "A name doesn't exist"
	qryDfn[7] = &testDBQueryParam{
		dbActName:  "GetByUName",
		ParamInput: "?username",
		qryDefn:    qrySelectNone,
		qryKVMap:   qrySelectNoneParam}

	return qryDfn
}

func testBadRunAPIFor(apiName, apiParamInput string, param *testDBQueryParam, dbQry eobedb.DBQueryInf) *APIError {
	api, err := GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")

	if _, dbAction := api.IsDBAction(); dbAction {
		api.SetDBInfo(dbQry, param.qryDefn)
	}

	_, apiErr := api.RunAPI(param.qryKVMap, nil)

	return apiErr
}

func testBadDBActionRun(paramIndx int, expected APIErrType) {
	dbQry, dbSrv := testConnectDB()
	defer dbSrv.StopDBServer()

	dbParaList := testGetBadDBActions()

	apiErr := testBadRunAPIFor(dbParaList[paramIndx].dbActName, dbParaList[paramIndx].ParamInput, dbParaList[paramIndx], dbQry)
	checkAPIErrorType(apiErr, expected)
}

func checkAPIErrorType(err *APIError, expected APIErrType) {
	if !err.HasError() || err.ErrType != expected {
		fmt.Println(err)
		panic("checkAPIErrorType failed")
	}

	fmt.Println(err)
}

//TestDBActionAPIInvalidQuery
//	go test -v  -run DBActionAPIInvalidQuery
func TestDBActionAPIInvalidQuery(t *testing.T) {
	testBadDBActionRun(0, CErrBadCallError)
}

//TestDBActionAPIEmptyParameter
//	go test -v  -run DBActionAPIEmptyParameter
func TestDBActionAPIEmptyParameter(t *testing.T) {
	testBadDBActionRun(1, CErrBadCallError)
}

//TestDBActionAPIWrongUpdate
//	go test -v  -run DBActionAPIWrongUpdate
func TestDBActionAPIWrongUpdate(t *testing.T) {
	testBadDBActionRun(2, CErrServerInternalError)
}

//TestDBActionAPIExpectValMissing
//	go test -v  -run DBActionAPIExpectValMissing
func TestDBActionAPIExpectValMissing(t *testing.T) {
	testBadDBActionRun(3, CErrBadCallError)
}

//TestDBActionAPINilParameter
//	go test -v  -run DBActionAPINilParameter
func TestDBActionAPINilParameter(t *testing.T) {
	testBadDBActionRun(4, CErrBadCallError)
}

//TestDBActionAPIPValEmpty
//	go test -v  -run DBActionAPIPValEmpty
func TestDBActionAPIPValEmpty(t *testing.T) {
	//Has no way to do too many verification for the DB parameters.
	testBadDBActionRun(5, CErrServerInternalError)
}

//TestDBActionAPIPTblEmpty
//	go test -v  -run DBActionAPIPTblEmpty
func TestDBActionAPIPTblEmpty(t *testing.T) {
	//Force to CErrServerInternalError as test result, since it's not a expected user error outside this test
	testBadDBActionRun(6, CErrServerInternalError)
}

//TestDBMultiRowsNoResult test connect to remote Server
//	go test -v  -run DBMultiRowsNoResult
func TestDBMultiRowsNoResult(t *testing.T) {
	const apiName = CAPIGetMultiRows
	apiParamInput := "GetByUName$$ ?username"

	dbQry, dbSrv := testConnectDB()
	defer dbSrv.StopDBServer()

	dbParaList := testGetBadDBActions()

	//Success but No result
	apiErr := testBadRunAPIFor(apiName, apiParamInput, dbParaList[7], dbQry)
	if apiErr.HasError() || apiErr.ErrType != CErrSuccess {
		panic(apiErr)
	}

	//Error update: for all this kinds of error need has ServerInternalError because they are unpredicatable.
	//	Because if this kinds of errors are checked, a SQL query is requred in this package.
	//	And that's totally not what I want here!
	apiParamInput = "UpdateNewItem$$ ?testfield1old; ?username, ?testfield1, ?testfield2, ?testfield3"
	apiErr = testBadRunAPIFor(apiName, apiParamInput, dbParaList[2], dbQry)
	checkAPIErrorType(apiErr, CErrServerInternalError)
}
