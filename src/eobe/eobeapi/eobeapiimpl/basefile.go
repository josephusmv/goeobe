package eobeapiimpl

import (
	"io/ioutil"
	"os"
)

type apiFileBase struct {
	apiBase
	//Specific data
	//init during construct
	fileMap map[string][]byte
}

func (afb *apiFileBase) doWriteFile(filepath string, bytes []byte) (err error) {
	err = os.Remove(filepath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	err = ioutil.WriteFile(filepath, bytes, 0666)
	return err
}

func (afb *apiFileBase) SetFileBytes(fileMap map[string][]byte) {
	afb.fileMap = fileMap
	return
}
