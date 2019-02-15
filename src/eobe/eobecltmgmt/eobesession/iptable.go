package eobesession

import (
	"time"
)

const cIntDefaultIpTableSize = 1024
const cIntDefaultIpTableCapacity = 1024 * 1024

//clientIPTable  single instance to hold all client IP binding.
//		this type will not concern whether to add, isduplicated an so on
//		this type only maintain IP table.
//		What if one IP with multiple clients? ---> use an IP+SID as unique hashID
type clientIPTable struct {
	cltSes map[string]*clientSession
	gr     genRandom
}

func (ipTble *clientIPTable) init() {
	ipTble.cltSes = make(map[string]*clientSession, cIntDefaultIpTableSize)
	ipTble.gr = genRandom{}
}

func (ipTble clientIPTable) find(cltID string, logger LoggerInf) *clientSession {
	pCS, found := ipTble.cltSes[cltID]
	if found {
		if pCS.getClientID(logger) == cltID {
			return pCS
		}

		//This is an error if CID as Index did not match CID in pCS
		logger.TraceError(cStrErrorClientIDError, cltID, pCS.getClientID(logger), pCS.ip)
		return nil
	}

	return nil
}

func (ipTble *clientIPTable) add(ip string, bindData BindDataInf, nowTime time.Time, logger LoggerInf) *clientSession {
	pCltSes := newClientSession(bindData, ip, &ipTble.gr, nowTime, logger)
	cid := pCltSes.genClientID(ip, nowTime, logger)

	ipTble.cltSes[cid] = pCltSes
	logger.TraceDev(cStrAddANewClient, ip, cid)
	return pCltSes
}

/*
func (ipTble clientIPTable) delete(ip, sid string, logger LoggerInf) *clientSession {
	pCS := ipTble.find(ip, sid, logger)
	if pCS != nil {
		logger.TraceDev(cStrDeleteClient, ip, pCS.pSid.sid())
		delete(ipTble.cltSes, ip)
		return pCS
	}
	return nil
}
*/

func (ipTble *clientIPTable) runExpireCheck(nowTime time.Time, logger LoggerInf) {
	var sem SessionExpireMgr
	for cid, pCS := range ipTble.cltSes {
		ip := pCS.ip
		sid := pCS.sidStr()

		if sem.RunExpire(pCS, nowTime, logger) {
			logger.TraceDev(cStrDeleteClient, ip, sid)
			delete(ipTble.cltSes, cid) //Must not call ipTble.delete, it involves too much logics
		}
	}
}
