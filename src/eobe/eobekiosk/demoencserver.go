package eobekiosk

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

const cStrLogFileName = "DemoEncServer.log"
const cStrLogPrefix = "ENC Server: "

type DemoEncServer struct {
	host, port, protocal string
	logger               *log.Logger
	key                  []byte
	crptyo               cryptoInf
}

//NewDemoEncServer Get a new DemoEncServer
func NewDemoEncServer(host, port, protocal, logdir string) *DemoEncServer {
	path := logdir + "/" + cStrLogFileName
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	logger := log.New(file, cStrLogPrefix, log.Ldate|log.Ltime)

	es := DemoEncServer{host: host, port: port, protocal: protocal, logger: logger}

	return &es
}

//InitCrypto Init crypto and Key
func (es *DemoEncServer) InitCrypto(algo string, keyPath string) (err error) {
	es.crptyo = newCryptoImpl(algo)
	es.key, err = es.crptyo.generateKey(keyPath, es.logger)
	return err
}

func (es *DemoEncServer) Run() {
	es.serverRun()
}

func (es *DemoEncServer) serverRun() {
	// Listen for incoming connections.
	l, err := net.Listen(es.protocal, es.host+":"+es.port)
	if err != nil {
		es.logger.Printf("Error listening: %s, program exit", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	// Close the listener when the application closes.
	es.logger.Printf("Key Server Listening TCP connection on " + es.host + ":" + es.port)

	//main accept loop
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			es.logger.Printf("Error accepting: %s", err.Error())
			break
		}
		// Handle connections in a new goroutine.
		go es.handleRequest(conn)
	}
}

func (es *DemoEncServer) handleRequest(conn net.Conn) {
	quit := false
	for !quit {
		//simple read
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			es.logger.Printf("Error when read from connection: %s", err.Error())
			break
		}

		buffer = bytes.Trim(buffer, "\x00")

		//eg: enc\ncontent\n
		var trgtMsg string
		var cmdMsg string
		indCmd := bytes.IndexAny(buffer, "\n")
		if indCmd > 0 && indCmd < len(buffer) {
			cmdMsg = string(buffer[:indCmd])
			trgtMsg = string(buffer[indCmd+1:])
		} else {
			trgtMsg = string(buffer)
		}
		es.logger.Printf("Kiosk: Got input: [%s]", trgtMsg)
		es.logger.Printf("Kiosk: Got input: [%v]", es.key)
		var resp string
		var cryptoErr error
		switch cmdMsg {
		case "enc":
			resp, cryptoErr = es.crptyo.encrypt(trgtMsg, es.key)
		case "dec":
			resp, cryptoErr = es.crptyo.decrypt(trgtMsg, es.key)
		default:
			resp = "ack"
			quit = true
		}

		if cryptoErr != nil {
			es.logger.Printf("Error When %s message: %s, error: %s", cmdMsg, trgtMsg, cryptoErr)
			return
		}

		//send response
		es.logger.Printf("WriteBack: <%s>", string(resp))
		resp = resp + "\n"

		n, errW := conn.Write([]byte(resp))
		if n != len(resp) || errW != nil {
			es.logger.Printf("Write %d bytes, error: %s", n, errW)
		}
	}

	conn.Close()
}
