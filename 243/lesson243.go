package main

import (
	"fmt"
	"sort"
)

/*
Медиана последовательности — это элемент,
который стоит в середине отсортированной последовательности.
Для заданной последовательности выведите медианы каждого его префикса.
*/

func main() {
	// Количество чисел
	var count int
	fmt.Scan(&count)

	// Считываем числа
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&numbers[i])
	}

	// Для каждого префикса вычисляем медиану
	for i := 1; i <= count; i++ {
		// Копируем префикс
		prefix := make([]int, i)
		copy(prefix, numbers[:i])

		// Сортируем префикс
		sort.Ints(prefix)

		// Находим медиану
		// (берём элемент с индексом (i-1)/2)
		median := prefix[(i-1)/2]

		// Выводим медиану
		fmt.Printf("%d ", median)
	}
	fmt.Println()
}
