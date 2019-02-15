package eobecltmgmt

import (
	"eobe/eobecltmgmt/eobesession"
	"net/http"
)

const (
	cStrCmdNewAccess    = "CmdNewAccess"
	cStrCmdLogin        = "CmdLogin"
	cStrCmdLogout       = "CmdLogout"
	cStrCmdAddCookie    = "CmdAddCookie"
	cStrCmdDelCookie    = "CmdDelCookie"
	cStrCmdReadCookie   = "CmdReadCookie"
	cStrCmdGetUpdates   = "CmdGetUpdates"
	cStrcmdSetBindData  = "cmdSetBindData"
	cStrcmdGetBindData  = "cmdGetBindData"
	cStrCmdFinishAccess = "CmdFinishAccess"
)

type commandMessage struct {
	actionCmd string
	inData    commandData
	retChan   chan resultMessage
}

type commandData struct {
	ip         string
	reqCookies []*http.Cookie
	cses       eobesession.ClientSessionInf
	bindData   eobesession.BindDataInf
	logger     eobesession.LoggerInf
	name       string //Cookie name
	value      string //Cookie Value, also used as SID or User Name
	expireDays int
}

type resultMessage struct {
	err         error
	cses        eobesession.ClientSessionInf //ClientAccess command result
	rsltCookie  *http.Cookie                 //ReadCookie command result
	diffCookies []*http.Cookie               //GetUpdateCookieList command result
	bindData    eobesession.BindDataInf      //GetBindData command result
}

type quitMsg struct {
	done chan bool
}
