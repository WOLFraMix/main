package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"b": 3, "c": 4, "d": 5}

	result := mergeMaps(m1, m2)
	// result будет равен map[string]int{"a": 1, "b": 5, "c": 7, "d": 5}
	fmt.Println(result)
}

// Объединяем мапы и складываем одинаковые ключи
func mergeMaps(m1, m2 map[string]int) map[string]int {
	// Если обе nil — возвращаем пустую карту
	if m1 == nil && m2 == nil {
		return make(map[string]int)
	}
	if m1 == nil {
		return m2
	}
	if m2 == nil {
		return m1
	}

	result := maps.Clone(m1)

	for key, value := range m2 {
		// Если ключ уже есть, складываем значения,
		// если нет — просто добавляем
		result[key] += value
	}

	return result
}
