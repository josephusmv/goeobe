package oldtemp

import (
	"fmt"
	"strconv"
	"strings"
)

type KeyValueMap map[string]string

//ActionParser initiated with an ModuleActionDef, do the Parameter Parse and result organization.
type ActionParser struct {
	ActionDef      *ModuleActionDef
	srchParamKey   []string
	srchParamValue []string
	updtParamKey   []string
	updtParamValue []string
}

//NewActionParser Get a new ActionParser. if QryKVMap is not Nil, then do the parse accordingly.
func NewActionParser(actionDefn *ModuleActionDef, QryKVMap KeyValueMap) (*ActionParser, error) {
	var err error
	ap := ActionParser{ActionDef: actionDefn}
	if QryKVMap != nil {
		err = ap.ParseParameterValues(QryKVMap)
	}

	if err != nil {
		return nil, err
	}

	return &ap, nil
}

//ParseParameterValues Parse qurey values into actual data query parameters
//	1. Read/Delete type actions has only condition parameters(srchParam)
//  2. Add type actions has only update parameters(updtParam)
//	3. Update type actions has both condition and update parameters(srchParam, updtParam)
//	Error will return the actual error, not log procedure required.
func (ap *ActionParser) ParseParameterValues(QryKVMap KeyValueMap) (err error) {
	if ap.ActionDef == nil {
		return fmt.Errorf(cStrActionDefineNotSetForParamParse)
	}

	qryType := strings.ToUpper(ap.ActionDef.QUERYTYPE)
	var srchParamMap, updtParamMap KeyValueMap

	switch qryType {
	case cStrKWSelect:
		fallthrough
	case cStrKWDelete:
		err = ap.parseSearch(QryKVMap)

	case cStrKWInsert:
		err = ap.parseUpdate(updtParamMap)

	case cStrKWUpdate:
		srchParamMap, updtParamMap = ap.catagoryParams(QryKVMap)
		err = ap.parseSearch(srchParamMap)
		err = ap.parseUpdate(updtParamMap)

	default:
		return fmt.Errorf(cStrInvalidQueryTypeForParamParse, ap.ActionDef.QUERYTYPE)
	}

	return
}

func (ap ActionParser) catagoryParams(paramMap KeyValueMap) (srchParamMap, updtParamMap KeyValueMap) {
	srchParamMap = make(KeyValueMap)
	updtParamMap = make(KeyValueMap)

	for k, v := range paramMap {
		if strings.HasSuffix(k, cStrKWSrch) {
			srchParamMap[k] = v
		} else if strings.HasSuffix(k, cStrKWUpdt) {
			updtParamMap[k] = v
		} else {
			continue
		}
	}

	return
}

//ap.ActionDef.PARAMETER may have duplicate names.
//ap.srchParamKey has only the real name user defined for DB query
func (ap *ActionParser) parseSearch(paramMap KeyValueMap) (err error) {
	paramLen := len(ap.ActionDef.PARAMETER)
	ap.srchParamKey = make([]string, paramLen)
	ap.srchParamValue = make([]string, paramLen)

	return ap.parseKeyValues(paramMap, ap.ActionDef.PARAMETER, ap.srchParamKey, ap.srchParamValue)
}

//ap.ActionDef.ExpectedFields may have duplicate names.(username___1)
//ap.updtParamKey has only the real name user defined for DB query.(username)
func (ap *ActionParser) parseUpdate(paramMap KeyValueMap) (err error) {
	paramLen := len(ap.ActionDef.ExpectedFields)
	ap.updtParamKey = make([]string, paramLen)
	ap.updtParamValue = make([]string, paramLen)

	return ap.parseKeyValues(paramMap, ap.ActionDef.ExpectedFields, ap.updtParamKey, ap.updtParamValue)
}

func (ap ActionParser) parseKeyValues(paramMap KeyValueMap, defnParams, keys, values []string) (err error) {

	//srchParam is not key indexed, it's only follow the order of ap.ActionDef.PARAMETER defined.
	for i, param := range defnParams {
		k, _ := ap.trimToRealname(param)
		v, found := paramMap[param]
		if found {
			keys[i] = k
			values[i] = v
		} else {
			err = fmt.Errorf(cStrNotEnoughParameterError, param)
			break
		}
	}
	return
}

func (ap ActionParser) trimToRealname(tusKey string) (result string, num int) {
	var err error

	if strings.HasSuffix(tusKey, cStrKWSrch) {
		sufx := strings.Index(tusKey, cStrKWSrch)
		tusKey = tusKey[:sufx]
	}

	if strings.HasSuffix(tusKey, cStrKWUpdt) {
		sufx := strings.Index(tusKey, cStrKWUpdt)
		tusKey = tusKey[:sufx]
	}

	pos := strings.Index(tusKey, cStrTrippleUnderScore)
	if pos <= 0 || pos+3 >= len(tusKey) {
		return tusKey, 0
	}

	num, err = strconv.Atoi(tusKey[pos+3:])
	if err != nil {
		num = -1
	}

	return tusKey[:pos], num
}

//GetSrchParamKeys ...
func (ap ActionParser) GetSrchParamKeys() []string {
	return ap.srchParamKey
}

//GetSrchParamValues ...
func (ap ActionParser) GetSrchParamValues() []string {
	return ap.srchParamValue
}

//GetUpdtParamKeys ...
func (ap ActionParser) GetUpdtParamKeys() []string {
	return ap.updtParamKey
}

//GetUpdtParamValues ...
func (ap ActionParser) GetUpdtParamValues() []string {
	return ap.updtParamValue
}
