package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func jsonExample() {
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
