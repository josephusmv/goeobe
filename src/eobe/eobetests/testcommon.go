package eobetests

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

/***********************************TIME DATE***********************************/
const cStrOSWindows = "windows"

const cStrFmtMACOS = "%02d%02d%04d" //date 01252018
//const cStrFmtWindows = "%04d-%02d-%02d"
const cStrFmtWindows = "%02d-%02d-%04d"
const layout = "2006-01-02"

func setOSDateAsStr(date string) {
	t, err := time.Parse(layout, date)
	if err != nil {
		panic(err.Error())
	}
	setOSDate(t)
}

func setOSDate(t time.Time) {
	var err error
	y, m, d := t.Date()
	month := int(m)
	if runtime.GOOS == cStrOSWindows {
		//str := fmt.Sprintf(cStrFmtWindows, y, month, d)
		str := fmt.Sprintf(cStrFmtWindows, month, d, y)
		fmt.Println("-----------> Time Change To:" + str)

		//Need Administrator permission to do so...
		err = exec.Command("cmd", "/C", "date "+str).Run()
	} else {
		str := fmt.Sprintf(cStrFmtMACOS, month, d, y)
		fmt.Println("-----------> Time Change To:" + str)

		//Need Administrator permission to do so...
		err = exec.Command("/bin/sh", "-c", "sudo date "+str).Run()
	}
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

/***********************************DB Connections***********************************/
const cStrDBConnMAC = "root:123456@/THDATABASE"
const cStrDBConnWIN = "root:3edc$RFV@/THDATABASE"
const cStrDBConn = cStrDBConnMAC
