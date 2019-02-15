package eobesession

import (
	"net/http"
	"time"
)

//SessionManager  encapsulate the user session behaviour
//	Wrap all functionalities to support call from both single threading or multi-threading.
type SessionManager struct {
	clientIPTable
}

func NewSessionManager() *SessionManager {
	var sesMgr SessionManager
	sesMgr.init()
	return &sesMgr
}

func (sesMgr *SessionManager) InitSessionManager() {
	sesMgr.init()
}

//When a Client access, parse the cookie list to get SID, and use IP+SID to found the client session info.
//	unless error, otherwise alway has an valid return
//	bindData: if found will ignore this one.
//	TODO: There is a potential security risck here: If a client send massive cookies with no cid, there will be massive new client added.
//	TODO: Solution is avoid this by an IP access table, denial and clean(DELETE) all possible mal-purpose IP
func (sesMgr *SessionManager) ClientAccess(ip string, reqClst []*http.Cookie, bindData BindDataInf, nowTime time.Time, logger LoggerInf) (ClientSessionInf, error) {
	if logger == nil {
		logger = NewDummyLogger()
	}

	var pCS *clientSession

	//Find CID cookie from client cookie list
	cidCookie := sesMgr.findCookie(cStrClientIDCookieName, reqClst)
	if cidCookie != nil {
		pCS = sesMgr.find(cidCookie.Value, logger)
	}

	//Either no SID cookie found(sidCookie == nil)
	//	OR ip + sid not found in IP table.
	if pCS == nil {
		pCS = sesMgr.add(ip, bindData, nowTime, logger)
	}

	if pCS == nil {
		err := logger.TraceError(cStrClientAccessFailed)
		return nil, err
	}

	newSid, _ := pCS.RefreshSID(true, nowTime, logger)
	logger.TraceDev(cStrNewSIDAfterRefresh, newSid)

	//Lock this pCS to not do the expire check.
	pCS.startRequest()
	return pCS, nil
}

//ClientAccessFinished	Unlock the pCltSes to enable expire check of this client
func (sesMgr *SessionManager) ClientAccessFinished(cses ClientSessionInf, nowTime time.Time, logger LoggerInf) {
	var pCltSes, ok = cses.(*clientSession)
	if !ok {
		logger.TraceError(cStrTypeAssertionClientSessionInfFailed)
		return
	}
	pCltSes.RefreshSID(false, nowTime, logger)
	pCltSes.finishRequest()
}

//GetUpdateCookieList ...
//		pCltSes the server(IPTable) stored  clientSession.
//		cookieList	client request stored cookies.
//		!!When this method called, means a request has finished.
func (sesMgr SessionManager) GetUpdateCookieList(cses ClientSessionInf, cookieList []*http.Cookie, logger LoggerInf) []*http.Cookie {
	newMap := cses.GetAllCookies()
	oldMap := make(map[string]*http.Cookie, len(cookieList))
	for _, pCkie := range cookieList {
		oldMap[pCkie.Name] = pCkie
	}

	return sesMgr.GetDifferntCookies(newMap, oldMap)
}

//GetDifferntCookies get a Cookie slice as differnt of two cookie map
//	Should save old before running API and compare with the new after running API
func (sesMgr SessionManager) GetDifferntCookies(new map[string]*http.Cookie, old map[string]*http.Cookie) []*http.Cookie {
	var minimalSize, maximalSize int

	if len(new) > len(old) {
		minimalSize = len(new) - len(old)
		maximalSize = len(new)
	} else {
		minimalSize = len(old) - len(new)
		maximalSize = len(old)
	}

	rslt := make([]*http.Cookie, minimalSize, maximalSize)
	currentIndx := 0

	//For each one in New, if it not in old or not eqaul to old.
	for k, newV := range new {
		oldV, found := old[k]
		if !found || sesMgr.isTwoCookieDiff(newV, oldV, true) {
			if currentIndx < minimalSize {
				rslt[currentIndx] = newV
				currentIndx++
			} else {
				rslt = append(rslt, newV)
			}
			continue
		}
	}

	//For each one in old, if it not in old or not eqaul to old.
	//mark as delete
	for k, oldV := range old {
		newV, found := new[k]
		if !found || sesMgr.isTwoCookieDiff(newV, oldV, false) { //delete if only in old
			delCookie := http.Cookie{Name: oldV.Name,
				Value:      oldV.Value,
				Path:       oldV.Path,
				Domain:     oldV.Domain,
				Expires:    cTimeLongTimesAgo,
				RawExpires: oldV.RawExpires,
				MaxAge:     -1,
				Secure:     oldV.Secure,
				HttpOnly:   oldV.HttpOnly,
				Raw:        oldV.Raw,
				Unparsed:   oldV.Unparsed}

			if currentIndx < minimalSize {
				rslt[currentIndx] = &delCookie
				currentIndx++
			} else {
				rslt = append(rslt, &delCookie)
			}
			continue
		}
	}

	return rslt
}

func (sesMgr SessionManager) isTwoCookieDiff(new *http.Cookie, old *http.Cookie, fullCheck bool) bool {
	if fullCheck {
		if new.Name != old.Name || new.Value != old.Value ||
			new.Path != old.Path || new.Domain != old.Domain ||
			new.RawExpires != old.RawExpires || new.MaxAge != old.MaxAge ||
			new.Secure != old.Secure || new.HttpOnly != old.HttpOnly ||
			new.Raw != old.Raw {
			return true
		}

		if !new.Expires.Equal(old.Expires) {
			return true
		}
	} else {
		if old.Name != new.Name {
			return true
		}
	}

	return false
}

//Login call login for ClientSessionInf
func (sesMgr *SessionManager) Login(cses ClientSessionInf, uname string, expireDays int, nowTime time.Time, logger LoggerInf) {
	if logger == nil {
		logger = NewDummyLogger()
	}

	var pCltSes, ok = cses.(*clientSession)
	if !ok {
		logger.TraceError(cStrTypeAssertionClientSessionInfFailed)
		return
	}

	pCltSes.Login(uname, expireDays, nowTime, logger)
	pCltSes.RefreshSID(true, nowTime, logger)
}

//Logout call logout for ClientSessionInf
func (sesMgr *SessionManager) Logout(cses ClientSessionInf, nowTime time.Time, logger LoggerInf) {
	if logger == nil {
		logger = NewDummyLogger()
	}

	var pCltSes, ok = cses.(*clientSession)
	if !ok {
		logger.TraceError(cStrTypeAssertionClientSessionInfFailed)
		return
	}

	pCltSes.Logout(nowTime, logger)
	pCltSes.RefreshSID(true, nowTime, logger)
}

//AddCookie call AddCookie for ClientSessionInf
func (sesMgr *SessionManager) AddCookie(name, value string, expireDays int, cses ClientSessionInf, nowTime time.Time, logger LoggerInf) {
	if logger == nil {
		logger = NewDummyLogger()
	}

	var pCltSes, ok = cses.(*clientSession)
	if !ok {
		logger.TraceError(cStrTypeAssertionClientSessionInfFailed)
		return
	}

	pCltSes.AddCookie(name, value, expireDays, nowTime, logger)
}

//AddCookie call AddCookie for ClientSessionInf
func (sesMgr *SessionManager) DelCookie(name string, cses ClientSessionInf, nowTime time.Time, logger LoggerInf) {
	if logger == nil {
		logger = NewDummyLogger()
	}

	var pCltSes, ok = cses.(*clientSession)
	if !ok {
		logger.TraceError(cStrTypeAssertionClientSessionInfFailed)
		return
	}

	pCltSes.DelCookie(name, nowTime, logger)
}

//ReadCookie call ReadCookie for ClientSessionInf
//    return the value and cookies
func (sesMgr *SessionManager) ReadCookie(name string, cses ClientSessionInf, logger LoggerInf) (string, *http.Cookie) {
	if logger == nil {
		logger = NewDummyLogger()
	}

	var pCltSes, ok = cses.(*clientSession)
	if !ok {
		logger.TraceError(cStrTypeAssertionClientSessionInfFailed)
		return "", nil
	}

	value, cookie := pCltSes.ReadCookie(name, logger)

	return value, pCltSes.copyCookie(cookie)
}

//SetBindData   ...
func (sesMgr *SessionManager) SetBindData(bindData BindDataInf, cses ClientSessionInf, logger LoggerInf) error {
	var pCltSes, ok = cses.(*clientSession)
	if !ok {
		logger.TraceError(cStrTypeAssertionClientSessionInfFailed)
		return nil
	}

	return pCltSes.SetBindData(bindData)
}

//SetBindData   ...
func (sesMgr *SessionManager) GetBindData(cses ClientSessionInf, logger LoggerInf) BindDataInf {
	var pCltSes, ok = cses.(*clientSession)
	if !ok {
		logger.TraceError(cStrTypeAssertionClientSessionInfFailed)
		return nil
	}

	return pCltSes.bindData()
}

//RunExpire   ...
func (sesMgr *SessionManager) RunExpire(nowTime time.Time, logger LoggerInf) {
	sesMgr.runExpireCheck(nowTime, logger)
}

func (sesMgr SessionManager) findCookie(name string, cookieList []*http.Cookie) *http.Cookie {
	for _, v := range cookieList {
		if name == v.Name {
			return v
		}
	}

	return nil
}
