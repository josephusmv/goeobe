package eobesession

import "net/http"

type clientData struct {
	pSid    *sessionID
	pUser   *userName
	cookies cookieList
	bd      BindDataInf
}

//**************************************************
//getters for cd.pSID
func (cd clientData) sidStr() string {
	if cd.pSid == nil {
		return ""
	}

	return cd.pSid.sid()
}

func (cd clientData) sidCookie() *http.Cookie {
	if cd.pSid == nil {
		return nil
	}

	_, pCoky := cd.pSid.get()
	return pCoky
}

//**************************************************
//Getters for  cd.pUser
func (cd clientData) userName() string {
	if cd.pUser == nil {
		return ""
	}

	uName, _ := cd.pUser.get()
	return uName
}

func (cd clientData) userNameCookie() *http.Cookie {
	if cd.pUser == nil {
		return nil
	}

	_, uCoky := cd.pUser.get()
	return uCoky
}

//**************************************************
//getters for cd.bd
func (cd *clientData) bindData() BindDataInf {
	if cd.bd == nil {
		cd.bd = &DummyBindDataImpl{}
	}

	return cd.bd
}

//**************************************************
//getters for cd.cookies
func (cd clientData) copyCookie(pSrc *http.Cookie) *http.Cookie {
	return cd.cookies.copyCookie(pSrc)
}

func (cd clientData) copyGenCookies() map[string]*http.Cookie {
	return cd.cookies.CopyAllCookies()
}

func (cd clientData) copyAllCookies() map[string]*http.Cookie {
	if cd.pSid == nil {
		return nil
	}

	allcookies := cd.copyGenCookies()

	if cd.pUser != nil {
		allcookies[cStrCookieNameCurrentUser] = cd.copyCookie(cd.userNameCookie())
	}

	allcookies[cStrCookieNameSessionID] = cd.copyCookie(cd.sidCookie())

	return allcookies
}

//**************************************************
//Update operations
func (cd *clientData) ClearAll() {
	cd.cookies.ClearAll()
	if cd.bd != nil {
		cd.bd.Clear()
	}
	cd.resetUser()
	cd.pSid = nil
}

func (cd *clientData) resetUser() {
	cd.pUser = nil
}
