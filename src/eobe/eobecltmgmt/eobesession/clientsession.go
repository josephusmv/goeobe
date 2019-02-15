package eobesession

import (
	"net/http"
	"time"
)

const cStrClientIDCookieName = "CID"

type clientSession struct {
	cltSessActions
}

func newClientSession(bd BindDataInf, ip string, pGR *genRandom, nowTime time.Time, logger LoggerInf) *clientSession {
	var cltSes clientSession

	cltSes.pGR = pGR
	cltSes.bd = bd
	cltSes.ip = ip
	cltSes.setSID(nowTime, logger)

	return &cltSes
}

func (cltSes *clientSession) Login(uname string, expireDays int, nowTime time.Time, logger LoggerInf) {
	loginExpire := nowTime.AddDate(0, 0, expireDays)
	cltSes.pUser = newUserName(uname, loginExpire)

	if cltSes.bd != nil {
		cltSes.bd.Login(uname)
	}

	logger.TraceDev(cStrLoginUserSuccess, uname, cltSes.sidStr())
}

func (cltSes *clientSession) Logout(nowTime time.Time, logger LoggerInf) {
	oldUname := cltSes.userName()
	cltSes.resetUser()

	if cltSes.bd != nil {
		cltSes.bd.Logout()
	}

	logger.TraceDev(cStrLogoutUserSuccess, oldUname, cltSes.sidStr())
}

//addCookie	Add cookie,
//	ignore if it's the SID cookie
//	Call log in if it's user cookie and user not logged in.
func (cltSes *clientSession) AddCookie(name, value string, expireDays int, nowTime time.Time, logger LoggerInf) {
	if name == cStrCookieNameSessionID {
		logger.TraceError(cStrAddCookieResvSIDError)
		return //Add SID cookie is not acceptable
	}

	if name == cStrCookieNameCurrentUser {
		if cltSes.pUser != nil {
			logger.TraceError(cStrAddCookieResvUserError)
			return //Add User Name cookie is not acceptable when Already login
		}

		cltSes.Login(value, expireDays, nowTime, logger)
		return
	}

	cookieExpire := nowTime.AddDate(0, 0, expireDays)
	pCookie := cltSes.cookies.AddByValues(name, value, cookieExpire)
	if pCookie == nil {
		logger.TraceError(cStrAddCookieFailed, name, value)
	} else {
		logger.TraceDev(cStrAddCookieSuccess, name, value)
	}
}

//delCookie	Delete cookie,
//	ignore if it's the SID cookie
//	Call log out if it's user cookie and user already logged in.
func (cltSes *clientSession) DelCookie(name string, nowTime time.Time, logger LoggerInf) {
	if name == cStrCookieNameSessionID {
		logger.TraceError(cStrDeleteCookieResvSIDError)
		return //Add SID cookie is not acceptable
	}

	if name == cStrCookieNameCurrentUser {
		if cltSes.pUser == nil {
			logger.TraceError(cStrDeleteCookieResvUserError)
			return //Add User Name cookie is not acceptable when Already login
		}

		cltSes.Logout(nowTime, logger)
		return
	}

	pCookie := cltSes.cookies.Delete(name)
	if pCookie == nil {
		logger.TraceInfo(cStrDeleteCookieNotFound, name)
	} else {
		logger.TraceDev(cStrDeleteCookieSuccess, name)
	}
}

//ReadCookie    return the value and cookies
func (cltSes clientSession) ReadCookie(name string, logger LoggerInf) (string, *http.Cookie) {
	if name == cStrCookieNameSessionID {
		return cltSes.sidStr(), cltSes.sidCookie()
	}

	if name == cStrCookieNameCurrentUser {
		return cltSes.userName(), cltSes.userNameCookie()
	}

	pCookie := cltSes.cookies.Read(name)
	if pCookie == nil {
		logger.TraceDev(cStrCookieNotFound, name, cltSes.sidStr())
		return "", nil
	}

	return pCookie.Value, pCookie
}
