package eobedb

import (
	"fmt"
)

const cStrBindVarSymbol = "?"

type baseActor struct {
}

func (ba baseActor) makeWhere(colNames, logcOps, mathOps []string) (string, error) {
	var stmt string

	if colNames == nil || len(colNames) == 0 || mathOps == nil || len(mathOps) == 0 {
		return "", nil
	}

	if len(colNames) != len(logcOps)+1 || len(colNames) != len(mathOps) {
		return "", fmt.Errorf(cstrWhereParamError)
	}

	stmt += cStrDBWhere

	for ind, colName := range colNames {
		stmt += colName + ba.convertOpToSymbol(mathOps[ind]) + cStrBindVarSymbol + " "
		if ind < len(logcOps) {
			stmt += logcOps[ind] + " "
		}
	}

	return stmt, nil
}

func (ba baseActor) convertOpToSymbol(operator string) string {
	switch operator {
	case CstrDBOptEqualTo:
		return "="
	case CstrDBOptNotEqualTo:
		return "!="
	case CstrDBOptLike:
		return " like "
	case CstrDBOptGreaterThan:
		return ">"
	case CstrDBOptSmallerThan:
		return "<"
	default:
		return "="
	}
}

func (qa baseActor) bindParameters(params []string, expects []string) []interface{} {
	lenPara := len(params)
	lenExpt := len(expects)
	//Interface slice to carry the string pointers
	dest := make([]interface{}, lenPara+lenExpt)
	for i := range params {
		dest[i] = &params[i] // Put pointers to each string in the interface slice
	}
	for i := range expects {
		dest[i+lenPara] = &expects[i]
	}

	return dest
}
