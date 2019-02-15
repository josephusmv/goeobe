package eobeapiimpl

import (
	"eobe/eobekiosk"
	"fmt"
	"testing"
	"time"
)

const CONN_HOST = "127.0.0.1"

//run full coverage test:
//		go test -cover -coverprofile cover.out
//		go tool cover -html=cover.out -o cover.html

//TestNetGetResponseFromRemote test connect to remote Server
//	go test -v  -run NetGetResponseFromRemote
func TestNetGetResponseFromRemote(t *testing.T) {
	const (
		CONN_TYPE  = "tcp"
		LOG_DIR    = "./"
		LOG_DIRbad = "./notexist"
	)
	const CONN_PORT = "8081"
	go runKioskServer(CONN_PORT)
	time.Sleep(time.Millisecond * 200)

	var anb apinetbase
	origin := "12345678abcdefg"
	testContentStr := "enc\n" + origin
	resp, err := anb.getResponseFromRemote(CONN_HOST, CONN_PORT, testContentStr)
	checkError(err, "getResponseFromRemote")
	fmt.Println(resp)

	testContentStr = "dec\n" + resp
	resp, err = anb.getResponseFromRemote(CONN_HOST, CONN_PORT, testContentStr)
	checkError(err, "getResponseFromRemote")
	fmt.Println(resp)

	if resp != origin {
		panic("Error result!!")
	}

	resp, err = anb.getResponseFromRemote(CONN_HOST, CONN_PORT, "fin")
	fmt.Println(resp)
	if resp != "ack" {
		panic("Error result!!")
	}
	time.Sleep(time.Millisecond * 200)
}

func runKioskServer(port string) {
	const (
		CONN_TYPE = "tcp"
		LOG_DIR   = "./"
		KEY_FILE  = "./key"
	)

	es := eobekiosk.NewDemoEncServer(CONN_HOST, port, CONN_TYPE, LOG_DIR)
	es.InitCrypto(eobekiosk.CStrAlgoAES, KEY_FILE)

	es.Run()

}

func checkError(err error, contextStr string) {
	if err != nil {
		fmt.Println(contextStr, "Error: ")
		fmt.Println(err.Error())
		panic(err)
	}
}

//TestNetRunRemoteEncAPI
//	go test -v  -run NetRunRemoteEncAPI
func TestNetRunRemoteEncAPI(t *testing.T) {
	go runKioskServer("9091")
	time.Sleep(time.Millisecond * 200)

	const pwdsrc = "ThisIsPWD123"
	const apiName = cStrRemoteEncrypt
	apiParamInput := fmt.Sprintf("%s, %s, %s, ?password", CONN_HOST, "9091", "enc")
	qryKVMap := make(map[string]string)
	qryKVMap["password"] = pwdsrc

	//encrypt
	api, err := GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")
	result, apiErr := api.RunAPI(qryKVMap, nil)
	checkAPIError(apiErr, "RunAPI() "+apiName+" error.")
	fmt.Println(result["encpassword"])

	//Decrypt
	apiParamInput = fmt.Sprintf("%s, %s, %s, ^encpassword", CONN_HOST, "9091", "dec")
	api, err = GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")
	result, apiErr = api.RunAPI(nil, result)
	checkAPIError(apiErr, "RunAPI() "+apiName+" error.")
	fmt.Println(result["decencpassword"])

	if result["decencpassword"] != pwdsrc {
		panic("got decrypt result: " + result["decencpassword"])
	}

	var anb apinetbase
	resp, err := anb.getResponseFromRemote(CONN_HOST, "9091", "fin")
	fmt.Println(resp)
	if resp != "ack" {
		panic("Error result!!")
	}

	//some error tests below
	apiParamInput = fmt.Sprintf("%s, %s, %s, ?notExist", CONN_HOST, "9091", "enc")
	api, err = GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")
	result, apiErr = api.RunAPI(nil, result)
	checkAPIErrorType(apiErr, CErrBadCallError)

	apiParamInput = fmt.Sprintf("%s, %s, %s", CONN_HOST, "9091", "enc")
	api, err = GetAPIImplementation(apiName, apiParamInput)
	if err == nil {
		panic("Leak parameter not match error")
	}

	apiParamInput = fmt.Sprintf("%s, %s, %s, ?password", CONN_HOST, "9092", "enc")
	api, err = GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")
	result, apiErr = api.RunAPI(qryKVMap, nil)
	checkAPIErrorType(apiErr, CErrServerInternalError)
}
