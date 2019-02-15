package eobesession

import (
	b64 "encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

func pannicErrors(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func printCookies(ckyLst []*http.Cookie) {
	for _, v := range ckyLst {
		var name string
		if v != nil {
			name = v.Name
		} else {
			name = "empty cookie"
		}
		fmt.Printf("--> Cookie %s: %v\n", name, v)
	}
}

//TestSingleThreadClientSession Smoke test for Single Thread
//	go test -v -cover -coverprofile cover.out -run SingleThreadClientSession
//	go tool cover -html=cover.out -o cover.html
func TestSingleThreadClientSession(t *testing.T) {
	var wg sync.WaitGroup
	var tm testMgr
	tm.Init(&wg)

	wg.Add(1)
	go listen(&tm)
	wg.Wait()
}

func listen(tm *testMgr) {
	http.Handle("/", tm)
	http.ListenAndServe(":8080", nil)
}

var tstUserName = "TSTADMIN"
var tstCookieNames = []string{"cName001", "cName002", "cName003", "cName004"}
var tstCookieValues = []string{"cValue001", "cValue002", "cValue003", "cValue004"}

type testMgr struct {
	sessMgr *SessionManager
	step    int
	ip      string
	logger  LoggerInf
	lastSID string
	wg      *sync.WaitGroup
	curTime time.Time
}

func (tm *testMgr) Init(wg *sync.WaitGroup) {
	tm.sessMgr = NewSessionManager()
	tm.step = 0
	tm.ip = "127.0.0.1"
	tm.logger = NewDummyLogger()
	tm.wg = wg
}

func (tm *testMgr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ckyLst, resultStr := tm.runTest(r)
		for _, v := range ckyLst {
			http.SetCookie(w, v)
		}
		fmt.Fprintf(w, resultStr+", Step: %d.", tm.step)
		return
	}

	if strings.HasSuffix(r.URL.Path, "js") {
		data, err := ioutil.ReadFile("./testres/jquery-3.2.1.min.js")
		if err != nil {
			return
		}
		w.Header().Add("Content-Type", "application/javascript")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}

	t, err := template.ParseFiles("./testres/addcookie.html")
	pannicErrors(err)
	t.Execute(w, nil)
}

func (tm *testMgr) runTest(r *http.Request) ([]*http.Cookie, string) {
	tm.step++
	switch tm.step {
	case 1:
		return tm.testStep1(r.Cookies())
	case 2:
		return tm.testStep2(r.Cookies())
	case 3:
		return tm.testStep3(r.Cookies())
	case 4:
		return tm.testStep4(r.Cookies())
	case 5:
		return tm.testStep5(r.Cookies())
	case 6:
		return tm.testStep6(r.Cookies())
	case 7:
		return tm.testStep7(r.Cookies())
	case 8:
		return tm.testStep8(r.Cookies())
	case 9:
		return tm.testStep9(r.Cookies())
	case 10:
		return tm.testStep10(r.Cookies())
	case 11:
		return tm.testStep11(r.Cookies())
	default:
		tm.wg.Done()
	}

	return nil, ""
}

func (tm *testMgr) testStep1(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	sm := tm.sessMgr
	nowTime := time.Now()

	fmt.Println("********** Test step 1: Add two cookie **********")
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, nowTime, tm.logger)
	pannicErrors(err)

	sm.AddCookie(tstCookieNames[0], tstCookieValues[0], 3, cltSess, nowTime, tm.logger)
	sm.AddCookie(tstCookieNames[1], tstCookieValues[1], 3, cltSess, nowTime, tm.logger)

	sm.ClientAccessFinished(cltSess, nowTime, tm.logger)
	tm.lastSID, _ = cltSess.GetClientInfo()

	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Add SID: %s, Add two Cookie: %s, %s", tm.lastSID, tstCookieNames[0], tstCookieNames[1])
}

func (tm *testMgr) verifyStep1Result(ckyLst []*http.Cookie) {
	var bCookie0, bCookie1, bSID bool
	for _, v := range ckyLst {
		if v.Name == tstCookieNames[0] && v.Value == tstCookieValues[0] {
			bCookie0 = true
			continue
		}
		if v.Name == tstCookieNames[1] && v.Value == tstCookieValues[1] {
			bCookie1 = true
			continue
		}
		if v.Name == cStrCookieNameSessionID && v.Value == tm.lastSID {
			bSID = true
			continue
		}
	}

	if len(ckyLst) != 1+3 || !bCookie0 || !bCookie1 || !bSID {
		printCookies(ckyLst)
		panic("verifyStep1Result failed")
	}
}

func (tm *testMgr) testStep2(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	tm.verifyStep1Result(ckyLst)

	sm := tm.sessMgr
	nowTime := time.Now()

	fmt.Println("\n\n********** Test step 2: Login **********")
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, nowTime, tm.logger)
	pannicErrors(err)

	sid, _ := cltSess.GetClientInfo()
	if sid != tm.lastSID {
		fmt.Printf("TEST Failed: sid = %s, last sid: %s", sid, tm.lastSID)
		panic("SID unepxected Changed.")
	}

	sm.Login(cltSess, tstUserName, 7, nowTime, tm.logger)
	sm.DelCookie(tstCookieNames[1], cltSess, nowTime, tm.logger)

	sm.ClientAccessFinished(cltSess, nowTime, tm.logger)
	tm.lastSID, _ = cltSess.GetClientInfo()

	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Refresh SID: %s,keep old Cookie: %s, login user %s", tm.lastSID, tstCookieNames[0], tstUserName)
}

func (tm *testMgr) verifyStep2Result(ckyLst []*http.Cookie) {
	var bCookie0, bUser, bSID bool
	for _, v := range ckyLst {
		if v.Name == tstCookieNames[0] && v.Value == tstCookieValues[0] {
			bCookie0 = true
			continue
		}
		if v.Name == cStrCookieNameCurrentUser && v.Value == tstUserName {
			bUser = true
			continue
		}
		if v.Name == cStrCookieNameSessionID && v.Value == tm.lastSID {
			bSID = true
			continue
		}
	}

	if len(ckyLst) != 1+3 || !bCookie0 || !bUser || !bSID {
		printCookies(ckyLst)
		panic("verifyStep2Result failed")
	}
}

func (tm *testMgr) testStep3(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	tm.verifyStep2Result(ckyLst)

	sm := tm.sessMgr
	nowTime := time.Now()

	fmt.Println("\n\n********** Test step 3: Logout **********")
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, nowTime, tm.logger)
	pannicErrors(err)

	sid, _ := cltSess.GetClientInfo()
	if sid != tm.lastSID {
		panic("SID unepxected Changed.")
	}

	sm.SetBindData(&DummyBindDataImpl{}, cltSess, tm.logger)

	value, cookie := sm.ReadCookie(tstCookieNames[0], cltSess, tm.logger)
	fmt.Printf("\tReadCookie: %s value: %s, cookie: %v", tstCookieNames[0], value, cookie)
	if value != tstCookieValues[0] {
		panic("Error read cookie!")
	}

	sm.Logout(cltSess, nowTime, tm.logger)

	sm.ClientAccessFinished(cltSess, nowTime, tm.logger)
	tm.lastSID, _ = cltSess.GetClientInfo()

	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Refresh SID: %s,keep old Cookie: %s, logout user.", tm.lastSID, tstCookieNames[0])
}

func (tm *testMgr) verifyStep3Result(ckyLst []*http.Cookie) {
	var bCookie0, bSID bool
	for _, v := range ckyLst {
		if v.Name == tstCookieNames[0] && v.Value == tstCookieValues[0] {
			bCookie0 = true
			continue
		}
		if v.Name == cStrCookieNameSessionID && v.Value == tm.lastSID {
			bSID = true
			continue
		}
	}

	if len(ckyLst) != 1+2 || !bCookie0 || !bSID {
		printCookies(ckyLst)
		panic("verifyStep3Result failed")
	}
}

func (tm *testMgr) testStep4(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	tm.verifyStep3Result(ckyLst)

	sm := tm.sessMgr
	nowTime := time.Now()
	tm.curTime = time.Now().AddDate(0, 0, 5)
	sm.RunExpire(tm.curTime, tm.logger) //expire the cookie added from step 1

	fmt.Println("\n\n********** Test step 4: Login then Expire Cookie **********")
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, nowTime, tm.logger)
	pannicErrors(err)

	sm.Login(cltSess, tstUserName, 7, nowTime, tm.logger)

	sm.ClientAccessFinished(cltSess, nowTime, tm.logger)
	tm.lastSID, _ = cltSess.GetClientInfo()

	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Refresh SID: %s, LOGOUT old Cookie: %s, login user %s.", tm.lastSID, tstCookieNames[0], tstUserName)
}

func (tm *testMgr) verifyStep4Result(ckyLst []*http.Cookie) {
	var bUser, bSID bool
	for _, v := range ckyLst {
		if v.Name == cStrCookieNameCurrentUser && v.Value == tstUserName {
			bUser = true
			continue
		}
		if v.Name == cStrCookieNameSessionID && v.Value == tm.lastSID {
			bSID = true
			continue
		}
	}

	if len(ckyLst) != 1+2 || !bUser || !bSID {
		printCookies(ckyLst)
		panic("verifyStep4Result failed")
	}
}

func (tm *testMgr) testStep5(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	tm.verifyStep4Result(ckyLst)

	sm := tm.sessMgr
	tm.curTime = tm.curTime.AddDate(0, 0, 8)
	sm.RunExpire(tm.curTime, tm.logger) //expire the cookie added from step 1

	fmt.Println("\n\n********** Test step 5: Expire Everything **********")
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, tm.curTime, nil)
	pannicErrors(err)

	//Just for coverage trick
	bd := sm.GetBindData(cltSess, tm.logger)
	bd2 := cltSess.GetBindData()
	fmt.Println(bd, bd2)
	errb := cltSess.SetBindData(nil)
	fmt.Println(errb)

	sm.ClientAccessFinished(cltSess, tm.curTime, tm.logger)
	tm.lastSID, _ = cltSess.GetClientInfo()

	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Refresh SID: %s, LOGOUT all other cookies.", tm.lastSID)
}

func (tm *testMgr) verifyStep5Result(ckyLst []*http.Cookie) {
	var bSID bool
	for _, v := range ckyLst {
		if v.Name == cStrCookieNameSessionID && v.Value == tm.lastSID {
			bSID = true
			continue
		}
	}

	if len(ckyLst) != 1+1 || !bSID {
		printCookies(ckyLst)
		panic("verifyStep5Result failed")
	}
}

func (tm *testMgr) testStep6(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	tm.verifyStep5Result(ckyLst)

	sm := tm.sessMgr
	tm.curTime = tm.curTime.AddDate(0, 0, 1)
	sm.RunExpire(tm.curTime, tm.logger) //expire the cookie added from step 1

	fmt.Println("\n\n********** Test step 6: Try to add SID and user cookie **********")
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, tm.curTime, tm.logger)
	pannicErrors(err)

	sm.AddCookie(cStrCookieNameSessionID, tstCookieValues[0], 3, cltSess, tm.curTime, tm.logger)
	sm.AddCookie(cStrCookieNameCurrentUser, tstUserName, 3, cltSess, tm.curTime, tm.logger)

	var user string
	sm.ClientAccessFinished(cltSess, tm.curTime, tm.logger)
	tm.lastSID, user = cltSess.GetClientInfo()

	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Refresh SID: %s, Login a user through Add cookie %s.", tm.lastSID, user)
}

func (tm *testMgr) verifyStep6Result(ckyLst []*http.Cookie) {
	var bSID, bUser bool
	for _, v := range ckyLst {
		if v.Name == cStrCookieNameSessionID && v.Value == tm.lastSID {
			bSID = true
			continue
		}
		if v.Name == cStrCookieNameCurrentUser && v.Value == tstUserName {
			bUser = true
			continue
		}
	}

	if len(ckyLst) != 1+2 || !bSID || !bUser {
		printCookies(ckyLst)
		panic("verifyStep6Result failed")
	}
}

func (tm *testMgr) testStep7(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	tm.verifyStep6Result(ckyLst)

	sm := tm.sessMgr
	tm.curTime = tm.curTime.AddDate(0, 0, 1)
	sm.RunExpire(tm.curTime, tm.logger) //expire the cookie added from step 1

	fmt.Println("\n\n********** Test step 7: Test Logout using DeleteCookie and read cookies **********")
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, tm.curTime, tm.logger)
	pannicErrors(err)

	sid, sCoky := sm.ReadCookie(cStrCookieNameSessionID, cltSess, tm.logger)
	fmt.Printf("SID info read from Cookie: %s, %v", sid, sCoky)
	uName, uCoky := sm.ReadCookie(cStrCookieNameCurrentUser, cltSess, tm.logger)
	fmt.Printf("User info read from Cookie: %s, %v", uName, uCoky)

	sm.DelCookie(cStrCookieNameSessionID, cltSess, tm.curTime, tm.logger)
	sm.DelCookie(cStrCookieNameCurrentUser, cltSess, tm.curTime, tm.logger)

	sm.ClientAccessFinished(cltSess, tm.curTime, tm.logger)
	tm.lastSID, _ = cltSess.GetClientInfo()

	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Refresh SID: %s, Logout user through delete cookie.", tm.lastSID)
}

func (tm *testMgr) verifyStep7Result(ckyLst []*http.Cookie) {
	var bSID bool
	for _, v := range ckyLst {
		if v.Name == cStrCookieNameSessionID && v.Value == tm.lastSID {
			bSID = true
			continue
		}
	}

	if len(ckyLst) != 1+1 || !bSID {
		printCookies(ckyLst)
		panic("verifyStep6Result failed")
	}
}

func (tm *testMgr) testStep8(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	tm.verifyStep7Result(ckyLst)

	sm := tm.sessMgr
	tm.curTime = tm.curTime.AddDate(0, 0, 1)
	sm.RunExpire(tm.curTime, tm.logger) //expire the cookie added from step 1

	fmt.Println("\n\n********** Test step 8: try to delete user cookie(no result shows) and login user, add two cookie **********")
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, tm.curTime, tm.logger)
	pannicErrors(err)

	sm.DelCookie(cStrCookieNameCurrentUser, cltSess, tm.curTime, tm.logger)
	sm.AddCookie(tstCookieNames[0], tstCookieValues[0], 3, cltSess, tm.curTime, tm.logger)
	sm.AddCookie(tstCookieNames[1], tstCookieValues[1], 1, cltSess, tm.curTime, tm.logger)
	sm.Login(cltSess, tstUserName, 3, tm.curTime, tm.logger)

	sm.ClientAccessFinished(cltSess, tm.curTime, tm.logger)
	tm.lastSID, _ = cltSess.GetClientInfo()

	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Refresh SID: %s, add two cookies: %s, %s, Login user", tm.lastSID, tstCookieNames[0], tstCookieNames[1])
}

func (tm *testMgr) verifyStep8Result(ckyLst []*http.Cookie) {
	var bSID, bUser, bCookie0, bCookie1 bool
	for _, v := range ckyLst {
		if v.Name == cStrCookieNameSessionID && v.Value == tm.lastSID {
			bSID = true
			continue
		}
		if v.Name == cStrCookieNameCurrentUser && v.Value == tstUserName {
			bUser = true
			continue
		}
		if v.Name == tstCookieNames[1] && v.Value == tstCookieValues[1] {
			bCookie1 = true
			continue
		}
		if v.Name == tstCookieNames[0] && v.Value == tstCookieValues[0] {
			bCookie0 = true
			continue
		}
	}

	if len(ckyLst) != 1+4 || !bSID || !bUser || !bCookie0 || !bCookie1 {
		printCookies(ckyLst)
		panic("verifyStep8Result failed")
	}
}

func (tm *testMgr) testStep9(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	tm.verifyStep8Result(ckyLst)

	sm := tm.sessMgr
	tm.curTime = tm.curTime.AddDate(0, 0, 2)
	sm.RunExpire(tm.curTime, tm.logger) //expire the cookie added from step 1

	fmt.Println("\n\n********** Test step 9: try to add user cookie with already login user and do a invalid login, add update cookie[0] **********")
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, tm.curTime, tm.logger)
	pannicErrors(err)

	_, oldUCoky := sm.ReadCookie(cStrCookieNameCurrentUser, cltSess, tm.logger)
	fmt.Printf("User info read from Cookie: %v", oldUCoky)

	sm.AddCookie(cStrCookieNameSessionID, tstCookieValues[0], 3, cltSess, tm.curTime, tm.logger)
	sm.AddCookie(cStrCookieNameCurrentUser, tstUserName, 10, cltSess, tm.curTime, tm.logger)

	sm.AddCookie(tstCookieNames[0], tstCookieValues[3], 9, cltSess, tm.curTime, tm.logger)
	sm.Login(cltSess, tstUserName, 3, tm.curTime, tm.logger)

	_, newUCoky := sm.ReadCookie(cStrCookieNameCurrentUser, cltSess, tm.logger)
	fmt.Printf("User info read from Cookie: %v", newUCoky)

	sm.ClientAccessFinished(cltSess, tm.curTime, tm.logger)
	tm.lastSID, _ = cltSess.GetClientInfo()

	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Refresh SID: %s, cookie %s value change to %s, user expire from %v to %v",
		tm.lastSID, tstCookieNames[0], tstCookieValues[3], oldUCoky.Expires, newUCoky.Expires)
}

func (tm *testMgr) verifyStep9Result(ckyLst []*http.Cookie) {
	var bSID, bUser, bCookie0 bool
	for _, v := range ckyLst {
		if v.Name == cStrCookieNameSessionID && v.Value == tm.lastSID {
			bSID = true
			continue
		}
		if v.Name == cStrCookieNameCurrentUser && v.Value == tstUserName {
			bUser = true
			continue
		}
		if v.Name == tstCookieNames[0] && v.Value == tstCookieValues[3] {
			bCookie0 = true
			continue
		}
	}

	if len(ckyLst) != 1+3 || !bSID || !bUser || !bCookie0 {
		printCookies(ckyLst)
		panic("verifyStep9Result failed")
	}
}

func (tm *testMgr) testStep10(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	tm.verifyStep9Result(ckyLst)

	sm := tm.sessMgr
	tm.curTime = tm.curTime.AddDate(0, 0, 3).Add(time.Hour * 2)
	sm.RunExpire(tm.curTime, tm.logger) //expire the cookie added from step 1

	fmt.Println("\n\n********** Test step 10: User expire: now is %v **********", tm.curTime)
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, tm.curTime, tm.logger)
	pannicErrors(err)

	sm.ClientAccessFinished(cltSess, tm.curTime, tm.logger)
	tm.lastSID, _ = cltSess.GetClientInfo()

	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Refresh SID: %s, user expire logout, but keep cookie 0", tm.lastSID)
}

func (tm *testMgr) verifyStep10Result(ckyLst []*http.Cookie) {
	var bSID, bCookie0 bool
	for _, v := range ckyLst {
		if v.Name == cStrCookieNameSessionID && v.Value == tm.lastSID {
			bSID = true
			continue
		}
		if v.Name == tstCookieNames[0] && v.Value == tstCookieValues[3] {
			bCookie0 = true
			continue
		}
	}

	if len(ckyLst) != 1+2 || !bSID || !bCookie0 {
		printCookies(ckyLst)
		panic("verifyStep10Result failed")
	}
}

func (tm *testMgr) testStep11(ckyLst []*http.Cookie) ([]*http.Cookie, string) {
	tm.verifyStep10Result(ckyLst)

	sm := tm.sessMgr
	tm.curTime = tm.curTime.AddDate(0, 0, 1).Add(time.Hour * 2)
	sm.RunExpire(tm.curTime, tm.logger)

	fmt.Println("\n\n********** Test step 11: Expire SID during request. **********")
	fmt.Println("---> Cookie From request: ")
	printCookies(ckyLst)

	cltSess, err := sm.ClientAccess(tm.ip, ckyLst, &DummyBindDataImpl{}, tm.curTime, tm.logger)
	pannicErrors(err)

	tm.curTime = tm.curTime.AddDate(0, 0, 3).Add(time.Hour * 2)
	sm.RunExpire(tm.curTime, tm.logger)

	sm.ClientAccessFinished(cltSess, tm.curTime, tm.logger)

	tm.curTime = tm.curTime.AddDate(0, 0, 3).Add(time.Hour * 2)
	sm.RunExpire(tm.curTime, tm.logger)

	tm.lastSID, _ = cltSess.GetClientInfo()

	tm.curTime = tm.curTime.AddDate(0, 0, 3).Add(time.Hour * 2)
	sm.RunExpire(tm.curTime, tm.logger)
	ckyLstDiff := sm.GetUpdateCookieList(cltSess, ckyLst, tm.logger)
	fmt.Println("\n---> Cookie will send to client: ")
	printCookies(ckyLstDiff)

	return ckyLstDiff, fmt.Sprintf("Expire SID many times during request processing, will get noting.", tm.lastSID)
}

//************************************************************************************************************************
//other small tests
//TestGenClientID
//	go test -v -run GenClientID
func TestGenClientID(t *testing.T) {
	var gr genRandom
	reslt := gr.genCientIDByIPAndSID("127.0.100.109", gr.getRandomSID(time.Now()))
	fmt.Println(reslt)
	deStr, _ := b64.StdEncoding.DecodeString(reslt)
	fmt.Println(string(deStr))
	reslt = gr.genCientIDByIPAndSID("2001:db8:85a3:8d3:1319:8a2e:370:7348", gr.getRandomSID(time.Now()))
	fmt.Println(reslt)
	deStr, _ = b64.StdEncoding.DecodeString(reslt)
	fmt.Println(string(deStr))
}
