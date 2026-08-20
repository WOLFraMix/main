package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{
		"a": 5,
		"b": 10,
	}
	m2 := map[string]int{
		"a": 50,
		"c": 10,
	}

	maps.Copy(m1, m2)
	fmt.Println(m1)

	m3 := maps.Clone(m1)
	fmt.Println(m3)
}
