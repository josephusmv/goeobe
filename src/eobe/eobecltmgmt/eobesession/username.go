package eobesession

import (
	"net/http"
	"time"
)

const cStrCookieNameCurrentUser = "USER"

type userName struct {
	uname     string
	pUserCkie *http.Cookie
}

func newUserName(uname string, expire time.Time) *userName {
	cookie := http.Cookie{
		Name:    cStrCookieNameCurrentUser,
		Value:   uname,
		Expires: expire,
		Path:    cStrCurrentPath}

	return &userName{uname: uname, pUserCkie: &cookie}
}

func (user userName) get() (string, *http.Cookie) {
	return user.uname, user.pUserCkie
}

func (user *userName) expireCheck(nowTime time.Time, logger LoggerInf) bool {
	var tCal TimeCalculator

	if tCal.TimeBefore(user.pUserCkie.Expires, nowTime) {
		logger.TraceDev(cStrUSRExpired, user.uname, user.pUserCkie.Expires.String(), nowTime.String())
		return true
	}

	return false
}
