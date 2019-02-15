package eobecltmgmt

import (
	"eobe/eobecltmgmt/eobesession"
	"fmt"
	"net/http"
	"testing"
)

type tstBindDataImpl struct {
}

func (tbd *tstBindDataImpl) Validate(name string, sid string, pCList map[string]*http.Cookie) bool {
	return true
}
func (tbd *tstBindDataImpl) Clear() {
	return
}
func (tbd *tstBindDataImpl) Login(name string) {

}
func (tbd *tstBindDataImpl) Logout() {

}

func panicErrors(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

//TestSmokeTestOneRequest
//	go test -v -cover -coverprofile cover.out -run SmokeTestOneRequest
//	go tool cover -html=cover.out -o cover.html
func TestSmokeTestOneRequest(t *testing.T) {
	cm := NewClientManager()
	cm.StartClientManagerServer()

	var user, value string
	var ckyLst []*http.Cookie
	reqCky := []*http.Cookie{}

	h := cm.NewHerald()
	err := h.NewAccess("127.0.0.1", reqCky, &tstBindDataImpl{}, pckgLogger{})
	panicErrors(err)

	//User smoke tests
	err = h.SignInUser("TSTADMIN", 7)
	panicErrors(err)

	user, err = h.GetCurrentUser()
	panicErrors(err)
	fmt.Println(user)
	err = h.SignOutUser()
	panicErrors(err)
	user, err = h.GetCurrentUser()
	panicErrors(err)
	fmt.Println(user)

	//Cookie smoke tests
	err = h.AddCookie("TSTCOOKIE", "VALUE001", 7)
	panicErrors(err)
	value, err = h.GetCurrentUser()
	panicErrors(err)
	fmt.Println(value)
	err = h.DeleteCookie("TSTCOOKIE")
	panicErrors(err)

	//get updates
	ckyLst, err = h.GetUpdatedCookieList()
	panicErrors(err)
	fmt.Println(ckyLst)

	//bind data smoke tests
	var bd eobesession.BindDataInf
	bd, err = h.GetBindData()
	panicErrors(err)
	fmt.Println(bd)
	err = h.SetBindData(&tstBindDataImpl{})
	panicErrors(err)

	err = h.FinishAccess()
	panicErrors(err)

	cm.StopClientManagerServer()
}
