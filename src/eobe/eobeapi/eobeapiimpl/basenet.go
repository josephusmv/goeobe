package eobeapiimpl

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type apinetbase struct {
	apiBase

	/* For this kind */

}

const cStrProtocalUDP = "udp"
const cStrProtocalTCP = "tcp"

func (api apinetbase) getResponseFromRemote(srvIP, srvPort string, contentStr string) (string, error) {
	connectStr := srvIP + cStrColon + srvPort
	conn, err := net.Dial(cStrProtocalTCP, connectStr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	//fmt.Println(contentStr)
	var n int
	n, err = fmt.Fprintf(conn, contentStr)
	if n != len(contentStr) || err != nil {
		return "", fmt.Errorf(cStrNetConnectionWriteError, contentStr)
	}

	reader := bufio.NewReader(conn)
	if reader == nil {
		return "", fmt.Errorf(cStrNetConnectionReadError, contentStr)
	}

	var resp string
	resp, err = reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	resp = strings.Trim(resp, cStrNewLine)
	resp = strings.TrimSpace(resp)

	return resp, nil
}
