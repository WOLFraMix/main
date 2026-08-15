package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  3,
		"banana": 20,
		"orange": 1000000,
	}

	for key, value := range m {
		fmt.Printf("Ключ: %s, Значение: %d\n", key, value)
	}

	s := []int{1, 2, 3, 4, 5, 1, 1, 2, 2, 2, 3, 3, 3, 3, -1, -2, -1, -3, -1, -2, 0, 0, 0, 0, 1}
	fmt.Println(CountMaxFrequency(s))
}

// максимальное количество вхождений в слайсе
func CountMaxFrequency(slice []int) (result int) {
	m := make(map[int]int)

	// ранжируем слайс и каждому ключу даём кол-во
	for _, v := range slice {
		m[v]++
	}
	// ранжируем map и вытягиваем самое большое кол-во
	for _, value := range m {
		if result < value {
			result = value
		}
	}

	return result
}
