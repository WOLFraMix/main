package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Print("Введите строку: ")
	fmt.Scan(&s)
	vowels := map[rune]bool{
		'а': true, 'е': true, 'ё': true, 'и': true, 'о': true,
		'у': true, 'ы': true, 'э': true, 'ю': true, 'я': true,
	}
	result := 0
	for _, v := range s {
		if vowels[unicode.ToLower(v)] {
			result++
		}
	}
	fmt.Println("Гласных:", result)
}
