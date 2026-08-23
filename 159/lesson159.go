package main

import "fmt"

// selectionSort сортирует слайс целых чисел
// по возрастанию методом выбора
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// Обмен элементов arr[i] и arr[min]
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func main() {
	data := []int{64, 25, 12, 22, 11}
	fmt.Println("Исходный массив:", data)
	selectionSort(data)
	fmt.Println("Отсортированный массив:", data)
}
