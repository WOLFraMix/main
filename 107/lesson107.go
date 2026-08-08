package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(cap(slice))

	slice = slice[:len(slice)-1] // удаление с конца
	fmt.Println(slice)
	fmt.Println(cap(slice))

	slice = slice[1:] // удаление с начала
	fmt.Println(slice)
	fmt.Println(cap(slice))

	index := 1
	slice = append(slice[:index], slice[index+1:]...) // удаление в середине
	fmt.Println(slice)
	fmt.Println(cap(slice))

	slice = slices.Clip(slice)
	fmt.Println(slice)
	fmt.Println(cap(slice))
}
