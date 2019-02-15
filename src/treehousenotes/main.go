package main

import (
	"eobe/eobecore"
	"eobe/eobekiosk"
	"fmt"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "8081"
	CONN_TYPE = "tcp"
	LOG_DIR   = "./ALLLOGS/"
)

func main() {
	ini := eobecore.InitEobe{CfgPath: "./res/config.yaml"}
	err := ini.Init()
	panicError(err)

	ec := eobecore.EobeCore{Ini: &ini}
	wg, err := ec.Start()
	panicError(err)

	es := eobekiosk.NewDemoEncServer(CONN_HOST, CONN_PORT, CONN_TYPE, LOG_DIR)
	if es == nil {
		panic("init enc server failed")
	}
	es.InitCrypto(eobekiosk.CStrAlgoAES)

	go es.Run()

	wg.Wait()
}

func panicError(err error) {
	if err != nil {
		fmt.Println("\n\t\t*********ERROR: " + err.Error() + "*********\n")
		panic(err)
	}
}
