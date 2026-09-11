package main

import (
	"fmt"
	"slices"
)

func main() {
	a := []int{1, 2, 3}
	b := a[:2:2]
	c := slices.Clone(a[:2])
	b[0] = 99
	b = append(b, 4)
	c = append(c, 5)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
