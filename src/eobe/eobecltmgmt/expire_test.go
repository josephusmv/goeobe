package eobecltmgmt

import (
	"fmt"
	"testing"
	"time"
)

type expireTst struct {
	cm      *ClientManager
	curTime time.Time
}

func (et expireTst) changeSystemDate(tm time.Time) {
	//call API to set system date, need administrator permission.
}

func (et *expireTst) changeToTheDay(days int) {
	et.curTime = et.curTime.AddDate(0, 0, days)
	et.changeSystemDate(et.curTime)
	return
}

func (et *expireTst) expireCookieTest() {
	//get today
	now := time.Now()
	et.curTime = now

	et.cm = NewClientManager()
	et.cm.StartClientManagerServer()

	//change to the 3rd day
	et.changeToTheDay(3)

	//restore to real date
	et.changeSystemDate(now)
}

func (et *expireTst) expireSIDTest() {
	//get today
	now := time.Now()
	et.curTime = now

	//change to the 3rd day
	et.changeToTheDay(3)

	//restore to real date
	et.changeSystemDate(now)

}

func (et *expireTst) expireUserTest() {
	//get today
	now := time.Now()
	et.curTime = now

	//change to the 3rd day
	et.changeToTheDay(10)

	//restore to real date
	et.changeSystemDate(now)

}

func (et *expireTst) expireClientIDTest() {
	//get today
	now := time.Now()
	et.curTime = now

	//change to 10 years 6 month and 15 days latter
	et.curTime = et.curTime.AddDate(10, 6, 15)

	//restore to real date
	et.changeSystemDate(now)

}

//TestExpires
//	go test -v -cover -coverprofile cover.out -run Expires
//	go tool cover -html=cover.out -o cover.html
func TestExpires(t *testing.T) {
	var et expireTst

	fmt.Printf("\n********************** Test Cookie Expire ******************\n")
	et.expireCookieTest()
	fmt.Printf("\n********************** Test SID Expire ******************\n")
	et.expireSIDTest()
	fmt.Printf("\n********************** Test User Expire ******************\n")
	et.expireUserTest()
	fmt.Printf("\n********************** Test Client Expire ******************\n")
	et.expireClientIDTest()

}
