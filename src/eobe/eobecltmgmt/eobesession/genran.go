package eobesession

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

type genRandom struct {
}

func (gr genRandom) getRandomString(length int) string {
	var result bytes.Buffer
	for i := 0; i < length; i++ {
		upperCharInt := 65 + rand.Intn(90-65)
		result.WriteString(string(upperCharInt))
	}
	return result.String()
}

func (gr genRandom) getRandomSID(timeSuffix time.Time) string {
	rand.Seed(timeSuffix.UTC().UnixNano() + time.Now().UnixNano())
	return fmt.Sprintf("%s%d", gr.getRandomString(8), time.Now().Nanosecond())
}

func (gr genRandom) genCientIDByIPAndSID(ip, sid string) string {
	hashID := ip + sid

	return b64.StdEncoding.EncodeToString([]byte(hashID))
}

/*
func (gr genRandom) isHexDigit(b byte) bool {
	if b >= 0x30 && b <= 0x39 {
		return true
	}

	if b >= 0x41 && b <= 0x46 {
		return true
	}

	if b >= 0x61 && b <= 0x66 {
		return true
	}

	return false
}
*/
