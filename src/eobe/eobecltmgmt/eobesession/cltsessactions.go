package eobesession

import (
	"net/http"
	"time"
)

const cIntTenYearsMeansNeverExpire = 36500

//Internal logics for Client Session
//	Just for seperates responsibilities.
//	(I just don't like C# style that putting class methods into different files)
type cltSessActions struct {
	ClientSessionImpl
	ip        string //purely for logging..
	inRequest bool
	pGR       *genRandom
}

func (cltSes *cltSessActions) setSID(nowTime time.Time, logger LoggerInf) {
	if cltSes.pSid != nil {
		return
	}

	sidExpire := nowTime.AddDate(0, 0, cIntCookieExpireDaysSID)
	cltSes.pSid = newSessionID(sidExpire, cltSes.pGR)
	logger.TraceDev(cStrDEVSetSIDNEW, cltSes.sidStr())
}

func (cltSes cltSessActions) getClientID(logger LoggerInf) string {
	pCookie := cltSes.cookies.Read(cStrClientIDCookieName)
	if pCookie == nil {
		logger.TraceError(cStrErrorClientIDNotFound, cltSes.ip)
		return ""
	}

	return pCookie.Value
}

//genClientID Should not call from other place except for ipTable.
func (cltSes *cltSessActions) genClientID(ip string, nowTime time.Time, logger LoggerInf) string {
	cidStr := cltSes.pGR.genCientIDByIPAndSID(ip, cltSes.sidStr())

	//Add a never expire cookie to record Client ID. entry expire is un related to this id.
	neverExpire := nowTime.AddDate(0, 0, cIntTenYearsMeansNeverExpire)
	pCookie := cltSes.cookies.AddByValues(cStrClientIDCookieName, cidStr, neverExpire)
	if pCookie == nil {
		logger.TraceError(cStrErrorFailedToAddClientID, cStrClientIDCookieName, cidStr)
	} else {
		logger.TraceDev(cStrAddCookieSuccess, cStrClientIDCookieName, cidStr)
	}

	return cidStr
}

func (cltSes *cltSessActions) startRequest() {
	cltSes.inRequest = true
}

func (cltSes *cltSessActions) finishRequest() {
	cltSes.inRequest = false
}

func (cltSes cltSessActions) isInRequest() bool {
	return cltSes.inRequest
}

//RefreshSID	get runned when: new client access comes, user login/logout
func (cltSes *clientSession) RefreshSID(justExtend bool, nowTime time.Time, logger LoggerInf) (string, *http.Cookie) {

	if cltSes.pSid == nil {
		cltSes.ClearAll()
		cltSes.setSID(nowTime, logger)
		return cltSes.sidStr(), cltSes.sidCookie()
	}

	var sidExpire time.Time

	if cltSes.pUser != nil {
		uCoky := cltSes.userNameCookie()
		if uCoky == nil {
			logger.TraceError(cStrFatalErrorInUserCookie, cltSes.userName())
			cltSes.resetUser()
			sidExpire = nowTime.AddDate(0, 0, cIntCookieExpireDaysSID)
		} else {
			sidExpire = uCoky.Expires.AddDate(0, 0, 1)
		}
	} else {
		sidExpire = nowTime.AddDate(0, 0, cIntCookieExpireDaysSID)
	}

	logger.TraceDev(cStrNewSIDRefresh, sidExpire)

	if justExtend {
		cltSes.pSid.exend(sidExpire)
	} else {
		cltSes.pSid.refresh(sidExpire)
	}

	return cltSes.sidStr(), cltSes.sidCookie()
}
