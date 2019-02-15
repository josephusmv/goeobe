package eobeapi

import (
	impl "eobe/eobeapi/eobeapiimpl"
	"fmt"
	"testing"
)

func tstCallSequence(calls []string, noQyrMap bool) CallSeqResult {
	dbQry, dbSrv := tstInitDBConn()
	defer dbSrv.StopDBServer()

	cm, hr := tstGetHarald()
	hrErr := hr.NewAccess("127.0.0.1", nil, nil, nil)
	tstCheckError(hrErr, "Call Herald.NewAccess error")
	defer cm.StopClientManagerServer()

	dbQryMap := tstGetDBActions()
	userConsts := tstGetUserConst()
	var qryParam map[string]string
	if !noQyrMap {
		qryParam = tstGetQueryParams()
	}

	cs := GetCallSequence(tst_loggerImpl{})
	cs.InitAPIFactory(userConsts, hr, dbQry, dbQryMap)
	cs.SetFileMap(testSetFile())

	return cs.Execute(qryParam, calls)
}

func tstDumpCallResult(csr CallSeqResult, exptErr impl.APIErrType) {
	if csr.ApiErr.ErrType != exptErr {
		fmt.Printf("\t Unexpected error: %s\n", csr.ApiErr.Error())
		panic(csr.ApiErr)
	}

	fmt.Println("\n-------- Dump CallSeqResult ------\n")
	fmt.Printf("API calls Single result: \n")
	for k, v := range csr.SingleRow {
		fmt.Printf("\t-->[%s]: %s \n", k, v)
	}

	if exptErr != impl.CErrSuccess {
		fmt.Printf("----> got error as expected: %s\n", csr.ApiErr.Error())
	}

	fmt.Printf("API calls Multi result: \n")
	fmt.Printf("\n----> Row names:\n\t")
	for _, v := range csr.MNames {
		fmt.Printf("%s,\t", v)
	}

	fmt.Println("\n----> Rows:\n")
	for _, v := range csr.MultiRow {
		fmt.Printf("\t %v\n", v)
	}
	fmt.Println("\n-------- Dump Done ------\n")

}

//TestCallSeqSmoke
//	go test -v  -run CallSeqSmoke
func TestCallSeqSmoke(t *testing.T) {
	callStr := []string{"(SaveAll)GetByUName: ?username"}

	callRslt := tstCallSequence(callStr, false)
	tstDumpCallResult(callRslt, impl.CErrSuccess)
}

//TestCallSeqRangeDB
//	go test -v  -run CallSeqRangeDB
func TestCallSeqRangeDB(t *testing.T) {
	callStr := []string{"GetRangeValues: ?srchkey",
		"(SaveAll)RangeGetRows: GetValuesByRange$$ ^testfield1"}

	callRslt := tstCallSequence(callStr, false)
	tstDumpCallResult(callRslt, impl.CErrSuccess)
}

//TestCallSeqFilterRows
//	go test -v  -run CallSeqFilterRows
func TestCallSeqFilterRows(t *testing.T) {
	callStr := []string{"GetByUName: ?username",
		"(SaveAll)FilterMultiRows: ?startIndx, ?count"}

	callRslt := tstCallSequence(callStr, false)
	tstDumpCallResult(callRslt, impl.CErrSuccess)
}

//TestCallSaveFile
//	go test -v  -run CallSaveFile
func TestCallSaveFile(t *testing.T) {
	callStr := []string{"SaveFile: ?FileName,  $LOCALStorePATH"}

	callRslt := tstCallSequence(callStr, false)
	tstDumpCallResult(callRslt, impl.CErrSuccess)
}

//TestCallSeqErros
//	go test -v  -run CallSeqErros
func TestCallSeqErros(t *testing.T) {
	callStr := []string{"(UnknownAction)ValidateInt: ^retAffectedRows, neq, $ValueZero, $InvalidLogin",
		"SaveFile: ?FileName,  $LOCALStorePATH"}

	callRslt := tstCallSequence(callStr, false)
	tstDumpCallResult(callRslt, impl.CErrServerInternalError)

	//No call string returns a success with empty results data
	callRslt = tstCallSequence(nil, false)
	tstDumpCallResult(callRslt, impl.CErrSuccess)

	callRslt = tstCallSequence(callStr, true)
	tstDumpCallResult(callRslt, impl.CErrServerInternalError)
}

//TestCallSeqDBActionWrong
//	go test -v  -run CallSeqDBActionWrong
func TestCallSeqDBActionWrong(t *testing.T) {
	callStr := []string{"UpdateNewItem: ?FileName;  $LOCALStorePATH"}

	callRslt := tstCallSequence(callStr, false)
	tstDumpCallResult(callRslt, impl.CErrServerInternalError)

}
