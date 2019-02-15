package eobetests

//This file contains the value verify for read data
//Need update hard code values if test file changes.

import (
	"eobe/eobeload"
	"fmt"
)

const cTstDBQueryActionName1 = "UpdatesAppTaskEntry"
const cTstExpectedWhereStr1 = "crtdt=?"
const cTstHTTPActionName = "LoginAdmin"
const cTstExpectedAPICallsIndex4 = "(SaveRslt)LogInUser: ?usr, ?expdays"
const cTstUserConstName = "UserForbidden"
const cTstExpectedUserConstValue = "User have not enough pervillidge to do this action."

//randomVerifyAllData Fill in expected data from file actually.
func randomVerifyAllData(daMap *eobeload.DBActionMap, haMap *eobeload.HTTPActionMap, uaMap *eobeload.UserConstsMap) (ok bool) {
	ok = true

	//Verify for DB Actions Map
	dm := daMap.GetData()
	qry, found := dm[cTstDBQueryActionName1]
	if !found {
		fmt.Printf("Loaded DB Action Data not contain expected: %s\n", cTstDBQueryActionName1)
		ok = false
	} else {
		if qry.WhereReadyStr != cTstExpectedWhereStr1 {
			fmt.Printf("Loaded DB WhereReadyStr(%s) not Match expected(%s).\n", qry.WhereReadyStr, cTstExpectedWhereStr1)
			ok = false
		} else {
			fmt.Printf(" Test passed for dm[%s].WhereReadyStr = (%s).\n", cTstDBQueryActionName1, qry.WhereReadyStr)
		}
	}

	//Verify for server resources Map
	hm := haMap.GetData()
	act, found := hm[cTstHTTPActionName]
	if !found {
		fmt.Printf("Loaded  server resources Data not contain expected: %s\n", cTstHTTPActionName)
		ok = false
	} else {
		if act.APICalls[4] != cTstExpectedAPICallsIndex4 {
			fmt.Printf("Loaded HTTP APICalls[4](%s) not Match expected(%s).\n", act.APICalls[4], cTstExpectedAPICallsIndex4)
			ok = false
		} else {
			fmt.Printf(" Test passed for hm[%s].APICalls[4] = (%s).\n", cTstHTTPActionName, act.APICalls[4])
		}
	}

	//Verify for  server resources Map
	ucm := uaMap.GetData()
	uc, found := ucm[cTstUserConstName]
	if !found {
		fmt.Printf("Loaded User const Data not contain expected: %s\n", cTstUserConstName)
		ok = false
	} else {
		if uc != cTstExpectedUserConstValue {
			fmt.Printf("Loaded User Const (%s) not Match expected(%s).\n", uc, cTstExpectedUserConstValue)
			ok = false
		} else {
			fmt.Printf(" Test passed for hm[%s] = (%s).\n", cTstUserConstName, uc)
		}
	}

	return ok
}
