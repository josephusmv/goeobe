package eobeapiimpl

import (
	"eobe/eobecltmgmt"
	"fmt"
	"testing"
)

func testGetHarald() (*eobecltmgmt.ClientManager, *eobecltmgmt.Herald) {
	cm := eobecltmgmt.NewClientManager("./LOGS")
	cm.StartClientManagerServer()
	return cm, cm.NewHerald()
}

//TestCltMgrSmokeTest
//	go test -v  -run CltMgrSmokeTest
func TestCltMgrSmokeTest(t *testing.T) {
	cm, h := testGetHarald()
	defer cm.StopClientManagerServer()

	h.NewAccess("127.0.0.1", nil, nil, nil)

	const testUserName = "userXYZ"
	var apiParamLogin = "?uname, ?expdays"

	qryKVMap := make(map[string]string)
	qryKVMap["uname"] = testUserName
	qryKVMap["expdays"] = "7"

	apiLgin, err := GetAPIImplementation(cStrLogInUser, apiParamLogin)
	checkError(err, "GetAPIImplementation() error.")
	apiLgin.SetDataSrc(nil, h)
	resultLgin, apiErrLgin := apiLgin.RunAPI(qryKVMap, nil)
	checkAPIError(apiErrLgin, "RunAPI() "+cStrLogInUser+" error.")
	fmt.Println(resultLgin)

	//run a duplicate login
	resultLgin, apiErrLgin = apiLgin.RunAPI(qryKVMap, nil)
	checkAPIError(apiErrLgin, "RunAPI() "+cStrLogInUser+" error.")
	fmt.Println(resultLgin)
	if resultLgin["retLoginUser"] != testUserName {
		panic("Wrong test result:" + resultLgin[cStrGetCurrentUserRetUserName])
	}

	apiCu, errC := GetAPIImplementation(cStrGetCurrentUser, "")
	checkError(errC, "GetAPIImplementation() error.")
	apiCu.SetDataSrc(nil, h)
	resultCU, apiCErr := apiCu.RunAPI(qryKVMap, nil)
	checkAPIError(apiCErr, "RunAPI() "+cStrLogInUser+" error.")
	fmt.Println(resultCU)
	if resultCU[cStrGetCurrentUserRetUserName] != testUserName {
		panic("Wrong test result:" + resultCU[cStrGetCurrentUserRetUserName])
	}

	apiLgot, err := GetAPIImplementation(cStrLogOutUser, "")
	checkError(err, "GetAPIImplementation() error.")
	apiLgot.SetDataSrc(nil, h)
	resultLgot, apiErrLgot := apiLgot.RunAPI(qryKVMap, nil)
	checkAPIError(apiErrLgot, "RunAPI() "+cStrLogOutUser+" error.")
	fmt.Println(resultLgot)

	resultCU, apiCErr = apiCu.RunAPI(qryKVMap, nil)
	checkAPIError(apiCErr, "RunAPI() "+cStrLogInUser+" error.")
	fmt.Println(resultCU)
	if resultCU[cStrGetCurrentUserRetUserName] != "" {
		panic("Wrong test result:" + resultCU[cStrGetCurrentUserRetUserName])
	}

	//bad try a login with error parameter
	apiParamLogin = "?uname, ?NOTEXISTED"
	apiLgin, err = GetAPIImplementation(cStrLogInUser, apiParamLogin)
	checkError(err, "GetAPIImplementation() error.")
	apiLgin.SetDataSrc(nil, h)
	resultLgin, apiErrLgin = apiLgin.RunAPI(qryKVMap, nil)
	checkAPIErrorType(apiErrLgin, CErrBadCallError)
}
