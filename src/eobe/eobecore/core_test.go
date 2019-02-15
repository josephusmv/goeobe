package eobecore

import (
	"fmt"
	"testing"
)

const cStrTestConfigFile = "./unittest/config.yaml"

//*********************************************************************
//***********************TEst cases
//TestSmokeTest SmokeTest
//	go test -v -cover -covermode=set -coverprofile cover.out -run SmokeTest
//	This is a manually test case need user to check HTML output.
//	The automatically case is located on ../../tests
func TestSmokeTest(t *testing.T) {
	ini := InitEobe{CfgPath: cStrTestConfigFile}
	err := ini.Init()
	panicError(err)

	ec := EobeCore{Ini: &ini}
	wg, err := ec.Start()
	panicError(err)
	wg.Wait()
}

func panicError(err error) {
	if err != nil {
		fmt.Println("\n\t\t*********ERROR: " + err.Error() + "*********\n")
		panic(err)
	}
}
