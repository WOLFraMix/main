package main

import "fmt"

func main() {
	m := make(map[string]map[string]int)

	m["fruits"] = map[string]int{
		"apple":  8,
		"banana": 10,
	}
	m["vegetables"] = map[string]int{
		"carrot": 2,
	}
	fmt.Println(m)
	fmt.Println(m["fruits"]["banana"])
}
