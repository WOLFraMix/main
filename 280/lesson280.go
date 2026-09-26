package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	fmt.Println("Квадрат:", n*n)
}
