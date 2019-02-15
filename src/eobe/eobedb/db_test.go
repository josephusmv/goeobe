package eobedb

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

//TestDBServerSmokeTest 	Smoke test for DB query and exec
//	go test -v -run DBServerSmokeTest
func TestDBServerSmokeTest(t *testing.T) {
	// Test 1 one single SELECT method.
	actionNames := []string{"GetByUName"}
	runNDBqueryTestForNGoroutines(actionNames, "This is a test query.", 1, t)

	// Test 2 one single INSERT method.
	actionNames = []string{"AddNewAPPTASKS"}
	runNDBqueryTestForNGoroutines(actionNames, "", 1, t)

	// Test 3 one single INSERT method.
	actionNames = []string{"ChangeUser1234"}
	runNDBqueryTestForNGoroutines(actionNames, "", 1, t)

	// Test 4 one single DELETE method.
	actionNames = []string{"RemoveUser1234"}
	runNDBqueryTestForNGoroutines(actionNames, "", 1, t)
}

// ***************************************
// Utilities for DB tests
func runNDBqueryTestForNGoroutines(actionNames []string, expected string, nroutines int, t *testing.T) {
	dbQry, dbSrv := runNewTestDBServer(t)
	qryMap := prepareQryMap()

	actLen := len(actionNames)
	if actLen <= 0 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(nroutines)
	for i := 0; i < nroutines; i++ {
		actInd := i % actLen
		qryData := qryMap[actionNames[actInd]]
		go fetchDBForAction(qryData, dbQry, expected, t, &wg)
	}
	wg.Wait()

	dbSrv.StopDBServer()
}

func runNewTestDBServer(t *testing.T) (dbQry DBQueryInf, dbSrv DBServerInf) {
	//init
	dbQry, dbSrv = RunNewDBServer("mysql", "root:123456@/THDATABASE", nil)
	dbSrv.SetDBOptions(OptionDBMaxConcurrencyInt, 80)
	err := dbSrv.Init()
	if err != nil {
		t.Fatal(err.Error())
	}
	return
}

func simpleVerifyTestResults(qryRslt QueryResult, expected string) (passed bool) {
	for _, row := range qryRslt.QueryRows {
		for _, val := range row {
			if val == expected {
				return true
			}
		}
	}

	return false
}

func fetchDBForAction(qryData QueryData, dbQry DBQueryInf, expected string, t *testing.T, wg *sync.WaitGroup) {
	qryRslt, err := dbQry.ExecDBAction(qryData, &DBLogger{})
	if err != nil {
		t.Fatal(err.Error())
		t.Fail()
	}

	if qryRslt.QueryErr != nil {
		t.Fatalf(qryRslt.QueryErr.Error())
		t.Fail()
	}

	if expected != "" && !simpleVerifyTestResults(qryRslt, expected) {
		t.Fail()
	}

	//t.Logf("Test Result: %v\n\n\n", qryRslt)
	wg.Done()
}

func prepareQryMap() (qryMap map[string]QueryData) {
	rand.Seed(time.Now().UTC().UnixNano())
	qryMap = make(map[string]QueryData)

	qry1 := QueryDefn{
		QueryActionName:  "GetByUName",
		TableName:        "APPTASKS_TBL",
		QueryType:        "SELECT",
		ParamterColNames: []string{"username", "title"},
		ParametrLogicOp:  []string{"AND"},
		ParamterOper:     []string{CstrDBOptEqualTo, CstrDBOptEqualTo},
		ExpectedColNames: []string{"id", "username", "title", "details", "crtdt"}}
	qryData1 := QueryData{
		QryActionDfn:    qry1,
		ParameterValues: []string{"testbenchmark", "This is a test query."},
		ExpectedValues:  nil}

	qryMap["GetByUName"] = qryData1

	qry2 := QueryDefn{
		QueryActionName:  "GetByUNameMultirows",
		TableName:        "APPTASKS_TBL",
		QueryType:        "SELECT",
		ParamterColNames: []string{"username", "crtdt"},
		ParametrLogicOp:  []string{"AND"},
		ParamterOper:     []string{CstrDBOptEqualTo, CstrDBOptGreaterThan},
		ExpectedColNames: []string{"username", "title", "crtdt"}}
	qryData2 := QueryData{
		QryActionDfn:    qry2,
		ParameterValues: []string{"testbenchmark", "2017-12-02 20:00:00"},
		ExpectedValues:  nil}

	qryMap["GetByUNameMultirows"] = qryData2

	qry3 := QueryDefn{
		QueryActionName:  "AddNewAPPTASKS",
		TableName:        "APPTASKS_TBL",
		QueryType:        "INSERT",
		ParamterColNames: []string{},
		ParametrLogicOp:  []string{},
		ParamterOper:     []string{},
		ExpectedColNames: []string{"title", "username", "crtdt", "details"}}
	qryData3 := QueryData{
		QryActionDfn:    qry3,
		ParameterValues: []string{},
		ExpectedValues:  []string{randomString(16), "user1234", "2017-12-03 20:00:00", randomString(32)}}

	qryMap["AddNewAPPTASKS"] = qryData3

	qry4 := QueryDefn{
		QueryActionName:  "RemoveUser1234",
		TableName:        "APPTASKS_TBL",
		QueryType:        "DELETE",
		ParamterColNames: []string{"username"},
		ParametrLogicOp:  []string{},
		ParamterOper:     []string{CstrDBOptEqualTo},
		ExpectedColNames: []string{}}
	qryData4 := QueryData{
		QryActionDfn:    qry4,
		ParameterValues: []string{"user1234"},
		ExpectedValues:  []string{}}

	qryMap["RemoveUser1234"] = qryData4

	qry5 := QueryDefn{
		QueryActionName:  "ChangeUser1234",
		TableName:        "APPTASKS_TBL",
		QueryType:        "UPDATE",
		ParamterColNames: []string{"username", "crtdt"},
		ParametrLogicOp:  []string{"AND"},
		ParamterOper:     []string{CstrDBOptEqualTo, CstrDBOptGreaterThan},
		ExpectedColNames: []string{"crtdt", "title"}}
	qryData5 := QueryData{
		QryActionDfn:    qry5,
		ParameterValues: []string{"user1234", "2016-12-03 20:00:00"},
		ExpectedValues:  []string{"2016-12-01 20:00:00", "Old already"}}

	qryMap["ChangeUser1234"] = qryData5

	qry11 := QueryDefn{
		QueryActionName:  "AddReadyWhere",
		TableName:        "APPTASKS_TBL",
		QueryType:        "INSERT",
		ParamterColNames: []string{},
		ParametrLogicOp:  []string{},
		ParamterOper:     []string{},
		ExpectedColNames: []string{"title", "username", "crtdt", "details"}}
	qryData11 := QueryData{
		QryActionDfn:    qry11,
		ParameterValues: []string{},
		ExpectedValues:  []string{randomString(16), "ReadyWhere", "2017-12-03 20:00:00", randomString(32)}}
	qryMap["AddReadyWhere"] = qryData11

	qry12 := QueryDefn{
		QueryActionName:  "QueryWithReadyWhere",
		TableName:        "APPTASKS_TBL",
		QueryType:        "SELECT",
		ParamterColNames: []string{},
		ParametrLogicOp:  []string{},
		ParamterOper:     []string{},
		WhereReadyStr:    "crtdt>? AND username=? OR username=?",
		ExpectedColNames: []string{"username", "crtdt"}}
	qryData12 := QueryData{
		QryActionDfn:    qry12,
		ParameterValues: []string{"2017-11-03 20:00:00", "User001", "ReadyWhere"},
		ExpectedValues:  []string{}}
	qryMap["QueryWithReadyWhere"] = qryData12

	qry13 := QueryDefn{
		QueryActionName:  "UpdateWithReadyWhere",
		TableName:        "APPTASKS_TBL",
		QueryType:        "UPDATE",
		ParamterColNames: []string{},
		ParametrLogicOp:  []string{},
		ParamterOper:     []string{},
		WhereReadyStr:    "crtdt>? AND username=? OR username=?",
		ExpectedColNames: []string{"title", "details", "crtdt"}}
	qryData13 := QueryData{
		QryActionDfn:    qry13,
		ParameterValues: []string{"2017-11-03 20:00:00", "ReadyWhere", "User001"},
		ExpectedValues:  []string{"New title", "；带给大家看过；", "2018-12-25 20:00:00"}}
	qryMap["UpdateWithReadyWhere"] = qryData13

	qry14 := QueryDefn{
		QueryActionName:  "QueryAfterUpdate",
		TableName:        "APPTASKS_TBL",
		QueryType:        "SELECT",
		ParamterColNames: []string{},
		ParametrLogicOp:  []string{},
		ParamterOper:     []string{},
		WhereReadyStr:    "crtdt=? AND username=?",
		ExpectedColNames: []string{"title", "username", "crtdt", "details"}}
	qryData14 := QueryData{
		QryActionDfn:    qry14,
		ParameterValues: []string{"2018-12-25 20:00:00", "ReadyWhere"},
		ExpectedValues:  []string{}}
	qryMap["QueryAfterUpdate"] = qryData14

	return
}

func randomString(len int) string {
	var result bytes.Buffer
	for i := 0; i < len; i++ {

		upperCharInt := 65 + rand.Intn(90-65)
		lowerCharInt := 97 + rand.Intn(122-97)
		punctList := []int{' ', ',', '.'}

		switch rand.Intn(30) {
		case 0:
			result.WriteString(string(punctList[0]))
		case 1:
			result.WriteString(string(punctList[1]))
		case 3:
			result.WriteString(string(punctList[2]))
		case 5, 6, 7, 8, 9, 10, 11:
			result.WriteString(string(upperCharInt))
		default:
			result.WriteString(string(lowerCharInt))
		}
	}
	return result.String()
}

//TestDBReadyWhereSmoke Smoke Test for customer provided where statement
//	go test -v -run DBReadyWhereSmoke
func TestDBReadyWhereSmoke(t *testing.T) {
	dbQry, dbSrv := runNewTestDBServer(t)
	qryMap := prepareQryMap()

	fmt.Println("\n\n **************Test 1 AddReadyWhere *****************")
	qryData := qryMap["AddReadyWhere"]
	runDBAction(qryData, dbQry, t)
	fmt.Println("*********************************************************")

	fmt.Println("\n\n **************Test 2 QueryWithReadyWhere *****************")
	qryData = qryMap["QueryWithReadyWhere"]
	runDBAction(qryData, dbQry, t)
	fmt.Println("*********************************************************")

	fmt.Println("\n\n **************Test 3 UpdateWithReadyWhere *****************")
	qryData = qryMap["UpdateWithReadyWhere"]
	runDBAction(qryData, dbQry, t)
	fmt.Println("*********************************************************")

	fmt.Println("\n\n **************Test 4 QueryAfterUpdate *****************")
	qryData = qryMap["QueryAfterUpdate"]
	runDBAction(qryData, dbQry, t)
	fmt.Println("*********************************************************")

	dbSrv.StopDBServer()
}

func runDBAction(qryData QueryData, dbQry DBQueryInf, t *testing.T) {
	qryRslt, err := dbQry.ExecDBAction(qryData, &DBLogger{})
	if err != nil {
		t.Fatal(err.Error())
		t.Fail()
	}

	fmt.Printf(" ---> Query %s Result: %d, %d\n", qryRslt.QueryActionName, qryRslt.AffectedRows, qryRslt.LastIndex)
	for i, v := range qryRslt.QueryRows {
		fmt.Printf(" ---> row %d: %v\n", i, v)
	}

	if qryRslt.QueryErr != nil {
		t.Fatalf(qryRslt.QueryErr.Error())
		t.Fail()
	}
}
