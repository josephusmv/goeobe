package eobeapiimpl

import (
	"os"
	"strings"
)

//saveFile
//	Define: SaveFile: filename, rootpath
//	Desc: Save a single file. if multiple files provided, look for the given name,
//			If not existed, then use the RANDOM one..
//			So user should make sure the file name is correct or only one file in the request.
//			All parameter must be variables
//	Return:
//		1. retResultFilePath: Full File Path with file name included.
//		2. Errors
type saveFile struct {
	apiFileBase
	//Specific data
	paramCount int
}

//CAPISaveFile API Name
const CAPISaveFile = "SaveFile"
const cStrSaveFileRetResultFilePath = "retResultFilePath"

//newAPIsaveFile API saveFile(paramSrcStr, paramRegexpStr) retIsValid
func newSaveFile(apiParamInput string) (ApiInf, error) {
	var api saveFile
	api.paramCount = 2

	api.apiRetrnnNames = []string{cStrSaveFileRetResultFilePath}

	return &api, api.parseParameter(apiParamInput, api.paramCount, CAPISaveFile)
}

const cStrSaveFileAPIParameterErr = "SaveFile API error: %s"

func (api *saveFile) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, *APIError) {
	//Initiate result map for error returns and set initial values, return non nil for CErrBadParameterError
	result := make(map[string]string)
	result[cStrSaveFileRetResultFilePath] = cStrEmpty
	result[cStrRetSuccess] = cStrFalse

	//Parse to get file srce name and dst local storage folder name
	values, err := api.getInputVarValues(qryKVMap, preCallRslts, api.paramCount, CAPIFilterMultiRowss)
	if err != nil {
		return nil, NewAPIErrorf(CErrBadCallError, cStrSaveFileAPIParameterErr, err.Error())
	}

	//Empty uploaded, must be bad request or has error in HTTP package
	if api.fileMap == nil || len(api.fileMap) == 0 {
		return result, NewAPIErrorf(CErrBadParameterError, cStrFileBytesAreNil)
	}

	if !strings.HasSuffix(values[1], cStrSlash) {
		values[1] = values[1] + cStrSlash
	}
	storePath := values[1] + cStrSlash
	os.MkdirAll(storePath, os.ModePerm)

	fileFullPath := storePath + values[0]

	//Found the specified file
	//If file name not found in update list, it must be bad request..
	bytes, found := api.fileMap[values[0]]
	if !found || bytes == nil {
		return result, NewAPIErrorf(CErrBadParameterError, cStrFileNameNotFoundinUploadListError, values[0])
	}

	err = api.doWriteFile(fileFullPath, bytes)
	if err != nil {
		return nil, NewAPIErrorf(CErrServerInternalError, cStrSaveFileAPIParameterErr, err.Error())
	}

	result[cStrSaveFileRetResultFilePath] = fileFullPath
	result[cStrRetSuccess] = cStrTure
	return result, ApiSuccess()
}
