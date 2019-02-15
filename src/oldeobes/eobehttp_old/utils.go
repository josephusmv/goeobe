package eobehttp

import (
	"net/http"
	"strings"
)

func combineFullPath(path, file string) (filePath string) {
	if path[len(path)-1] == '/' {
		filePath = path + file
	} else {
		filePath = path + "/" + file
	}
	return filePath
}

func parseIPPort(r *http.Request) (ipAddr, portNum string) {

	//fmt.Println(r.RemoteAddr)
	ipPort := strings.Split(r.RemoteAddr, ":")
	for i := 0; i < len(ipPort); i++ {
		if i == 0 {
			ipAddr = ipPort[i]
		} else {
			portNum = ipPort[i]
		}
	}

	if ipAddr == "" || strings.HasPrefix(ipAddr, cStrIPAddrLocalhostRemoteStart) {
		ipAddr = cStrIPAddrLocalhost
	}

	return ipAddr, portNum
}
