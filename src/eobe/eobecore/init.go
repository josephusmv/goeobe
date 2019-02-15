package eobecore

import (
	"eobe/eobecltmgmt"
	"eobe/eobedb"
	"fmt"
)

type InitEobe struct {
	CfgPath string
	cfgHld  CfgHolder
	//DB
	dbQry *eobedb.DBQueryInf
	dbSrv *eobedb.DBServerInf
	// client manager
	cm *eobecltmgmt.ClientManager
	//module map
	mdlList    []*httpModule
	rootModule string
}

//Init  all the global data required.
func (ini *InitEobe) Init() error {
	err := ini.readCfg()
	if err != nil {
		return err
	}

	err = ini.initDB()
	if err != nil {
		return err
	}

	ini.initClientManager()

	err = ini.loadAllModule()
	if err != nil {
		return err
	}
	return nil
}

func (ini *InitEobe) readCfg() error {
	return ini.cfgHld.LoadConfiguration(ini.CfgPath)
}

func (ini *InitEobe) initDB() error {
	logDevLevel, logFolder := ini.cfgHld.GetLogInfo()
	dbQry, dbSrv := eobedb.RunNewDBServer(ini.cfgHld.GetDBDescriptions())
	dbSrv.SetDBOptions(eobedb.OptionEnableDEVLog, logDevLevel)
	dbSrv.SetDBOptions(eobedb.OptionDBLogRoot, logFolder)
	dbSrv.SetDBOptions(eobedb.OptionDBMaxConcurrencyInt, 80)
	err := dbSrv.Init()
	if err != nil {
		return err
	}
	ini.dbQry = &dbQry
	ini.dbSrv = &dbSrv

	return nil
}

func (ini *InitEobe) initClientManager() {
	enableDEV, logFolder := ini.cfgHld.GetLogInfo()
	if logFolder != "" {
		vStrLogFilePath = logFolder + cStrSingleSlash
	}
	vEnableDevLogs = enableDEV

	eobecltmgmt.DbgEmergencyLogFlag = enableDEV
	ini.cm = eobecltmgmt.NewClientManager(vStrLogFilePath)
	ini.cm.StartClientManagerServer()
}

func (ini *InitEobe) loadAllModule() error {
	module, indActs, rootFolder, rootModule := ini.cfgHld.GetModuleInfo()
	if module == nil || indActs == nil || rootFolder == "" || len(indActs) != len(module) {
		return fmt.Errorf(cStrModuleListInvalid)
	}

	ini.rootModule = rootModule

	ini.mdlList = make([]*httpModule, len(module))
	failNames := make(map[string]error)
	for i, name := range module {
		var err error
		ini.mdlList[i], err = ini.loadModule(name, rootFolder)
		if err != nil {
			failNames[name] = err
			continue
		}
		ini.mdlList[i].indxActionsName = indActs[i]
	}

	if len(failNames) > 0 {
		errSummary := cStrLoadModuleErrorList
		for mdl, err := range failNames {
			errSummary += fmt.Sprintf(cStrLoadModuleErrorItem, mdl, err)
		}
		return fmt.Errorf(errSummary)
	}

	return nil
}

func (ini *InitEobe) loadModule(name, rootPath string) (*httpModule, error) {
	var hm = httpModule{name: name, rootPath: rootPath}
	err := hm.loadAll(cStrHTTPActionFile, cStrDBActionFile, cStrUserConstFile)
	if err != nil {
		return nil, err
	}
	return &hm, nil
}

//GetTemplateNames return the HTML Template file name list
func (ini *InitEobe) GetHTMLInfo() string {
	return ini.cfgHld.GetHTMLInfo()
}
