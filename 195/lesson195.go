package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	// В примере все числа от 0 до 99, отрицательные и трехзначные мы не рассматриваем
	nums := []int{5, 12, 3, 45, 8, 20}

	slices.SortFunc(nums, func(a, b int) int {
		// Определяем, является ли число двузначным (от 10 до 99)
		isADouble := a >= 10 && a < 100
		isBDouble := b >= 10 && b < 100

		// Однозначные (false) идут раньше двузначных (true)
		if isADouble != isBDouble {
			if !isADouble {
				return -1
			}
			return 1
		}

		// Внутри группы сортируем по убыванию
		return cmp.Compare(b, a)
	})

	fmt.Println(nums) // [8 5 3 45 20 12]
}
