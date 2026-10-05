package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	name := ""
	age := 0
	address := ""

	fmt.Scan(&name, &age, &address)
	p := Person{
		Name:    name,
		Age:     age,
		Address: address,
	}

	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
