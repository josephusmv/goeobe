package eobehttp

import (
	"strings"
)

//parse results
const (
	cIntFileFound      = 0x0001
	cIntRedirectIndex  = 0x0002
	cIntTryRespFetcher = 0x0003
)

type urlParser struct {
	page404FName string
	indexAction  string
	rr           resourceReaderInf
}

//parseURL the valid URL is <workfolder>/resources
func (up urlParser) parseURL(rawURL string, logger HttpLoggerInf) (string, int) {
	rootPath, workFolder := up.rr.getPathes() //top level root path and working folder name

	//root path: should do redirect...
	if rawURL == cStrSlash || rawURL == "" {
		return up.redirectToIndexAction()
	}

	var localPath string

	//Directly try to find rawURL from root file system
	localPath = rootPath + rawURL
	valid, isdir := up.rr.isFileExisted(localPath, logger)
	if valid && !isdir {
		return localPath, cIntFileFound
	}
	if valid && isdir { //meaningless to access a local folder, redirect to work folder's index.html
		return up.redirectToIndexAction()
	}

	//raw URL is started with work folder
	localPath = strings.Trim(rawURL, cStrSlash)
	if strings.HasPrefix(localPath, workFolder) {
		resPath := combineFullPath(rootPath, rawURL) //file path in the local FS
		valid, isdir := up.rr.isFileExisted(resPath, logger)
		if valid && !isdir {
			return resPath, cIntFileFound
		}
		if valid && isdir { //Raw folder access is not acceptable, redirect to work folder's index.html
			return up.redirectToIndexAction()
		}
	}

	//Assuming that raw URL is a resources file name(with sub-folder names) under workpath
	workPath := combineFullPath(rootPath, workFolder)
	localPath = combineFullPath(workPath, rawURL)
	valid, isdir = up.rr.isFileExisted(localPath, logger)
	if valid && !isdir {
		return localPath, cIntFileFound
	}
	if valid && isdir { //Raw folder access is not acceptable, redirect to work folder's index.html
		return up.redirectToIndexAction()
	}

	//All attempts fails, let respFetcher try...
	//HTTP pkg should not try to modify this rawURL again.
	//let respFetcher find target from it!!!!
	return rawURL, cIntTryRespFetcher
}

func (up urlParser) redirectToIndexAction() (rediURL string, code int) {
	return up.indexAction, cIntRedirectIndex
}

func (up urlParser) serachResInDir(resPath string, srchPath string) (string, bool, bool) {
	var absolutePath string
	var found, isdir bool

	return absolutePath, found, isdir
}
