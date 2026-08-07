package main

import "fmt"

func main() {
	s1 := []int{1, 2, 6, 11, 8}
	s2 := []int{1, 2, 3, 4, 5}
	s3 := []int{}
	s4 := []int{5, 10, 15}
	s5 := []int{-5, -10, 20, 15}

	fmt.Println(PlayWithSlice(s1))
	fmt.Println(PlayWithSlice(s2))
	fmt.Println(PlayWithSlice(s3))
	fmt.Println(PlayWithSlice(s4))
	fmt.Println(PlayWithSlice(s5))
}

func PlayWithSlice(slice []int) (result []int) {
	// result = slice - не создаёт копию
	result = make([]int, len(slice))
	copy(result, slice)

	if len(slice) == 0 {
		return
	}

	insert1 := 100
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] >= 10 {
			result = append(result[:i+1], append([]int{insert1}, result[i+1:]...)...)
			break
		}
	}

	insert2 := 500
	sum := 0
	for i := 0; i < len(result); i++ {
		sum += result[i]
	}
	if sum > 100 {
		result = append(result, insert2)
	}

	insert3 := 1000
	even := 0 // чётные
	odd := 0  // нечётные
	for _, v := range slice {
		if v%2 == 0 {
			even++
		}
		if v%2 != 0 {
			odd++
		}
	}
	if even > odd {
		result = append([]int{insert3}, result...)
	}

	return
}
