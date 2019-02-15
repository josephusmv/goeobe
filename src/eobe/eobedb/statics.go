package eobedb

//String Constants
const cstrDBSuccessfullyOpened = " ****** DB Successfully opened for: %s\n"
const cstrDBSuccessfullyInited = "DB Successfully Initiated, start DB worker.\n"
const cstrQueryActionNameNotFound = "could not found given query name: %s\n"
const cstrExchangeData = "Exchange Data: %v\n"
const cstrDBServerStopGood = " ****** DB server properly stopped.\n"
const cstrDBServerStopBad = " ****** DB server stopped for reason: %s.\n"
const cstrTypeCastError = "Error recongnize type: %T.\n"
const cstrDBConnectionError = " !!!! DB connection error: %s.\n"
const cstrDBOpenError = "DB open error, DB type: %s, ConnStr: %s, Error: %s\n"
const cstrUnknown = "Unknown"
const cstrDBActorInitMsg = "Got Query: %v - %v\n"
const cstrWhereParamError = "Invalid number of WHERE condition parameters\n"
const cstrInvalidExecType = "Invalid Execution type here: %s\n"
const cstrDBPrepareError = "db.Prepare error: %s.\n"
const cstrDBQueryError = "db.Query error: %s.\n"
const cstrDBQueryColError = "Get queried columns error: %s.\n"
const cstrDBQueryColCountError = "UnExpected query Length: %d.\n"
const cstrDBQueryColNameError = "UnExpected query Name mismatch: %s - %s.\n"
const cstrDBQueryFetchError = "Fetch DB result Error: %s.\n"
const cstrDBExeError = "db.Exec error: %s.\n"
const cstrDBRowsAffectedError = "RowsAffected error: %s.\n"
const cstrLastInsertIDError = "db.LastInsertId error: %s.\n"
const cstrPreparedStmt = "\n--- SQL Statment: \n---\t%s.\n"
const cstrDBQueryActionResultInfo = "\n--- DB Query summary: \n---\t%s\n--- Bind values: \n---\t%v, \n--- Fetched %d rows.\n--- Done.\n\n"
const cstrDBExecActionResultInfo = "\n--- DB Exec summary: \n---\t%s\n--- Bind values: \n---\t%v\n---\t%v, \n--- Affected %d rows, last index: %d.\n--- Done.\n\n"
const cStrDetailErrorInfoForDBExec = "\n\tStmt: %s, \n\texpected values: %v, \n\twhere params: %s, \n\tDB driver error: %s"

const cStrSetDBLogPath = "Set DB log Path to: %s"

//DBServer Options
const (
	OptionDBMaxConcurrencyInt = 0x000000F1
	OptionDBLogRoot           = 0x000000F2
	OptionEnableDEVLog        = 0x000000F3
)
const DefaultOptionDBMaxConcurrencyInt = 30
const DefaultLogRoot = "./LOGS"

//DB Operations lists
const CstrDBOptGreaterThan = "gt"
const CstrDBOptSmallerThan = "lt"
const CstrDBOptEqualTo = "eq"
const CstrDBOptNotEqualTo = "neq"
const CstrDBOptLike = "like"

//Key Words
const cStrDBFROM = " FROM "
const cStrDBWhere = " WHERE "
const CStrDBINSERT = "INSERT"
const CStrDBDELETE = "DELETE"
const CStrDBSELECT = "SELECT"
const CStrDBUPDATE = "UPDATE"
