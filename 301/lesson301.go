package main

import "fmt"

func main() {
	fmt.Println(fibonacci(100))
}

// fibonacci возвращает слайс чисел Фибоначчи, не превышающих maxValue.
func fibonacci(n int) []int {
	result := []int{}
	// Обработка некорректных значений: для отрицательных чисел возвращаем пустой слайс
	if n < 0 {
		return result
	}

	// Первые два числа последовательности
	f1, f2 := 0, 1

	// Добавляем 0, если он входит в диапазон (maxValue >= 0)
	if f1 <= n {
		result = append(result, f1)
	}

	// Генерируем следующие числа, пока они не превысят maxValue
	for f2 <= n {
		result = append(result, f2)
		f1, f2 = f2, f1+f2
	}

	return result
}
