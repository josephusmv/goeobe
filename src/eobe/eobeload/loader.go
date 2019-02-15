package eobeload

import (
	"fmt"
)

//Loader
type LoaderFactory struct {
	dbLder   *DBActionLoader
	httpLder *HTTPActionLoader
	ucLder   *UserConstsLoader
}

//NewLoaderFactory ...
func NewLoaderFactory() *LoaderFactory {
	var lf LoaderFactory
	lf.dbLder = NewDBActionLoader()
	lf.httpLder = NewHTTPActionLoader()
	lf.ucLder = NewUserConstsLoader()

	return &lf
}

//LoadDBActMap ...
func (lf *LoaderFactory) LoadDBActMap(fileName string) (*DBActionMap, error) {
	if lf.dbLder == nil {
		return nil, fmt.Errorf(cStrLoaderFactoryIError)
	}
	return lf.dbLder.LoadFromFile(fileName)
}

//LoadHTTPActMap ...
func (lf *LoaderFactory) LoadHTTPActMap(fileName string) (*HTTPActionMap, error) {
	if lf.httpLder == nil {
		return nil, fmt.Errorf(cStrLoaderFactoryIError)
	}
	return lf.httpLder.LoadFromFile(fileName)
}

//LoadUserCnstMap
func (lf *LoaderFactory) LoadUserCnstMap(fileName string) (*UserConstsMap, error) {
	if lf.ucLder == nil {
		return nil, fmt.Errorf(cStrLoaderFactoryIError)
	}
	return lf.ucLder.LoadFromFile(fileName)
}

//LoadAll ...
func (lf *LoaderFactory) LoadAll(df, hf, uf string) (*DBActionMap, *HTTPActionMap, *UserConstsMap, error) {
	if lf.dbLder == nil || lf.httpLder == nil || lf.ucLder == nil {
		return nil, nil, nil, fmt.Errorf(cStrLoaderFactoryIError)
	}

	daMap, dErr := lf.dbLder.LoadFromFile(df)
	if dErr != nil {
		return nil, nil, nil, fmt.Errorf(cStrLoaderFactoryGError, dErr.Error())
	}

	haMap, hErr := lf.httpLder.LoadFromFile(hf)
	if hErr != nil {
		return nil, nil, nil, fmt.Errorf(cStrLoaderFactoryGError, hErr.Error())
	}

	ucMap, uErr := lf.ucLder.LoadFromFile(uf)
	if uErr != nil {
		return nil, nil, nil, fmt.Errorf(cStrLoaderFactoryGError, uErr.Error())
	}

	return daMap, haMap, ucMap, nil
}

//GetErrorList ...
func (lf *LoaderFactory) GetErrorList() (errList []error) {
	if lf.dbLder == nil || lf.httpLder == nil || lf.ucLder == nil {
		return nil
	}

	errList = append(errList, lf.httpLder.errList...)
	errList = append(errList, lf.dbLder.errList...)
	errList = append(errList, lf.ucLder.errList...)

	return
}
