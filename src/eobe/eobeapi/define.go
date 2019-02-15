package eobeapi

import (
	impl "eobe/eobeapi/eobeapiimpl"
)

//CallSeqResult struct holding Call Sequence Result
type CallSeqResult struct {
	SingleRow map[string]string
	MNames    []string
	MultiRow  [][]string
	ApiErr    *impl.APIError
}
