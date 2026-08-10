package main

import (
	"fmt"
	"slices"
)

func main() {
	slice1 := [][]int{
		{3, 1, 4, 1},
		{2, 2, 2},
		{5, 0, 6, 3, -8, 1},
		{4, 6, 8, 2},
	}

	slice2 := [][]int{
		{1, 2, 3},
		{2, 4},
		{6, 0},
		{5, -5, 5},
	}

	slice3 := [][]int{
		{10, 3, 5},
		{1, 4, 6},
		{7, 2},
		{8, 9},
	}

	slice4 := [][]int{
		{3, 1, 4},
		{1, 5, 9},
		{2, 6, 5},
		{0},
	}

	slice5 := [][]int{
		{5, 3, 1},
		{2, 4},
		{6, 8, 10},
		{7, 9, 11},
	}

	slice6 := [][]int{}

	slice7 := [][]int{
		{5},
		{3},
		{8},
		{1},
	}

	slice8 := [][]int{
		{3, 1, 4, 1, 0},
		{2, 2, 2},
		{5, 0, 6, 0, 3, -8, 1},
		{11, 4, 0, 6, 1, 0, 8, -5, 2, 0, 7},
	}

	magicSort(slice1)
	magicSort(slice2)
	magicSort(slice3)
	magicSort(slice4)
	magicSort(slice5)
	magicSort(slice6)
	magicSort(slice7)
	magicSort(slice8)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
	fmt.Println(slice7)
	fmt.Println(slice8)
}

// magicSort сортирует двумерный слайс:
// 1) Внешний слайс — по возрастанию суммы элементов внутреннего слайса
// 2) Каждый внутренний слайс — ноль в начале, затем чётные по убыванию, затем нечётные по убыванию
func magicSort(matrix [][]int) {
	// Сначала сортируем каждый внутренний слайс
	for i := range matrix {
		sortInnerSlice(matrix[i])
	}

	// Сортируем каждый внешний слайс
	slices.SortFunc(matrix, func(a, b []int) int {
		sumA := sumSlice(a)
		sumB := sumSlice(b)
		return sumA - sumB
	})

}

func sortInnerSlice(s []int) {
	slices.SortFunc(s, func(a, b int) int {
		// Ноль всегда раньше всего
		if a == 0 && b == 0 {
			return 0
		}
		if a == 0 {
			return -1 // a раньше
		}
		if b == 0 {
			return 1 // b раньше
		}

		aEven := a%2 == 0
		bEven := b%2 == 0

		// Если разная чётность: чётное раньше нечётного
		if aEven && !bEven {
			return -1
		}
		if !aEven && bEven {
			return 1
		}

		// Одинаковая чётность: сортируем по убыванию
		return b - a
	})
}

func sumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
