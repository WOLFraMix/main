package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите N: ")
	fmt.Scan(&n)
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("Сумма:", sum)
}
