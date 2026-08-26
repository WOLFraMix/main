package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID        int    `json:"id" db:"database_id"`
	FirstName string `json:"username" db:"first_name"`
	Email     string `json:"email" db:"email"`
}

func main() {
	u := User{
		ID:        15,
		FirstName: "Stepan",
		Email:     "happy@mail.ru",
	}

	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Поле: %s, JSON: %s, DB: %s.\n", field.Name, field.Tag.Get("json"), field.Tag.Get("db"))
	}
}
