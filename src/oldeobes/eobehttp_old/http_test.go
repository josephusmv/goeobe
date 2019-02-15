package eobehttp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

/* All test cases:
 *	go test -v -run HttpSmokeGetURL
 *	go test -v -run Http404Page
 *	go test -v -run HttpPostSmokeTest
 */

const cTestStrhtmlRoot = "httptest_res"
const cTestStrIndexPath = cTestStrhtmlRoot + "/test/index.tmplt"
const cTestStrDynamicPath = cTestStrhtmlRoot + "/test/dYnAmic.tmplt"
const cTest404Path = cTestStrhtmlRoot + "/test/page404.html"
const cTestStrIndexno404Path = cTestStrhtmlRoot + "/testno404/index.html"
const cTestTempIndexActionName = "GetIndex"
const cTestTempDynamicActionName = "GetDynamicPage"

/* *************************************************************
 * Tests for GET Method
 * *************************************************************/
//TestHttpSmokeGetURL Smoke test for getting URL test, inlcudes:
//	1. Get root
//	2. Get Module root
//	3. Get static html
//	4. Get Dynamic not index.html
//	5. Get static css
//	6. Get JSON Values
//	go test -v -run HttpSmokeGetURL
func TestHttpSmokeGetURL(t *testing.T) {
	//Test 1 Get root
	////fmt.Println("\n\n******Test 1. Get root *******")
	//t.Logf("\n\n******Test output: Test 1. Get root  *******")
	doTestSmokeGetURLAndValidate("/", http.StatusOK, t)
	t.Logf("***** Test 1 Passed ********\n\n")

	//Test 2. Get Module root
	////fmt.Println("\n\n******Test	2. Get Module root *******")
	//t.Logf("\n\n******Test output:	2. Get Module root *******")
	doTestSmokeGetURLAndValidate("/tEsT", http.StatusOK, t)
	t.Logf("***** Test 2 Passed ********\n\n")

	//Test 3. Get static html
	//t.Logf("\n\n******Test output:	3. Get static html  *******")
	////fmt.Println("\n\n******Test	3. Get static html *******")
	doTestSmokeGetURLAndValidate("/tEst/stAtic.html", http.StatusOK, t)
	t.Logf("***** Test 3 Passed ********\n\n")

	//4. Get Dynamic not index.html
	////fmt.Println("\n\n******Test	4. Get Dynamic not index.html *******")
	//t.Logf("\n\n******Test output:	4. Get Dynamic not index.html   *******")
	doTestSmokeGetURLAndValidate("/tEst/GetDynamicPage", http.StatusOK, t)
	t.Logf("***** Test 4 Passed ********\n\n")

	//5. Get static css
	////fmt.Println("\n\n******Test	5. Get static css *******")
	//t.Logf("\n\n******Test output:	5. Get static css   *******")
	doTestSmokeGetURLAndValidate("/tEst/css/style.css", http.StatusOK, t)
	t.Logf("***** Test 5 Passed ********\n\n")

	//6. Get JSON Values
	////fmt.Println("\n\n******Test	6. Get JSON Values *******")
	//t.Logf("\n\n******Test output:	6. Get JSON Values   *******")
	doTestSmokeGetURLAndValidate("TEST/DoEchoTest", http.StatusOK, t)
	t.Logf("***** Test 6 Passed ********\n\n")

	////fmt.Println("\n\n******Test output done. *******")
	//t.Logf("\n\n****** Done: Test Passed.  *******")
}

func doTestSmokeGetURLAndValidate(url string, expectedCode int, t *testing.T) {
	actionList := []string{cTestTempIndexActionName, "DoEchoTest", cTestTempDynamicActionName}

	httpTemplates := []string{cTestStrIndexPath, cTest404Path, cTestStrDynamicPath}
	reqSvr := createRequestServer("TEST", httpTemplates, actionList, t)
	code := DoTestGetMethod(url, reqSvr, t)
	if code != expectedCode {
		t.Errorf("Failed!!!!! handler returned wrong status code: got %v want %v",
			code, expectedCode)
		t.Fail()
	}
}

//TestHttp404Page Smoke test for data exchange
//	go test -v -run Http404Page
func TestHttp404Page(t *testing.T) {
	actionList := []string{"DoEchoTest"}
	//Test invalid page when has 404 templates
	////fmt.Println("\n\n******Test 1 log output *******")
	httpTemplatesWith404Page := []string{cTestStrIndexPath, cTest404Path}
	reqSvr := createRequestServer("TEST", httpTemplatesWith404Page, actionList, t)
	code := DoTestGetMethod("/InvalidPage", reqSvr, t)
	if code != http.StatusNotFound {
		t.Errorf("Failed!!!!! handler returned wrong status code: got %v want %v",
			code, http.StatusNotFound)
		t.Fail()
	}

	//Test invalid page when not has 404 templates
	////fmt.Println("\n\n******Test 2 log output *******")
	httpTemplatesWithOut404Page := []string{cTestStrIndexno404Path}
	reqSvr = createRequestServer("TESTNO404", httpTemplatesWithOut404Page, actionList, t)
	code = DoTestGetMethod("/InvalidPage", reqSvr, t)
	if code != http.StatusNotFound {
		t.Errorf("Failed!!!!! handler returned wrong status code: got %v want %v",
			code, http.StatusNotFound)
		t.Fail()
	}
	////fmt.Println("\n\n******End of Test*******")
}

func DoTestGetMethod(url string, reqSvr *RequestServer, t *testing.T) int {
	req := createGetRequest(url, t)
	rsp := httptest.NewRecorder()
	//run test
	handler := http.Handler(reqSvr)
	handler.ServeHTTP(rsp, req)
	// Check the response body is what we expect.
	//t.Logf("\nRespons code: \n\t[%v]", rsp.Code)
	//t.Logf("\nRespons Header: \n\t[%v]", rsp.Header())
	//t.Logf("\nRespons Body: \n\t[%v]", rsp.Body.String())

	return rsp.Code
}

func createGetRequest(url string, t *testing.T) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("Key...%d", i)
		value := fmt.Sprintf("value...%d", i)
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	return req
}

func createRequestServer(module string, templates, actionList []string, t *testing.T) *RequestServer {
	//Make Request Server
	contentType := CStrExpectedHTML
	body := "ok smoke test"
	rspFcther := RespFetcher{tst: t, ContentType: contentType, body: body}

	var rfInf RespFetcherInf
	rfInf = rspFcther
	rs := NewRequestServer(module, cTestStrhtmlRoot, cTestTempIndexActionName, &rfInf, nil)
	rs.Init(templates, actionList)

	return rs
}

type RespFetcher struct {
	tst         *testing.T
	HTMLfile    string
	ContentType string
	body        string
}

func (rf RespFetcher) FetchResponse(req RequestData) (rsp ResponseData, err error) {
	rf.tst.Logf("\nFetchResponse() : \n\t\tRequest data get, content:\n\t\t %v", req)
	rsp.ContentType = rf.ContentType
	rsp.Body = []byte(rf.body)
	rsp.HTMLfile = rf.HTMLfile

	var tp TemplateData
	tp.Data = [][]string{
		[]string{"Value1111", "Value1112", "Vao & bar斯卡拉lue1113"},
		[]string{"V卡alue1222", "Value2222", "Value22卡拉23"},
		[]string{"Value3卡拉331", "Value3卡332", "Value卡拉3333"},
	}

	rsp.Body = make([]byte, tp.GetSize())
	tp.Read(rsp.Body)

	switch req.QueryTarget {
	case cTestTempIndexActionName:
		rsp.HTMLfile = "index.tmplt"
		rsp.ContentType = cStrTextHTML
	case cTestTempDynamicActionName:
		rsp.HTMLfile = "dYnAmic.tmplt"
		rsp.ContentType = cStrTextHTML

	case "DoEchoTest":
		fallthrough
	default:
		rsp.HTMLfile = ""
		rsp.ContentType = cStrAppJSON
	}
	return
}

/* *************************************************************
 * Tests for POST Method
 * *************************************************************/
//TestHttpPostSmokeTest Smoke test for data exchange
//	go test -v -run HttpPostSmokeTest
func TestHttpPostSmokeTest(t *testing.T) {
	DoTestPostMethod("TEST/DoEchoTest", http.StatusOK, t)
}

func DoTestPostMethod(url string, expectedCode int, t *testing.T) {
	actionList := []string{cTestTempIndexActionName, "DoEchoTest", cTestTempDynamicActionName}
	httpTemplates := []string{cTestStrIndexPath, cTest404Path, cTestStrDynamicPath}

	reqSvr := createRequestServer("TEST", httpTemplates, actionList, t)

	req := createPostRequest(url, t)
	rsp := httptest.NewRecorder()

	//run test
	handler := http.Handler(reqSvr)
	handler.ServeHTTP(rsp, req)
	// Check the response body is what we expect.
	//t.Logf("\nRespons code: \n\t[%v]", rsp.Code)
	//t.Logf("\nRespons Header: \n\t[%v]", rsp.Header())
	//t.Logf("\nRespons Body: \n\t[%v]", rsp.Body.String())

	if rsp.Code != expectedCode {
		t.Errorf("Failed!!!!! handler returned wrong status code: got %v want %v",
			rsp.Code, expectedCode)
		t.Fail()
	}
}

func createPostRequest(urlStr string, t *testing.T) *http.Request {
	req, err := http.NewRequest("POST", urlStr, nil)
	if err != nil {
		t.Fatal(err)
	}

	Values := []string{"Value1abcde", "foo & bar斯卡拉夫"}

	req.PostForm = url.Values{}
	req.PostForm.Add("resp1", Values[0])
	req.PostForm.Add("resp2", Values[1])
	req.PostForm.Encode()

	return req
}
