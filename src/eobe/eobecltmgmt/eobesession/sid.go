package eobesession

import (
	"net/http"
	"time"
)

const cStrCookieNameSessionID = "SID"
const cIntCookieExpireDaysSID = 1

type sessionID struct {
	idStr    string
	pSIDCkie *http.Cookie
	pGR      *genRandom
}

func newSessionID(expire time.Time, pGR *genRandom) *sessionID {
	var sid sessionID

	sid.pGR = pGR
	sid.idStr = sid.pGR.getRandomSID(expire)
	sid.setIDCookie(sid.idStr, expire)

	return &sid
}

//Reset Expire and generate a new SID string
func (sid *sessionID) refresh(expire time.Time) {
	sid.idStr = sid.pGR.getRandomSID(expire)
	sid.setIDCookie(sid.idStr, expire)
}

func (sid *sessionID) exend(expire time.Time) {
	sid.pSIDCkie.Expires = expire
}

func (sid *sessionID) setIDCookie(idStr string, expire time.Time) {
	sid.pSIDCkie = &http.Cookie{
		Name:    cStrCookieNameSessionID,
		Value:   idStr,
		Expires: expire,
		Path:    cStrCurrentPath}
}

func (sid sessionID) sid() string {
	return sid.idStr
}

func (sid sessionID) get() (string, *http.Cookie) {
	return sid.idStr, sid.pSIDCkie
}

//expireCheck   return true if SID has expired.
func (sid *sessionID) expireCheck(nowTime time.Time, logger LoggerInf) bool {
	var tCal TimeCalculator

	if tCal.TimeBefore(sid.pSIDCkie.Expires, nowTime) {
		logger.TraceDev(cStrSIDExpired, sid.idStr, sid.pSIDCkie.Expires.String(), nowTime.String())
		return true
	}

	return false
}
