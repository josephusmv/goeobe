package eobereqhdler

import (
	"eobe/eobecltmgmt"
	"eobe/eobedb"
	"eobe/eobehttp"
	"eobe/eobeload"
	"fmt"
	"testing"
)

const cStrDBConnMAC = "root:123456@/THDATABASE"
const cStrDBConnWIN = "root:3edc$RFV@/THDATABASE"
const cStrDBConn = cStrDBConnMAC

const cTestIndexActionName = "GetIndex"
const cTestDefnRootPath = "./testres/"
const cStrTestDBActionFile = "tstdbact.yaml"
const cStrTestHTTPActionFile = "tsthttpact.yaml"
const cStrTestUserConstFile = "userconsts.yaml"
const cStrTestLogFileName = "tstLog.log"

func prepareReqst() eobehttp.RequestData {
	var req eobehttp.RequestData
	req.IP = "192.168.0.1"
	req.Port = "8080"
	req.Module = "RHTEST"
	req.QueryTarget = cTestIndexActionName
	req.QueryKeyValueMap = getQryKVMap()
	req.Logger = newhttpLoggerImpl(cStrTestLogFileName)

	return req
}

func getQryKVMap() map[string]string {
	qryKVMap := make(map[string]string)
	qryKVMap["username"] = "Mike"
	qryKVMap["userpwd"] = "mike_enc"
	qryKVMap["expdays"] = "7"
	return qryKVMap
}

//**********************************************************
//newRequestHandler call func: loadDefines,
func newRequestHandler() (*RequestHandler, eobedb.DBServerInf) {
	haMap, daMap, ucMap := loadDefines()
	dbQry, dbSrv := initDB()
	cm := &eobecltmgmt.ClientManager{}
	cm.StartClientManagerServer()

	rh := RequestHandler{cm: cm,
		haMap: haMap,
		daMap: daMap, ucMap: ucMap,
		dbQry: &dbQry}

	return &rh, dbSrv
}
func loadDefines() (map[string]*eobeload.HTTPActionDefn, map[string]*eobedb.QueryDefn, map[string]string) {
	df := cTestDefnRootPath + cStrTestDBActionFile
	hf := cTestDefnRootPath + cStrTestHTTPActionFile
	uf := cTestDefnRootPath + cStrTestUserConstFile
	lfct := eobeload.NewLoaderFactory()
	daMap, haMap, ucMap, err := lfct.LoadAll(df, hf, uf)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return haMap.GetData(), daMap.GetData(), ucMap.GetData()
}

func initDB() (eobedb.DBQueryInf, eobedb.DBServerInf) {
	dbQry, dbSrv := eobedb.RunNewDBServer("mysql", cStrDBConn)
	dbSrv.SetDBOptions(eobedb.OptionDBMaxConcurrencyInt, 80)
	err := dbSrv.Init()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return dbQry, dbSrv
}

//*********************************************************************
//***********************TEst cases
//TestSmokeTest SmokeTest
//	go test -v -run SmokeTest
//	go tool cover -html=cover.out -o cover.html
func TestSmokeTest(t *testing.T) {
	rh, dbSrv := newRequestHandler()
	reqst := prepareReqst()

	resp, err := rh.FetchResponse(reqst)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("\n************ Result response *****************")
	fmt.Println(resp)
	fmt.Println("**************** end test **********************\n")

	dbSrv.StopDBServer()
	rh.cm.StopClientManagerServer()
}
