package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name        string
	Email       string
	dateOfBirth time.Time
}

func main() {
	p := Person{
		Name:  "Иван",
		Email: "ivan@yandex.ru",
	}
	fmt.Println(p)

	ps := NewPerson("Иван", "ivan@yandex.ru", 2000, 12, 1)
	fmt.Println(ps.Name, ps.Email)

	ps.Name = "Пётр"
	fmt.Println(ps.Name)
	fmt.Println(ps)
}

func NewPerson(name, email string, dobYear, dobMonth, dobDay int) Person {
	return Person{
		Name:        name,
		Email:       email,
		dateOfBirth: time.Date(dobYear, time.Month(dobMonth), dobDay, 0, 0, 0, 0, time.UTC),
	}
}
