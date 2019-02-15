package eobeapiimpl

//API keywords
const cStrColon = ":"
const cStrComma = ","
const cStrSemiColon = ";"
const cStrSlash = "/"
const cStrNewLine = "\n"
const cStrEmpty = ""
const cStrSpace = " "
const cStrNotApplicable = "NA"
const cStrMultiRowDBActSepe = "$$"
const (
	cIntSingleResult   = 0x0871 //save to CallSeqResult.SingleRow
	cIntMultiRowResult = 0x0872 //Save to CallSeqResult.MNames and MultiRow
	cIntFilterRows     = 0x0873 //Save to a temporary filter variable.(only for MultiRow)
	cIntAsRange        = 0x0875 //Save to a temporary range variable.(only for MultiRow)
	cIntNotSave        = 0x0876
)

const cStrTure = "true"
const cStrFalse = "false"
const cStrRetSuccess = "retSuccess"

const cStrHexPrefix = "0x"
const cStrEqualTo = "eq"
const cStrGreaterThan = "gt"
const cStrSmallerThan = "lt"
const cStrGreaterEqualThan = "ge"
const cStrSmallerEqualThan = "le"
const cStrNotEqualTo = "neq"
const cStrSymbolEqualTo = "="
const cStrSymbolGreaterThan = ">"
const cStrSymbolSmallerThan = "<t"
const cStrSymbolGreaterEqualThan = ">="
const cStrSymbolSmallerEqualThan = "<="
const cStrSymbolNotEqualTo = "!="
const COPSaveRslt = "(SaveRslt)"
const COPSaveRows = "(SaveRows)"
const COPFilterRows = "(FilterRows)"
const COPAsRange = "(AsRange)"

//API Name list, and parameter name list, sort by API catagory
const CStrWriteCookie = "APIWriteCookie"
const CStrCookieName = "CookieName"
const CStrCookieValue = "CookieValue"
const CStrCookieExpire = "CookieExpire"
const CStrCookieDefaultExpire = 7
const cStrCookieAPIErrorParameter = CStrWriteCookie + " invalid parameter error: "

//Error Messages
const cStrEmptyErrorStr = "%s"
const cStrErrorAPICallFormat = "Malformat: %s - API call should be as format: APIName: ?Param1, ^PAram2, UParAM3."
const cStrErrorInput = "input lastAPIResults should be initiated."
const cStrFailedWhenExe = "failed when executing: %s, error: %s"
const cStrParameterNotFound = "parameter %s not found"
const cStrParameterCountError = "parameter for API %s accept only %d parameters"
const cStrUnrecognizedAPICall = "unrecognized API call: %s"
const cStrDBInterfaceImplError = "unimplemented DB query interface"
const cStrErrorParsingDBAction = "Error parsing DB Action parameter: %s"
const cStrUnexpectedDBAction = "Unexpected DB Action: %s"
const cStrDBAParamMismatch = "DB Action Parameters Mismatch: %v, %v"
const cStrInvalidParameterTypeMore = "Invalid Parameter Type for API call: %s, input value: [%s]"
const cStrErrorGetMultiRowsFormat = "API GetMultiRows call should be as format: GetMultiRows: DBActionName " + cStrMultiRowDBActSepe + " Expected; WHERPARAMS"
const cStrErrorParameterTypeError = "API %s Parameter Type Error: Parameter: %s, value: %s"
const cStrCookieServerInternalError = "Cookie Server Internal Error: %s"
const cStrNetConnectionWriteError = "Error when write to server for: %s"
const cStrNetConnectionReadError = "Error when read for response for: %s"
const cStrFetchRespError = "Error when fetch response for %s: %v"
const cStrDEVLogGetUser = "Get user result: [%s]"

//Others
const cStrTimeYDMLayout = "20060102"
