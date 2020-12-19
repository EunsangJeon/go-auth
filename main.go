package main

import (
	"net/http"
)

type person struct {
	First string
}

func main() {
	// jsonExample()

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

}

func bar(w http.ResponseWriter, r *http.Request) {

}
