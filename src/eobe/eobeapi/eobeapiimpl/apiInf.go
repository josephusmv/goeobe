package eobeapiimpl

import (
	cltmgrcmd "eobe/eobecltmgmt"
	"eobe/eobedb"
	"eobe/eobehttp"
)

type stepResult struct {
	names    []string
	rsltRows [][]string
}

//ApiInf implementation provided by this package
type ApiInf interface {
	RunAPI(map[string]string, map[string]string) (map[string]string, *APIError)
	GetResultRows() ([]string, [][]string)
	SetDataSrc(map[string]string, *cltmgrcmd.Herald)
	SetRangeSource(mNames []string, mrows [][]string)
	SetFilterSource(mNames []string, mrows [][]string)
	SetFileBytes(map[string][]byte)
	SetLogger(logger eobehttp.HttpLoggerInf)
	IsDBAction() (string, bool)
	SetDBInfo(dbQry eobedb.DBQueryInf, dbActDefn eobedb.QueryDefn) error
}
