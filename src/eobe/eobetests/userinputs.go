package eobetests

import (
	"fmt"
	"io/ioutil"
)

type userInputs struct {
	actionName string
	userName   string
}

func (srt *userInputs) setActionName(actionname string) {
	srt.actionName = actionname
}
func (srt *userInputs) getRootPath() string {
	return "tstdata/usertst/"
}

func (srt *userInputs) getActionName() string {
	return srt.actionName
}
func (srt *userInputs) getClientInfo() (ip, port, sid string) {
	switch srt.userName {
	case "Mike":
		return "192.168.1.101", "8050", ""
	case "Alex":
		return "192.168.2.101", "8070", ""
	case "Kate":
		return "192.168.134.121", "8070", ""
	case "Jung":
		return "192.168.201.121", "8020", ""
	case "Brow":
		return "192.158.201.121", "8120", ""
	case "tommy":
		return "193.168.201.131", "8029", ""
	default:
	}
	return "192.168.0.1", "80", ""
}

//If invalid acion given, the call seq will report: "error:input lastAPIResults should be initiated."
func (srt *userInputs) getRequestMap() map[string]string {
	switch srt.actionName {
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
		fmt.Printf("Invalid test action: %s", srt.actionName)
		return nil
	}
}

func (srt *userInputs) LoginUser() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	switch srt.userName {
	case "Mike":
		rqstKeyValueMap["usr"] = "Mike"
		rqstKeyValueMap["pwd"] = "mike_enc"
		rqstKeyValueMap["expdays"] = "3"
	case "Alex":
		rqstKeyValueMap["usr"] = "Alex"
		rqstKeyValueMap["pwd"] = "alex_enc"
		rqstKeyValueMap["expdays"] = "7"
	case "Kate":
		rqstKeyValueMap["usr"] = "Kate"
		rqstKeyValueMap["pwd"] = "kate_enc"
		rqstKeyValueMap["expdays"] = "7"
	case "Jung":
		rqstKeyValueMap["usr"] = "Jung"
		rqstKeyValueMap["pwd"] = "jung_enc"
		rqstKeyValueMap["expdays"] = "7"
	case "Brow":
		rqstKeyValueMap["usr"] = "Brow"
		rqstKeyValueMap["pwd"] = "brow_enc"
		rqstKeyValueMap["expdays"] = "7"
	case "tommy":
		rqstKeyValueMap["usr"] = "tommy"
		rqstKeyValueMap["pwd"] = "tommy_enc"
		rqstKeyValueMap["expdays"] = "7"
	default:
	}
	return rqstKeyValueMap
}

func (srt *userInputs) FetchAPIList() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	switch srt.userName {
	case "Mike":
		rqstKeyValueMap["sidx"] = "0"
		rqstKeyValueMap["count"] = "100"
	case "Alex":
		rqstKeyValueMap["sidx"] = "0"
		rqstKeyValueMap["count"] = "100"
	case "Kate":
		rqstKeyValueMap["sidx"] = "0"
		rqstKeyValueMap["count"] = "100"
	case "Jung":
		rqstKeyValueMap["sidx"] = "0"
		rqstKeyValueMap["count"] = "100"
	case "Brow":
		rqstKeyValueMap["sidx"] = "0"
		rqstKeyValueMap["count"] = "100"
	case "tommy":
		rqstKeyValueMap["sidx"] = "0"
		rqstKeyValueMap["count"] = "100"
	default:
	}
	return rqstKeyValueMap
}

func (srt *userInputs) FetchAllAPIParams() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	return rqstKeyValueMap
}

func (srt *userInputs) FetchAllAPIRslts() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	return rqstKeyValueMap
}

func (srt *userInputs) ValidateCrntUsr() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	switch srt.userName {
	case "Mike":
		rqstKeyValueMap["user"] = "Mike"
	case "Alex":
		rqstKeyValueMap["user"] = "Alex"
	default:
	}
	return rqstKeyValueMap
}

func (srt *userInputs) LogoutUser() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	return rqstKeyValueMap
}

func (srt *userInputs) UploadFile() map[string]string {
	rqstKeyValueMap := make(map[string]string)
	rqstKeyValueMap["filename"] = "fileuploaded.jpg"
	return rqstKeyValueMap
}

const cStrLocalFileSource = "sourceimg.jpg"

func (srt *userInputs) loadFile() map[string][]byte {
	fdata, err := ioutil.ReadFile(cStrLocalFileSource)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fileMap := make(map[string][]byte)
	fileMap[cStrLocalFileSource] = fdata

	return fileMap
}
