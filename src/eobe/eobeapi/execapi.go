package eobeapi

import (
	"eobe/eobehttp"
)

const cStrRunsAPI = "Run API: %s"

type callExec struct {
	fileMap map[string][]byte
	logger  eobehttp.HttpLoggerInf
}

func (ce callExec) runExecute(callsList []*apiDefine, qryKVMap map[string]string) (rslts CallSeqResult) {
	//Make the intermediate result
	tempResults := CallSeqResult{}
	tempResults.SingleRow = make(map[string]string)

	//Execute call sequence
	for _, apiCall := range callsList {
		ce.logger.TraceDev(cStrRunsAPI, apiCall.apiName)

		apiCall.doStepPreAction(ce, tempResults)

		stepRslt, apiErr := apiCall.implementation.RunAPI(qryKVMap, tempResults.SingleRow)
		if apiErr.HasError() {
			rslts.SingleRow = stepRslt //Error may need some error info stored in result returned
			rslts.ApiErr = apiErr
			rslts.MNames = nil
			rslts.MultiRow = nil
			return rslts
		}

		//save temp results, this should be done for all steps
		for k, v := range stepRslt {
			tempResults.SingleRow[k] = v
		}
		tempResults.MNames, tempResults.MultiRow = apiCall.implementation.GetResultRows()
		tempResults.ApiErr = apiErr

		//Modify tempResults for some action, and modify the final result for return
		//	use step result as result.
		rslts = apiCall.doStepPostAction(ce, &rslts, stepRslt)
	}

	return rslts
}
