package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	m := map[string]int{
		"banana":     2,
		"apple":      1,
		"grapefruit": 3,
		"cherry":     1,
	}
	invertedMap := invertMap(m)
	printMap(invertedMap)
}

// инвертируем входную map
func invertMap(m map[string]int) (result map[int][]string) {
	result = make(map[int][]string)

	for key, value := range m {
		// создаём новый слайс для текущего значения
		str := result[value]
		// добавляем ключ в слайс
		str = append(str, key)
		// записываем слайс обратно в map
		result[value] = str
	}

	return result
}

func printMap(m map[int][]string) {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Println("{")
	for _, k := range keys {
		sort.Strings(m[k]) // сортируем строки внутри слайса

		var parts []string
		for _, s := range m[k] {
			parts = append(parts, fmt.Sprintf(`"%s"`, s))
		}
		list := "[" + strings.Join(parts, ", ") + "]"

		fmt.Printf("  %d: %s,\n", k, list)
	}
	fmt.Println("}")
}
