//This file contains test cases for user operation as a system admin.
package eobetests

import (
	"eobe/eobecltmgmt"
	"eobe/eobedb"
	"fmt"
	"sync"
	"testing"
	"time"
)

//TestUserActions 	Test manage users as a system admin smoke test, with no error scenario.
//	go test -v -run UserActions
func TestUserActions(t *testing.T) {
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

	totalUser := 6
	usrTst := make([]*userTst, 6)
	usrTst[0] = createUserTest("Mike", cm, dbQry)
	usrTst[1] = createUserTest("Alex", cm, dbQry)
	usrTst[2] = createUserTest("Kate", cm, dbQry)
	usrTst[3] = createUserTest("Jung", cm, dbQry)
	usrTst[4] = createUserTest("Brow", cm, dbQry)
	usrTst[5] = createUserTest("tommy", cm, dbQry)

	var wg sync.WaitGroup

	wg.Add(totalUser)
	//Test four user for LoginUser
	for i, ust := range usrTst {
		if i >= totalUser {
			break
		}
		go runUserWithWait(t, ust, "LoginUser", &wg)
	}
	wg.Wait()

	wg.Add(totalUser)
	//Test four user for viewAllAPIInfo
	for i, ust := range usrTst {
		if i >= totalUser {
			break
		}
		go viewAllAPIInfo(t, ust, &wg)
	}
	wg.Wait()

	wg.Add(totalUser)
	//Test four user for LogoutUser
	for i, ust := range usrTst {
		if i >= totalUser {
			break
		}
		go runUserWithWait(t, ust, "LogoutUser", &wg)
	}
	wg.Wait()

	dbSrv.StopDBServer()
	cm.StopClientManagerServer()
}

func viewAllAPIInfo(t *testing.T, usrTst *userTst, wg *sync.WaitGroup) {
	runUserWithWait(t, usrTst, "FetchAPIList", nil)
	runUserWithWait(t, usrTst, "FetchAllAPIParams", nil)
	runUserWithWait(t, usrTst, "FetchAllAPIParams", nil)
	if wg != nil {
		wg.Done()
	}
}

func createUserTest(userName string, cm *eobecltmgmt.ClientManager, dbQry eobedb.DBQueryInf) *userTst {
	usrOut := usrOutputs{
		userName: userName}
	usrIn := userInputs{
		userName: userName}
	usrTest := userTst{
		user:     userName,
		tstIn:    &usrIn,
		tstOut:   &usrOut,
		cm:       cm,
		dbQry:    dbQry,
		exptRlst: true}

	return &usrTest
}

func runUserWithWait(t *testing.T, usrTst *userTst, actionName string, wg *sync.WaitGroup) {
	usrTst.runUserTest(t, actionName)
	if wg != nil {
		wg.Done()
	}
}

type userTst struct {
	user     string
	tstIn    testInputInf
	tstOut   TstOutInf
	cm       *eobecltmgmt.ClientManager
	dbQry    eobedb.DBQueryInf
	exptRlst bool
}

func (srt *userTst) runUserTest(t *testing.T, actionName string) {
	fmt.Printf("\n****************************Step: %s for %s*********************************\n", actionName, srt.user)
	srt.tstIn.setActionName(actionName)
	if doTestAction(t, srt.tstIn, srt.tstOut, srt.cm, srt.dbQry) != srt.exptRlst {
		estr := fmt.Sprintf("\n\n\n********* Step: Step: %s for %s failed, stopped ******************\n\n\n", actionName, srt.user)
		t.Fail()
		panic(estr)
	}
}

func (srt *userTst) runUserTestMayFaile(t *testing.T, actionName string, exptRlst bool) {
	srt.exptRlst = exptRlst
	srt.runUserTest(t, actionName)
}
