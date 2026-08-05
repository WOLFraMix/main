package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{10, 11, 12}
	s2 := []int{1, 2, 3, 4, 5}
	copy(s1, s2)
	fmt.Println(s1)

	s3 := []int{10, 11, 12}
	s4 := []int{1, 2, 3, 4, 5}
	copy(s4, s3)
	fmt.Println(s4)

	s5 := []int{10, 11, 12}
	s6 := s5
	fmt.Println(s5, s6)
	s6[0] = 666
	fmt.Println(s5, s6)

	s7 := []int{100, 110, 120}
	s8 := make([]int, len(s7))
	copy(s8, s7)
	fmt.Println(s7, s8)
	s8[0] = 777
	fmt.Println(s7, s8)

	s9 := slices.Clone(s7)
	fmt.Println(s9)
}
