package eobereqhdler

import (
	"eobe/eobeapi"
	"eobe/eobehttp"
)

type respHTMLMaker struct {
}

func (rtm respHTMLMaker) makeResponse(apiRslt eobeapi.CallSeqResult, actualResName string, logger eobehttp.HttpLoggerInf) (eobehttp.ResponseData, error) {
	var resp eobehttp.ResponseData
	var err error

	resp.Body = nil
	resp.ContentType = cStrTextHTML
	resp.HTMLTmpltName = actualResName
	resp.HTMLTmpltData.KVMap = apiRslt.SingleRow
	resp.HTMLTmpltData.Rows = apiRslt.MultiRow

	return resp, err
}
