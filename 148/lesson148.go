package main

import (
	"fmt"
	"maps"
)

func main() {
	m111 := map[string]int{
		"a": 5,
		"b": 10,
	}
	m222 := map[string]int{
		"a": 5,
		"b": 10,
	}
	fmt.Println(maps.Equal(m111, m222))

	m1 := map[string][]int{
		"a": []int{5, 2},
		"b": []int{20, 11},
	}
	m2 := map[string][]int{
		"a": {1, 5, 1},
		"b": {11, 20},
	}

	result := maps.EqualFunc(m1, m2, func(v1, v2 []int) bool {
		sum1 := 0
		for _, v := range v1 {
			sum1 += v
		}
		sum2 := 0
		for _, v := range v2 {
			sum2 += v
		}
		return sum1 == sum2
	})
	fmt.Println(result)
}
