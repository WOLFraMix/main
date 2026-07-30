package main

import "fmt"

func main() {
	var slice1 []int // пустой слайс
	fmt.Println("slice1 == nil:", slice1 == nil)

	slice2 := []int{6, 4, 5, 9, 7} // слайс имеет длину и вместимость
	fmt.Printf("slice2: %d, len: %d, cap(capacity): %d\n", slice2, len(slice2), cap(slice2))

	slice3 := make([]int, 5, 10) // создаём слайс с нужными параметрами: длина и вместимость
	fmt.Printf("slice3: %d, len: %d, cap(capacity): %d\n", slice3, len(slice3), cap(slice3))
}
