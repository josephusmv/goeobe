package eobedb

import (
	"database/sql"
	"fmt"
)

type queryActor struct {
	baseActor
}

func (qa queryActor) doDBAction(db *sql.DB, qryDefn QueryDefn, exData exchgnData) QueryResult {
	var rslt QueryResult
	rslt.QueryActionName = qryDefn.QueryActionName

	stmt := qa.makeSelectStatment(qryDefn.TableName, qryDefn.ExpectedColNames)

	var where string
	if len(qryDefn.WhereReadyStr) > 0 {
		where = cStrDBWhere + qryDefn.WhereReadyStr
	} else {
		var wErr error
		where, wErr = qa.makeWhere(qryDefn.ParamterColNames, qryDefn.ParametrLogicOp, qryDefn.ParamterOper)
		if wErr != nil {
			rslt.QueryErr = wErr
			return rslt
		}
	}

	stmt += where
	exData.callerLogger.TraceDev(cstrPreparedStmt, stmt)

	stmtOut, pErr := db.Prepare(stmt)
	defer func(stmtOut *sql.Stmt) {
		if stmtOut != nil {
			stmtOut.Close()
		}
	}(stmtOut)

	if pErr != nil {
		rslt.QueryErr = exData.callerLogger.TraceError(cstrDBPrepareError, pErr.Error())
		return rslt
	}

	inputParams := qa.bindParameters(exData.qryValue.ParameterValues, nil)
	rows, qryErr := stmtOut.Query(inputParams...)
	if qryErr != nil {
		//exec error needs further outputs:
		errStr := fmt.Sprintf(cStrDetailErrorInfoForDBExec, stmt, exData.qryValue.ExpectedValues, exData.qryValue.ParameterValues, qryErr.Error())
		rslt.QueryErr = exData.callerLogger.TraceError(cstrDBQueryError, errStr)
		return rslt
	}

	cols, cErr := rows.Columns()
	if cErr != nil {
		rslt.QueryErr = exData.callerLogger.TraceError(cstrDBQueryColError, cErr.Error())
		return rslt
	} else if len(qryDefn.ExpectedColNames) != 0 && len(cols) != len(qryDefn.ExpectedColNames) {
		rslt.QueryErr = exData.callerLogger.TraceError(cstrDBQueryColCountError, len(cols))
		return rslt
	} else {
		//verify the expected col names is exactly same as this.
		for ind, val := range qryDefn.ExpectedColNames {
			if cols[ind] != val {
				rslt.QueryErr = exData.callerLogger.TraceError(cstrDBQueryColNameError, val, cols[ind])
				return rslt
			}
		}
	}

	resultList := [][]string{}
	affectedRows, fErr := qa.fetchFromRows(qryDefn, rows, &resultList)
	if fErr != nil {
		rslt.QueryErr = exData.callerLogger.TraceError(cstrDBQueryFetchError, fErr.Error())
		return rslt
	}

	rslt.AffectedRows = int64(affectedRows)
	rslt.QueryRows = resultList

	exData.callerLogger.TraceDev(cstrDBQueryActionResultInfo, stmt, exData.qryValue.ParameterValues, rslt.AffectedRows)
	return rslt
}

func (qa queryActor) fetchFromRows(qryDefn QueryDefn, rows *sql.Rows, pResultList *[][]string) (count int, err error) {
	var cols []string
	if len(qryDefn.ExpectedColNames) != 0 {
		cols = qryDefn.ExpectedColNames
	} else {
		cols, _ = rows.Columns() //ignore the error since if we get here, this function should already be called.
	}
	// Result slice string.
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	//for all rows
	count = 0
	for rows.Next() {
		result := make([]string, len(cols))
		err = rows.Scan(dest...)
		if err != nil {
			return count, fmt.Errorf("failed to scan row, Error: %s" + err.Error())
		}

		for i := range cols {
			if len(rawResult) <= i {
				break
			}
			if rawResult[i] == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(rawResult[i])
			}
		}

		*pResultList = append(*pResultList, result)
		count++
	}
	if rows.Err() != nil {
		return
	}
	return
}

func (qa queryActor) makeSelectStatment(tablename string, expected []string) string {
	stmt := "SELECT "

	if len(expected) <= 0 {
		stmt += "*"
	} else {
		for inx, val := range expected {
			if inx != 0 {
				stmt += ", "
			}
			stmt += val
		}
	}
	stmt += " FROM " + tablename

	return stmt
}
