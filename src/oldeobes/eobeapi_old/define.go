package eobeapi

import (
	"eobe/eobeapi/eobeapiimpl"
)

type apiDefine struct {
	apiName        string
	options        int
	implementation eobeapiimpl.ApiInf
}
