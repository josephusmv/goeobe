package eobesession

import "time"

//SessionExpireMgr ...
type SessionExpireMgr struct {
}

//RunExpire Run Expire check for a specific clientSession
func (sem SessionExpireMgr) RunExpire(cltSes *clientSession, nowTime time.Time, logger LoggerInf) bool {
	//cltSes is in request is set from SessionManager.ClientAccess and cleared from SessionManager.ClientAccessFinished
	//During that procedure, lock this client session to not do the expire session.
	if cltSes.isInRequest() {
		return false
	}

	var userForcedInvalid bool
	if cltSes.bd != nil {
		userForcedInvalid = !cltSes.bd.Validate(cltSes.userName(), cltSes.sidStr(), cltSes.GetAllCookies())
	}

	if cltSes.pSid == nil {
		cltSes.ClearAll()
		return true
	}

	if cltSes.pSid.expireCheck(nowTime, logger) || userForcedInvalid {
		cltSes.ClearAll()
		return true
	}

	if cltSes.pUser != nil && cltSes.pUser.expireCheck(nowTime, logger) {
		cltSes.Logout(nowTime, logger)
	}

	cltSes.cookies.expireCheck(nowTime, logger)

	return false
}
