package eobedb

import (
	"database/sql"
	"fmt"
)

type execActor struct {
	baseActor
}

func (ea execActor) doDBAction(db *sql.DB, qryDefn QueryDefn, exData exchgnData) QueryResult {
	var rslt QueryResult
	rslt.QueryActionName = qryDefn.QueryActionName

	var where string

	if len(qryDefn.WhereReadyStr) > 0 {
		where = cStrDBWhere + qryDefn.WhereReadyStr
	} else {
		var wErr error
		where, wErr = ea.makeWhere(qryDefn.ParamterColNames, qryDefn.ParametrLogicOp, qryDefn.ParamterOper)
		if wErr != nil {
			rslt.QueryErr = exData.callerLogger.TraceError(wErr.Error())
			return rslt
		}
	}

	var stmt string
	switch qryDefn.QueryType {
	case "INSERT":
		stmt = ea.makeInsertStatement(qryDefn.TableName, qryDefn.ExpectedColNames)
	case "DELETE":
		stmt = stmtWordsDeleteFrom + qryDefn.TableName + " " + where
	case "UPDATE":
		stmt = ea.makeUpdateStatement(qryDefn.TableName, qryDefn.ExpectedColNames) + " " + where
	default:
		rslt.QueryErr = exData.callerLogger.TraceError(cstrInvalidExecType, qryDefn.QueryType)
		return rslt
	}

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

	inputParams := ea.bindParameters(exData.qryValue.ExpectedValues, exData.qryValue.ParameterValues)
	exeResult, exeErr := stmtOut.Exec(inputParams...)
	if exeErr != nil {
		//exec error needs further outputs:
		errStr := fmt.Sprintf(cStrDetailErrorInfoForDBExec, stmt, exData.qryValue.ExpectedValues, exData.qryValue.ParameterValues, exeErr.Error())
		rslt.QueryErr = exData.callerLogger.TraceError(cstrDBExeError, errStr)
		return rslt
	}
	affRows, arErr := exeResult.RowsAffected()
	lsID, liErr := exeResult.LastInsertId()
	if arErr != nil {
		rslt.QueryErr = exData.callerLogger.TraceError(cstrDBRowsAffectedError, arErr.Error())
		return rslt
	}
	if liErr != nil {
		rslt.QueryErr = exData.callerLogger.TraceError(cstrLastInsertIDError, liErr.Error())
		return rslt
	}

	rslt.AffectedRows = affRows
	rslt.LastIndex = lsID

	exData.callerLogger.TraceDev(cstrDBExecActionResultInfo, stmt, exData.qryValue.ExpectedValues, exData.qryValue.ParameterValues, rslt.AffectedRows, rslt.LastIndex)
	return rslt
}

const stmtWordsDeleteFrom = "DELETE FROM "

const stmtWordsInsertINTO = "INSERT INTO "
const stmtWordsleftBrace = " ("
const stmtWordsRightBrace = ")"
const stmtWordsvalueChar = ") VALUES("

func (ea execActor) makeInsertStatement(tablename string, columns []string) string {
	stmt := stmtWordsInsertINTO + tablename + stmtWordsleftBrace
	for ind, col := range columns {
		if ind != 0 {
			stmt += ", "
		}
		stmt += col
	}
	stmt += stmtWordsvalueChar

	for i := range columns {
		if i != 0 {
			stmt += ", " + cStrBindVarSymbol
		} else {
			stmt += cStrBindVarSymbol
		}
	}

	stmt += stmtWordsRightBrace

	return stmt
}

const stmtWordsUpdateChar = "UPDATE "
const stmtWordsSETChar = " SET "
const stmtEqualSymbol = "="

func (ea execActor) makeUpdateStatement(tablename string, columns []string) string {
	stmt := stmtWordsUpdateChar + tablename + stmtWordsSETChar

	for i, val := range columns {
		if i != 0 {
			stmt += ", " + val + stmtEqualSymbol + cStrBindVarSymbol
		} else {
			stmt += val + stmtEqualSymbol + cStrBindVarSymbol
		}
	}

	return stmt
}
