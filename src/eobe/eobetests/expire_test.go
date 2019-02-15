//This file contains test cases for user operation as a system admin.
package eobetests

import (
	"eobe/eobecltmgmt"
	"eobe/eobedb"
	"testing"
	"time"
)

//TestUserExpires 	Test for user login expires.
//	Reuse utilities from TestUserActions(), adjust system time, need to update time before running this test
//	Base line time: "2019-01-01" --> Mike login
//		--> time change to "2019-01-05" --> Alex login
//		--> all user do action as expected.
//		--> Time change to  "2019-01-10" : Mike Expire
//		--> Mike expect fail, and alex still success
//		--> login Mike again, Time change to  "2019-01-15"
//		--> alex expect fail, and Mike still success
//	go test -v -run UserExpires
func TestUserExpires(t *testing.T) {
	//Keep real system date
	realNow := time.Now()
	//Mike log in on : "2019-01-01"
	setOSDateAsStr("2018-01-01")

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

	usrTst := make([]*userTst, 2)
	usrTst[0] = createUserTest("Mike", cm, dbQry) //Mike expire days is 3
	usrTst[1] = createUserTest("Alex", cm, dbQry) //Alex expire days is 7

	runUserWithWait(t, usrTst[0], "LoginUser", nil)
	runUserWithWait(t, usrTst[1], "LoginUser", nil)
	usrTst[0].runUserTestMayFaile(t, "ValidateCrntUsr", true)
	usrTst[1].runUserTestMayFaile(t, "ValidateCrntUsr", true)

	setOSDateAsStr("2018-01-06")
	time.Sleep(2 * time.Second)
	usrTst[0].runUserTestMayFaile(t, "ValidateCrntUsr", false)
	usrTst[1].runUserTestMayFaile(t, "ValidateCrntUsr", true)

	usrTst[0].runUserTestMayFaile(t, "LoginUser", true)
	setOSDateAsStr("2018-01-08")
	time.Sleep(2 * time.Second)
	usrTst[0].runUserTestMayFaile(t, "ValidateCrntUsr", true)
	usrTst[1].runUserTestMayFaile(t, "ValidateCrntUsr", false)

	dbSrv.StopDBServer()
	cm.StopClientManagerServer()

	//recover real system date
	setOSDate(realNow)
}
