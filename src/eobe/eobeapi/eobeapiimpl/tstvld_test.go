package eobeapiimpl

import (
	"fmt"
	"testing"
)

//TestValidateInt test connect to remote Server
//	go test -v  -run ValidateInt
func TestValidateInt(t *testing.T) {
	const apiName = cStrValidateInt
	var apiParamInput = "^intforvalidate, gt, $ONE_HUNDRED, $FORBIDEN_MSG"

	precalrelts := make(map[string]string)
	ucStr := make(map[string]string)
	ucStr["FORBIDEN_MSG"] = "\"Failed Int Validation.\""

	api, err := GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")

	//smoke test
	precalrelts["intforvalidate"] = "190"
	ucStr["ONE_HUNDRED"] = "100"
	api.SetDataSrc(ucStr, nil)
	result, apiErr := api.RunAPI(nil, precalrelts)
	checkAPIError(apiErr, "RunAPI() "+apiName+" error.")
	fmt.Println(result)

	//validation failed
	precalrelts["intforvalidate"] = "90"
	ucStr["ONE_HUNDRED"] = "100"
	api.SetDataSrc(ucStr, nil)
	result, apiErr = api.RunAPI(nil, precalrelts)
	checkAPIErrorType(apiErr, CErrRunValidateFailure)
	fmt.Println(result)

	//invalid src int
	precalrelts["intforvalidate"] = "aStr"
	result, apiErr = api.RunAPI(nil, precalrelts)
	checkAPIErrorType(apiErr, CErrBadCallError)
	fmt.Println(result)

	//invalid dst int
	precalrelts["intforvalidate"] = "90"
	ucStr["ONE_HUNDRED"] = "aStr"
	result, apiErr = api.RunAPI(nil, precalrelts)
	checkAPIErrorType(apiErr, CErrBadCallError)
	fmt.Println(result)

	//invalid compare keyword
	apiParamInput = "^intforvalidate, INVALID, $ONE_HUNDRED, $FORBIDEN_MSG"
	api, err = GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")
	precalrelts["intforvalidate"] = "190"
	ucStr["ONE_HUNDRED"] = "100"
	api.SetDataSrc(ucStr, nil)
	result, apiErr = api.RunAPI(nil, precalrelts)
	checkAPIErrorType(apiErr, CErrBadCallError)
	fmt.Println(result)

	//nonexisted - src
	apiParamInput = "^notexstedSrc, lt, $ONE_HUNDRED, $FORBIDEN_MSG"
	api, err = GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")
	api.SetDataSrc(ucStr, nil)
	result, apiErr = api.RunAPI(nil, precalrelts)
	checkAPIErrorType(apiErr, CErrBadCallError)
	fmt.Println(result)

	//nonexisted - dst
	apiParamInput = "^intforvalidate, lt, $notexstedSrc, $FORBIDEN_MSG"
	api, err = GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")
	api.SetDataSrc(ucStr, nil)
	result, apiErr = api.RunAPI(nil, precalrelts)
	checkAPIErrorType(apiErr, CErrBadCallError)
	fmt.Println(result)
}
