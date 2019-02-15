package eobekiosk

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	CONN_HOST  = "127.0.0.1"
	CONN_PORT  = "8081"
	CONN_TYPE  = "tcp"
	LOG_DIR    = "./"
	LOG_DIRbad = "./notexist"
	KEY_FILE   = "./key"
)

//TestRunDemoServer  Run the demo Enc server.
//	go test -v -run RunDemoServer
func TestRunDemoServer(t *testing.T) {
	es := NewDemoEncServer(CONN_HOST, CONN_PORT, CONN_TYPE, LOG_DIR)
	es.InitCrypto(CStrAlgoAES, KEY_FILE)

	es.serverRun()
}

//TestDemoEncServer   Smoke test for DB query and exec
//	go test -v -cover -coverprofile cover.out  -run DemoEncServer
//	go tool cover -html=cover.out -o cover.html
func TestDemoEncServer(t *testing.T) {
	es := NewDemoEncServer(CONN_HOST, CONN_PORT, CONN_TYPE, LOG_DIRbad)
	if es == nil {
		es = NewDemoEncServer(CONN_HOST, CONN_PORT, CONN_TYPE, LOG_DIR)
	}
	es.InitCrypto(CStrAlgoAES, KEY_FILE)

	go es.serverRun()
	time.Sleep(time.Millisecond)
	runCltTests()
}

const TEST_RUNS = 100
const TEST_DELAYRUNS = 1

func runCltTests() {
	var wg sync.WaitGroup
	wg.Add(TEST_RUNS * TEST_DELAYRUNS)
	for k := 0; k < TEST_DELAYRUNS; k++ {
		for i := 0; i < TEST_RUNS; i++ {
			go cltConn(i, &wg)
		}
		time.Sleep(time.Millisecond * 200)
	}
	wg.Wait()
}

const cStrNetConnectionWriteError = "Error when write to server error: %s"
const cStrNetConnectionReadError = "Error when read for response error: %s"
const cStrNewLine = "\n"

func cltConn(index int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	oriStr := fmt.Sprintf("MSG_%s_%d", getGID(), index)
	contentStr := "enc\n" + oriStr

	encrStr := sendAndFetch(contentStr, conn)

	contentStr = "dec\n" + encrStr
	decrStr := sendAndFetch(contentStr, conn)

	ack := sendAndFetch("fin\n", conn)
	if ack != "ack" {
		panic(ack)
	}

	if decrStr != oriStr {
		fmt.Println("contentStr: <", oriStr, ">\n---> Length:", len(oriStr))
		fmt.Println("resp: <", decrStr, ">")
		panic(oriStr)
	}
	conn.Close()

	conn, err = net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	sendThenDirectClose("", conn)

	return
}

func getGID() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return string(b)
}

func sendAndFetch(rqst string, conn net.Conn) (resp string) {
	n, err := fmt.Fprintf(conn, rqst)
	if n != len(rqst) {
		fmt.Printf("write length mismatch.")
		panic(rqst)
	}
	if err != nil {
		fmt.Printf(cStrNetConnectionWriteError, err.Error())
		panic(rqst)
	}

	reader := bufio.NewReader(conn)
	if reader == nil {
		fmt.Printf(cStrNetConnectionReadError, err.Error())
		panic(rqst)
	}

	resp, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ReadString error: %s", err.Error())
		panic(rqst)
	}

	resp = strings.Trim(resp, cStrNewLine)
	resp = strings.TrimSpace(resp)

	return resp
}

func sendThenDirectClose(rqst string, conn net.Conn) {
	n, err := fmt.Fprintf(conn, rqst)
	if n != len(rqst) {
		fmt.Printf("write length mismatch.")
		panic(rqst)
	}
	if err != nil {
		fmt.Printf(cStrNetConnectionWriteError, err.Error())
		panic(rqst)
	}

	conn.Close()
}

//TestCryptoAES   Smoke test for DB query and exec
//	go test -v -run CryptoAES
func TestCryptoAES(t *testing.T) {
	const testStr = "f3qt#$^YWERt326ESRYQ%#^Y&*#@@$Q%GDFBG"
	var caes cryptoAES

	key, err := caes.generateKey("./key", &log.Logger{})
	if err != nil {
		panic(err)
	}

	result, err := caes.encrypt(testStr, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	result, err = caes.decrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	if result != testStr {
		fmt.Println(result, "\n", testStr)
		panic(result)
	}

	fmt.Println("Test Result Passed: ", result)
}

//TestDemoEncServerMultiRoutine
//	go test -v  -run DemoEncServerMultiRoutine
func TestDemoEncServerMultiRoutine(t *testing.T) {
	es := NewDemoEncServer(CONN_HOST, CONN_PORT, CONN_TYPE, LOG_DIR)
	es.InitCrypto(CStrAlgoAES, KEY_FILE)

	go es.serverRun()
	time.Sleep(time.Millisecond)

	var wg sync.WaitGroup
	wg.Add(2)
	go cltConnMulti(&wg)
	go cltConnMulti(&wg)
	wg.Wait()
}

func cltConnMulti(wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	oriStr := "Mickey"
	contentStr := "enc\n" + oriStr

	encrStr := sendAndFetch(contentStr, conn)
	fmt.Println(encrStr)

	contentStr = "dec\n" + encrStr
	decrStr := sendAndFetch(contentStr, conn)

	ack := sendAndFetch("fin\n", conn)
	if ack != "ack" {
		panic(ack)
	}

	if decrStr != oriStr {
		fmt.Println("contentStr: <", oriStr, ">\n---> Length:", len(oriStr))
		fmt.Println("resp: <", decrStr, ">")
		panic(oriStr)
	}
	conn.Close()

	return
}
