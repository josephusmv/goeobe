package eobehttp

const cStrHTTPMainLoggerFileName = "EobeHttpLog"
const cStrIndexTMPLT = "index" + cStrSuffixTmplt
const cStrIndexHTML = "index" + cStrSuffixHTML
const cStrSlash = "/"
const cStrContentType = "Content-Type"
const cStrCharSet = "charset"
const cStrDefaultCharSetUTF8 = "utf-8"
const cStrIPAddrLocalhostRemoteStart = "["
const cStrIPAddrLocalhost = "127.0.0.1"

const cStrProgramLogicError = "Server Internal logic error at: %s:%d"
const cStrIndexHTMLNotFound = cStrIndexHTML + " is mandatory for init Request Server"
const cStrErrorNoTemplates = "Failed to add template provided, interanl error."
const cStrErrorParseTemplates = "Failed to parse template provided, interanl error: %s"
const cStrErrorMustTemplates = "Failed to Must template provided, interanl error."
const cStrActionListInvalid = "Action list is for filtering invalid actions, and is mandatory."
const cStrIllegalReqServer = "Illegal initiated Request Server, Please create Request server using function: NewRequestServer"
const cStrReadModuleDirError = "Read error for module dir %s, error: %s."
const cStrDevErrorDebug2Param = "Error: %s, %s"
const cStrBadRequest = "Bad request: %s"
const cStrStatusInternalServerError = "Internal Error: %s"
const cStrParseQueryFailedError = "Parse Query failed, URL: %s, error: "
const cStrInvalidHTMLTemplateFileError = "Invalid Template File %s, delete from list.\n"
const cStrExcuteTemplateError = "Excute template %s error %s"

const cStrInvalidPostRequestError = "Post Request must be a request for an Action Target, URL: %s"
const cStrStatusNotFoundError = "<html><head><title>404 Not Found</title><head><body><h1>404 Not Found. %s</h1><body></html>"
const cStrParsePostFormError = "Parse post request form error: %s"
const cStrParseMultipartFormError = "Parse post request multipart formerror: %s"

const cStrNewRequest = "Get request from: %s:%s root URL: %s"

const cStrRequireRespFetcherError = "Implementation of Resp Fetcher Interface is mandatory."

const cStrNULLURL = "URL is null.%s" //should return http.cStrStatusInternalServerError
const cStrEmptyTemplt = "template file list is empty"
const cStrDEVLoadTmplateFile = "LoadTmplateFile: %s.\n"

//content types
const cStrTextHTML = "text/html"
const cStrTextCSS = "text/css"
const cStrAppJS = "application/javascript"
const cStrAppJSON = "application/json"
const cStrAppXML = "application/xml"
const cStrImgPNG = "image/png"
const cStrImgJPG = "image/jpg"
const cStrImgSVG = "image/svg+xml"
const cStrImgICON = "image/x-icon"
const cStrAppType = "application"
const cStrMultipart = "multipart"
const cStrMultipartForm = "multipart/form-data"

//suffixes:
const cStrSuffixTmplt = ".tmplt"
const cStrSuffixCSS = ".css"
const cStrSuffixJS = ".js"
const cStrSuffixPNG = ".png"
const cStrSuffixJPG = ".jpg"
const cStrSuffixSVG = ".svg"
const cStrSuffixICO = ".ico"
const cStrSuffixHTML = ".html"

const (
	cIntURLOK         = 7300
	cIntURLRootDomain = 7301
	cIntURLRootModule = 7302
	cIntURLBadModule  = 7304
	cIntURLNil        = 7399
)

const cStrDefaultHTML404 = `<html>
<head>
	<title> Paget Not found - 404</title>
</head>
<body>
	<h1> We cannot found the requested resources. </h1>
	<tr>
	<td>Cannot found requested URL: %s <td>
	</tr>
<body>
</html>`

//CStrExpectedXML EXPECTEDRESP = RESP_XML
const CStrExpectedXML = "RESP_XML"

//CStrExpectedJSON EXPECTEDRESP = RESP_JSON
const CStrExpectedJSON = "RESP_JSON"

//CStrExpectedHTML EXPECTEDRESP = RESP_HTML
const CStrExpectedHTML = "RESP_HTML"

//VStr404PageFileName file name for 404Page, user could modify this file but no security provided about this file
// If specified 404 file is not found, will jump to index.html page, with random Module(depends on route).
var VStr404PageFileName = "page404.html"

/*DEV log strings*/
const cStrLoadFileSuccess = "Read resources file: %s for Module: %s"
const cStrInitRequestServer = "**** Init Request Server for module: %s ****"

const (
	cIntDefaultMaxMemory = 32 << 20 // 32 MB
)

//CStrUploadFileInKey w
const CStrUploadFileInKey = "HasUploadFileInKey"
