package main

import "fmt"

func main() {
	s1 := []int{5, 7, 2, 8}
	s2 := []int{-1, -5, -3}
	s3 := []int{}

	fmt.Println(Max(s1))
	fmt.Println(Max(s2))
	fmt.Println(Max(s3))
}

func Max(slice []int) (int, error) {
	// ищем максимальное значение в слайсе
	if len(slice) == 0 { // если слайс пустой, возвращаем ошибку
		return 0, fmt.Errorf("slice is nil or empty")
	}
	max := slice[0] // приравниваем к любому значению чтобы избежать нуля
	for i := 0; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}
