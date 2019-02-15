package eobereqhdler

import (
	"eobe/eobeapi"
	"eobe/eobehttp"
	"fmt"
)

const cFmtStringJSON = `{"singlerow":{%s}, "multirows":{%s}}`
const cFmtStrComma = `,`
const cFmtStringJSONPair = `"%s":"%s"`
const cFmtStrNames = `"names":[%s]`
const cFmtStrRow = `"row%d":[%s]`
const cFmtStrValue = `"%s"`

//JSON response rules:
//	first row is from SingleRow, directly using KV map
//	from second row, all MultiRow will filled...???TODO???.
type respJSONMaker struct {
}

func (rtm respJSONMaker) makeResponse(apiRslt eobeapi.CallSeqResult, actualResName string, logger eobehttp.HttpLoggerInf) (eobehttp.ResponseData, error) {
	var resp eobehttp.ResponseData
	var err error

	logger.TraceDev(cStrDEVDebugResultJSONString, len(apiRslt.SingleRow), len(apiRslt.MNames), len(apiRslt.MultiRow))

	resp.ContentType = cStrAppJSON
	resp.HTMLTmpltName = ""
	//Still carry the Pointer to HTTP package, even not used for now.
	resp.HTMLTmpltData.KVMap = apiRslt.SingleRow
	resp.HTMLTmpltData.Rows = apiRslt.MultiRow

	resp.Body = []byte(rtm.makeJSONString(apiRslt.SingleRow, apiRslt.MNames, apiRslt.MultiRow))
	return resp, err
}

func (rtm respJSONMaker) makeJSONString(SingleRow map[string]string, MNames []string, MultiRow [][]string) string {
	var result string
	var signleJSON, multiJSON string

	if SingleRow != nil && len(SingleRow) > 0 {
		signleJSON = rtm.makeSingleRowJSONString(SingleRow)
	}

	if len(MNames) == 0 && len(MultiRow) == 0 /*|| (len(MultiRow) > 0 && len(MultiRow[0]) != len(MNames))*/ {
		return fmt.Sprintf(cFmtStringJSON, signleJSON, multiJSON)
	}

	if len(MNames) > 0 && len(result) > 0 {
		result += cFmtStrComma
	}

	if len(MNames) > 0 {
		multiJSON = rtm.makeMultiRowJSONString(MNames, MultiRow)
	}

	return fmt.Sprintf(cFmtStringJSON, signleJSON, multiJSON)
}

func (rtm respJSONMaker) makeSingleRowJSONString(SingleRow map[string]string) string {
	if len(SingleRow) <= 0 {
		return ""
	}
	var result string
	total := len(SingleRow)
	for k, v := range SingleRow {
		result += fmt.Sprintf(cFmtStringJSONPair, k, v)
		if total > 1 {
			result += cFmtStrComma
			total--
		}
	}
	return result
}

func (rtm respJSONMaker) makeMultiRowJSONString(MNames []string, MultiRow [][]string) string {
	var nStr, rStr string
	for i, v := range MNames {
		nStr += fmt.Sprintf(cFmtStrValue, v)
		if i < len(MNames)-1 {
			nStr += cFmtStrComma
		}
	}
	nStr = fmt.Sprintf(cFmtStrNames, nStr)

	for i, row := range MultiRow {
		var oneRowStr string
		for j, v := range row {
			oneRowStr += fmt.Sprintf(cFmtStrValue, v)
			if j < len(row)-1 {
				oneRowStr += cFmtStrComma
			}
		}
		rStr += fmt.Sprintf(cFmtStrRow, i, oneRowStr)
		if i < len(MultiRow)-1 {
			rStr += cFmtStrComma
		}
	}

	return nStr + cFmtStrComma + rStr
}
