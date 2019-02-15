package eobecltmgmt

import (
	"eobe/eobecltmgmt/eobesession"
	"net/http"
)

//Herald herald design parttern.
//	Encapsulate all ClientManager's responsibility here, hide all ClientManager's details to outter multithreads
type Herald struct {
	cm      *ClientManager
	cmdData commandData
}

//internal!!
func (h *Herald) setLogger(logger eobesession.LoggerInf) {
	if h.cmdData.logger == nil {
		if logger == nil {
			h.cmdData.logger = eobesession.NewDummyLogger()
		} else {
			h.cmdData.logger = logger
		}
	}
}

//NewAccess
func (h *Herald) NewAccess(ip string, reqCookies []*http.Cookie, bindData eobesession.BindDataInf, logger eobesession.LoggerInf) error {
	h.cmdData.ip = ip
	h.cmdData.reqCookies = reqCookies

	if bindData == nil {
		h.cmdData.bindData = &eobesession.DummyBindDataImpl{}
	} else {
		h.cmdData.bindData = bindData
	}

	h.setLogger(logger)

	rsltMsg := h.cm.sendCommand(cStrCmdNewAccess, h.cmdData)
	if rsltMsg.err != nil {
		logger.TraceError("Execute New Access Command Failed with error: %s", rsltMsg.err.Error())
		return rsltMsg.err
	}

	h.cmdData.cses = rsltMsg.cses
	return nil
}

//SignInUser
func (h *Herald) SignInUser(user string, expireDays int, logger eobesession.LoggerInf) error {
	h.cmdData.name = CStrKeyWordCurrentUser
	h.cmdData.value = user
	h.cmdData.expireDays = expireDays

	h.setLogger(logger)

	rsltMsg := h.cm.sendCommand(cStrCmdLogin, h.cmdData)
	if rsltMsg.err != nil {
		h.cmdData.logger.TraceError("Execute SignIn User(%s) Command Failed with error: %s", user, rsltMsg.err.Error())
		return rsltMsg.err
	}

	return nil
}

//SignOutUser Let API Layer do the conversion from cookie to values. we just do record and management
//	Return: SID and Cookie list of the given SID
func (h *Herald) SignOutUser(logger eobesession.LoggerInf) error {
	h.setLogger(logger)
	rsltMsg := h.cm.sendCommand(cStrCmdLogout, h.cmdData)
	if rsltMsg.err != nil {
		h.cmdData.logger.TraceError("Execute SignOut User Command Failed with error: %s", rsltMsg.err.Error())
		return rsltMsg.err
	}
	return nil
}

//AddCookie ...
func (h *Herald) AddCookie(cookieName, cookieValue string, expireDays int, logger eobesession.LoggerInf) error {
	h.cmdData.name = cookieName
	h.cmdData.value = cookieValue
	h.cmdData.expireDays = expireDays

	h.setLogger(logger)
	rsltMsg := h.cm.sendCommand(cStrCmdAddCookie, h.cmdData)
	if rsltMsg.err != nil {
		h.cmdData.logger.TraceError("Execute Add Cookie Command Failed with error: %s", rsltMsg.err.Error())
		return rsltMsg.err
	}

	return nil
}

//DeleteCookie ...
func (h *Herald) DeleteCookie(cookieName string, logger eobesession.LoggerInf) error {
	h.cmdData.name = cookieName

	h.setLogger(logger)
	rsltMsg := h.cm.sendCommand(cStrCmdDelCookie, h.cmdData)
	if rsltMsg.err != nil {
		h.cmdData.logger.TraceError("Execute Delete Cookie Command Failed with error: %s", rsltMsg.err.Error())
		return rsltMsg.err
	}

	return nil
}

//ReadCookie  ..
func (h *Herald) ReadCookie(cookieName string, logger eobesession.LoggerInf) (*http.Cookie, error) {
	h.cmdData.name = cookieName
	h.setLogger(logger)
	rsltMsg := h.cm.sendCommand(cStrCmdReadCookie, h.cmdData)
	if rsltMsg.err != nil {
		h.cmdData.logger.TraceError("Execute Delete Cookie Command Failed with error: %s", rsltMsg.err.Error())
		return nil, rsltMsg.err
	}

	return rsltMsg.rsltCookie, nil
}

//GetCurrentUser  ..
func (h *Herald) GetCurrentUser(logger eobesession.LoggerInf) (string, error) {
	h.setLogger(logger)
	ckie, err := h.ReadCookie(CStrKeyWordCurrentUser, logger)
	if err != nil || ckie == nil {
		return "", err
	}

	return ckie.Value, err
}

//GetUpdatedCookieList  ..
func (h *Herald) GetUpdatedCookieList(logger eobesession.LoggerInf) ([]*http.Cookie, error) {
	h.setLogger(logger)
	rsltMsg := h.cm.sendCommand(cStrCmdGetUpdates, h.cmdData)
	if rsltMsg.err != nil {
		h.cmdData.logger.TraceError("Try to get updated cookie list Failed with error: %s", rsltMsg.err.Error())
		return nil, rsltMsg.err
	}

	return rsltMsg.diffCookies, nil
}

//GetBindData  ..
func (h *Herald) GetBindData(logger eobesession.LoggerInf) (eobesession.BindDataInf, error) {
	h.setLogger(logger)
	rsltMsg := h.cm.sendCommand(cStrcmdGetBindData, h.cmdData)
	if rsltMsg.err != nil {
		h.cmdData.logger.TraceError("Try to get bind data Failed with error: %s", rsltMsg.err.Error())
		return nil, rsltMsg.err
	}

	return rsltMsg.bindData, nil
}

//SetBindData  ..
func (h *Herald) SetBindData(bd eobesession.BindDataInf, logger eobesession.LoggerInf) error {
	h.cmdData.bindData = bd

	h.setLogger(logger)
	rsltMsg := h.cm.sendCommand(cStrcmdSetBindData, h.cmdData)
	if rsltMsg.err != nil {
		h.cmdData.logger.TraceError("Try to set bind data Failed with error: %s", rsltMsg.err.Error())
		return rsltMsg.err
	}

	return nil
}

//FinishAccess  ..
func (h *Herald) FinishAccess(logger eobesession.LoggerInf) error {
	h.setLogger(logger)
	rsltMsg := h.cm.sendCommand(cStrCmdFinishAccess, h.cmdData)
	if rsltMsg.err != nil {
		h.cmdData.logger.TraceError("Execute finish access command Failed with error: %s", rsltMsg.err.Error())
		return rsltMsg.err
	}

	return nil
}
