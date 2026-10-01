package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// Пример конвертация строки из формата camelСase в формат snake_case
var (
	// Example: match "rN" in "userName"
	camelRe = regexp.MustCompile(`[a-z][A-Z]`)
)

// fixCase gets a string in the format "aB" and return "a_b"
func fixCase(s string) string {
	return fmt.Sprintf("%c_%c", s[0], unicode.ToLower(rune(s[1])))
}

// camelToLower turns "camelCase" to "camel_case"
func camelToLower(s string) string {
	return camelRe.ReplaceAllStringFunc(s, fixCase)
}

func main() {
	fmt.Println(camelToLower("userName")) // user_name

	fmt.Println(toCamelCase("user_name")) // userName
}

// Пример конвертации строки из формата snake_case в формат camelСase:
var (
	// Матчит подчеркивания и следующую за ним букву
	snakeRe = regexp.MustCompile(`_([a-z])`)
)

// toCamelCase конвертирует "user_name" в "userName"
func toCamelCase(s string) string {
	return snakeRe.ReplaceAllStringFunc(s, func(match string) string {
		// Берем букву после подчеркивания и делаем её заглавной
		letter := match[1:] // убираем "_"
		return strings.ToUpper(letter)
	})
}
