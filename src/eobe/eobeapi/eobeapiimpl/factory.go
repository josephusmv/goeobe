package eobeapiimpl

import (
	"fmt"
	"strings"
)

func GetAPIImplementation(apiName string, apiParamInput string) (api ApiInf, err error) {

	switch apiName {
	/**************************/
	case CAPIGetMultiRows:
		detailParams := strings.Split(apiParamInput, cStrMultiRowDBActSepe)
		if len(detailParams) != 2 {
			return nil, fmt.Errorf(cStrErrorGetMultiRowsFormat)
		}
		name := strings.Trim(detailParams[0], cStrSpace)
		paramInput := strings.Trim(detailParams[1], cStrSpace)
		api, err = getDBAPIImplementation(name, paramInput, CAPIGetMultiRows)
	/**************************/
	case CAPIRangeGetRows:
		detailParams := strings.Split(apiParamInput, cStrMultiRowDBActSepe)
		if len(detailParams) != 2 {
			return nil, fmt.Errorf(cStrErrorGetMultiRowsFormat)
		}
		name := strings.Trim(detailParams[0], cStrSpace)
		paramInput := strings.Trim(detailParams[1], cStrSpace)
		api, err = getDBAPIImplementation(name, paramInput, CAPIRangeGetRows)
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
	case CAPISaveFile:
		api, err = newSaveFile(apiParamInput)
	/**************************/
	/*Add New Support APIes here!*/
	default: //DB Action
		api, err = getDBAPIImplementation(apiName, apiParamInput, apiName)
	}

	return api, err
}

//Attention!! Do not delete Error!! it will needed After DB cache introduced and DB refactoried.
func getDBAPIImplementation(dbActionName string, apiParamInput string, apiName string) (ApiInf, error) {
	switch apiName {
	case CAPIGetMultiRows:
		return newgetMultiRows(dbActionName, apiParamInput)
	case CAPIRangeGetRows:
		return newRangeGetRows(dbActionName, apiParamInput)
	default:
		return newdbActionExec(dbActionName, apiParamInput)
	}
}
