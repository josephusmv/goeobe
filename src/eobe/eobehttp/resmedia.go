package eobehttp

import "strings"

const cStrLoadMediaFilesForFolder = "Read media files under folder: %s."

type mediaFileManager struct {
	resReaderBase
	mediaFilesMap map[string][]byte
}

func (mfm *mediaFileManager) readMediaFiles(rootPath string, logger HttpLoggerInf) error {
	if mfm.mediaFilesMap == nil {
		mfm.mediaFilesMap = make(map[string][]byte)
	}
	logger.TraceDev(cStrLoadMediaFilesForFolder, rootPath)
	return mfm.readFile(rootPath, logger, mfm.isMediafile, mfm.saveMediaFileBytes)
}

//check is supported file suffix for this struct
func (mfm mediaFileManager) isMediafile(filename string) bool {
	if strings.HasSuffix(filename, cStrSuffixPNG) {
		return true
	}
	if strings.HasSuffix(filename, cStrSuffixJPG) {
		return true
	}
	if strings.HasSuffix(filename, cStrSuffixSVG) {
		return true
	}
	if strings.HasSuffix(filename, cStrSuffixICO) {
		return true
	}
	if strings.HasSuffix(filename, cStrSuffixBMP) {
		return true
	}
	if strings.HasSuffix(filename, cStrSuffixJPEG) {
		return true
	}

	return false
}

func (mfm *mediaFileManager) saveMediaFileBytes(fname string, fbytes []byte, logger HttpLoggerInf) {
	logger.TraceDev(cStrLoadMeidaFileSuccess, fname, len(fbytes))
	mfm.mediaFilesMap[fname] = fbytes
}
