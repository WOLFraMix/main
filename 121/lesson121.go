package main

import "fmt"

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	fmt.Println(RemoveDuplicates(input))
}

func RemoveDuplicates(input []string) []string {
	// cоздаётся пустая map, которая будет хранить уже встреченные строки
	// будем помечать строки, которые уже видели, значением true
	seen := make(map[string]bool)
	// в результат запишем слайс без дубликатов
	result := make([]string, 0, len(input))

	for _, v := range input {
		if seen[v] != true { // если мы ещё не видели эту строку
			seen[v] = true             // то добавляем её
			result = append(result, v) // а в результат добавляем значение
		}
	}
	return result
}
