package eobetests

import (
	"eobe/eobeapi"
	"eobe/eobecltmgmt"
	"eobe/eobedb"
	"testing"
)

func doTestAction(t *testing.T, inputInf testInputInf, outputInf TstOutInf, cm *eobecltmgmt.ClientManager, dbQry eobedb.DBQueryInf) bool {
	apiTester, err := newAPITester(inputInf.getRootPath(), cm, dbQry)
	if err != nil || apiTester == nil {
		t.Fail()
		return false
	}
	apiTester.setFileBytes(inputInf.loadFile())
	ok := apiTester.runHTTPRequest(inputInf.getActionName(), inputInf.getRequestMap(), inputInf)
	if !ok {
		return false
	}

	ok = apiTester.verifyResult(inputInf.getActionName(), outputInf)
	if !ok {
		t.Fail()
		return false
	}

	apiTester.finishTests()

	return true
}

type testInputInf interface {
	setActionName(string)
	getClientInfo() (ip, port, sid string)
	getActionName() string
	getRootPath() string
	getRequestMap() map[string]string
	loadFile() map[string][]byte
}

//TstOutInf ...also used as TLTst inputs.
type TstOutInf interface {
	GetOutput(action string) *eobeapi.CallSeqResult
}
