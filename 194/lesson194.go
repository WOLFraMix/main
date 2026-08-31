package main

import (
	"cmp"
	"fmt"
	"slices"
	"unicode/utf8"
)

func main() {
	words := []string{"Соня", "Боня", "Павел", "Елена"}

	slices.SortFunc(words, func(a, b string) int {
		// Сначала сравниваем длины
		lenCmp := cmp.Compare(utf8.RuneCountInString(a), utf8.RuneCountInString(b))

		// Если длины разные, возвращаем результат сравнения длин
		if lenCmp != 0 {
			return lenCmp
		}

		// Если длины одинаковые, сравниваем сами строки лексикографически, по Unicode
		return cmp.Compare(a, b)
	})

	fmt.Println(words) // [Боня Соня Елена Павел]
}
