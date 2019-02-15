package eobeload

import (
	"fmt"
	"testing"
)

const cTstDBActYaml = "tstdbact.yaml"
const cTstHTTPActYaml = "tsthttpact.yaml"

//TestDBActionLoadSmoke Smoke Test for customer provided where statement
//	go test -v -run DBActionLoadSmoke
func TestDBActionLoadSmoke(t *testing.T) {
	ldr := NewDBActionLoader()
	abAct, err := ldr.LoadFromFile(cTstDBActYaml)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
		return
	}
	dumpDBActionMap(abAct, ldr.errList, "DBActionLoadSmoke results: ")
}

func dumpDBActionMap(abAct *DBActionMap, errList []error, dumptitle string) {
	fmt.Printf("********** %s **********\n", dumptitle)
	for i, v := range errList {
		fmt.Printf("\t\t--->Error[%d]: %s\n", i, v)
	}
	for k, v := range abAct.qryMap {
		fmt.Printf("\t\t--->QueryDefn[%s]: %s\n", k, v)
	}
	fmt.Printf("************************************\n")

}

//TestHTTPActionLoadSmoke Smoke Test for customer provided where statement
//	go test -v -run HTTPActionLoadSmoke
func TestHTTPActionLoadSmoke(t *testing.T) {
	ldr := NewHTTPActionLoader()
	acts, err := ldr.LoadFromFile(cTstHTTPActYaml)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
		return
	}
	dumpHTTPActionMap(acts, ldr.errList, "HTTPActionLoadSmoke results: ")
}

func dumpHTTPActionMap(abAct *HTTPActionMap, errList []error, dumptitle string) {
	fmt.Printf("********** %s **********\n", dumptitle)
	for i, v := range errList {
		fmt.Printf("\t\t--->Error[%d]: %s\n", i, v)
	}
	for k, v := range abAct.actMap {
		fmt.Printf("\t\t--->QueryDefn[%s]: %s\n", k, v)
	}
	fmt.Printf("************************************\n")

}

const cTstUserConstYaml = "userconsts.yaml"

//TestUserConstsLoadSmoke Smoke Test for customer provided where statement
//	go test -v -run UserConstsLoadSmoke
func TestUserConstsLoadSmoke(t *testing.T) {
	ldr := NewUserConstsLoader()
	acts, err := ldr.LoadFromFile(cTstUserConstYaml)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
		return
	}
	dumpUserConstsMap(acts, ldr.errList, "UserConstsLoadSmoke results: ")
}

func dumpUserConstsMap(abAct *UserConstsMap, errList []error, dumptitle string) {
	fmt.Printf("********** %s **********\n", dumptitle)
	for i, v := range errList {
		fmt.Printf("\t\t--->Error[%d]: %s\n", i, v)
	}
	for k, v := range abAct.ucstMap {
		fmt.Printf("\t\t--->User Constant %s = %s\n", k, v)
	}
	fmt.Printf("************************************\n")

}

//TestLoadFactorySmoke Smoke Test for customer provided where statement
//	go test -v -run LoadFactorySmoke
func TestLoadFactorySmoke(t *testing.T) {
	lf := NewLoaderFactory()
	daMap, haMap, uaMap, err := lf.LoadAll(cTstDBActYaml, cTstHTTPActYaml, cTstUserConstYaml)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
		return
	}

	dumpDBActionMap(daMap, nil, "DBActionLoadSmoke results: ")
	dumpHTTPActionMap(haMap, nil, "HTTPActionLoadSmoke results: ")
	dumpUserConstsMap(uaMap, nil, "UserConstsLoadSmoke results: ")

	fmt.Printf("\t**********Errors*************\n")
	for i, v := range lf.GetErrorList() {
		fmt.Printf("\t\t--->Error[%d]: %s\n", i, v)
	}

}
