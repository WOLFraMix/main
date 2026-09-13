package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		a[i] = v
	}
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		b[i] = v
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(a[i] + b[i])
	}
	fmt.Println()
}
