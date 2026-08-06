package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5, 6}
	subSlice := original[1:4] // ОТ включительно, ДО не включая
	fmt.Println(subSlice)

	str := "Ку-ку!"
	str1 := str[:4]
	fmt.Println(str1) // байты 0, 1, 2, 3 = 2 ру буквы

	slice := make([]int, 5, 10)
	copy(slice, []int{1, 2, 3, 4, 5})
	fmt.Println(slice)
	sSlice := slice[2:8] // брать значения можно только в пределах capacity
	fmt.Println(sSlice)
	sSlice[1] = 333
	fmt.Println(slice, sSlice)
	sSlice[5] = 100
	fmt.Println(slice, sSlice)

	numbers := []int{10, 20, 30, 40, 50}
	ptr := &numbers[2]
	fmt.Println("Значение:", *ptr)
	fmt.Println("Адрес в памяти:", ptr)
	*ptr = 100
	fmt.Println(numbers)

	numbers = append(numbers, 60) // новый слайс, новый адрес
	*ptr = 444
	fmt.Println(numbers)
}
