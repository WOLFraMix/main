package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}
