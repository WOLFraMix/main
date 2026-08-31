package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func (p Person) SayHello() {
	fmt.Printf("%s приветствует вас!", p.Firstname)
}

func (Person) Struct() string {
	return "Person"
}

func main() {
	var p *Person
	fmt.Println(p == nil)
}
