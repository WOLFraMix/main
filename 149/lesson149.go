package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string][]int{
		"a": {0, 2, 3},
		"b": {4, 5, 6},
	}
	m2 := map[string][]int{
		"a": {3, 1, 2},
		"b": {6, 5, 4},
	}
	m3 := map[string][]int{
		"a": {-5, -3},
		"b": {-5, -3, 0},
	}
	m4 := map[string][]int{
		"a": {-5, -1},
		"b": {},
	}

	fmt.Println(CompareMaxValues(m1, m2))
	fmt.Println(CompareMaxValues(m3, m4))
	fmt.Println(CompareMaxValues(m2, m3))
}

func CompareMaxValues(m1, m2 map[string][]int) bool {
	result := maps.EqualFunc(m1, m2, func(v1, v2 []int) bool {
		maxV1 := -9999
		for _, v := range v1 {
			if v > maxV1 {
				maxV1 = v
			}
		}
		maxV2 := -9999
		for _, v := range v2 {
			if v > maxV2 {
				maxV2 = v
			}
		}
		return maxV1 == maxV2
	})
	return result
}
