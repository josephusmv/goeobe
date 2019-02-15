//This file contains test cases for user operation as a system admin.
package eobetests

import (
	"eobe/eobecltmgmt"
	"eobe/eobedb"
	"fmt"
	"testing"
)

//TestAsAdminUserMgmt 	Test manage users as a system admin smoke test, with no error scenario.
//	go test -v -run AsAdminUserMgmt
func TestAsAdminUserMgmt(t *testing.T) {
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
	//time.Sleep(time.Second)

	var tstOut AdminTestsOutput
	var tstIn admInputs

	// test 1.1: LoginAdmin

	fmt.Println("\n****************************Step: LoginAdmin*********************************")
	tstIn.actionName = "LoginAdmin"
	if !doTestAction(t, &tstIn, &tstOut, cm, dbQry) {
		t.Errorf("\n\n\n********* Step: LoginAdmin failed, stopped ******************\n\n\n")
		return
	}
	//time.Sleep(time.Second)

	// test 1.2: ViewAllUsers
	fmt.Println("\n****************************Step: ViewAllUsers*********************************")
	tstIn.actionName = "ViewAllUsers"
	if !doTestAction(t, &tstIn, &tstOut, cm, dbQry) {
		t.Errorf("\n\n\n********* Step: ViewAllUsers failed, stopped ******************\n\n\n")
		return
	}
	//time.Sleep(time.Second)

	// test 1.3:  Check one user detail
	fmt.Println("\n****************************Step: ViewUserDetails*********************************")
	tstIn.actionName = "ViewUserDetails"
	if !doTestAction(t, &tstIn, &tstOut, cm, dbQry) {
		t.Errorf("\n\n\n********* Step: ViewUserDetails failed, stopped ******************\n\n\n")
		return
	}
	//time.Sleep(time.Second)

	// test 1.4:   Add a new user
	fmt.Println("\n****************************Step: AddNewUserByAdmin*********************************")
	tstIn.actionName = "AddNewUserByAdmin"
	if !doTestAction(t, &tstIn, &tstOut, cm, dbQry) {
		t.Errorf("\n\n\n********* Step: AddNewUserByAdmin failed, stopped ******************\n\n\n")
		return
	}
	//time.Sleep(time.Second)

	// test 1.5:   Modify a user
	fmt.Println("\n****************************Step: ModifyUserByAdmin*********************************")
	tstIn.actionName = "ModifyUserByAdmin"
	if !doTestAction(t, &tstIn, &tstOut, cm, dbQry) {
		t.Errorf("\n\n\n********* Step: ModifyUserByAdmin failed, stopped ******************\n\n\n")
		return
	}
	//time.Sleep(time.Second)

	// test 1.6:   Delete a user
	fmt.Println("\n****************************Step: DeleteUserByAdmin*********************************")
	tstIn.actionName = "DeleteUserByAdmin"
	if !doTestAction(t, &tstIn, &tstOut, cm, dbQry) {
		t.Errorf("\n\n\n********* Step: DeleteUserByAdmin failed, stopped ******************\n\n\n")
		return
	}
	//time.Sleep(time.Second)

	// test 1.7:   Logout admin
	fmt.Println("\n****************************Step: LogoutUser*********************************")
	tstIn.actionName = "LogoutUser"
	if !doTestAction(t, &tstIn, &tstOut, cm, dbQry) {
		t.Errorf("\n\n\n********* Step: LogoutUser failed, stopped ******************\n\n\n")
		return
	}
	//time.Sleep(time.Second)

	dbSrv.StopDBServer()
	cm.StopClientManagerServer()
}
