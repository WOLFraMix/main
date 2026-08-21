package main

import "fmt"

// binarySearchSqrt находит наибольшее целое число x, такое что x*x <= target
func binarySearchSqrt(target int) int {
	if target < 0 {
		return -1 // для отрицательных чисел целочисленный корень не определён
	}
	if target == 0 || target == 1 {
		return target
	}

	left := 0
	right := target
	result := 0

	for left <= right {
		middle := left + (right-left)/2 // избегаем потенциального переполнения
		square := middle * middle

		if square == target {
			return middle
		} else if square > target {
			right = middle - 1
		} else {
			// square < target: middle может быть ответом, но попробуем найти больше
			result = middle
			left = middle + 1
		}
	}

	return result
}

func main() {
	testValues := []int{0, 1, 4, 8, 9, 15, 16, 24, 25, 100, 2147395600}

	for _, val := range testValues {
		sqrt := binarySearchSqrt(val)
		fmt.Printf("Целочисленный корень из %d = %d (проверка: %d^2 = %d)\n",
			val, sqrt, sqrt, sqrt*sqrt)
	}
}
