package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Введите строку: ")
	fmt.Scan(&s)
	runes := []rune(strings.ToLower(s))
	j := len(runes) - 1
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[j] {
			fmt.Println("Нет")
			return
		}
		j--
	}
	fmt.Println("Да")
}
