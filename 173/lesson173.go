package main

import (
	"fmt"
	"sort"
)

// groupAnagrams группирует строки по анаграммам.
func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	// Ключом в мапе служит отсортированная версия строки.
	for _, word := range strs {
		// Преобразуем строку в срез рун для корректной сортировки
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedWord := string(runes)

		// Добавляем исходное слово в группу по отсортированному ключу
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	// Преобразуем значения мапы в срез срезов для возврата результата
	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := groupAnagrams(input)
	fmt.Println(output)
	// Примерный вывод (порядок групп и слов внутри групп может отличаться):
	// [[eat tea ate] [tan nat] [bat]]
}
