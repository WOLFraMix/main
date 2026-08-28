package main

import "fmt"

func main() {
	s := []int{5, 2, 9, 4, 2}
	fmt.Println(UniqueSlice(s))
}

// UniqueSlice проверяет все ли значения в слайсе уникальные
func UniqueSlice(s []int) bool {
	m := make(map[int]struct{}, len(s))

	for i := 0; i < len(s); i++ {
		val := s[i]
		if _, ok := m[val]; ok {
			return false
		}
		// важны только ключи, поэтому
		// присваиваем пустую структуру для значения
		m[val] = struct{}{}
	}
	return true
}
