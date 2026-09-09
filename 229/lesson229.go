package main

import (
	"fmt"
)

func main() {
	a := [10]int{}
	var sum int
	for i := range a {
		var n int
		fmt.Scan(&n)
		a[i] = n
		sum += n
	}
	avg := sum / 10
	for i := range a {
		if a[i] > avg {
			fmt.Print(a[i])
			if i < len(a) {
				fmt.Print(" ")
			}
		}
	}
}
