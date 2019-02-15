package eobeapiimpl

import (
	cltmgrcmd "eobe/eobecltmgmt"
	"eobe/eobehttp"
)

type stepResult struct {
	names    []string
	rsltRows [][]string
}

type ApiInf interface {
	RunAPI(map[string]string, map[string]string) (map[string]string, error)
	GetResultRows([]string, [][]string) ([]string, [][]string)
	SetDataSrc(map[string]string, *cltmgrcmd.Herald)
	SetRangeSource(mNames []string, mrows [][]string)
	SetFilterSource(mNames []string, mrows [][]string)
	SetFileBytes(map[string][]byte)
	SetLogger(logger eobehttp.HttpLoggerInf)
}
