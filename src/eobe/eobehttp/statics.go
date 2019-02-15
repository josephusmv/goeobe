package eobehttp

const cStrHTTPMainLoggerFileName = "EobeHttpLog"
const cStrIndexTMPLT = "index" + cStrSuffixTmplt
const cStrIndexHTML = "index" + cStrSuffixHTML
const cStrPageNotFound404 = "PageNotFound404" + cStrSuffixHTML
const cStrSlash = "/"
const cStrQuestionair = "?"
const cStrDot = "."
const cStrContentType = "Content-Type"
const cStrIPAddrLocalhostRemoteStart = "["
const cStrIPAddrLocalhost = "127.0.0.1"

const cStrInvalidPage404FileError = "Invalid Page404 File %s, error: %s."
const cStrIndexHTMLNotFound = cStrIndexHTML + " is mandatory for init Request Server"
const cStrErrorNoTemplates = "Failed to add template provided, interanl error."
const cStrErrorParseTemplates = "Failed to parse template provided, interanl error: %s"
const cStrErrorMustTemplates = "Failed to Must template provided, interanl error."
const cStrActionListInvalid = "Action list is for filtering invalid actions, and is mandatory."
const cStrPackageNotInit = "eobehtpp Package not initiated, cannot init request server"
const cStrIllegalReqServer = "Illegal initiated Request Server, Please create Request server using function: NewRequestServer"
const cStrReadModuleDirError = "Read error for module dir %s, error: %s."
const cStrDevErrorDebug2Param = "Error: %s, %s"
const cStrGenError = "general error"
const cStrBadRequest = "Bad request: %s"
const cStrBadRequestWithErr = "request: %s, error %s, error code: %d"
const cStrStatusInternalServerError = "Server Internal Error: %s"
const cStrStatusInternalServerErrorWithErr = "Server Internal Error for request: %s, error: %s"
const cStrParseQueryFailedError = "Parse Query failed, URL: %s, error: "
const cStrInvalidHTMLTemplateFileError = "Invalid Template File %s, delete from list.\n"
const cStrExcuteTemplateError = "Excute template %s error %s"

const cStrInvalidPostRequestError = "Post Request must be a request for an Action Target, URL: %s"
const cStrStatusNotFoundError = "<html><head><title>404 Not Found</title><head><body><h1>404 Not Found. %s</h1><body></html>"
const cStrParsePostFormError = "Parse post request form error: %s"
const cStrParseMultipartFormError = "Parse post request multipart formerror: %s"
const cStrDEVRedirectURL = "Redirect URL to: %s"

const cStrOpenHTTPConfigFileError = "Open HTTP Config File %s Error: %s"
const cStrParseHTTPConfigFileError = "Parse HTTP Config File %s Error: %s"

const cStrNewRequest = "Get request from: %s:%s root URL: %s"

const cStrRequireRespFetcherError = "Implementation of Resp Fetcher Interface is mandatory."

const cStrNULLURL = "URL is null.%s" //should return http.cStrStatusInternalServerError
const cStrEmptyTemplt = "template file list is empty"
const cStrDEVLoadTmplateFile = "Load Template File path: %s, file name: %s.\n"

const cStrAccepct = "Accept"

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
const cStrSuffixJPEG = ".jpeg"
const cStrSuffixJPG = ".jpg"
const cStrSuffixSVG = ".svg"
const cStrSuffixICO = ".ico"
const cStrSuffixBMP = ".bmp"
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

/*DEV log strings*/
const cStrLoadResFileSuccess = "Read resources file: [%s]: len: %d."
const cStrLoadMeidaFileSuccess = "Read Media file: [%s]: len: %d."
const cStrInitRequestServer = "**** Init Request Server for module: %s successfully. ****"

const (
	cIntDefaultMaxMemory = 32 << 20 // 32 MB
)
