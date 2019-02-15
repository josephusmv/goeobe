package eobetests

import (
	"eobe/eobeapi"
	"eobe/eobecltmgmt"
	"eobe/eobedb"
	"eobe/eobeload"
	"fmt"
)

type apitester struct {
	haMap   map[string]*eobeload.HTTPActionDefn
	daMap   map[string]*eobedb.QueryDefn
	ucMap   map[string]string
	dbQry   eobedb.DBQueryInf
	cm      *eobecltmgmt.ClientManager
	fileMap map[string][]byte
	result  eobeapi.CallSeqResult
}

func newAPITester(rootpath string, cm *eobecltmgmt.ClientManager, dbQry eobedb.DBQueryInf) (*apitester, error) {
	df := rootpath + cStrTestDBActionFile
	hf := rootpath + cStrTestHTTPActionFile
	uf := rootpath + cStrTestUserConstFile
	lf := eobeload.NewLoaderFactory()
	daMap, haMap, uaMap, err := lf.LoadAll(df, hf, uf)
	if err != nil {
		fmt.Printf("Load All test data error: %s", err.Error())
		return nil, err
	}

	ok := randomVerifyAllData(daMap, haMap, uaMap)
	if !ok {
		return nil, err
	}

	at := &apitester{haMap: haMap.GetData(), daMap: daMap.GetData(), ucMap: uaMap.GetData()}
	at.cm = cm

	at.dbQry = dbQry

	return at, nil
}

func (at *apitester) runHTTPRequest(actionName string, rqstKeyValueMap map[string]string, inputInf testInputInf) bool {
	ok := true

	haDfn, found := at.haMap[actionName]
	if !found {
		fmt.Printf("\n *** ERROR: Cannot found Action [%s] from Server Resources map: %v.\n", actionName, at.haMap)
		return false
	}

	fmt.Printf("\n****************Run Action [%s]****************\n\t\tQuery Values:\n", haDfn.ActionName)
	for k, v := range rqstKeyValueMap {
		fmt.Printf("\t\t\t%s:%s\n", k, v)
	}
	fmt.Printf("\t\tExpected Response type: %s\n", haDfn.ExpectedResponse.ExpectedFMT)
	fmt.Printf("\t\tExpected Response file: %s\n", haDfn.ExpectedResponse.HTMLFile)

	fmt.Printf("\t\tAPI Sequence:\n")
	for i, v := range haDfn.APICalls {
		fmt.Printf("\t\t\tstep-%d:%s\n", i, v)
	}

	ok = at.doRunTest(haDfn, rqstKeyValueMap, inputInf)

	fmt.Printf("\n****************Run Action [%s] Done.****************\n\n", haDfn.ActionName)
	return ok
}

func (at *apitester) setFileBytes(files map[string][]byte) {
	at.fileMap = files
}

func (at *apitester) doRunTest(haDfn *eobeload.HTTPActionDefn, rqstKeyValueMap map[string]string, inputInf testInputInf) bool {
	hrld := at.cm.NewHerald()

	hrld.NewAccess("127.0.0.1", nil, nil, nil)

	//Run API call sequence
	cs := eobeapi.GetCallSequence(nil)
	if cs == nil {
		fmt.Printf("\n *** ERROR: eobeapi.GetCallSequence error:%s\n", "cs is nil")
		return false
	}

	cs.InitAPIFactory(at.dbQry, at.daMap, at.ucMap, hrld)

	var err error
	if err != nil {
		fmt.Printf("\n *** ERROR: CallSeq.PrepareForAPICalls error:%s\n", err.Error())
		return false
	}

	cs.SetFileMap(at.fileMap)

	at.result, err = cs.Execute(rqstKeyValueMap, haDfn.APICalls)
	if err != nil {
		fmt.Printf("\n *** ERROR: CallSeq.Execute error:%s\n", err.Error())
		return false
	}

	at.dumpCallSeqResult(at.result)
	hrld.FinishAccess()

	return true
}

func (at *apitester) verifyResult(actionName string, expected TstOutInf) bool {
	fmt.Printf("\n**************** Verify Results for %s. \n", actionName)
	expectedResult := expected.GetOutput(actionName)

	if expectedResult == nil {
		fmt.Printf("\n *** ERROR: Nil expected result\n")
		return false
	}

	for ek, ev := range expectedResult.SingleRow {
		rv, found := at.result.SingleRow[ek]

		if ev == "MUSTIGNORE" || rv == "MUSTIGNORE" {
			continue
		}
		if !found || rv != ev {
			fmt.Printf("\n *** ERROR: Verify SingleRow failed, expect: [%s], got: [%s]\n", rv, ev)
			return false
		}
		fmt.Printf("\t\t --> verify for SingleRow %s: %s, ok.\n", ek, ev)
	}
	fmt.Printf("\t --> verify for SingleRow, ok.\n")

	ok := compareTwoStringSlice(expectedResult.MNames, at.result.MNames)
	if !ok {
		fmt.Printf("\n *** ERROR: Verify MNames failed.\n")
		return false
	}
	fmt.Printf("\t --> verify for MNames, ok.\n")

	fmt.Printf("\t --> verify for MultiRow: \n")
	for i, exptvalues := range expectedResult.MultiRow {
		ok = compareTwoStringSlice(exptvalues, at.result.MultiRow[i])
		if !ok {
			fmt.Printf("\n *** ERROR: Verify MultiRow failed.\n")
			return false
		}
		fmt.Printf("\t\t --> verify for row %d: %v, ok.\n", i, at.result.MultiRow[i])
	}
	fmt.Printf("\t --> verify for MultiRow, ok.\n")

	fmt.Printf("\n**************** Verify Results for %s DONE \n\n\n ", actionName)
	return true
}

func (at *apitester) finishTests() {

}

func (at *apitester) dumpCallSeqResult(rslts eobeapi.CallSeqResult) {
	fmt.Println("\n   .....................Dump Results: ")
	fmt.Println("\n\t --------> SingleRow: ")
	for k, v := range rslts.SingleRow {
		fmt.Printf("\t\t---> SingleRow[%s]%s\n", k, v)
	}

	fmt.Println("\n\t --------> MultiRow: ")
	for i, row := range rslts.MultiRow {
		fmt.Printf("\t\t---> Row:%d\n", i)
		for j, vv := range rslts.MNames {
			fmt.Printf("\t\t\t---> %s:%s\n", vv, row[j])
		}
	}

	fmt.Println("\n   .....................Dump results done. ")
}

func compareTwoStringSlice(src, dst []string) bool {
	if src == nil {
		return true
	}

	if len(dst) != len(src) {
		return false
	}

	if len(dst) == 0 {
		return true
	}

	for i, v := range src {
		if v == "MUSTIGNORE" || dst[i] == "MUSTIGNORE" {
			continue
		}
		if dst[i] != v {
			fmt.Printf("\n *** ERROR: Testing failed, expect: %s, got: %s\n", v, dst[i])
			return false
		}
	}

	return true
}
