package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Printf("Введите три числа: ")
	fmt.Scan(&a, &b, &c)
	max := a
	if max < b {
		max = b
	}
	if max < c {
		max = c
	}
	fmt.Println("Максимум:", max)
}
