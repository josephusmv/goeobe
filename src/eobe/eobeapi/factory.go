package eobeapi

import (
	impl "eobe/eobeapi/eobeapiimpl"
	cltmgrcmd "eobe/eobecltmgmt"
	"eobe/eobedb"
	"eobe/eobehttp"
	"fmt"
	"strings"
)

const cIntMinimalAPIStrLen = 5
const cStrErrAPICallFormat = "Malformat API call stirng: %s"
const cStrErrInvalidActionOption = "The action option is invalid for call: %s"

type apiFactory struct {
	dbQry       eobedb.DBQueryInf
	dbActDfnMap map[string]*eobedb.QueryDefn
	userConsts  map[string]string
	hr          *cltmgrcmd.Herald
}

func (af apiFactory) makeNewAPI(callStr string, logger eobehttp.HttpLoggerInf) (*apiDefine, error) {
	apiDefn, funTexts, errPars := af.parseCallString(callStr, logger)
	if errPars != nil {
		return nil, errPars
	}

	apiName := funTexts[0]
	apiParamInput := funTexts[1]

	api, err := impl.GetAPIImplementation(apiName, apiParamInput)
	if api == nil || err != nil {
		return nil, err
	}

	api.SetDataSrc(af.userConsts, af.hr)
	api.SetLogger(logger)

	if dbActionName, dbAction := api.IsDBAction(); dbAction {
		err = af.initDBAction(api, dbActionName)
	}

	apiDefn.apiName = apiName
	apiDefn.implementation = api

	return apiDefn, err
}

func (af apiFactory) parseCallString(fullCallStr string, logger eobehttp.HttpLoggerInf) (*apiDefine, []string, error) {
	callStr, action := af.getActionOption(fullCallStr)
	if callStr == "" {
		return nil, nil, fmt.Errorf(cStrErrInvalidActionOption, fullCallStr)
	}

	apiDefn := newApiDefine(callStr, action, logger)
	if apiDefn == nil {
		return nil, nil, fmt.Errorf(cStrErrInvalidActionOption, fullCallStr)
	}

	funTexts := strings.Split(callStr, cStrColon)
	if funTexts == nil || len(funTexts) != 2 {
		return nil, nil, fmt.Errorf(cStrErrAPICallFormat, fullCallStr)
	}

	return apiDefn, funTexts, nil
}

func (af apiFactory) getActionOption(fullCallStr string) (callStr string, actionStr string) {
	inLeftBrc := strings.Index(fullCallStr, cStrLeftBrace)
	inRightBrc := strings.Index(fullCallStr, cStrRightBrace)

	if inLeftBrc == inRightBrc && inRightBrc == -1 {
		return fullCallStr, ""
	}

	if inLeftBrc == -1 || inRightBrc == -1 {
		return "", "" //error
	}

	callStr = fullCallStr[inRightBrc+1:]
	actionStr = fullCallStr[inLeftBrc : inRightBrc+1]

	if len(callStr) < cIntMinimalAPIStrLen {
		return "", "" //error
	}

	return
}

func (af apiFactory) initDBAction(api impl.ApiInf, dbActionName string) error {
	dbAction, found := af.dbActDfnMap[dbActionName]
	if !found {
		return fmt.Errorf(cStrUnexpectedDBAction, dbActionName)
	}

	api.SetDBInfo(af.dbQry, *dbAction)
	return nil
}
