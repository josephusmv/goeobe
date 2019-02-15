package eobeload

//Keywords
const cStrHash = "#"
const cStrActionSeperateKeyWord = "########"
const cStrAction = "ACTION"
const cStrQryActionName = "QueryActionName"
const cStrResourceName = "ResourceName"
const cStrNewLine = "\n"
const cStrLineFeedNewLine = "\n\r"
const cStrHTMLTemplateDefine = "HTMLTemplateDefine"
const cStrCommaSymbol = ","
const cStrSemiColon = ":"
const cStrSpace = " "
const cStrTab = "\t"
const cStr2QuotaMark = "\""
const cStrHTTP = "HTTP"
const cStrDB = "DB"
const cStrUserConst = "UserConst"
const cCharSharp = '#'
const cChareNewLine = '\n'

//Errors
const cStrSuccessLoadCfgPromotes = "Successfully load config file: %s\n"
const cStrOpenConfigFileFailed = "open file %s failed with error: %s\n"
const cStrLoadConfigDefineFailed = "unmarshal config failed, src: %s, error: %s\n"
const cStrSuccessLoadPromotes = "Successfully load Action Define: %s\n"
const cStrOpenModuleDefineFileError = "Open Module Define file %s failed. Error: %s\n"
const cStrLoadActionDefineFailed = "unmarshal Action define failed, src: %s, error: %s\n"
const cStrLoadDBActionDefineFailed = "unmarshal DB Action define failed, src: %s, error: %s\n"
const cStrLoadResourcesDefineFailed = "unmarshal Server Resources define failed, src: %s, error: %s\n"

const cStrInvalidUconstLine = "invalid user const line: %s"
const cStrHtmlTemplateFileDoesNotExist = "HTML Template File Does not exist, path: %s\n"

const cStrActionDefineNotSetForParamParse = "Action define is null.\n"
const cStrInvalidQueryTypeForParamParse = "Invalid Query Type: %s.\n"
const cStrNotEnoughParameterError = "Cannot found value for: %s, invalid parameter map."

const cStrLoaderFactoryIError = "LoaderFactory must be initiated using call to NewLoaderFactory()."
const cStrLoaderFactoryGError = "Loader Error: %s."

//For UPDATE operation, the parameter must has below two suffix
//Otherwise, the query value will be discarded.
//Example Values: usrname_srch, title_updt...
const cStrKWSrch = "_srch"
const cStrKWUpdt = "_updt"

//For multiple Presents of one same key, use TrippleUnderScore("___") with a number to distinguish.
//This TrippleUnderScore format is defined in module define and should used in query values
//During the parsing, the tripple underscore will be dicarded.
//For example:
//		user___1, user___2 used in Define and HTTP request.
//		but when use in DB, will be user, user. only order guarantee the correct value assigned.
const cStrTrippleUnderScore = "___"

const cStrKWSelect = "SELECT"
const cStrKWDelete = "DELETE"
const cStrKWInsert = "INSERT"
const cStrKWUpdate = "UPDATE"
