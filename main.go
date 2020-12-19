package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var key []byte

func main() {
	pass := "123456789"
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = comparePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}

	log.Println("Logged in!")

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

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating hashed password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}

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
