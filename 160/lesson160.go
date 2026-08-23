package main

import "fmt"

// insertionSort сортирует слайс целых чисел по возрастанию,
// используя логику "последовательные обмены".
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				// Обмен соседних элементов
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
}

func main() {
	data := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("До сортировки:", data)
	insertionSort(data)
	fmt.Println("После сортировки:", data)
}
