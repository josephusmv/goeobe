package eobehttp

import "net/http"

//No reentrances, only for one single request processing
type headerWritter struct {
	//mandatory
	valCacheControl string

	//optional
	valXSSProtection string
	valXContent      string
	valCharSet       string
	valHSTS          string
}

func (hdrw *headerWritter) setDefaultRespHeaders() {
	hdrw.valCacheControl = cStrDefaultValCacheCtrlNOT
	hdrw.valXSSProtection = cStrDefaultValXSSProtection
	hdrw.valXContent = cStrDefaultValXContent
	hdrw.valCharSet = cStrDefaultValCharSetUTF8
	hdrw.valHSTS = cStrDefaultValMustHTTPS
}

func (hdrw *headerWritter) addDefaultRespHeaders(w http.ResponseWriter) {
	pkgcfg := getPckConfig()

	serverStr := pkgcfg.ServerName + cStrDot + pkgcfg.ServerVersion
	w.Header().Add(cStrDefaultKeyServer, serverStr)
	w.Header().Add(cStrDefaultKeyCacheCtrl, hdrw.valCacheControl)

	if pkgcfg.AddXSSProtectionHeader {
		w.Header().Add(cStrDefaultKeyXSSProtection, hdrw.valXSSProtection)
	}
	if pkgcfg.AddXContentHeader {
		w.Header().Add(cStrDefaultKeyXContent, hdrw.valXContent)
	}
	if pkgcfg.AddCharSetHeader {
		w.Header().Add(cStrDefaultKeyCharSet, pkgcfg.ValHeaderCharSet)
	}
	if pkgcfg.AddHSTSHeader {
		w.Header().Add(cStrDefaultKeyHSTS, hdrw.valHSTS)
	}
}

func (rspw *responseWritter) setCacheControl(cc, lastModify, etag string) {
	rspw.valCacheControl = cc
}

//*************************************
// Header constants
// Not supported for now:
//		1. ETag -- we only do simple Cache-Control,  and don't involves more complexities for cache issues.
const cStrDefaultKeyServer = "Server"
const cStrDefaultValServer = "eobe"

const cStrDefaultKeyXSSProtection = "X-XSS-Protection"
const cStrDefaultValXSSProtection = "1; mode=block"

const cStrDefaultKeyXContent = "X-Content-Type-Options"
const cStrDefaultValXContent = "nosniff"

const cStrDefaultKeyCharSet = "charset"
const cStrDefaultValCharSetUTF8 = "utf-8"

const cStrDefaultKeyHSTS = "Strict-Transport-Security"
const cStrDefaultValMustHTTPS = "max-age=31536000; includeSubDomains"

const cStrDefaultKeyCacheCtrl = "Cache-Control"
const cStrDefaultValCacheCtrlNOT = "no-cache, no-store, must-revalidate"
const cStrDefaultValCacheCtrlSTATIC = "public, max-age=31536000"
