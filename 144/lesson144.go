package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "Да здравствует прекрасный язык, да здравствует golang!"
	text = strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, text)

	words := strings.Split(text, " ")
	fmt.Println(words)

	wordsCount := make(map[string]int)
	for _, word := range words {
		wordsCount[word]++
	}
	fmt.Println(wordsCount)

	for word, count := range wordsCount {
		fmt.Printf("Слово %q встречается %d раз.\n", word, count)
	}
}
