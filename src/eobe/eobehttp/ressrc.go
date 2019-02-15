package eobehttp

import (
	"io/ioutil"
	"strings"
)

const cStrLoadSrcFilesForFolder = "Read source files under folder: %s."

type sourceFileManager struct {
	resReaderBase
	srcFilesMap map[string][]byte
}

func (sfm *sourceFileManager) readSrcFiles(rootPath string, logger HttpLoggerInf) error {
	if sfm.srcFilesMap == nil {
		sfm.srcFilesMap = make(map[string][]byte)
	}
	logger.TraceDev(cStrLoadSrcFilesForFolder, rootPath)
	err := sfm.readFile(rootPath, logger, sfm.isSourcefile, sfm.saveSrcFileBytes)
	if err != nil {
		return err
	}

	//load 404 file
	pkgcfg := getPckConfig()
	if pkgcfg.Page404FilePath != "" {
		fdata, err := ioutil.ReadFile(pkgcfg.Page404FilePath)
		if err != nil {
			logger.TraceDev(cStrInvalidPage404FileError, pkgcfg.Page404FilePath, err.Error())
			return err
		}
		sfm.saveSrcFileBytes(pkgcfg.Page404FilePath, fdata, logger)
	}

	return nil
}

//check is supported file suffix for this struct
func (sfm sourceFileManager) isSourcefile(filename string) bool {
	if strings.HasSuffix(filename, cStrSuffixCSS) {
		return true
	}
	if strings.HasSuffix(filename, cStrSuffixJS) {
		return true
	}
	if strings.HasSuffix(filename, cStrSuffixHTML) {
		return true
	}

	return false
}

func (sfm *sourceFileManager) saveSrcFileBytes(fname string, fbytes []byte, logger HttpLoggerInf) {
	logger.TraceDev(cStrLoadResFileSuccess, fname, len(fbytes))
	sfm.srcFilesMap[fname] = fbytes
}
