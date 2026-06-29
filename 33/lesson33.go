package main

import (
	"fmt"
)

func main() {
	var firstName string
	var lastName string
	var age int
	fmt.Println("Введите ваше имя и фамилию через пробел: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Println("Введите ваш возраст: ")
	fmt.Scan(&age)
	result := age - 5
	fmt.Printf("Приятно познакомиться, %s. Я 5 лет назад познакомился с человеком, у которого тоже фамилия %s, вам тогда было %d. Как молоды мы были!", firstName, lastName, result)
}
