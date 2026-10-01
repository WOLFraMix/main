package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
На вход программы подается строка.
Удалите из строки все символы, кроме букв и цифр.
*/

// \W - любой символ, кроме букв и цифр
// \w - любая буква или цифра

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	re := regexp.MustCompile(`\w`)

	matches := re.FindAllStringSubmatch(input, -1)
	for _, v := range matches {
		fmt.Print(strings.Join(v, ""))
	}
}
