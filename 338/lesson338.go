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
	s := ""
	fmt.Scan(&s)

	var p Person
	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		fmt.Println("Некорректное значение json!")
		return
	}
	fmt.Println(p)
}
