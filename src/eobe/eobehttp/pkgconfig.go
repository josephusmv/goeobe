package eobehttp

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml-2/yaml-2"
)

const cStrConfigErrorHTTPSKeyInfo = "Configuration invalid: Must set HTTPS keys if NoHTTPS is false"

//PkgConfig maintain all the READONLY configurations
type PkgConfig struct {
	initiated          bool
	ServerName         string   `yaml:"ServerName"`
	ServerVersion      string   `yaml:"ServerVersion"`
	Page404FilePath    string   `yaml:"Page404FilePath"`
	TemplateFileList   []string `yaml:"TemplateFileList"`
	PreloadStaticFiles bool     `yaml:"PreloadStaticFiles"`
	DevMode            bool     `yaml:"DevMode"`

	//Inititate options
	HasHTTP      bool   `yaml:"HasHTTP"`
	HTTPPort     string `yaml:"HTTPPort"`
	NoHTTPS      bool   `yaml:"NoHTTPS"`
	HTTPSPort    string `yaml:"HTTPSPort"`
	CertFilePath string `yaml:"CertFilePath"`
	KeyFilePath  string `yaml:"KeyFilePath"`

	//HTTP response headers
	AddXSSProtectionHeader bool   `yaml:"AddXSSProtectionHeader"`
	AddXContentHeader      bool   `yaml:"AddXContentHeader"`
	AddCharSetHeader       bool   `yaml:"AddCharSetHeader"`
	ValHeaderCharSet       string `yaml:"ValHeaderCharSet"`
	AddHSTSHeader          bool   `yaml:"AddHSTSHeader"`
}

var glbPckgConfig = PkgConfig{initiated: false}

//LoadHTTPPackageConfiguration Load global package configurations
func LoadHTTPPackageConfiguration(fName string) (*PkgConfig, error) {
	content, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, fmt.Errorf(cStrOpenHTTPConfigFileError, fName, err.Error())
	}

	err = yaml.Unmarshal([]byte(content), &glbPckgConfig)
	if err != nil {
		return nil, fmt.Errorf(cStrParseHTTPConfigFileError, fName, err.Error())
	}

	if glbPckgConfig.HTTPPort == "" {
		glbPckgConfig.HTTPPort = "8080"
	}
	if glbPckgConfig.HTTPSPort == "" {
		glbPckgConfig.HTTPPort = "4030"
	}

	if !glbPckgConfig.NoHTTPS &&
		(glbPckgConfig.CertFilePath == "" || glbPckgConfig.KeyFilePath == "") {
		return nil, fmt.Errorf(cStrConfigErrorHTTPSKeyInfo)
	}

	glbPckgConfig.initiated = true

	return &glbPckgConfig, err
}

func getPckConfig() PkgConfig {
	return glbPckgConfig
}
