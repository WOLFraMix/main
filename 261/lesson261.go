package main

import "fmt"

/*
Дана строка s,
найдите длину самой длинной подстроки
без повторяющихся символов.
*/

func lengthOfLongestSubstring(s string) int {
	// Карта для хранения индексов каждого символа
	charIndex := make(map[rune]int)
	maxLen := 0
	start := 0 // Начальный индекс текущей подстроки

	for i, ch := range s {
		lastIdx, found := charIndex[ch]
		if found && lastIdx >= start {
			// если символ повторяется,
			// то перемещаемся на следующую позицию
			// после последнего появления
			start = lastIdx + 1
		}
		charIndex[ch] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}

	return maxLen
}

func main() {
	s := "abcabcbb"                          // abc
	fmt.Println(lengthOfLongestSubstring(s)) // Вывод: 3
}
