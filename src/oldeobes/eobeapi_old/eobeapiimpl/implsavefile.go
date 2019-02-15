package eobeapiimpl

import (
	"fmt"
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
func newSaveFile(apiParamInput string) (*ApiInf, error) {
	var api saveFile
	api.paramCount = 2

	api.apiRetrnnNames = []string{cStrSaveFileRetResultFilePath}
	api.parseParameter(apiParamInput)

	if len(api.apiValueVarInput) != api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, CAPISaveFile, api.paramCount)
	}

	//convert to interface.
	var retIf ApiInf
	retIf = &api

	return &retIf, nil
}

func (api *saveFile) RunAPI(qryKVMap map[string]string, preCallRslts map[string]string) (map[string]string, error) {
	//Value[0] filename, value[1] path
	values, err := api.getInputVarValues(qryKVMap, preCallRslts)
	if err != nil {
		return nil, err
	}

	if len(values) < api.paramCount {
		return nil, fmt.Errorf(cStrParameterCountError, CAPISaveFile, api.paramCount)
	}

	if !strings.HasSuffix(values[1], cStrSlash) {
		values[1] = values[1] + cStrSlash
	}

	result := make(map[string]string)
	//Regardless of the byte source, the file name must be user specified.
	fileFullPath := values[1] + values[0]

	//1. Look for the given name
	bytes, found := api.fileMap[values[0]]
	if !found {
		//2. Use a random file in the map if not found
		for _, fileB := range api.fileMap {
			bytes = fileB
			break //just run onece
		}
	}

	//merge two different but same error here
	if len(api.fileMap) == 0 || bytes == nil {
		result[cStrSaveFileRetResultFilePath] = cStrEmpty
		result[cStrRetSuccess] = cStrFalse
		return result, nil //Don't give error if file map is empty, user may be just want to upload nothing.
	}

	err = api.doWriteFile(fileFullPath, bytes)
	if err != nil {
		return nil, err
	}

	result[cStrSaveFileRetResultFilePath] = fileFullPath
	result[cStrRetSuccess] = cStrTure
	return result, nil
}
