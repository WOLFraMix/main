package main

import "fmt"

type User struct {
	Firstname string
	Lastname  string
	Birthyear int
}

func main() {
	user := User{
		Firstname: "Stepan",
		Lastname:  "Bgan",
		Birthyear: 1999,
	}

	fmt.Println(user)
	fmt.Printf("%+v\n", user)
	fmt.Println(user.Lastname)
}
