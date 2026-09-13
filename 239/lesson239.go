package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)

	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		s[i] = v
	}
	for i := len(s) - 1; i >= 0; i-- {
		if i >= 0 && i < len(s)-1 {
			fmt.Print(" ")
		}
		fmt.Print(s[i])
	}
	fmt.Println()
}
