package main

import "fmt"

type Name struct {
	First string
	Last  string
}

type User struct {
	Name      Name
	Birthyear int
}

func main() {
	name := Name{
		First: "Stepan",
		Last:  "Bgan",
	}
	user := User{
		Name:      name,
		Birthyear: 1999,
	}
	fmt.Printf("%+v", user)
	fmt.Println(user.Name.First)

	user.Name.First = "Степашка"
	fmt.Printf("%+v", user)
}
