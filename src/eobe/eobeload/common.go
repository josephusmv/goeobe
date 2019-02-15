package eobeload

import (
	"bytes"
	"strings"
)

type commonBase struct {
	errList []error
}

func (cb commonBase) formatYamlBytes(actionByte []byte, aName string) ([]byte, bool) {
	if len(actionByte) <= 0 {
		return nil, true
	}

	//skip lines with ######## started, e.g.: ######## Sth...
	dbgStr := string(actionByte)
	for !strings.HasPrefix(dbgStr, aName) {
		crlfPos := bytes.Index(actionByte, []byte(cStrNewLine))
		if crlfPos > 0 && crlfPos < len(actionByte)-1 {
			actionByte = actionByte[crlfPos:]
			actionByte = bytes.Trim(actionByte, cStrNewLine)
		} else {
			return nil, true
		}
		dbgStr = string(actionByte)
	}

	actionByte = bytes.Trim(actionByte, cStrNewLine)
	if len(actionByte) <= 0 {
		return nil, true
	}

	if len(actionByte) <= len(aName) {
		return nil, true
	}

	return actionByte, false
}
