package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	numbers := []int{5, 2, 8, 1}

	// По возрастанию (на англ. это ASC, от слова Ascending)
	slices.SortFunc(numbers, func(a, b int) int {
		return cmp.Compare(a, b) // a < b вернет -1, значит a встанет раньше
	})
	fmt.Println(numbers) // [1 2 5 8]

	// По убыванию (на англ. это DESC, от слова Descending), здесь мы просто поменяем аргументы местами :)
	slices.SortFunc(numbers, func(a, b int) int {
		return cmp.Compare(b, a) // Если b < a, вернет -1, значит `a` (большее) встанет раньше
	})
	fmt.Println(numbers) // [8 5 2 1]
}
