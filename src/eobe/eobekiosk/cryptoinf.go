package eobekiosk

import (
	"log"
)

type cryptoInf interface {
	generateKey(keyPath string, logger *log.Logger) (key []byte, err error)
	encrypt(input string, key []byte) (output string, err error)
	decrypt(input string, key []byte) (output string, err error)
}

func newCryptoImpl(algo string) cryptoInf {
	switch algo {
	case "aes":
		return &cryptoAES{}
	default:
		return nil
	}
}
