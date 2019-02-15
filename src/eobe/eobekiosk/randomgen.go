package eobekiosk

import (
	"bytes"
	"fmt"
	mrand "math/rand"
	"strconv"
	"time"
)

type randomGenerator struct {
}

func (rg randomGenerator) generateRandomBytes(byteLen int) []byte {
	randomStrLen := byteLen - strconv.IntSize/8
	now := time.Now()
	mrand.Seed(time.Now().UTC().UnixNano())
	keyStr := rg.randomString(randomStrLen) + "%d"
	keyStr = fmt.Sprintf(keyStr, now.Nanosecond())

	output := []byte(keyStr)
	redBytesCount := len(keyStr) - byteLen
	if redBytesCount > 0 {
		output = output[redBytesCount:]
	}

	return []byte(output)
}

func (rg randomGenerator) randomString(len int) string {
	var result bytes.Buffer
	for i := 0; i < len; i++ {

		upperCharInt := 65 + mrand.Intn(90-65)
		lowerCharInt := 97 + mrand.Intn(122-97)
		punctList := []int{'@', '#', '_'}

		switch mrand.Intn(30) {
		case 0:
			result.WriteString(string(punctList[0]))
		case 1:
			result.WriteString(string(punctList[1]))
		case 3:
			result.WriteString(string(punctList[2]))
		case 5, 6, 7, 8, 9, 10, 11:
			result.WriteString(string(upperCharInt))
		default:
			result.WriteString(string(lowerCharInt))
		}
	}
	return result.String()
}
