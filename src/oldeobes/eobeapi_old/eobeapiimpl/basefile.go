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
	//fmt.Println("LOG that a file deleted!") //Change to logger.

	err = ioutil.WriteFile(filepath, bytes, 0666)
	return err
}

func (afb *apiFileBase) SetFileBytes(fileMap map[string][]byte) {
	afb.fileMap = fileMap
	return
}
