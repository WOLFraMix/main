package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Alice", Age: 25}
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%s\n", jsonBytes) // -> {"Name":"Alice","Age":25}

	jsonString := `{"name":"Bob","age":30}`
	err = json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(person.Name, person.Age) // -> Bob 30
}
