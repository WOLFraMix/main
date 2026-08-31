package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	slice := []int{3, 1, 4, 1, 10, 0, 8, 5, 7, 15, 16, 2}
	sortMagic(slice)
	fmt.Println(slice)
}

// sortMagic сортирует слайс целых чисел в порядке убывания,
// при этом четные числа идут перед нечетными.
func sortMagic(s []int) {
	slices.SortFunc(s, func(a, b int) int {
		if a%2 == 0 && b%2 == 0 {
			return cmp.Compare(b, a)
		}
		if a%2 == 0 && b%2 != 0 {
			return -1
		}
		if a%2 != 0 && b%2 == 0 {
			return 1
		}
		return cmp.Compare(b, a)
	})
}
