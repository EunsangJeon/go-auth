package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "John",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encoded bad data:", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("Decoded bad data:")
	}

	log.Println("Person:", p1)
}

func usingJSONExample() {
	p1 := person{
		First: "John",
	}

	p2 := person{
		First: "Paul",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Print JSON", string(bs))

	xp2 := []person{}

	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Back to Go data struct", xp2)
}
