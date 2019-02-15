package eobeapi

//API keywords
const cStrColon = ":"
const cStrComma = ","
const cStrSemiColon = ";"
const cStrNewLine = "\n"
const cStrSpace = " "
const cStrNotApplicable = "NA"
const cStrMultiRowDBActSepe = "$$"
const cStrLeftBrace = "("
const cStrRightBrace = ")"

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

//API Name list, and parameter name list, sort by API catagory
const CStrWriteCookie = "APIWriteCookie"
const CStrCookieName = "CookieName"
const CStrCookieValue = "CookieValue"
const CStrCookieExpire = "CookieExpire"
const CStrCookieDefaultExpire = 7
const cStrCookieAPIErrorParameter = CStrWriteCookie + " invalid parameter error: "

//Error Messages
const cStrEmptyErrorStr = "%s"
const cStrErrorInput = "input lastAPIResults should be initiated."
const cStrFailedWhenExe = "failed when executing: %s, error: %s"
const cStrParameterNotFound = "parameter %s not found"
const cStrParameterCountError = "parameter for API %s accept only %d parameters"
const cStrUnrecognizedAPICall = "unrecognized API call: %s"
const cStrDBInterfaceImplError = "unimplemented DB query interface"
const cStrErrorParsingDBAction = "Error parsing DB Action parameter: %s"
const cStrUnexpectedDBAction = "Unexpected DB Action: %s"
const cStrDBAParamMismatch = "DB Action Parameters Mismatch: %v, %v"
const cStrInvalidParameterType = "Invalid Parameter Type for API call: %s"
const cStrInvalidParameterTypeMore = "Invalid Parameter Type for API call: %s, input value: [%s]"
const cStrErrorGetMultiRowsFormat = "API GetMultiRows call should be as format: GetMultiRows: DBActionName " + cStrMultiRowDBActSepe + " Expected; WHERPARAMS"
const cStrErrorParameterTypeError = "API %s Parameter Type Error: Parameter: %s, value: %s"
const cStrCookieServerInternalError = "Cookie Server Internal Error"
const cStrNetConnectionWriteError = "Error when write to server for: %s"
const cStrNetConnectionReadError = "Error when read for response for: %s"
const cStrFetchRespError = "Error when fetch response for %s: %v"

//DEV traces
const cStrRunPresteps = "Run Pre-step for %s"
const cStrRunPoststepsSnglSave = "Run Post-step for %s: save single result"
const cStrRunPoststepsMultSave = "Run Post-step for %s: save Multiple Rows result"
const cStrRunPoststepFilter = "Run Post-step for %s: save Filter multi-rows result"

//Others
const cStrTimeYDMLayout = "20060102"
