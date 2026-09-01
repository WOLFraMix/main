package main

import "fmt"

type UserID int

type User struct {
	ID        UserID
	Firstname string
}

func (u UserID) String() string {
	return fmt.Sprintf("UserID - %d:", u)
}

type ProductID int

type Product struct {
	ID    ProductID
	Title string
}

func (p ProductID) String() string {
	return fmt.Sprintf("ProductID - %d:", p)
}

func main() {
	user := User{
		ID:        3,
		Firstname: "Stepan",
	}
	prod := Product{
		ID:    55,
		Title: "Pencil",
	}

	fmt.Println(user)
	fmt.Println(prod)

	// разные типы чтобы их нельзя было случайно смешать
	// или неправильно присвоить:
	// например - user.ID = prod.ID
}
