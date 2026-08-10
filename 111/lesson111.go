package main

import (
	"fmt"
	"slices"
)

func main() {
	ints := []int{4, 2, 3, 1, 7}
	slices.Sort(ints)
	fmt.Println(ints)

	strings := []string{"banana", "cherry", "apple"}
	slices.Sort(strings)
	fmt.Println(strings)

	numbers := []int{4, 2, 7, 1, 3, 5, 7, 2, 8, 9}
	slices.SortFunc(numbers, func(a, b int) int {
		return a - b
		// тоже самое что:
		/*
			if a > b {
				return 1
			}
			if a < b {
				return -1
			}
			return 0
		*/
		// в обратном порядке сортировка b - a
	})
	fmt.Println(numbers)
}
