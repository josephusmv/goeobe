package eobesession

import (
	"net/http"
	"time"
)

//cookieList This struct should only be operated by serialized way
//	cookieList contains list for only one client, may not very large, so need only one MAP is enough.
type cookieList struct {
	cookies map[string]*http.Cookie
}

//Helpers - copyCookie
func (cl cookieList) copyCookie(pSrc *http.Cookie) (pDst *http.Cookie) {
	if pSrc == nil {
		return nil
	}

	pDst = &http.Cookie{}
	pDst.Name = pSrc.Name
	pDst.Value = pSrc.Value
	pDst.Path = pSrc.Path
	pDst.Domain = pSrc.Domain
	pDst.Expires = pSrc.Expires
	pDst.RawExpires = pSrc.RawExpires
	pDst.MaxAge = pSrc.MaxAge
	pDst.Secure = pSrc.Secure
	pDst.HttpOnly = pSrc.HttpOnly
	pDst.Raw = pSrc.Raw
	pDst.Unparsed = pSrc.Unparsed
	return
}

//Helpers - newCookie
func (cl cookieList) newCookie(pOrgin *http.Cookie, name, value string, pExpire *time.Time) *http.Cookie {
	var pCookie *http.Cookie

	if pOrgin == nil {
		pCookie = &http.Cookie{}
	} else {
		pCookie = pOrgin
	}

	pCookie.Name = name
	pCookie.Value = value
	pCookie.Expires = *pExpire
	pCookie.Path = cStrCurrentPath

	return pCookie
}

//isEmptyList return: Is empty cookielist and make a new map
func (cl *cookieList) isEmptyList() bool {
	if cl.cookies == nil {
		cl.cookies = make(map[string]*http.Cookie)
		return true
	}

	if len(cl.cookies) == 0 {
		return true
	}

	return false
}

//Read Make a copy and return result
func (cl *cookieList) Read(name string) *http.Cookie {
	if cl.isEmptyList() {
		return nil
	}

	pCookie, found := cl.cookies[name]
	if !found || pCookie == nil {
		return nil
	}

	return cl.copyCookie(pCookie)
}

func (cl *cookieList) Add(name string, pCookie *http.Cookie) *http.Cookie {
	cl.isEmptyList() //make a map if nil
	cl.cookies[name] = pCookie
	return pCookie
}

func (cl *cookieList) AddByValues(name, value string, expire time.Time) *http.Cookie {
	cl.isEmptyList() //make a map if nil

	//IF existed, call update
	_, found := cl.cookies[name]
	if found {
		return cl.Update(name, value, expire)
	}

	pCookie := cl.newCookie(nil, name, value, &expire)

	return cl.Add(name, pCookie)
}

func (cl *cookieList) Update(name, value string, expire time.Time) *http.Cookie {
	cl.isEmptyList()

	//IF non-existed, call Add By Name
	pCookie, found := cl.cookies[name]
	if !found {
		return cl.AddByValues(name, value, expire)
	}

	pCookie = cl.newCookie(pCookie, name, value, &expire)

	return pCookie
}

func (cl *cookieList) Delete(name string) *http.Cookie {
	if cl.isEmptyList() {
		return nil
	}

	pCookie, found := cl.cookies[name]
	if !found {
		return nil
	}

	pCopy := cl.copyCookie(pCookie)

	delete(cl.cookies, name)

	return pCopy
}

func (cl cookieList) CopyAllCookies() map[string]*http.Cookie {
	allCookies := make(map[string]*http.Cookie, len(cl.cookies))

	for key, pCoky := range cl.cookies {
		allCookies[key] = cl.copyCookie(pCoky)
	}

	return allCookies
}

func (cl *cookieList) ClearAll() {
	for k := range cl.cookies {
		delete(cl.cookies, k)
	}
}

func (cl *cookieList) expireCheck(nowTime time.Time, logger LoggerInf) {
	var tCal TimeCalculator
	for k, v := range cl.cookies {
		if tCal.TimeBefore(v.Expires, nowTime) {
			logger.TraceDev(cStrCookieExpired, k, v.Expires.String(), nowTime.String())
			delete(cl.cookies, k)
		}
	}
}
