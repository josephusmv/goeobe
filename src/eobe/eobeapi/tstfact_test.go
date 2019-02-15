package eobeapi

import (
	"fmt"
	"testing"
)

//TestFactorySmoke test connect to remote Server
//	go test -v  -run FactorySmoke
func TestFactorySmoke(t *testing.T) {
	dbQry, dbSrv := tstInitDBConn()
	defer dbSrv.StopDBServer()

	cm, hr := tstGetHarald()
	err := hr.NewAccess("127.0.0.1", nil, nil, nil)
	tstCheckError(err, "Call Herald.NewAccess error")
	defer cm.StopClientManagerServer()

	dbQryMap := tstGetDBActions()
	userConsts := tstGetUserConst()

	factory := apiFactory{
		dbQry:       dbQry,
		dbActDfnMap: dbQryMap,
		userConsts:  userConsts,
		hr:          hr}

	var tstAPICallStr = "ValidateInt: ^retAffectedRows, neq, $ValueZero, $InvalidLogin"
	apiDefn, err := factory.makeNewAPI(tstAPICallStr, nil)
	tstDumpAPIDefine(apiDefn)
	tstCheckError(err, "Call factory.makeNewAPI error")

	//cover all step actions
	tstAPICallStr = "RangeGetRows: GetByUName$$ ^testfield1"
	apiDefn, err = factory.makeNewAPI(cStrSaveRslt+tstAPICallStr, nil)
	tstDumpAPIDefine(apiDefn)
	tstCheckError(err, "Call factory.makeNewAPI error")

	tstAPICallStr = "RangeGetRows: GetByUName$$ ^testfield1"
	apiDefn, err = factory.makeNewAPI(cStrSaveRows+tstAPICallStr, nil)
	tstDumpAPIDefine(apiDefn)
	tstCheckError(err, "Call factory.makeNewAPI error")

	tstAPICallStr = "RangeGetRows: GetByUName$$ ^testfield1"
	apiDefn, err = factory.makeNewAPI(tstAPICallStr, nil)
	tstDumpAPIDefine(apiDefn)
	tstCheckError(err, "Call factory.makeNewAPI error")

	tstAPICallStr = "RangeGetRows: GetByUName$$ ^testfield1"
	apiDefn, err = factory.makeNewAPI(tstAPICallStr, nil)
	tstDumpAPIDefine(apiDefn)
	tstCheckError(err, "Call factory.makeNewAPI error")

	tstAPICallStr = cStrAPISaveFile + ": ?FileName,  ?LOCALStorePATH"
	apiDefn, err = factory.makeNewAPI(tstAPICallStr, nil)
	tstDumpAPIDefine(apiDefn)
	tstCheckError(err, "Call factory.makeNewAPI error")

	//Unknown action string
	tstAPICallStr = "(UnknownAction)ValidateInt: ^retAffectedRows, neq, $ValueZero, $InvalidLogin"
	apiDefn, err = factory.makeNewAPI(tstAPICallStr, nil)
	fmt.Println(err.Error())

	//Unknown action string
	tstAPICallStr = "UnknownActionValidateInt: ^retAffecte)dRows, neq, $ValueZero, $InvalidLogin"
	apiDefn, err = factory.makeNewAPI(tstAPICallStr, nil)
	fmt.Println(err.Error())

	//no call string
	apiDefn, err = factory.makeNewAPI(tstAPICallStr, nil)
	fmt.Println(err.Error())

	//invalid call string
	tstAPICallStr = "ValidateInt: ^retAffectedRows, neq, $ValueZero: $InvalidLogin"
	apiDefn, err = factory.makeNewAPI(tstAPICallStr, nil)
	fmt.Println(err.Error())

	//invalid DB
	tstAPICallStr = "NOSUCHDBAction: ^ABC; ?DFGH"
	apiDefn, err = factory.makeNewAPI(tstAPICallStr, nil)
	fmt.Println(err.Error())

	//range rows with wrong DB action calls
	tstAPICallStr = "RangeGetRows: GetByUName"
	apiDefn, err = factory.makeNewAPI(tstAPICallStr, nil)
	fmt.Println(err.Error())

	//range rows with wrong DB action calls
	tstAPICallStr = "(SaveAll)"
	apiDefn, err = factory.makeNewAPI(tstAPICallStr, nil)
	fmt.Println(err.Error())
}
