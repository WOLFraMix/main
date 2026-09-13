package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s []int
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		if v%2 == 0 {
			s = append(s, v)
		}
	}
	fmt.Println(s)
}
