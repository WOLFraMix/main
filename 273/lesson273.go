package main

import (
	"fmt"
	"math"
)

/*
Простое число —  натуральное число,
имеющее ровно два различных натуральных делителя,
т. е. число N является простым,
если оно отлично от 1
и делится без остатка только на 1 и на само N.
*/

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

// isPrime проверяет простое ли число
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for d := 3; d <= limit; d += 2 {
		if n%d == 0 {
			return false
		}
	}
	return true
}
