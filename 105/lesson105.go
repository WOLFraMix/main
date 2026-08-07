package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4) // вставка в начало
	fmt.Println(slice)

	slice = append([]int{0}, slice...) // вставка в конец
	fmt.Println(slice)

	index := 3
	value := 100
	before := slice[:index]
	after := append([]int{value}, slice[index:]...)
	slice = append(before, after...) // вставка в середину
	fmt.Println(slice)
}
