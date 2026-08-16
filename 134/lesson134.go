package main

import "fmt"

func mergeSortedArrays(arr1, arr2 []int) []int {
	mergedArray := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			mergedArray = append(mergedArray, arr1[i])
			i++
		} else {
			mergedArray = append(mergedArray, arr2[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы из arr1, если они есть
	for i < len(arr1) {
		mergedArray = append(mergedArray, arr1[i])
		i++
	}
	// Добавляем оставшиеся элементы из arr2, если они есть
	for j < len(arr2) {
		mergedArray = append(mergedArray, arr2[j])
		j++
	}

	return mergedArray
}

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8, 10}
	result := mergeSortedArrays(arr1, arr2)
	fmt.Println(result) // Вывод: [1 2 3 4 5 6 7 8 10]
}
