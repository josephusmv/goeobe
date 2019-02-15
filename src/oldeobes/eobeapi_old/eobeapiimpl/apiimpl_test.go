package eobeapiimpl

import (
	"fmt"
	"testing"
)

const (
	CONN_HOST  = "127.0.0.1"
	CONN_PORT  = "8081"
	CONN_TYPE  = "tcp"
	LOG_DIR    = "./"
	LOG_DIRbad = "./notexist"
)

//TestGetResponseFromRemote test connect to remote Server
//	go test -v  -run GetResponseFromRemote
//	go tool cover -html=cover.out -o cover.html
func TestGetResponseFromRemote(t *testing.T) {
	var anb apinetbase
	origin := "12345678abcdefg\n"
	testContentStr := "enc\n" + origin
	resp, err := anb.getResponseFromRemote(CONN_HOST, CONN_PORT, testContentStr)
	checkError(err, "getResponseFromRemote")
	fmt.Println(resp)

	testContentStr = "dec\n" + resp
	resp, err = anb.getResponseFromRemote(CONN_HOST, CONN_PORT, testContentStr)
	checkError(err, "getResponseFromRemote")
	fmt.Println(resp)

	if resp != origin {
		panic("Error result!!")
	}
}

func checkError(err error, contextStr string) {
	if err != nil {
		fmt.Println(contextStr, "Error: ")
		fmt.Println(err.Error())
		panic(err)
	}
}
