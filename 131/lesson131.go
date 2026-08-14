package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 7,
	}

	fmt.Println(m)

	value, exists := m["orange"]
	if exists {
		fmt.Println("Значение:", value)
	} else {
		fmt.Println("Значение по ключу не найдено")
	}

	delete(m, "banana")
	fmt.Println(m)

	m["strawberry"] = 13
	fmt.Println(m)
}
