package main

import "fmt"

// bubbleSort реализует пузырьковую сортировку
func bubbleSort(arr []int) {
	sorted := false

	for !sorted {
		sorted = true
		// Проходим по всем элементам, кроме последнего
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				// Меняем элементы местами
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
}

func main() {
	data := []int{5, 3, 8, 4, 2, 1}
	fmt.Println("До сортировки:", data)

	bubbleSort(data)

	fmt.Println("После сортировки:", data)
}
