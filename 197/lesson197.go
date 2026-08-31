package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	data := [][]int{
		{3, 1, 4, 1},
		{2, 2, 2},
		{5, 0, 6, 3, -8, 1},
		{4, 6, 8, 2},
	}
	sortBySum(data)
	fmt.Println(data)
}

// sortBySum сортирует двумерный слайс целых чисел по возрастанию суммы элементов
// во внутренних слайсах. Сортировка стабильна, содержимое внутренних слайсов не изменяется.
func sortBySum(data [][]int) {
	slices.SortStableFunc(data, func(a, b []int) int {
		var sumA int
		var sumB int
		for _, v := range a {
			sumA += v
		}
		for _, v := range b {
			sumB += v
		}
		return cmp.Compare(sumA, sumB)
	})
}
