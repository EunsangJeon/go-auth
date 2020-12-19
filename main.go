package main

import (
	"encoding/base64"
	"fmt"
)

type person struct {
	First string
}

func main() {
	// jsonExample()

	// http.HandleFunc("/encode", foo)
	// http.HandleFunc("/decode", bar)
	// http.ListenAndServe(":8080", nil)

	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
}
