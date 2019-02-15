package eobeapiimpl

import "fmt"

//RemoteEncrypt
//	Define: RemoteEncrypt(ip, port, enc/dec, src)
//		expected Encrypt server: udp only(for alphat),
//		Only accept printable string.
//		Each response should use deliminator: '\n' and also as seperator of method and content.
//		Request format:
//			enc/dec\ncontent\n
//		Further server implementation requirements....
//	Return:
//		1. Map: map[name]values	temporary variables
//		3. Errors
type remoteEncrypt struct {
	apinetbase
	paramCount int
}

//API RemoteEncrypt
const cStrRemoteEncrypt = "RemoteEncrypt"
const cStrRemoteEncryptRetencr = "enc/decSRC_Name" //for example: input: pws for encrypt, return includes a "encpwd"
const cStrOKConfirm = "ok"

//newAPIRemoteEncryptAPI
func newAPIRemoteEncrypt(apiParamInput string) (*ApiInf, error) {
	var api remoteEncrypt
	api.paramCount = 4

	//api.apiParamNames = []string{cStrValidateStrRexParamSrc, cStrValidateStrRexParamReg}
	api.apiRetrnnNames = []string{cStrRemoteEncryptRetencr}
	api.parseParameter(apiParamInput)

	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrRemoteEncrypt, api.paramCount)
	}

	var retIf ApiInf
	retIf = &api

	return &retIf, nil
}

func (api *remoteEncrypt) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, cStrRemoteEncrypt, api.paramCount)
	}

	// ip should support both literal and variable
	ip, found := api.getParamValue(api.apiValueVarInput[0], qryKVMap, preCallRslts)
	if !found {
		ip = api.apiValueVarInput[0]
	}

	// port should support both literal and variable
	port, found := api.getParamValue(api.apiValueVarInput[1], qryKVMap, preCallRslts)
	if !found {
		port = api.apiValueVarInput[1]
	}

	//method support only literal
	method := api.apiValueVarInput[2]

	//target support only variable
	target, found := api.getParamValue(api.apiValueVarInput[3], qryKVMap, preCallRslts)
	if !found {
		return nil, fmt.Errorf(cStrParameterNotFound, api.apiValueVarInput[3])
	}

	reqStr := method + cStrNewLine + target + cStrNewLine

	rsltMap := make(map[string]string)
	resStr, err := api.getResponseFromRemote(ip, port, reqStr)
	if len(resStr) == 0 || err != nil {
		return nil, fmt.Errorf(cStrFetchRespError, target, err)
	}

	retVarName := method + api.getParamName(api.apiValueVarInput[3])
	rsltMap[retVarName] = resStr

	return rsltMap, nil
}
