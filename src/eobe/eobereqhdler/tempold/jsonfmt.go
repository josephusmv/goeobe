package oldtemp

import (
	"fmt"
)

const cStrFMTJSON = "RESP_JSON"

/*Example row data: "rows":["r1":["value1","value2","value3"], "r2":["value1","value2","value3"]*/
const cStrJSONFMTStringMain = ` {"lastindx":%d,"rowcount":%d,"rows":{%s}}`
const cStrJSONFMTStringRow = `"r%d":[%s]`
const cStrJSONFMTStringRowValueItem = `"%s"`

type jsonFmt struct{}

func (jf jsonFmt) DoAction(rouwCount, lastindx int64, srcRows [][]string) ([]byte, error) {
	var rowsValueStr string
	var resBytes []byte

	if rouwCount != int64(len(srcRows)) {
		return nil, fmt.Errorf(cStrRowCountValueLenNotMatch, rouwCount, len(srcRows))
	}

	for i, v := range srcRows {
		rowStr := jf.makeOneRow(int64(i), v)
		rowsValueStr += rowStr
		if i+1 != len(srcRows) {
			rowsValueStr += cStrCommaSymbol
		}
	}

	resStr := fmt.Sprintf(cStrJSONFMTStringMain, lastindx, rouwCount, rowsValueStr)
	vLen := len(resStr)

	resBytes = make([]byte, vLen)

	copy(resBytes[:], resStr)

	return resBytes, nil
}

func (jf jsonFmt) makeOneRow(rowIndx int64, rowData []string) string {
	var rowValue string

	for i, v := range rowData {
		value := fmt.Sprintf(cStrJSONFMTStringRowValueItem, v)
		rowValue += value
		if i+1 != len(rowData) {
			rowValue += cStrCommaSymbol
		}
	}

	rowStr := fmt.Sprintf(cStrJSONFMTStringRow, rowIndx, rowValue)

	return rowStr
}
