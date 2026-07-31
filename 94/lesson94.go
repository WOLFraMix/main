package main

import "fmt"

func main() {
	fmt.Println("Сумма:", sum(1, 2, 3, 4, 5))

	s := []int{3, 4, 5, 3, 4, 5, 3, 4, 5}
	fmt.Println("Сумма:", sum(s...)) // значения в функцию

	s1 := []int{1, 2, 3}
	s2 := []int{6, 5, 4}
	s1 = append(s1, s2...) // значения из массива
	fmt.Println(s1)
}

func sum(values ...int) int { // какое-то количество значений
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
