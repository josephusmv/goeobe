package eobecltmgmt

const cStrErrorSessionIDMismatch = "The Session ID is mismatch, previous: %s, new: %s"
const cStrErrorAddNewCookie = "Failed to Add New Cookie"

const cStrErrorFailedToDoAction = "Failed to Do Action: %d"

///logs
const cStrError = "ERROR: %s.\n"
const cStrWorkerRunning = "Worker started, now: %s.\n"
const cStrCommandRecvd = "Got command %s.\n"
const cStrCommandRslt = "Command result: %s, %s, %s.\n"

const cStrCmdRcvdNewAccess = "Got Command NewAccess Recieved, parameter: %s, %s, %s.\n"
const cStrCmdRcvdLogin = "Got Command Login Recieved, parameter: %s, %s, %s, %s, %d.\n"
const cStrCmdRcvdLogout = "Got Command Logout Recieved, parameter: %s, %s, %s, %s.\n"
const cStrCmdRcvdAddCookie = "Got Command AddCookie Recieved, parameter: %s, %s, %s, %s, %s, %d.\n"
const cStrCmdRcvdDeleteCookie = "Got Command DeleteCookie Recieved, parameter: %s, %s, %s, %s.\n"
const cStrCmdRcvdGetCookie = "Got Command GetCookie Recieved, parameter: %s, %s, %s.\n"
const cStrRunExpireCheckDetails = "Checking Expire for %s --- %s:  expire: %s, now: %s.\n"
const cStrClientExpired = "Client %s --- %s expired.\n"
const cStrDumpIpClientTable = "Dump Client table, looking for ID: %s\n"
const cVerboseIPUserTableInfo = "\t\t--> %s %s %s %s %b"

const cStrDumpCommandResultDetails = "Dump Command Result Details: %s, %s\n"
const cVerboseCookieInfo = "\t\t--> %s %s %s %s"

const cStrSlash = "\\/"
const cStrSingleSlash = "/"

//DEV Logs
const cStrDEVCommandRecvd = "Got command %s.\n"

//Error Logs
const cStrErrorUnreconizedCmd = "Error: Unreconized Message %s.\n"

//Info Logs
const cStrQuitWorker = "Got Quit command, quit worker."
const cStrInfoInitWorker = "Worker initiated."
const cStrInfoWorkerStarted = "Worker goroutine started. "
const cStrInfoWorkerExited = "Worker goroutine exited. "
const cStrRunExpireCheck = "Run Expire Check, now: %s."
const cStrWokerGoRoutineExited = "Woker Goroutine properly exited."
