package eobetests

import (
	"eobe/eobeapi"
	"fmt"
)

//AdminTestsOutput ...
type AdminTestsOutput struct {
}

//GetOutput ...
func (at *AdminTestsOutput) GetOutput(action string) *eobeapi.CallSeqResult {
	switch action {
	case "LoginAdmin":
		return at.loginAdmin()
	case "ViewAllUsers":
		return at.viewAllUsers()
	case "ViewUserDetails":
		return at.viewUserDetails()
	case "AddNewUserByAdmin":
		return at.addNewUserByAdmin()
	case "ModifyUserByAdmin":
		return at.ModifyUserByAdmin()
	case "DeleteUserByAdmin":
		return at.DeleteUserByAdmin()
	case "LogoutUser":
		return at.LogoutUser()
	default:
		fmt.Printf("Invalid test action: %s", action)
		return nil
	}
}

func (at *AdminTestsOutput) loginAdmin() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.MNames = nil
	csrslt.MultiRow = nil

	csrslt.SingleRow["retSuccess"] = "true"
	csrslt.SingleRow["retLoginUser"] = "sadmin"

	return &csrslt
}

func (at *AdminTestsOutput) viewAllUsers() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.MNames = []string{"username", "permission"}

	csrslt.MultiRow = [][]string{
		[]string{"sadmin", "65535"},
		[]string{"freshboy", "1"},
		[]string{"tommy", "4095"},
		[]string{"Alex", "255"},
		[]string{"Kate", "255"},
		[]string{"Mike", "255"},
		[]string{"Brow", "255"},
		[]string{"Jung", "255"},
		[]string{"John", "255"}}

	return &csrslt
}

func (at *AdminTestsOutput) viewUserDetails() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.SingleRow["Address"] = "Building-1, some place for test"
	csrslt.SingleRow["crtdttmcrtdt"] = "2014-09-30 11:29:59"
	csrslt.SingleRow["retAffectedRows"] = "1"
	csrslt.SingleRow["retLastIndex"] = "0"
	csrslt.SingleRow["username"] = "Mike"
	csrslt.SingleRow["age"] = "32"
	csrslt.SingleRow["phone1"] = "014-255-1995663-01"
	csrslt.SingleRow["phone2"] = "410-215-9395643-01"

	csrslt.MNames = nil
	csrslt.MultiRow = nil

	return &csrslt
}

func (at *AdminTestsOutput) addNewUserByAdmin() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.SingleRow["username"] = "Lee"
	csrslt.SingleRow["age"] = "54"
	csrslt.SingleRow["phone1"] = "014-255-1995663-01"
	csrslt.SingleRow["phone2"] = "410-215-9395643-01"
	csrslt.SingleRow["Address"] = "Building-2, some place for test Lee"
	csrslt.SingleRow["crtdttmcrtdt"] = "MUSTIGNORE"
	csrslt.SingleRow["retAffectedRows"] = "1"
	csrslt.SingleRow["retLastIndex"] = "0"

	csrslt.MNames = nil
	csrslt.MultiRow = nil

	return &csrslt
}

func (at *AdminTestsOutput) ModifyUserByAdmin() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.SingleRow["username"] = "Lee123"
	csrslt.SingleRow["age"] = "59"
	csrslt.SingleRow["phone1"] = "014-255-8888888-01"
	csrslt.SingleRow["phone2"] = "410-215-9395643-01"
	csrslt.SingleRow["Address"] = "Building-2, some place for test 3333"
	csrslt.SingleRow["crtdttmcrtdt"] = "MUSTIGNORE"
	csrslt.SingleRow["retAffectedRows"] = "1"
	csrslt.SingleRow["retLastIndex"] = "MUSTIGNORE"

	csrslt.MNames = nil
	csrslt.MultiRow = nil

	return &csrslt
}

func (at *AdminTestsOutput) DeleteUserByAdmin() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.SingleRow["retAffectedRows"] = "1"

	csrslt.MNames = nil
	csrslt.MultiRow = nil

	return &csrslt
}

func (at *AdminTestsOutput) LogoutUser() *eobeapi.CallSeqResult {
	var csrslt eobeapi.CallSeqResult
	csrslt.SingleRow = make(map[string]string)
	csrslt.SingleRow["retCurrentUserName"] = ""

	csrslt.MNames = nil
	csrslt.MultiRow = nil

	return &csrslt
}
