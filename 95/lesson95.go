package main

import "fmt"

func main() {
	fmt.Println(filterEven(1, 2, 3, 4, 5, 6))
	fmt.Println(filterEven(0, -2, -3, -4, 5, 6))
	fmt.Println(filterEven())
}

func filterEven(values ...int) (result []int) {
	// оставляем только чётные числа
	result = []int{}

	for i := 0; i < len(values); i++ {
		if values[i]%2 == 0 { // если чётное
			result = append(result, values[i]) // добавляем в слайс
		}
	}
	return result
}
