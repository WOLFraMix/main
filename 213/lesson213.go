package main

import (
	"fmt"
	"sort"
)

// twoPointers проверяет,
// есть ли в слайсе два числа,
// сумма которых равна target.
func twoPointers(a []int, target int) bool {
	sort.Ints(a) // Сортируем.
	// Инициализируем два указателя:
	left := 0
	right := len(a) - 1

	// Двигаем указатели, пока они не встретятся.
	for left < right {
		s := a[left] + a[right]
		if s == target {
			return true
		} else if s < target {
			left++
		} else {
			right--
		}
	}
	return false
}

func main() {
	// Пример использования функции twoPointers.

	nums := []int{3, 5, 1, 8, 4, 2}
	target := 9

	fmt.Printf("Слайс: %v\n", nums)
	fmt.Printf("Целевая сумма: %d\n", target)

	if twoPointers(nums, target) {
		fmt.Println("Найдена пара чисел, дающая целевую сумму.")
	} else {
		fmt.Println("Пара чисел с целевой суммой не найдена.")
	}

	// Ещё один тестовый пример: когда пары нет.
	nums2 := []int{1, 2, 3}
	target2 := 10

	fmt.Printf("\nСлайс: %v\n", nums2)
	fmt.Printf("Целевая сумма: %d\n", target2)

	if twoPointers(nums2, target2) {
		fmt.Println("Найдена пара чисел, дающая целевую сумму.")
	} else {
		fmt.Println("Пара чисел с целевой суммой не найдена.")
	}
}
