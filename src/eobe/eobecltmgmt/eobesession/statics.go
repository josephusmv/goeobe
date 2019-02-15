package eobesession

import "time"

//Global Variables for debug.
var FlagEnableDevLog = true
var FlagEnablePanicNoAssertion = false //some (nearly impossisble) panics will be checked.

//Key words
const cStrCurrentPath = "/"
const cStrCurrentPathLocal = "./"
const cStrSlash = "\\/"
const cStrSingleSlash = "/"

//Error messages
const cStrAddCookieResvSIDError = "Add Cookie Error: " + cStrCookieNameSessionID + " is an reserved cookie name"
const cStrAddCookieResvUserError = "Add Cookie Error: " + cStrCookieNameCurrentUser + " is an reserved cookie name"
const cStrAddCookieFailed = "Failed to Add Cookie: %s - %s"
const cStrDeleteCookieResvSIDError = "Delete Cookie Error: " + cStrCookieNameSessionID + " is an reserved cookie name"
const cStrDeleteCookieResvUserError = "Delete Cookie Error: " + cStrCookieNameCurrentUser + " is an reserved cookie name"
const cStrDeleteCookieNotFound = "Try to Delete but not found: %s"
const cStrClientAccessFailed = "Client Access Failed Error, fatal error."
const cStrTypeAssertionClientSessionInfFailed = "ClientSessionInf must be get from SessionManager"
const cStrErrorNULLBINDDATA = "Nil bind Data is not acceptable"
const cStrFatalErrorInUserCookie = "Fatal Error In User Cookie : User %s has nil cookie."
const cStrErrorClientIDError = "FATAL - Client ID Mismatch, debug info: CID from client and Index: %s, CID from Client Session: %s, Client IP: %s"
const cStrErrorClientIDNotFound = "FATAL - Client ID Not Found for IP: %s"
const cStrErrorFailedToAddClientID = "FATAL - Failed to Add ClientID Cookie: %s - %s"

//DEV Message
const cStrAddCookieSuccess = "Successfully Add Cookie: %s - %s"
const cStrLoginUserSuccess = "Successfully Login User: %s, New SID: %s"
const cStrLogoutUserSuccess = "Successfully Logout User: %s, New SID: %s"
const cStrDeleteCookieSuccess = "Successfully Deleted Cookie: %s"
const cStrGetUserInfoNotLogin = "Client: %s is not logged in yet."
const cStrCookieNotFound = "Cookie %s for client %s not found"
const cStrCookieExpired = "Cookie %s has expired at %s, now: %s"
const cStrSIDExpired = "SID %s has expired at %s, now: %s"
const cStrUSRExpired = "User %s has expired at %s, now: %s"
const cStrAddANewClient = "Has added a new client, ip: %s, client id: %s"
const cStrDeleteClient = "Delete client, ip: %s, sid: %s"
const cStrDEVSetSIDNEW = "Create a New SID: %s"
const cStrNewSIDAfterRefresh = "New SID after refresh: %s"
const cStrNewSIDRefresh = "New SID expire: %v"

//Time
const cStrTimeFmts = "2006-Jan-02"

var cTimeLongTimesAgo, _ = time.Parse(cStrTimeFmts, "1970-Jan-01")

const CStrKeyWordSessionID = cStrCookieNameSessionID
const CStrKeyWordCurrentUser = cStrCookieNameCurrentUser
const CStrKeyWordClientID = cStrClientIDCookieName
