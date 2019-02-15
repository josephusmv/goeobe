package eobeload

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

//UserConstsLoader eobeload.UserConstsLoader
type UserConstsLoader struct {
	commonBase
}

const cStrDuplicatedUserConst = "duplicated %s User Const: %s"

//NewUserConstsLoader eobeload.NewUserConstsLoader
func NewUserConstsLoader() *UserConstsLoader {
	var usrCnstLoader UserConstsLoader

	return &usrCnstLoader
}

//LoadFromFile ...
func (loader *UserConstsLoader) LoadFromFile(filename string) (*UserConstsMap, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf(cStrOpenConfigFileFailed, filename, err.Error())
	}

	var usrCnsts UserConstsMap
	usrCnsts.ucstMap = make(map[string]string)

	bytesArry := bytes.Split(content, []byte(cStrNewLine))
	for _, actionByte := range bytesArry {
		line := string(actionByte)
		line = strings.Trim(line, cStrLineFeedNewLine)
		if len(line) <= 0 {
			continue
		}

		if line[0] == cCharSharp {
			continue
		}

		scPos := strings.Index(line, cStrSemiColon)
		if scPos <= 0 || scPos >= len(line)-1 {
			loader.errList = append(loader.errList, fmt.Errorf(cStrInvalidUconstLine, line))
			continue
		}

		key := line[0:scPos]
		key = strings.Trim(key, cStrSemiColon+cStrSpace+cStrTab)
		value := line[scPos:]
		value = strings.Trim(value, cStrSemiColon+cStrSpace+cStrTab+cStr2QuotaMark)

		if _, found := usrCnsts.ucstMap[key]; found {
			loader.errList = append(loader.errList, fmt.Errorf(cStrDuplicatedUserConst, cStrUserConst, key))
		}

		usrCnsts.ucstMap[key] = value
	}

	return &usrCnsts, nil
}

//UserConstsMap  ....
type UserConstsMap struct {
	ucstMap map[string]string //make this internal to let no outside threads change it.
}

//GetData  Cannot guarantee the lower layer is readonly.
func (um UserConstsMap) GetData() map[string]string {
	return um.ucstMap
}
