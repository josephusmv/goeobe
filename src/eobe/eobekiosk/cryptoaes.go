package eobekiosk

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const CIntBytesLen = 32
const CStrAlgoAES = "aes"

type cryptoAES struct {
	randomGenerator

	gcm cipher.AEAD
}

func (caes *cryptoAES) generateKey(keyPath string, logger *log.Logger) ([]byte, error) {

	var newKey []byte
	kf, err := os.OpenFile(keyPath, os.O_RDONLY, 0666)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if err == nil { //read key from key file.
		newKey, err = ioutil.ReadFile(keyPath)
		if err != nil {
			kf.Close()
			return nil, err
		}
	}

	//cannot defer this, need to release file.
	kf.Close()

	if newKey == nil {
		newKey = caes.generateRandomBytes(CIntBytesLen)
		err = ioutil.WriteFile(keyPath, newKey, 0666)
		if err != nil {
			logger.Printf("Write key to file %s error: %s", keyPath, err.Error())
			return nil, err
		}
	}

	var c cipher.Block
	c, err = aes.NewCipher(newKey)
	if err != nil {
		logger.Println("Error When Create New Cipher")
		return nil, err
	}

	caes.gcm, err = cipher.NewGCM(c)
	if err != nil {
		logger.Println("Error When Create New GCM")
		return nil, err
	}

	return newKey, err
}

func (caes *cryptoAES) encryptBytes(input, key []byte) (output []byte, err error) {
	nonce := make([]byte, caes.gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return caes.gcm.Seal(nonce, nonce, input, nil), nil
}

func (caes *cryptoAES) decryptBytes(input, key []byte) (output []byte, err error) {
	nonceSize := caes.gcm.NonceSize()
	if len(input) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := input[:nonceSize], input[nonceSize:]
	return caes.gcm.Open(nil, nonce, ciphertext, nil)
}

func (caes *cryptoAES) encrypt(input string, key []byte) (output string, err error) {
	encBytes, err := caes.encryptBytes([]byte(input), key)
	if err != nil {
		return "", err
	}

	output = base64.StdEncoding.EncodeToString(encBytes)

	return
}

func (caes *cryptoAES) decrypt(input string, key []byte) (output string, err error) {
	srcBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		fmt.Println(input)
		return "", fmt.Errorf("Invalid input for decrypt, more: %s", err.Error())
	}

	var reslt []byte
	reslt, err = caes.decryptBytes(srcBytes, key)

	return string(reslt), err
}
