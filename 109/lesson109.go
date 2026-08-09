package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Оригинальный слайс:", s)

	changeSlice(s)
	fmt.Println("Слайс после изменений:", s)

	fmt.Println("Слайс 2 после изменений:", changeSlice2(s))

}

func changeSlice(slice []int) {
	slice[len(slice)-1] = 100
	slice = append(slice, -1)
	slice[0] = 500
	fmt.Println("Слайс после append создался новый:", slice)
}

func changeSlice2(slice []int) []int {
	slice = append(slice, -1)
	slice[0] = 500
	fmt.Println("Слайс после append создался новый:", slice)
	return slice
}
