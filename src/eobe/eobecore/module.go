package eobecore

import (
	"eobe/eobedb"
	"eobe/eobeload"
	"strings"
)

type httpModule struct {
	name, rootPath  string
	indxActionsName string
	haMap           map[string]*eobeload.HTTPActionDefn
	daMap           map[string]*eobedb.QueryDefn
	ucMap           map[string]string
}

func (hm *httpModule) loadAll(haFile, daFile, ucFile string) error {
	path := strings.Trim(hm.rootPath, cStrSlash) + cStrSingleSlash + strings.ToLower(hm.name) + cStrSingleSlash
	hf := path + haFile
	df := path + daFile
	uf := path + ucFile

	lfct := eobeload.NewLoaderFactory()
	daSt, haSt, ucSt, err := lfct.LoadAll(df, hf, uf)
	if err != nil {
		return err
	}

	hm.daMap = daSt.GetData()
	hm.haMap = haSt.GetData()
	hm.ucMap = ucSt.GetData()

	return nil
}
