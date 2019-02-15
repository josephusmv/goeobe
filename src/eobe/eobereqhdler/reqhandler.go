package eobereqhdler

import (
	"eobe/eobeapi"
	impl "eobe/eobeapi/eobeapiimpl"
	"eobe/eobecltmgmt"
	"eobe/eobedb"
	"eobe/eobehttp"
	"eobe/eobeload"
	"fmt"
	"strings"
)

//RequestHandler implements eobehttp.RespFetcherInf
//	Owns:
//	1. eobecltmgmt.ClientManager instance as a pointer
//	2. All three define maps for data definition
//	3. eobedb.DBQueryInf for query
type RequestHandler struct {
	cm    *eobecltmgmt.ClientManager
	haMap map[string]*eobeload.HTTPActionDefn
	daMap map[string]*eobedb.QueryDefn
	ucMap map[string]string
	dbQry *eobedb.DBQueryInf
}

//NewRequestHandler ...
func NewRequestHandler(cm *eobecltmgmt.ClientManager,
	haMap map[string]*eobeload.HTTPActionDefn,
	daMap map[string]*eobedb.QueryDefn,
	ucMap map[string]string,
	dbQry *eobedb.DBQueryInf) *RequestHandler {
	var rh = RequestHandler{
		cm:    cm,
		haMap: haMap,
		daMap: daMap,
		ucMap: ucMap,
		dbQry: dbQry}

	return &rh
}

func (rh *RequestHandler) parseTarget(fullTarget string) string {
	var target string
	tIdx := strings.LastIndex(fullTarget, cStrSlash)
	if tIdx > 0 && tIdx < len(fullTarget) {
		target = strings.Trim(fullTarget[tIdx+1:], cStrSlash)
	} else {
		target = strings.Trim(fullTarget, cStrSlash) //this is already best efforts... will not works for most times.
	}

	return target
}

func (rh *RequestHandler) evalAcceptFormat(acceptStr string) string {
	if strings.Index(acceptStr, eobehttp.CHTTPTypeStrHTML) >= 0 {
		return eobehttp.CStrExpectedHTML
	}

	if strings.Index(acceptStr, eobehttp.CHTTPTypeStrXML) >= 0 {
		return eobehttp.CStrExpectedXML
	}

	return eobehttp.CStrExpectedJSON
}

//Return error should be BadRequest!~!!
func (rh *RequestHandler) evalCallType(req eobehttp.RequestData, actDefn eobeload.HTTPActionDefn) (apiCallList []string, resName string, err error) {
	//!!!!Caution!!! Make sure APE CallSequence Do NOT return any unnessary errors.
	//	Error are error!!, execute something fail are faile, don;t miss up!!
	var unsupported bool
	switch req.HTTPMethod {
	case "GET":
		apiCallList = actDefn.MethodGET.APICalls
		resName = actDefn.MethodGET.ActualResources
		unsupported = actDefn.MethodGET.UnSupport

	case "POST":
		apiCallList = actDefn.MethodPOST.APICalls
		resName = actDefn.MethodPOST.ActualResources
		unsupported = actDefn.MethodPOST.UnSupport
	case "UPDATE":
		apiCallList = actDefn.MethodUPDATE.APICalls
		resName = actDefn.MethodUPDATE.ActualResources
		unsupported = actDefn.MethodUPDATE.UnSupport
	case "DELETE":
		apiCallList = actDefn.MethodDELETE.APICalls
		resName = actDefn.MethodDELETE.ActualResources
		unsupported = actDefn.MethodDELETE.UnSupport
	default:
		apiCallList = actDefn.MethodGET.APICalls
		resName = actDefn.MethodGET.ActualResources
		unsupported = actDefn.MethodGET.UnSupport
	}

	if unsupported {
		return nil, "", fmt.Errorf(cStrBadRequestMethod, req.HTTPMethod, req.QueryTarget)
	}

	return
}

//FetchResponse implements eobehttp.RespFetcherInf.FetchResponse method
//	responsible for trigger manage all procedures from one front-end client.
func (rh *RequestHandler) FetchResponse(req eobehttp.RequestData) (resp eobehttp.ResponseData, err error) {
	logger := req.Logger

	//Firstly, Get Herald for Client Managment
	hrld := rh.cm.NewHerald()
	err = hrld.NewAccess(req.IP, req.CookieList, nil, req.Logger)
	if err != nil {
		logger.TraceDev(cStrClientManagmentErrorNewAccess, req.QueryTarget, err.Error())
		return resp, fmt.Errorf(cStrBadRequest, req.QueryTarget)
	}

	//Secondly, get resources Define
	target := rh.parseTarget(req.QueryTarget)
	actDefn, found := rh.haMap[target]
	if !found || target == "" {
		//!!This is not an error, any remote inlegal call will lead to this.
		resp.APIErr = impl.NewAPIErrorf(impl.CErrBadParameterError, cStrBadRequest, req.QueryTarget)
		logger.TraceDev(cStrDEVInvalidAction, req.QueryTarget)
		return resp, nil
	}

	/*************/
	//Runn API
	apiList, resName, mthdErr := rh.evalCallType(req, *actDefn)
	if mthdErr != nil { //evalCallType only return error when trying to use undefined method, must be badrequest.
		resp.APIErr = impl.NewAPIErrorf(impl.CErrBadParameterError, mthdErr.Error())
		return resp, nil //return error will lead to 500 error, but this is not!
	}
	apiRslt := rh.doRunAPI(req, apiList, hrld)
	/*************/

	/*************/
	//Make a response
	if apiRslt.ApiErr.HasError() {
		resp = eobehttp.ResponseData{APIErr: apiRslt.ApiErr}
	} else {
		respType := rh.evalAcceptFormat(req.Accept)
		resp, err = doMakeResponse(apiRslt, resName, respType, logger)
	}
	/*************/

	//End work-1: After get a valid response using doFinishResponse
	//	Get an updated coookie list to send to client side
	//	And notify the client management module to finish this access.
	resp.CookieList, err = hrld.GetUpdatedCookieList(req.Logger)
	if err != nil {
		logger.TraceDev(cStrClientManagmentErrorNewAccess, req.QueryTarget, err.Error())
		return resp, fmt.Errorf(cStrBadRequest, req.QueryTarget)
	}

	err = hrld.FinishAccess(req.Logger)
	if err != nil {
		logger.TraceDev(cStrClientManagmentErrorNewAccess, req.QueryTarget, err.Error())
		return resp, fmt.Errorf(cStrBadRequest, req.QueryTarget)
	}

	return resp, err
}

func (rh *RequestHandler) doRunAPI(req eobehttp.RequestData, apiCallList []string, hrld *eobecltmgmt.Herald) eobeapi.CallSeqResult {
	logger := req.Logger
	cs := eobeapi.GetCallSequence(logger)
	cs.InitAPIFactory(rh.ucMap, hrld, *rh.dbQry, rh.daMap)

	cs.SetFileMap(req.UploadedFiles)

	return cs.Execute(req.QueryKeyValueMap, apiCallList)
}
