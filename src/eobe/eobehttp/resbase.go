package eobehttp

import "io/ioutil"

type canAccept func(string) bool
type saveFileBytes func(string, []byte, HttpLoggerInf)

type resReaderBase struct {
}

func (rb *resReaderBase) readFile(rootPath string, logger HttpLoggerInf, accept canAccept, save saveFileBytes) error {
	files, derr := ioutil.ReadDir(rootPath)
	if derr != nil {
		return logger.TraceError(cStrReadModuleDirError, rootPath, derr)
	}

	for _, file := range files {
		filePath := combineFullPath(rootPath, file.Name())

		if file.IsDir() {
			rb.readFile(filePath, logger, accept, save) //ignore errors
			continue
		}

		//filter unsupported file types.
		if !accept(file.Name()) {
			continue
		}

		if file.Size() <= 0 {
			continue
		}

		fdata, err := ioutil.ReadFile(filePath)
		if err != nil {
			logger.TraceDev(cStrReadModuleDirError, file.Name(), err.Error())
			continue
		}

		save(file.Name(), fdata, logger)
	}

	return nil
}
