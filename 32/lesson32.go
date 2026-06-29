package main

import (
	"fmt"
)

func main() {
	var firstName string
	fmt.Print("Введите имя: ")
	fmt.Scan(&firstName) // передаём указатель на переменную чтобы считывать и менять её значение
	fmt.Printf("Ваше имя: %s\n", firstName)
	// fmt.Scanln(&firstName, &lastName) - читает всю строчку, для считывания нескольких переменных через пробел
}
