package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите N: ")
	fmt.Scan(&n)
	sieve := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if !sieve[i] {
			for j := i * i; j <= n; j += i {
				sieve[j] = true
			}
		}
	}
	for i := range sieve {
		if i >= 2 && !sieve[i] {
			if i > 2 {
				fmt.Print(" ")
			}
			fmt.Print(i)
		}
	}
}
