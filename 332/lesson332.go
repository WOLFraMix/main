package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите N: ")
	fmt.Scan(&n)
	fmt.Println(isPrime(n))
}

func isPrime(n int) string {
	if n < 2 {
		return "Составное"
	}
	if n == 2 {
		return "Простое"
	}
	if n%2 == 0 {
		return "Составное"
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return "Составное"
		}
	}
	return "Простое"
}
