package main

import (
	"eobe/eobecore"
	"fmt"
)

func main() {
	ini := eobecore.InitEobe{CfgPath: "./res/config.yaml"}
	err := ini.Init()
	panicError(err)

	ec := eobecore.EobeCore{Ini: &ini}
	wg, err := ec.Start()
	panicError(err)

	go serverRunTCP()

	wg.Wait()
}

func panicError(err error) {
	if err != nil {
		fmt.Println("\n\t\t*********ERROR: " + err.Error() + "*********\n")
		panic(err)
	}
}
