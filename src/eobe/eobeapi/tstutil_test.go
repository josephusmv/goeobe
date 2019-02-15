package eobeapi

import (
	"eobe/eobecltmgmt"
	"eobe/eobedb"
	"fmt"
	"io/ioutil"
)

//run full coverage test:
//		go test -cover -coverprofile cover.out
//		go tool cover -html=cover.out -o cover.html

const cTstStrDBType = "mysql"
const cTstDBConnStr = "root:123456@/THDATABASE"

func tstInitDBConn() (eobedb.DBQueryInf, eobedb.DBServerInf) {
	dbQry, dbSrv := eobedb.RunNewDBServer(cTstStrDBType, cTstDBConnStr)
	if dbQry == nil || dbSrv == nil {
		tstCheckError(fmt.Errorf("Nil DB Server"), "testConnectDB() create DB server failed.")
	}

	dbSrv.SetDBOptions(eobedb.OptionDBMaxConcurrencyInt, 80)
	dbSrv.SetDBOptions(eobedb.OptionDBLogRoot, "./LOGS")
	err := dbSrv.Init()
	tstCheckError(err, "testConnectDB() DB init failed.")

	return dbQry, dbSrv
}

func tstGetHarald() (*eobecltmgmt.ClientManager, *eobecltmgmt.Herald) {
	cm := eobecltmgmt.NewClientManager("./LOGS")
	cm.StartClientManagerServer()
	return cm, cm.NewHerald()
}

func tstGetUserConst() map[string]string {
	userConsts := make(map[string]string)
	userConsts["FORBID_ACCESS"] = "The specified action is not allowed."
	userConsts["ACCESSLEVE"] = "100"
	userConsts["LOCALStorePATH"] = "./eobeapiimpl/uploaded"
	return userConsts
}

func tstCheckError(err error, info string) {
	if err != nil {
		fmt.Println(info)
		panic(err)
	}
}

func tstDumpAPIDefine(adfn *apiDefine) {
	println("API Name:", adfn.apiName)
	println("PreAction:", adfn.preAction)
	println("postAction:", adfn.postAction)
	println("")
}

func tstGetQueryParams() map[string]string {
	qryMap := make(map[string]string)
	qryMap["username"] = "sadmin"
	qryMap["srchkey"] = "2014"
	qryMap["startIndx"] = "7"
	qryMap["count"] = "7"
	qryMap["FileName"] = "statics.go"
	return qryMap
}

func tstGetDBActions() map[string]*eobedb.QueryDefn {
	const cTstStrTableName = "DBTEST_TBL"
	qryDfnMap := make(map[string]*eobedb.QueryDefn)

	//SELECT
	qrySelect := eobedb.QueryDefn{
		QueryActionName:  "GetByUName",
		TableName:        cTstStrTableName,
		QueryType:        "SELECT",
		WhereReadyStr:    "username=?",
		ExpectedColNames: []string{"id", "username", "testfield1", "testfield3"}}
	qryDfnMap["GetByUName"] = &qrySelect

	//INSERT
	qryInsert := eobedb.QueryDefn{
		QueryActionName:  "InsertNewItem",
		TableName:        cTstStrTableName,
		QueryType:        "INSERT",
		ExpectedColNames: []string{"testfield1", "testfield2", "testfield3", "username"}}
	qryDfnMap["InsertNewItem"] = &qryInsert

	//UPDATE
	qryUpdate := eobedb.QueryDefn{
		QueryActionName:  "UpdateNewItem",
		TableName:        cTstStrTableName,
		QueryType:        "UPDATE",
		WhereReadyStr:    "testfield1 LIKE ?",
		ExpectedColNames: []string{"username", "testfield1", "testfield2", "testfield3"}}
	qryDfnMap["UpdateNewItem"] = &qryUpdate

	//Insert Before Delete
	qryInsertPreDelete := eobedb.QueryDefn{
		QueryActionName:  "InsertOneForDelete",
		TableName:        cTstStrTableName,
		QueryType:        "INSERT",
		ExpectedColNames: []string{"testfield1", "testfield2", "testfield3", "username"}}
	qryDfnMap["InsertOneForDelete"] = &qryInsertPreDelete

	//Delete
	qryDelete := eobedb.QueryDefn{
		QueryActionName: "DeleteByUName",
		TableName:       cTstStrTableName,
		QueryType:       "DELETE",
		WhereReadyStr:   "testfield1 LIKE ?"}
	qryDfnMap["DeleteByUName"] = &qryDelete

	//get range
	qryRangeDB := eobedb.QueryDefn{
		QueryActionName:  "GetRangeValues",
		TableName:        "DBRANGETEST_TBL",
		QueryType:        "SELECT",
		WhereReadyStr:    "srchkey=?",
		ExpectedColNames: []string{"testfield1"}}
	qryDfnMap["GetRangeValues"] = &qryRangeDB

	//get range
	qryGetByRange := eobedb.QueryDefn{
		QueryActionName:  "GetValuesByRange",
		TableName:        cTstStrTableName,
		QueryType:        "SELECT",
		WhereReadyStr:    "testfield1=?",
		ExpectedColNames: []string{"id", "username", "testfield1", "testfield3"}}
	qryDfnMap["GetValuesByRange"] = &qryGetByRange

	return qryDfnMap
}

type tst_loggerImpl struct{}

func (lgger tst_loggerImpl) TraceError(format string, a ...interface{}) error {
	fmt.Printf(format, a...)
	return fmt.Errorf(format, a...)
}

func (lgger tst_loggerImpl) TraceDev(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (lgger tst_loggerImpl) TraceInfo(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func testSetFile() map[string][]byte {
	const testFilePath1 = "statics.go"
	bytes1, err := ioutil.ReadFile(testFilePath1)
	if err != nil {
		panic("Open file " + testFilePath1 + "failed.")
	}

	const testFilePath2 = "factory.go"
	var bytes2 []byte
	bytes2, err = ioutil.ReadFile(testFilePath2)
	if err != nil {
		panic("Open file " + testFilePath2 + "failed.")
	}

	bytesMap := make(map[string][]byte)
	bytesMap[testFilePath1] = bytes1
	bytesMap[testFilePath2] = bytes2

	return bytesMap
}
