package eobehttp

import (
	"html/template"
	"os"
	"strings"
)

const cStrLoadTmpltFiles = "Read template files, expect total: %d files."

type templatesManager struct {
	htmlTemplates    *template.Template
	htmlTemplatesStr []string
}

func (tm *templatesManager) readTemplatesFiles(templates []string, logger HttpLoggerInf) error {
	logger.TraceDev(cStrLoadTmpltFiles, len(templates))
	return tm.readTmplt(templates, logger)
}

//initHTMLTemplates Validate the given template is ok or not, to ensure there is no file missing during parse.
func (tm *templatesManager) readTmplt(templates []string, logger HttpLoggerInf) error {
	for _, v := range templates {

		if _, err := os.Stat(v); err != nil {
			logger.TraceError(cStrInvalidHTMLTemplateFileError, v)
		} else {
			tm.htmlTemplatesStr = append(tm.htmlTemplatesStr, v)
		}
		logger.TraceDev(cStrDEVLoadTmplateFile, v, v)
	}

	if len(tm.htmlTemplatesStr) <= 0 {
		return logger.TraceError(cStrErrorNoTemplates) //a Warning error
	}

	tmps, err := template.ParseFiles(tm.htmlTemplatesStr...)
	if tmps == nil || err != nil {
		return logger.TraceError(cStrErrorParseTemplates, err.Error())
	}

	tm.htmlTemplates = template.Must(tmps, err)
	if tm.htmlTemplates == nil {
		return logger.TraceError(cStrErrorMustTemplates)
	}
	return nil
}

/****************************************************************************/
const cStrLoadTmpltFilesDev = "Read template files for Dev mode, expect total: %d files."

type templatesDevManager struct {
	tmpltMap map[string]string
}

func (tm *templatesDevManager) readTemplatesFilesDev(templates []string, logger HttpLoggerInf) error {
	tm.tmpltMap = make(map[string]string)
	logger.TraceDev(cStrLoadTmpltFilesDev, len(templates))
	for _, fullPath := range templates {
		//do file exist verify first, to expose error earlier
		if _, err := os.Stat(fullPath); err != nil {
			logger.TraceError(cStrInvalidHTMLTemplateFileError, fullPath)
			continue
		}

		var fname string
		indx := strings.LastIndex(fullPath, cStrSlash)
		if indx >= 0 && indx < len(fullPath) {
			fname = fullPath[indx:]
			fname = strings.Trim(fname, cStrSlash)
		} else {
			fname = fullPath
		}

		tm.tmpltMap[fname] = fullPath

		logger.TraceDev(cStrDEVLoadTmplateFile, fullPath, fname)
	}
	return nil
}
