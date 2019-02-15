package eobeload

import (
	"bytes"
	"eobe/eobedb"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml-2/yaml-2"
)

const cStrDuplicatedDBAction = "duplicated %s DB Action define: %s"

//DBActionLoader eobeload.DBActionLoader
type DBActionLoader struct {
	commonBase
}

//NewDBActionLoader eobeload.NewDBActionLoader
func NewDBActionLoader() *DBActionLoader {
	var dbActLoader DBActionLoader

	return &dbActLoader
}

//LoadFromFile ...
func (loader *DBActionLoader) LoadFromFile(filename string) (*DBActionMap, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf(cStrOpenConfigFileFailed, filename, err.Error())
	}

	var dbActs DBActionMap
	dbActs.qryMap = make(map[string]*eobedb.QueryDefn)

	actBytesArry := bytes.Split(content, []byte(cStrActionSeperateKeyWord))

	for _, actionByte := range actBytesArry {
		var skip bool
		actionByte, skip = loader.formatYamlBytes(actionByte, cStrQryActionName)
		if skip {
			continue
		}

		actDfn, err := loader.loadFromYaml(actionByte)
		if err != nil {
			loader.errList = append(loader.errList, err)
			continue
		}

		if actDfn == nil || actDfn.QueryActionName == "" {
			continue
		}

		if _, found := dbActs.qryMap[actDfn.QueryActionName]; found {
			loader.errList = append(loader.errList, fmt.Errorf(cStrDuplicatedDBAction, cStrDB, actDfn.QueryActionName))
		}

		dbActs.qryMap[actDfn.QueryActionName] = actDfn
	}

	return &dbActs, nil
}

func (loader DBActionLoader) loadFromYaml(data []byte) (*eobedb.QueryDefn, error) {
	var err error
	var qDfn eobedb.QueryDefn

	bytes.Trim(data, cStrLineFeedNewLine)
	if len(data) <= 0 {
		return nil, nil
	}

	err = yaml.Unmarshal([]byte(data), &qDfn)
	if err != nil {
		return nil, fmt.Errorf(cStrLoadDBActionDefineFailed, string(data), err.Error())
	}

	return &qDfn, err
}

//DBActionMap eobedb.QueryDefn storage class.
type DBActionMap struct {
	qryMap map[string]*eobedb.QueryDefn //make this internal to let no outside threads change it.

}

//GetData  Cannot guarantee the lower layer is readonly.
func (dm DBActionMap) GetData() map[string]*eobedb.QueryDefn {
	return dm.qryMap
}
