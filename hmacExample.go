package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

var key []byte

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)

	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error in signMessage while hashing message: %w", err)
	}

	signature := h.Sum(nil)

	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSig while getting signature of msg: %w", err)
	}

	same := hmac.Equal(newSig, sig)

	return same, nil
}

func usingHMACExample() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	testMsg := []byte("Hello, HMAC!")
	log.Println("testMsg: ", base64.StdEncoding.EncodeToString(testMsg))
	log.Println("testMsg: ", hex.EncodeToString(testMsg))
	testMsgSig, err := signMessage(testMsg)
	if err != nil {
		log.Println(err)
	}
	log.Println("testMsg: ", base64.StdEncoding.EncodeToString(testMsg))
	log.Println("testMsg: ", hex.EncodeToString(testMsg))
	log.Println("testMsgSig: ", base64.StdEncoding.EncodeToString(testMsgSig))
	log.Println("testMsg: ", hex.EncodeToString(testMsgSig))
}
