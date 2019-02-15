package eobeapiimpl

//getCurrentUser
//	Define: GetCurrentUser()
//		will get an empty string if not login
//	Return:
//		1. Map: map[retCurrentUserName]UserName
//		2. Errors
type getCurrentUser struct {
	apiBase
	//Specific data
	paramCount int
}

//API GetCurrentUser
const cStrGetCurrentUser = "GetCurrentUser"
const cStrGetCurrentUserRetUserName = "retCurrentUserName"

//newGetCurrentUser API getCurrentUser(paramSrcStr, paramRegexpStr) retIsValid
func newGetCurrentUser(apiParamInput string) (ApiInf, error) {
	var api getCurrentUser
	api.paramCount = 0

	//api.apiParamNames = []string{cStrValidateStrRexParamSrc, cStrValidateStrRexParamReg}
	api.apiRetrnnNames = []string{cStrGetCurrentUserRetUserName}

	return &api, api.parseParameter(apiParamInput, api.paramCount, cStrGetCurrentUser)
}

func (api *getCurrentUser) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, *APIError) {
	uName := api.getCurrentUser()

	rsltMap := make(map[string]string)
	rsltMap[cStrGetCurrentUserRetUserName] = uName

	return rsltMap, ApiSuccess()
}
