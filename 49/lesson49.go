package main

import (
	"fmt"
	"strings"
)

func main() {
	var pass string
	fmt.Println("Введите пароль: ")
	fmt.Scan(&pass)
	isSecured := securePassword(pass)
	fmt.Println(isSecured)
}

// securePassword validates whether a password meets basic security requirements:
// 1. Length is at least 8 characters.
// 2. Contains no whitespace.
// 3. Contains both uppercase and lowercase letters.
func securePassword(pass string) bool {
	return len([]rune(pass)) >= 8 &&
		!strings.Contains(pass, " ") &&
		pass != strings.ToLower(pass) &&
		pass != strings.ToUpper(pass)
}
