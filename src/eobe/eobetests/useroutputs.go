package eobetests

import (
	"eobe/eobeapi"
	"fmt"
)

//usrOutputs ...
type usrOutputs struct {
	userName string
}

//GetOutput ...
func (srt *usrOutputs) GetOutput(action string) *eobeapi.CallSeqResult {
	switch action {
	case "LoginUser":
		return srt.LoginUser()
	case "FetchAPIList":
		return srt.FetchAPIList()
	case "FetchAllAPIParams":
		return srt.FetchAllAPIParams()
	case "FetchAllAPIRslts":
		return srt.FetchAllAPIRslts()
	case "ValidateCrntUsr":
		return srt.ValidateCrntUsr()
	case "LogoutUser":
		return srt.LogoutUser()
	case "UploadFile":
		return srt.UploadFile()
	default:
		fmt.Printf("Invalid test action: %s", action)
		return nil
	}
}

func (srt *usrOutputs) LoginUser() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)

	csrslt.SingleRow["retSuccess"] = "true"
	switch srt.userName {
	case "Mike":
		csrslt.SingleRow["retLoginUser"] = "Mike"
		csrslt.SingleRow["permission"] = "255"
		csrslt.SingleRow["retAffectedRows"] = "1"
		csrslt.SingleRow["retSuccess"] = "true"
		csrslt.SingleRow["username"] = "Mike"
	case "Alex":
		csrslt.SingleRow["retLoginUser"] = "Alex"
		csrslt.SingleRow["permission"] = "255"
		csrslt.SingleRow["retAffectedRows"] = "1"
		csrslt.SingleRow["retSuccess"] = "true"
		csrslt.SingleRow["username"] = "Alex"
	case "Kate":
		csrslt.SingleRow["retLoginUser"] = "Kate"
		csrslt.SingleRow["permission"] = "255"
		csrslt.SingleRow["retAffectedRows"] = "1"
		csrslt.SingleRow["retSuccess"] = "true"
		csrslt.SingleRow["username"] = "Kate"
	case "Jung":
		csrslt.SingleRow["retLoginUser"] = "Jung"
		csrslt.SingleRow["permission"] = "255"
		csrslt.SingleRow["retAffectedRows"] = "1"
		csrslt.SingleRow["retSuccess"] = "true"
		csrslt.SingleRow["username"] = "Jung"
	case "Brow":
		csrslt.SingleRow["retLoginUser"] = "Brow"
		csrslt.SingleRow["permission"] = "255"
		csrslt.SingleRow["retAffectedRows"] = "1"
		csrslt.SingleRow["retSuccess"] = "true"
		csrslt.SingleRow["username"] = "Brow"
	case "tommy":
		csrslt.SingleRow["retLoginUser"] = "tommy"
		csrslt.SingleRow["permission"] = "4095"
		csrslt.SingleRow["retAffectedRows"] = "1"
		csrslt.SingleRow["retSuccess"] = "true"
		csrslt.SingleRow["username"] = "tommy"
	default:
	}

	csrslt.MNames = nil
	csrslt.MultiRow = nil

	return &csrslt
}

func (srt *usrOutputs) FetchAPIList() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.MNames = nil
	csrslt.MultiRow = nil
	return &csrslt
}

func (srt *usrOutputs) FetchAllAPIParams() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.MNames = nil
	csrslt.MultiRow = [][]string{
		[]string{"compareInt", "ParamLeft", "14", "12", "1", "left source for compare, only accepty variables from: ^/?/$"},
		[]string{"compareInt", "ParamOprt", "1", "3", "2", "operator for compare, only accepty literal string, like: eq, gt, lt, ge...."},
		[]string{"compareInt", "ParamRght", "15", "12", "3", "Right source for compare, accepty both literal and variables from: ^/?/$"},
		[]string{"ValidateInt", "ParamLeft", "14", "12", "1", "left source for Validate, only accepty variables from: ^/?/$"},
		[]string{"ValidateInt", "ParamOprt", "1", "3", "2", "operator for Validate, only accepty literal string, like: eq, gt, lt, ge...."},
		[]string{"ValidateInt", "ParamRght", "15", "12", "3", "Right source for Validate, accepty both literal and variables from: ^/?/$"}}

	return &csrslt
}

func (srt *usrOutputs) FetchAllAPIRslts() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.MNames = nil
	csrslt.MultiRow = [][]string{
		[]string{"compareInt", "retCmpIntResult", "Return lower case \"true\" or \"false\" as a result."},
		[]string{"ValidateInt", "retVldIntResult", "Return lower case \"true\" or \"false\" as a result."}}

	return &csrslt
}

func (srt *usrOutputs) LogoutUser() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.SingleRow["retCurrentUserName"] = ""

	csrslt.MNames = nil
	csrslt.MultiRow = nil

	return &csrslt
}

func (srt *usrOutputs) ValidateCrntUsr() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)

	switch srt.userName {
	case "Alex":
		csrslt.SingleRow["retIsValid"] = "true"
		csrslt.SingleRow["retCurrentUserName"] = "Alex"
	case "Mike":
		csrslt.SingleRow["retIsValid"] = "true"
		csrslt.SingleRow["retCurrentUserName"] = "Mike"
	default:
	}
	csrslt.MNames = nil
	csrslt.MultiRow = nil

	return &csrslt
}

func (srt *usrOutputs) UploadFile() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.SingleRow["retResultFilePath"] = "./fileuploaded.jpg"
	csrslt.SingleRow["retSuccess"] = "true"

	csrslt.MNames = nil
	csrslt.MultiRow = nil

	return &csrslt
}
