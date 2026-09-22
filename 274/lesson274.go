package main

import "fmt"

func main() {
	var a, b int
	var n int
	fmt.Scan(&a, &b)
	if a > b {
		n = b
	} else {
		n = a
	}
	for i := n; i > 0; i-- {
		if a%i == 0 {
			if b%i == 0 {
				fmt.Println(i)
				return
			}
		}
	}
}
