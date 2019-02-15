package eobeapiimpl

import (
	cltmgrcmd "eobe/eobecltmgmt"
	"eobe/eobedb"
	"eobe/eobehttp"
	"fmt"
	"strings"
)

type ApiFactory struct {
	dbQry       eobedb.DBQueryInf
	dbActDfnMap map[string]*eobedb.QueryDefn
	userConsts  map[string]string
	hr          *cltmgrcmd.Herald
	logger      eobehttp.HttpLoggerInf
}

func (af *ApiFactory) SetLogger(logger eobehttp.HttpLoggerInf) {
	af.logger = logger
}

func (af *ApiFactory) PrepareForAPICalls(userConsts map[string]string, hr *cltmgrcmd.Herald) error {
	af.userConsts = userConsts
	af.hr = hr
	return nil
}

func (af *ApiFactory) PrepareForDBActions(dbQry eobedb.DBQueryInf, dbActMapmap map[string]*eobedb.QueryDefn) error {
	af.dbQry = dbQry
	af.dbActDfnMap = dbActMapmap
	return nil
}

func (af ApiFactory) MakeNewAPI(callStr string) (string, *ApiInf, int, error) {
	funTexts, saveTo := af.parsePrefix(callStr)
	if len(funTexts) != 2 {
		return "", nil, saveTo, fmt.Errorf(cStrErrorAPICallFormat, callStr)
	}

	var api *ApiInf
	var err error
	apiName := funTexts[0]
	apiParamInput := funTexts[1]
	switch apiName {
	case cStrValidateStrEqualName:
		api, err = newAPIValidateStrEqual(apiParamInput)
	/**************************/
	case cStrValidateStrRexName:
		api, err = newAPIValidateStrRegX(apiParamInput)
	/**************************/
	case cStrCompareInt:
		api, err = newCompareInt(apiParamInput)
	/**************************/
	case CAPIGetMultiRows:
		detailParams := strings.Split(apiParamInput, cStrMultiRowDBActSepe)
		if len(detailParams) != 2 {
			return apiName, nil, saveTo, fmt.Errorf(cStrErrorGetMultiRowsFormat)
		}
		name := strings.Trim(detailParams[0], cStrSpace)
		paramInput := strings.Trim(detailParams[1], cStrSpace)
		api, err = af.getDBAPIDefine(name, paramInput, CAPIGetMultiRows)
	/**************************/
	case CAPIRangeGetRows:
		detailParams := strings.Split(apiParamInput, cStrMultiRowDBActSepe)
		if len(detailParams) != 2 {
			return apiName, nil, saveTo, fmt.Errorf(cStrErrorGetMultiRowsFormat)
		}
		name := strings.Trim(detailParams[0], cStrSpace)
		paramInput := strings.Trim(detailParams[1], cStrSpace)
		api, err = af.getDBAPIDefine(name, paramInput, CAPIRangeGetRows)
	/**************************/
	case cStrLogInUser:
		api, err = newLogInUser(apiParamInput)
	/**************************/
	case cStrValidateInt:
		api, err = newValidateInt(apiParamInput)
	/**************************/
	case cStrRemoteEncrypt:
		api, err = newAPIRemoteEncrypt(apiParamInput)
	/**************************/
	case cStrLogOutUser:
		api, err = newLogOutUser(apiParamInput)
	/**************************/
	case cStrGetCurrentUser:
		api, err = newGetCurrentUser(apiParamInput)
	/**************************/
	case CAPIFilterMultiRowss:
		api, err = newFilterMultiRows(apiParamInput)
	/**************************/
	case cStrGetServerDTTM:
		api, err = newgetServerDTTM(apiParamInput)
	/**************************/
	case CAPISaveFile:
		api, err = newSaveFile(apiParamInput)
	/**************************/
	/*Add New Support APIes here!*/
	default: //DB Action
		api, err = af.getDBAPIDefine(apiName, apiParamInput, apiName)
	}

	if api == nil || *api == nil {
		return apiName, nil, saveTo, err
	}

	(*api).SetDataSrc(af.userConsts, af.hr)
	(*api).SetLogger(af.logger)

	return apiName, api, saveTo, err
}

func (af ApiFactory) parsePrefix(callStr string) (funTexts []string, saveTo int) {
	saveTo = cIntNotSave

	if strings.HasPrefix(callStr, COPSaveRslt) {
		saveTo = cIntSingleResult
		callStr = string(callStr[len(COPSaveRslt):])
	}

	if strings.HasPrefix(callStr, COPSaveRows) {
		saveTo = cIntMultiRowResult
		callStr = string(callStr[len(COPSaveRows):])
	}

	if strings.HasPrefix(callStr, COPFilterRows) {
		saveTo = cIntFilterRows
		callStr = string(callStr[len(COPFilterRows):])
	}

	if strings.HasPrefix(callStr, COPAsRange) {
		saveTo = cIntAsRange
		callStr = string(callStr[len(COPAsRange):])
	}

	funTexts = strings.Split(callStr, cStrColon)

	return
}

func (af ApiFactory) getDBAPIDefine(dbActionName string, apiParamInput string, apiName string) (*ApiInf, error) {
	dbAction, found := af.dbActDfnMap[dbActionName]
	if !found {
		return nil, fmt.Errorf(cStrUnexpectedDBAction, dbActionName)
	}

	switch apiName {
	case CAPIGetMultiRows:
		return newgetMultiRows(apiParamInput, af.dbQry, *dbAction)
	case CAPIRangeGetRows:
		return newRangeGetRows(apiParamInput, af.dbQry, *dbAction)
	default:
		return newdbActionExec(apiParamInput, af.dbQry, *dbAction)
	}

}
