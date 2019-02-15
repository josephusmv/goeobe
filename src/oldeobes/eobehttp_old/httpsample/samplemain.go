package main

import (
	"eobe/eobehttp"
	httpsample3 "eobe/eobehttp/httpsample/fileupld"
	"net/http"
)

const cSampleModuleName = "sample"
const cSampleStrhtmlRoot = "./res"
const cSampleStrIndexPath = cSampleStrhtmlRoot + "/sample/index.tmplt"
const cSampleStrDynamicPath = cSampleStrhtmlRoot + "/sample/dYnAmic.tmplt"
const cSample404Path = cSampleStrhtmlRoot + "/sample/page404.html"
const cSampleStrIndexno404Path = cSampleStrhtmlRoot + "/testno404/index.html"
const cSampleTempIndexActionName = "GetIndex"
const cSampleTempDynamicActionName = "GetDynamicPage"

//RunHttpRequestServer sample codes main entrance, use a main package to call this function
func runRequestServerSample(respFetcher eobehttp.RespFetcherInf) {
	//log factory
	lf := loggerFactoryImpl{}
	//Template files
	tf := []string{cSampleStrIndexPath, cSampleStrDynamicPath, cSample404Path}

	//ctor
	rs := eobehttp.NewRequestServer(cSampleModuleName, cSampleStrhtmlRoot, lf)

	//init
	rs.Init(cSampleTempIndexActionName, tf, respFetcher)

	//run request server
	http.Handle("/", rs)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//Simple example
	rf2 := &httpsample3.RespFetcherImpl{}

	runRequestServerSample(rf2)

}
