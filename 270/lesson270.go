package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var m int
	var min int = math.MaxInt
	var max int = math.MinInt
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		if m > max {
			max = m
		}
		if m < min {
			min = m
		}
	}

	fmt.Printf("Наименьшее число: %d\n", min)
	fmt.Printf("Наибольшее число: %d\n", max)
}
