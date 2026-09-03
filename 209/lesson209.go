package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	arr := make([]int, n)
	totalSum := 0
	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			return
		}
		totalSum += arr[i]
	}

	// Количество множеств равно n - 1
	k := n - 1

	// Проверка 1: Общая сумма должна делиться на количество множеств без остатка
	if totalSum%k != 0 {
		fmt.Println("NO")
		return
	}

	targetSum := totalSum / k

	// Подсчет элементов, равных целевой сумме
	countEqual := 0
	var remaining []int

	for _, val := range arr {
		if val == targetSum {
			countEqual++
		} else {
			remaining = append(remaining, val)
		}
	}

	// Проверка 2: Должно быть ровно n-2 элемента, равных targetSum
	// Это означает, что в слайсе remaining должно остаться ровно 2 элемента
	if countEqual == n-2 && len(remaining) == 2 {
		// Проверка 3: Сумма оставшихся двух элементов должна быть равна targetSum
		if remaining[0]+remaining[1] == targetSum {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")
}
