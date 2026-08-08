package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{1, 2, 3, 4, 5, 11}
	s3 := []int{1, 2, 3, 4, 5, 6, 11}
	s4 := []int{}
	s5 := []int{1, 2, 3, 4, 5, 11, 12}

	fmt.Println(DeletingFromSlice(s1))
	fmt.Println(DeletingFromSlice(s2))
	fmt.Println(DeletingFromSlice(s3))
	fmt.Println(DeletingFromSlice(s4))
	fmt.Println(DeletingFromSlice(s5))
}

func DeletingFromSlice(slice []int) (result []int) {
	result = make([]int, len(slice))
	copy(result, slice)

	if len(slice) == 0 {
		return
	}

	count := 0
	if result[len(result)-1] >= 10 {
		result = result[:len(result)-1]
		count++
	}
	if cap(result) > 5 {
		result = append(result[:2], result[3:]...)
		count++
	}
	if len(result) >= 1 && count == 2 {
		result = result[1:]
	}
	result = slices.Clip(result)

	return
}
