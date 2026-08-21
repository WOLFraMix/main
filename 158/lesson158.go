package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string][]int{
		"a": {5, 1},
		"b": {20, 11},
		"c": {1, 3, 9},
		"d": {6, 111, 5},
	}
	m2 := map[string][]int{
		"a": {1, 2, 3},
		"b": {1, 1, 1, 1},
		"c": {20, 30, -100},
		"d": {7},
	}

	RemoveSlicesBySum(m1)
	fmt.Println(m1)
	RemoveSlicesBySum(m2)
	fmt.Println(m2)
}

func RemoveSlicesBySum(m map[string][]int) {
	maps.DeleteFunc(m, func(key string, value []int) bool {
		sum := 0
		for _, v := range value {
			sum += v
		}
		if sum > 6 {
			return true
		}
		return false
	})
}
