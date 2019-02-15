package eobeapiimpl

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"testing"
)

func testSetFile(api ApiInf) {
	const testFilePath1 = "statics.go"
	bytes1, err := ioutil.ReadFile(testFilePath1)
	if err != nil {
		panic("Open file " + testFilePath1 + "failed.")
	}

	const testFilePath2 = "factory.go"
	var bytes2 []byte
	bytes2, err = ioutil.ReadFile(testFilePath2)
	if err != nil {
		panic("Open file " + testFilePath2 + "failed.")
	}

	bytesMap := make(map[string][]byte)
	bytesMap[testFilePath1] = bytes1
	bytesMap[testFilePath2] = bytes2
	api.SetFileBytes(bytesMap)
}

//TestMiscSaveFile
//	go test -v  -run MiscSaveFile
func TestMiscSaveFile(t *testing.T) {
	const saveFilePath = "./uploaded"
	const apiName = CAPISaveFile
	apiParamInput := "?FileName,  ?LOCALStorePATH"
	qryKVMap := make(map[string]string)
	qryKVMap["FileName"] = "statics.go"
	qryKVMap["LOCALStorePATH"] = saveFilePath

	api, err := GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")

	// Call api.SetFileBytes(bytesMap)
	testSetFile(api)

	//smoke test
	result, apiErr := api.RunAPI(qryKVMap, nil)
	checkAPIError(apiErr, "RunAPI() "+apiName+" error.")
	fmt.Println(result)

	//*********************************************************
	//Error scenarios
	//cause an write file error by open the target file.
	if runtime.GOOS == "windows" {
		file, err := os.Open(saveFilePath + "/" + "statics.go")
		if err != nil {
			panic("TEST facility error: " + err.Error())
		}
		result, apiErr = api.RunAPI(qryKVMap, nil)
		checkAPIErrorType(apiErr, CErrServerInternalError)
		file.Close()
	} else {
		qryKVMap["LOCALStorePATH"] = "/" //linux cannot access root, this will make dowrite fails
		result, apiErr = api.RunAPI(qryKVMap, nil)
		checkAPIErrorType(apiErr, CErrServerInternalError)
	}

	//non existed file in bytes -  save fi
	qryKVMap["FileName"] = "notexisted.go"
	result, apiErr = api.RunAPI(qryKVMap, nil)
	checkAPIErrorType(apiErr, CErrBadParameterError)
	fmt.Println(result)

	//Set File Map to nil
	api.SetFileBytes(nil)
	result, apiErr = api.RunAPI(qryKVMap, nil)
	checkAPIErrorType(apiErr, CErrBadParameterError)
	fmt.Println(result)

	//input has non-existed values
	apiParamInput = "?NON-Exsted,  ?LOCALStorePATH"
	api, err = GetAPIImplementation(apiName, apiParamInput)
	checkError(err, "GetAPIImplementation() error.")
	testSetFile(api)
	result, apiErr = api.RunAPI(qryKVMap, nil)
	checkAPIErrorType(apiErr, CErrBadCallError)
	fmt.Println(result)
}
