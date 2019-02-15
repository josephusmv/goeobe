package eobeload

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml-2/yaml-2"
)

const cStrDuplicatedResource = "duplicated %s resource define: %s"

//HTTPActionDefn  ...
type HTTPActionDefn struct {
	ResourceName string `yaml:"ResourceName"`
	MethodGET    struct {
		UnSupport       bool     `yaml:"UnSupport"`
		QueryParameters []string `yaml:"QueryParameters"`
		ActualResources string   `yaml:"ActualResources"`
		SupportRespType []string `yaml:"SupportRespType"`
		APICalls        []string `yaml:"APICalls"`
	} `yaml:"MethodGET"`
	MethodPOST struct {
		UnSupport       bool     `yaml:"UnSupport"`
		QueryParameters []string `yaml:"QueryParameters"`
		ActualResources string   `yaml:"ActualResources"`
		SupportRespType []string `yaml:"SupportRespType"`
		APICalls        []string `yaml:"APICalls"`
	} `yaml:"MethodPOST"`
	MethodDELETE struct {
		UnSupport       bool     `yaml:"UnSupport"`
		QueryParameters []string `yaml:"QueryParameters"`
		ActualResources string   `yaml:"ActualResources"`
		SupportRespType []string `yaml:"SupportRespType"`
		APICalls        []string `yaml:"APICalls"`
	} `yaml:"MethodDELETE"`
	MethodUPDATE struct {
		UnSupport       bool     `yaml:"UnSupport"`
		QueryParameters []string `yaml:"QueryParameters"`
		ActualResources string   `yaml:"ActualResources"`
		SupportRespType []string `yaml:"SupportRespType"`
		APICalls        []string `yaml:"APICalls"`
	} `yaml:"MethodUPDATE"`
}

//HTTPActionLoader eobeload.HTTPActionLoader
type HTTPActionLoader struct {
	commonBase
}

//NewHTTPActionLoader eobeload.NewHTTPActionLoader
func NewHTTPActionLoader() *HTTPActionLoader {
	var httpActLoader HTTPActionLoader

	return &httpActLoader
}

//LoadFromFile ...
func (loader *HTTPActionLoader) LoadFromFile(filename string) (*HTTPActionMap, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf(cStrOpenConfigFileFailed, filename, err.Error())
	}

	var httpActs HTTPActionMap
	httpActs.actMap = make(map[string]*HTTPActionDefn)

	actBytesArry := bytes.Split(content, []byte(cStrActionSeperateKeyWord))

	for _, actionByte := range actBytesArry {
		var skip bool
		actionByte, skip = loader.formatYamlBytes(actionByte, cStrResourceName)
		if skip {
			continue
		}

		actDfn, err := loader.loadFromYaml(actionByte)
		if err != nil {
			loader.errList = append(loader.errList, err)
			continue
		}

		if actDfn == nil || actDfn.ResourceName == "" {
			continue
		}

		if _, found := httpActs.actMap[actDfn.ResourceName]; found {
			loader.errList = append(loader.errList, fmt.Errorf(cStrDuplicatedResource, cStrHTTP, actDfn.ResourceName))
		}

		httpActs.actMap[actDfn.ResourceName] = actDfn
	}

	return &httpActs, nil
}

func (loader HTTPActionLoader) loadFromYaml(data []byte) (*HTTPActionDefn, error) {
	var err error
	var httpDfn HTTPActionDefn

	bytes.Trim(data, cStrLineFeedNewLine)
	if len(data) <= 0 {
		return nil, nil
	}

	err = yaml.Unmarshal([]byte(data), &httpDfn)
	if err != nil {
		return nil, fmt.Errorf(cStrLoadResourcesDefineFailed, string(data), err.Error())
	}

	//if a request is not for trigguring an API nor easking for a resources with expected format, it's invalid
	if (httpDfn.MethodGET.APICalls == nil || len(httpDfn.MethodGET.APICalls) == 0) &&
		(httpDfn.MethodGET.SupportRespType == nil || len(httpDfn.MethodGET.SupportRespType) == 0) &&
		httpDfn.MethodGET.ActualResources == "" {
		httpDfn.MethodGET.UnSupport = true
	}
	if (httpDfn.MethodPOST.APICalls == nil || len(httpDfn.MethodPOST.APICalls) == 0) &&
		(httpDfn.MethodPOST.SupportRespType == nil || len(httpDfn.MethodPOST.SupportRespType) == 0) &&
		httpDfn.MethodPOST.ActualResources == "" {
		httpDfn.MethodPOST.UnSupport = true
	}
	if (httpDfn.MethodDELETE.APICalls == nil || len(httpDfn.MethodDELETE.APICalls) == 0) &&
		(httpDfn.MethodDELETE.SupportRespType == nil || len(httpDfn.MethodDELETE.SupportRespType) == 0) &&
		httpDfn.MethodDELETE.ActualResources == "" {
		httpDfn.MethodDELETE.UnSupport = true
	}
	if (httpDfn.MethodUPDATE.APICalls == nil || len(httpDfn.MethodUPDATE.APICalls) == 0) &&
		(httpDfn.MethodUPDATE.SupportRespType == nil || len(httpDfn.MethodUPDATE.SupportRespType) == 0) &&
		httpDfn.MethodUPDATE.ActualResources == "" {
		httpDfn.MethodUPDATE.UnSupport = true
	}

	//Pre-format
	for i := range httpDfn.MethodGET.APICalls {
		httpDfn.MethodGET.APICalls[i] = strings.Trim(httpDfn.MethodGET.APICalls[i], cStrSpace+cStrTab+cStrNewLine)
	}
	for i := range httpDfn.MethodPOST.APICalls {
		httpDfn.MethodPOST.APICalls[i] = strings.Trim(httpDfn.MethodPOST.APICalls[i], cStrSpace+cStrTab+cStrNewLine)
	}
	for i := range httpDfn.MethodDELETE.APICalls {
		httpDfn.MethodDELETE.APICalls[i] = strings.Trim(httpDfn.MethodDELETE.APICalls[i], cStrSpace+cStrTab+cStrNewLine)
	}
	for i := range httpDfn.MethodUPDATE.APICalls {
		httpDfn.MethodUPDATE.APICalls[i] = strings.Trim(httpDfn.MethodUPDATE.APICalls[i], cStrSpace+cStrTab+cStrNewLine)
	}

	return &httpDfn, err
}

//HTTPActionMap HTTPActionDefn storage class.
type HTTPActionMap struct {
	actMap map[string]*HTTPActionDefn //make this internal to let no outside threads change it.
}

//GetData  Cannot guarantee the lower layer is readonly.
func (hm HTTPActionMap) GetData() map[string]*HTTPActionDefn {
	return hm.actMap
}
