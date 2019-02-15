package eobecore

import (
	"eobe/eobehttp"
	"eobe/eobereqhdler"
	"fmt"
	"net/http"
	"sync"
)

//StartEobe ...
type EobeCore struct {
	Ini     *InitEobe
	httpCfg *eobehttp.PkgConfig
}

//Start run the HTTP server and returns a waitgroup to the client to wait.
//	user should call wg.wait after return.
func (ec *EobeCore) Start() (*sync.WaitGroup, error) {
	if ec.Ini == nil {
		return nil, fmt.Errorf(cStrMandatoryInitEobeRequired)
	}

	rsMap := make([]*eobehttp.RequestServer, len(ec.Ini.mdlList)) //[modulename] ReqServer
	failNames := make(map[string]error)
	for i, module := range ec.Ini.mdlList {
		name := module.name
		pRs, err := ec.createRequestServerForModule(name, module)
		if err != nil {
			failNames[name] = err
			continue
		}
		rsMap[i] = pRs
	}

	if len(failNames) > 0 || len(rsMap) <= 0 {
		errSummary := cStrLoadModuleErrorList
		for mdl, err := range failNames {
			errSummary += fmt.Sprintf(cStrLoadModuleErrorItem, mdl, err)
		}
		return nil, fmt.Errorf(errSummary)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go ec.runAllRequestServer(rsMap, &wg)

	return &wg, nil
}

//run in another thread to listen to http, for not blocking caller.
func (ec *EobeCore) createRequestServerForModule(name string, module *httpModule) (*eobehttp.RequestServer, error) {
	pRh := eobereqhdler.NewRequestHandler(ec.Ini.cm,
		module.haMap,
		module.daMap,
		module.ucMap,
		ec.Ini.dbQry)

	var err error
	ec.httpCfg, err = eobehttp.LoadHTTPPackageConfiguration(ec.Ini.GetHTMLInfo())
	if err != nil {
		return nil, err
	}

	rs := eobehttp.NewRequestServer(name, module.rootPath, loggerFactoryImpl{})
	err = rs.Init(pRh, module.indxActionsName)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Load and init Request server: %s done\n", name)
	if name == ec.Ini.rootModule {
		fmt.Printf("\t-->Request server: %s is root module.\n", name)
		rs.SetRootModule()
	}

	return rs, nil
}

func (ec *EobeCore) runAllRequestServer(rsList []*eobehttp.RequestServer, wg *sync.WaitGroup) {
	defer func(wg *sync.WaitGroup) {
		if wg != nil {
			wg.Done()
		}
	}(wg)

	if len(rsList) < 0 {
		fmt.Printf(cStrRequestServerListLengthInvalid)
		panic(cStrRequestServerListLengthInvalid)
	}

	http.Handle(cStrSingleSlash, rsList[0]) //register root once, randomly....

	for _, rs := range rsList {
		name := rs.GetModuleName()
		http.Handle(cStrSingleSlash+name+cStrSingleSlash, rs)
	}

	var wg2 sync.WaitGroup
	hashttp := ec.httpCfg.HasHTTP
	httpPort := ec.httpCfg.HTTPPort
	nohttps := ec.httpCfg.NoHTTPS
	httpsPort := ec.httpCfg.HTTPSPort

	if !nohttps {
		wg2.Add(1)
		go ec.serverTLS(cStrComma+httpsPort, ec.httpCfg.CertFilePath, ec.httpCfg.KeyFilePath, &wg2)
	}

	if hashttp {
		wg2.Add(1)
		go ec.serverHTTP(cStrComma+httpPort, &wg2)
	}

	wg2.Wait()
}

func (ec *EobeCore) serverHTTP(addr string, wg *sync.WaitGroup) {
	defer func(wg *sync.WaitGroup) {
		if wg != nil {
			wg.Done()
		}
	}(wg)
	fmt.Printf("Listen And Serve for HTTP: %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

}

func (ec *EobeCore) serverTLS(addr, certFile, keyFile string, wg *sync.WaitGroup) {
	defer func(wg *sync.WaitGroup) {
		if wg != nil {
			wg.Done()
		}
	}(wg)
	fmt.Printf("Listen And Serve for HTTP over TLS, cert: %s, key: %s, %s\n", certFile, keyFile, addr)
	err := http.ListenAndServeTLS(addr, certFile, keyFile, nil)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

}
