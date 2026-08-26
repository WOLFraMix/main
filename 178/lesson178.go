package main

import (
	"errors"
	"fmt"
	"log"
)

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

// Функция-конструктор
func NewPerson(firstname string, lastname string, age int) (*Person, error) {
	if age <= 0 {
		return nil, errors.New("age must be > 0")
	}

	return &Person{
		Firstname: firstname,
		Lastname:  lastname,
		Age:       age,
	}, nil
}

func main() {
	p1 := Person{Firstname: "Stepan", Age: 27}
	fmt.Printf("%+v\n", p1)

	p2, err := NewPerson("Slim", "Shady", 37)
	if err != nil {
		log.Fatalf("unable to create person: %v", err)
	}
	fmt.Printf("%+v\n", p2)
}
