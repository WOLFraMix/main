package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	runes := []rune(str)
	n := len(runes)

	// Идём только до середины: i < n/2
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		runes[i], runes[j] = runes[j], runes[i]
	}

	fmt.Println(string(runes))
}
