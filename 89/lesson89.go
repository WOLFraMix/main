package main

import "fmt"

func main() {
	numbers := []int{6, 4, 5, 9, 7}
	fmt.Println(numbers)

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("Сумма элементов слайса:", sum)

	for i := range numbers {
		numbers[i] *= 10
	}
	fmt.Println(numbers)
}
