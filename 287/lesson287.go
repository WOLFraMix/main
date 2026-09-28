package main

import "fmt"

func main() {
	n := 0
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}
