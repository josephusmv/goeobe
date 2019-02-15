package eobehttp

import (
	"html/template"
	"io/ioutil"
	"os"
)

type resourceReaderInf interface {
	findFile(string) (bool, []byte)
	openDiskFile(string) (bool, []byte)
	getTemplates() *template.Template
	getTemplatesDev() map[string]string
	getPathes() (string, string)
	isFileExisted(string, HttpLoggerInf) (bool, bool) //should be no real Disk IO in real scenario
}

type resourcesManager struct {
	rootPath   string
	workFolder string
	sourceFileManager
	mediaFileManager
	templatesManager
	templatesDevManager
}

func (rmgr *resourcesManager) readResources(rootPath, workFolder string, templates []string, logger HttpLoggerInf) {
	rmgr.rootPath = rootPath
	rmgr.workFolder = workFolder
	workPath := combineFullPath(rootPath, workFolder)
	if getPckConfig().PreloadStaticFiles {
		rmgr.readSrcFiles(workPath, logger)
		rmgr.readMediaFiles(workPath, logger)
	}

	if getPckConfig().DevMode {
		rmgr.readTemplatesFilesDev(templates, logger)
	} else {
		rmgr.readTemplatesFiles(templates, logger)
	}
}

func (rmgr *resourcesManager) findFile(fname string) (bool, []byte) {
	if rmgr.isMediafile(fname) {
		fbytes, found := rmgr.mediaFilesMap[fname]
		return found, fbytes
	}

	if rmgr.isSourcefile(fname) {
		fbytes, found := rmgr.srcFilesMap[fname]
		return found, fbytes
	}

	return false, nil
}

//Try not call this, security leaking.
func (rmgr *resourcesManager) openDiskFile(fname string) (bool, []byte) {
	fbytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return false, nil
	}

	return true, fbytes
}

func (rmgr *resourcesManager) getTemplates() *template.Template {
	return rmgr.htmlTemplates
}

func (rmgr *resourcesManager) getTemplatesDev() map[string]string {
	return rmgr.tmpltMap
}

func (rmgr *resourcesManager) getPathes() (rootPath string, workFolder string) {
	return rmgr.rootPath, rmgr.workFolder
}

func (rmgr *resourcesManager) isFileExisted(fname string, logger HttpLoggerInf) (bool, bool) {
	stat, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return false, false
	}

	if stat == nil {
		logger.TraceError("Get File State Error: %s", err.Error())
		return false, false
	}

	if stat.IsDir() {
		return true, true
	}

	return true, false
}
