package eobecltmgmt

import (
	"eobe/eobecltmgmt/eobesession"
	"time"
)

//DbgEmergencyLogFlag This is an flag switch for verbose trace logs in emergency situation.
var DbgEmergencyLogFlag = true

const cIntExpireCheckInterval = time.Second * 120

type worker struct {
	smgr     *sessMgrWrapper
	cmdChan  chan commandMessage
	quitChan chan quitMsg
	wlogger  pckgLogger //logger for woker thread, get from package main entrance.
}

func (w *worker) init() {
	w.wlogger.TraceInfo(cStrInfoInitWorker)
	w.smgr = newSessMgrWrapper()
	w.cmdChan = make(chan commandMessage)
	w.quitChan = make(chan quitMsg)

	go w.run()
}

//run should be private, only worker Init could start it.
func (w *worker) run() {
	defer func(w *worker) {
		close(w.cmdChan)
		close(w.quitChan)
		w.wlogger.TraceInfo(cStrInfoWorkerExited)
	}(w)

	w.wlogger.TraceInfo(cStrInfoWorkerStarted)

	timeDoCheck := time.Now().Add(cIntExpireCheckInterval)
	var qmsg quitMsg
	var quit bool
	for !quit {
		now := time.Now()
		select {
		case cmdMsg := <-w.cmdChan:
			w.doActionAsCmded(cmdMsg, now)
		case qmsg = <-w.quitChan:
			w.wlogger.TraceInfo(cStrQuitWorker)
			quit = true
			break
		default:
			tc := eobesession.TimeCalculator{}
			if tc.TimeBefore(timeDoCheck, now) {
				w.wlogger.TraceInfo(cStrRunExpireCheck, now.String())
				w.smgr.RunExpire(now, w.wlogger)
				timeDoCheck = now.Add(cIntExpireCheckInterval)
			}
		}
	}
	qmsg.done <- true
}

func (w *worker) doActionAsCmded(cmdMsg commandMessage, nowTime time.Time) {
	var outResult *resultMessage
	inData := cmdMsg.inData

	inData.logger.TraceDev(cStrCommandRecvd, cmdMsg.actionCmd)
	switch cmdMsg.actionCmd {
	case cStrCmdNewAccess:
		outResult = w.smgr.cmdClientAccess(inData, nowTime)
	case cStrCmdLogin:
		w.smgr.cmdLogin(inData, nowTime)
	case cStrCmdLogout:
		w.smgr.cmdLogout(inData, nowTime)
	case cStrCmdAddCookie:
		w.smgr.cmdAddCookie(inData, nowTime)
	case cStrCmdDelCookie:
		w.smgr.cmdDelCookie(inData, nowTime)
	case cStrCmdReadCookie:
		outResult = w.smgr.cmdReadCookie(inData)
	case cStrCmdGetUpdates:
		outResult = w.smgr.cmdGetUpdateCookieList(inData)
	case cStrcmdSetBindData:
		outResult = w.smgr.cmdSetBindData(inData)
	case cStrcmdGetBindData:
		outResult = w.smgr.cmdGetBindData(inData)
	case cStrCmdFinishAccess:
		w.smgr.cmdFinishAccess(inData, nowTime)
	default:
		inData.logger.TraceError(cStrErrorUnreconizedCmd, cmdMsg.actionCmd)
	}

	if outResult == nil {
		cmdMsg.retChan <- resultMessage{}
	} else {
		cmdMsg.retChan <- *outResult
	}

	close(cmdMsg.retChan)
}

func (w *worker) sendQuit() {
	var qmsg quitMsg
	qmsg.done = make(chan bool)

	w.quitChan <- qmsg

	if <-qmsg.done {
		w.wlogger.TraceInfo(cStrWokerGoRoutineExited)
	}

	close(qmsg.done)
}

func (w *worker) sendCommand(actionCmd string, inData commandData) *resultMessage {
	var cMsg commandMessage
	cMsg.actionCmd = actionCmd
	cMsg.inData = inData
	cMsg.retChan = make(chan resultMessage)

	w.cmdChan <- cMsg

	rsltMsg := <-cMsg.retChan
	return &rsltMsg
}
