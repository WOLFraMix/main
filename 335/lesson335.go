package main

import (
	"encoding/json"
	"fmt"
)

// Пример сложной структуры с тегами json
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age,omitempty"`
	Address struct {
		City string `json:"city,omitempty"`
		Zip  string `json:"zip,omitempty"`
	} `json:"address"`
	Hobbies []string `json:"hobbies,omitempty"`
}

func main() {
	// Сериализация структуры в JSON
	p := Person{
		Name:    "Alice",
		Hobbies: []string{"reading", "traveling"},
	}

	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	// {"name":"Alice","address":{},"hobbies":["reading","traveling"]}

	// Десериализация JSON в структуру
	var q Person
	err = json.Unmarshal(data, &q)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", q)
	// {Name:Alice Age:0 Address:{City: Zip:} Hobbies:[reading traveling]}
}
