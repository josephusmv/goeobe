package eobeapiimpl

import (
	"fmt"
)

//APIErrType ...
type APIErrType int

const (
	//CErrSuccess  No error happens
	CErrSuccess APIErrType = iota
	//CErrServerInternalError  Some server internal logic error happens, should log
	CErrServerInternalError
	//CErrBadCallError  An server error possiblly  lead by inproper call or parameter define, should log
	CErrBadCallError
	//CErrBadParameterError Bad parameter from request, front end caused error, only DEV log.
	CErrBadParameterError
	//CErrRunValidateFailure  a validation step fails, , only DEV log.
	CErrRunValidateFailure
	//CErrGenericError  generic uncatagoried error, should log
	CErrGenericError
)

const cStrAPIErrorPrefix = "Run API error: "

//APIError ...
type APIError struct {
	error
	ErrType APIErrType
}

func NewAPIErrorf(errType APIErrType, format string, a ...interface{}) *APIError {
	err := fmt.Errorf(format, a...)
	apiErr := APIError{
		error:   err,
		ErrType: errType}

	return &apiErr
}

func ApiSuccess() *APIError {
	return &APIError{ErrType: CErrSuccess}
}

func (e APIError) HasError() bool {
	if e.ErrType == CErrSuccess {
		return false
	}
	return true
}

//Error Strings
//DB errors
const cStrDBParseParameterErr = "Parse Parameter error for DB Action: %s, error: %s"
const cStrDBInvalidQueryDefn = "Query define for DB action is invalid"
const cStrErrorFromEobeDB = "Error when executing DB Actions, DB module reported: %s"
