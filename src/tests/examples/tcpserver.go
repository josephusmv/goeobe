package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "8081"
	CONN_TYPE = "tcp"
)

/* ******************** TCP ******************************** */

func serverRunTCP() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Key Server Listening TCP connection on " + CONN_HOST + ":" + CONN_PORT)

	//main accept loop
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	//simple read
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		panic(err)
	}

	recvMsg := string(buffer)

	//eg: enc\ncontent\n
	inx := strings.Index(recvMsg, "\n")
	recvMsg = recvMsg[inx+1:]
	inx = strings.Index(recvMsg, "\n")
	recvMsg = recvMsg[:inx]

	recvMsg = strings.Trim(recvMsg, "\n")
	//fmt.Printf("Message from client: [%s]\n", recvMsg)

	//send response
	resp := recvMsg + "_enc\n"
	conn.Write([]byte(resp))
	conn.Close()
}
