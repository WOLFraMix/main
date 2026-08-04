package main

import "fmt"

func main() {
	s1 := [][]int{
		{5, 6, 7, 8, 0},
		{9, 10, 11, 12, 14, 16, 18, 20},
		{},
		{7, 8, 9, 22, 48, -16, -4},
		{10, 11},
	}
	fmt.Println(replaceEvenOnEvenIndices(s1))
}

func replaceEvenOnEvenIndices(slice [][]int) [][]int {
	result := make([][]int, len(slice))

	for i, row := range slice {
		evenRow := make([]int, len(row))

		for j, v := range row {
			if j%2 == 0 && v%2 == 0 {
				evenRow[j] = 0
			} else {
				evenRow[j] = v
			}
			result[i] = evenRow
		}
	}
	return result
}
