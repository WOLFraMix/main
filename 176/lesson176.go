package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	person := Person{
		Firstname: "Майкл",
		Lastname:  "Джексон",
		Age:       50,
	}
	person.SayHello()
	fmt.Println(person.Struct())
}

func (p Person) SayHello() {
	fmt.Printf("%s приветствует вас!\n", p.Firstname)
}

func (Person) Struct() string {
	return "Person"
}
