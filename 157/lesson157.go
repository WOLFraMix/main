package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string][]int{
		"a": {5, 2},
		"b": {20, 11},
		"c": {1, 3, 9},
		"d": {6, 111, 5},
	}
	fmt.Println(m)

	maps.DeleteFunc(m, func(key string, value []int) bool {
		for _, v := range value {
			if v == 5 {
				return true
			}
		}
		return false
	})
	fmt.Println(m)
}
