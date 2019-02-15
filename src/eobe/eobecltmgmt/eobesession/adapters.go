package eobesession

import "net/http"

type BindDataInf interface {
	Validate(name string, sid string, pCList map[string]*http.Cookie) bool //The session entry will be droped, if this returns false.
	Clear()                                                                //Notify to this entry is going to be deleted.
	Login(name string)
	Logout()
}

//ClientSessionInf enable caller to debug dump
type ClientSessionInf interface {
	GetClientInfo() (string, string)
	GetAllCookies() map[string]*http.Cookie
	GetBindData() BindDataInf
	SetBindData(BindDataInf) error //Return error if nil interface
}

//LoggerInf Let User to implement this interface.
// If no concrete implementation provided, will use stdout(fmt.Print) as default.
type LoggerInf interface {
	TraceError(format string, a ...interface{}) error
	TraceDev(format string, a ...interface{})
	TraceInfo(format string, a ...interface{})
}
