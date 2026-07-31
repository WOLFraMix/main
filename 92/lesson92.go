package main

import "fmt"

func main() {
	s := []int{6, 4, 5, 9, 7}
	fmt.Println(s, len(s), cap(s))
	s = append(s, 11, 22)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 30, 31, 32, 33)
	fmt.Println(s, len(s), cap(s))
}
