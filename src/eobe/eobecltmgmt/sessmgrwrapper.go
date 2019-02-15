package eobecltmgmt

import (
	"eobe/eobecltmgmt/eobesession"
	"time"
)

type sessMgrWrapper struct {
	eobesession.SessionManager
}

func newSessMgrWrapper() *sessMgrWrapper {
	var sm sessMgrWrapper
	sm.InitSessionManager()
	return &sm
}

func (smw *sessMgrWrapper) cmdClientAccess(inData commandData, nowTime time.Time) *resultMessage {
	var rsltMsg resultMessage
	rsltMsg.cses, rsltMsg.err = smw.ClientAccess(inData.ip, inData.reqCookies, inData.bindData, nowTime, inData.logger)
	return &rsltMsg
}

func (smw *sessMgrWrapper) cmdLogin(inData commandData, nowTime time.Time) {
	smw.Login(inData.cses, inData.value, inData.expireDays, nowTime, inData.logger)
}

func (smw *sessMgrWrapper) cmdLogout(inData commandData, nowTime time.Time) {
	smw.Logout(inData.cses, nowTime, inData.logger)
}

func (smw *sessMgrWrapper) cmdAddCookie(inData commandData, nowTime time.Time) {
	smw.AddCookie(inData.name, inData.value, inData.expireDays, inData.cses, nowTime, inData.logger)
}

func (smw *sessMgrWrapper) cmdDelCookie(inData commandData, nowTime time.Time) {
	smw.DelCookie(inData.name, inData.cses, nowTime, inData.logger)
}

func (smw *sessMgrWrapper) cmdReadCookie(inData commandData) *resultMessage {
	var rsltMsg resultMessage
	_, rsltMsg.rsltCookie = smw.ReadCookie(inData.name, inData.cses, inData.logger)
	return &rsltMsg
}

func (smw *sessMgrWrapper) cmdGetUpdateCookieList(inData commandData) *resultMessage {
	var rsltMsg resultMessage
	rsltMsg.diffCookies = smw.GetUpdateCookieList(inData.cses, inData.reqCookies, inData.logger)
	return &rsltMsg
}

func (smw *sessMgrWrapper) cmdSetBindData(inData commandData) *resultMessage {
	var rsltMsg resultMessage
	rsltMsg.err = smw.SetBindData(inData.bindData, inData.cses, inData.logger)
	return &rsltMsg
}

func (smw *sessMgrWrapper) cmdGetBindData(inData commandData) *resultMessage {
	var rsltMsg resultMessage
	rsltMsg.bindData = smw.GetBindData(inData.cses, inData.logger)
	return &rsltMsg
}

func (smw *sessMgrWrapper) cmdFinishAccess(inData commandData, nowTime time.Time) {
	smw.ClientAccessFinished(inData.cses, nowTime, inData.logger)
}
