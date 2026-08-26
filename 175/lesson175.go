package main

import (
	"fmt"
)

type Name struct {
	First string
	Last  string
}

type User struct {
	Name      Name
	BirthYear int
}

func (u User) Greet() {
	fmt.Printf("Приветствую. Я %s %s, %d года рождения.\n", u.Name.First, u.Name.Last, u.BirthYear)
}

func (u *User) changeName(name Name) {
	u.Name = name
}

func main() {
	name := Name{
		First: "Степан",
		Last:  "Бган",
	}
	user := User{
		Name:      name,
		BirthYear: 1999,
	}
	fmt.Printf("%+v\n", user)
	user.Greet()

	newName := Name{
		First: "Кто-то",
		Last:  "Другой",
	}
	user.changeName(newName)
	user.Greet()
}
