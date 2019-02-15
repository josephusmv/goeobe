package eobetests

import (
	"fmt"
)

type admInputs struct {
	actionName string
}

func (admt *admInputs) getClientInfo() (ip, port, sid string) {
	return "192.168.1.101", "8050", ""
}
func (admt *admInputs) setActionName(actionname string) {
	admt.actionName = actionname
}
func (admt *admInputs) getRootPath() string {
	return "tstdata/admintst/"
}

func (admt *admInputs) getActionName() string {
	return admt.actionName
}

//If invalid acion given, the call seq will report: "error:input lastAPIResults should be initiated."
func (admt *admInputs) getRequestMap() map[string]string {
	switch admt.actionName {
	case "LoginAdmin":
		return admt.loginAdmin()
	case "ViewAllUsers":
		return admt.viewAllUsers()
	case "ViewUserDetails":
		return admt.viewUserDetails()
	case "AddNewUserByAdmin":
		return admt.AddNewUserByAdmin()
	case "ModifyUserByAdmin":
		return admt.ModifyUserByAdmin()
	case "DeleteUserByAdmin":
		return admt.DeleteUserByAdmin()
	case "LogoutUser":
		return admt.LogoutUser()
	default:
		fmt.Printf("Invalid test action: %s", admt.actionName)
		return nil
	}
}

func (admt *admInputs) loginAdmin() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	rqstKeyValueMap["usr"] = "sadmin"
	rqstKeyValueMap["pwd"] = "superadmin"
	rqstKeyValueMap["expdays"] = "7"
	return rqstKeyValueMap
}

func (admt *admInputs) viewAllUsers() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	rqstKeyValueMap["sidx"] = "0"
	rqstKeyValueMap["count"] = "100"
	return rqstKeyValueMap
}

func (admt *admInputs) viewUserDetails() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	rqstKeyValueMap["username"] = "Mike"
	return rqstKeyValueMap
}

func (admt *admInputs) AddNewUserByAdmin() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	rqstKeyValueMap["username"] = "Lee"
	rqstKeyValueMap["userpwd"] = "lee"
	rqstKeyValueMap["permission"] = "32768"
	rqstKeyValueMap["age"] = "54"
	rqstKeyValueMap["phone1"] = "014-255-1995663-01"
	rqstKeyValueMap["phone2"] = "410-215-9395643-01"
	rqstKeyValueMap["address"] = "Building-2, some place for test Lee"
	return rqstKeyValueMap
}

func (admt *admInputs) ModifyUserByAdmin() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	rqstKeyValueMap["oldusername"] = "Lee"
	rqstKeyValueMap["username"] = "Lee123"
	rqstKeyValueMap["userpwd"] = "lee123"
	rqstKeyValueMap["permission"] = "32769"
	rqstKeyValueMap["age"] = "59"
	rqstKeyValueMap["phone1"] = "014-255-8888888-01"
	rqstKeyValueMap["phone2"] = "410-215-9395643-01"
	rqstKeyValueMap["address"] = "Building-2, some place for test 3333"
	return rqstKeyValueMap
}

func (admt *admInputs) DeleteUserByAdmin() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	rqstKeyValueMap["username"] = "Lee123"
	return rqstKeyValueMap
}

func (admt *admInputs) LogoutUser() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	return rqstKeyValueMap
}

func (admt *admInputs) loadFile() map[string][]byte {
	return nil
}
