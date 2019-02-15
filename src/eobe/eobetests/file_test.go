//This file contains test cases for user operation as a system admin.
package eobetests

import (
	"eobe/eobecltmgmt"
	"eobe/eobedb"
	"testing"
	"time"
)

//TestFileUpload 	Test for user login expires.
//	go test -v -run FileUpload
func TestFileUpload(t *testing.T) {
	//init DB
	dbQry, dbSrv := eobedb.RunNewDBServer("mysql", cStrDBConn)
	dbSrv.SetDBOptions(eobedb.OptionDBMaxConcurrencyInt, 80)
	err := dbSrv.Init()
	if err != nil {
		t.Errorf("\n\n\n********* DB init failed, stopped ******************\n\n\n")
		return
	}

	//init client manager
	cm := &eobecltmgmt.ClientManager{}
	cm.StartClientManagerServer()

	//start UDP server for simulate encrypt srverr
	go serverRunTCP()
	time.Sleep(time.Second)

	usrTst := createUserTest("Mike", cm, dbQry)

	usrTst.runUserTestMayFaile(t, "UploadFile", true)

	dbSrv.StopDBServer()
	cm.StopClientManagerServer()
}
