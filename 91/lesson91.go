package main

import "fmt"

func main() {
	fmt.Println(fib(20))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	a := 0
	b := 1
	var result int

	for i := 2; i <= n; i++ {
		result = a + b
		a = b
		b = result
	}

	return result
}
