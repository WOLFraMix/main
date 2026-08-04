package main

import (
	"fmt"
)

func main() {
	s1 := [][]int{{1, 2}, {3, 4}}
	fmt.Println(mirrorMatrix(s1))
}

// отзеркаливаем каждую строку двумерного слайса по горизонтали
func mirrorMatrix(matrix [][]int) [][]int {
	// создаём новый слайс для результата
	result := make([][]int, len(matrix))

	for i, row := range matrix {
		// создаём новую строку нужной длины
		mirroredRow := make([]int, len(row))

		// заполняем строку в обратном порядке
		for j, val := range row {
			// длина слайса - 1 - индекс = индекс наоборот
			mirroredRow[len(row)-1-j] = val
		}
		result[i] = mirroredRow
	}
	return result
}
