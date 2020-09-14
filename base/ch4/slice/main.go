package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name string "json:\"name\""
	Age  int    `json:"age,omitempty"`
}

func main() {
	p1 := person{
		Name: "ll",
		Age:  0,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(b))
}
