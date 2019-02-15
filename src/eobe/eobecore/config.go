package eobecore

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml-2/yaml-2"
)

const cStrOpenConfigFileFailed = "Failed to open Configuration file:%s, error: %s"
const cStrLoadConfigDefineFailed = "Parse Configuration failed, from file: %s, error: %s"
const cStrConfigurationInvalid = "Configuration invalid: missing mandatory field: %s"
const cStrConfigErrorHTTPPkgeError = "Configuration invalid: must specifiy the HTTPPkgCfgFile in config.yaml fail."

//ConfigDefn - YAML format matching the main configuration file input.
//	Matching: <command line input>.yaml
type ConfigDefn struct {
	//mandatory
	DBName    string `yaml:"DBName"`    //Example: mysql
	DBConnStr string `yaml:"DBConnStr"` //Example: root:123456@/THDATABASE

	//mandatory
	RootFolder     string `yaml:"RootFolder"`
	HTTPPkgCfgFile string `yaml:"HTTPPkgCfgFile"`

	//optional
	EnableDEVLog bool   `yaml:"EnableDEVLog"`
	LOGFolder    string `yaml:"LOGFolder"`

	//Mandatory
	RootModule     string   `yaml:"RootModule"`
	Modules        []string `yaml:"Modules"`
	IndexResources []string `yaml:"IndexResources"`
}

//CfgHolder ...
type CfgHolder struct {
	cfg ConfigDefn
}

//LoadConfiguration Load configuration from a given file
func (ch *CfgHolder) LoadConfiguration(cfgFilePath string) error {
	content, err := ioutil.ReadFile(cfgFilePath)
	if err != nil {
		return fmt.Errorf(cStrOpenConfigFileFailed, cfgFilePath, err.Error())
	}

	err = yaml.Unmarshal([]byte(content), &ch.cfg)
	if err != nil {
		return fmt.Errorf(cStrLoadConfigDefineFailed, cfgFilePath, err.Error())
	}

	//fill some optional fields:
	if ch.cfg.LOGFolder == "" {
		ch.cfg.LOGFolder = "./LOGS"
	}

	switch {
	case ch.cfg.HTTPPkgCfgFile == "":
		return fmt.Errorf(cStrConfigurationInvalid, "HTTPPkgCfgFile")
	case ch.cfg.Modules == nil:
		return fmt.Errorf(cStrConfigurationInvalid, "Modules")
	case ch.cfg.RootFolder == "":
		return fmt.Errorf(cStrConfigurationInvalid, "RootFolder")
	case ch.cfg.IndexResources == nil:
		return fmt.Errorf(cStrConfigurationInvalid, "IndexResources")
	case len(ch.cfg.IndexResources) != len(ch.cfg.Modules):
		return fmt.Errorf(cStrConfigurationInvalid, "IndexResources")
	case ch.cfg.DBConnStr == "":
		return fmt.Errorf(cStrConfigurationInvalid, "DBConnStr")
	case ch.cfg.DBName == "":
		return fmt.Errorf(cStrConfigurationInvalid, "DBName")
	}

	return nil
}

//GetDBDescriptions  return DB type and DB connection string
func (ch CfgHolder) GetDBDescriptions() (string, string) {
	return ch.cfg.DBName, ch.cfg.DBConnStr
}

//GetModuleInfo  return Module name and module folder array
func (ch CfgHolder) GetModuleInfo() ([]string, []string, string, string) {
	return ch.cfg.Modules, ch.cfg.IndexResources, ch.cfg.RootFolder, ch.cfg.RootModule
}

//GetLogInfo  return Debug log related infor
func (ch CfgHolder) GetLogInfo() (bool, string) {
	return ch.cfg.EnableDEVLog, ch.cfg.LOGFolder
}

//GetTemplateFileList  ,...
func (ch CfgHolder) GetHTMLInfo() string {
	return ch.cfg.HTTPPkgCfgFile
}
