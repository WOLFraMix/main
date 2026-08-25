package main

import "fmt"

type User struct {
	Firstname string
}

func main() {
	user := User{
		Firstname: "Степан",
	}
	fmt.Println(user)

	modifyValue(user)
	fmt.Println(user)

	modifyValueWithPointer(&user)
	fmt.Println(user)
}

func modifyValue(u User) { // сюда передаётся только копия
	u.Firstname = "Андрей"
}

func modifyValueWithPointer(u *User) {
	u.Firstname = "Андрей"
}
