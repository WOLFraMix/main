package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "hello world"
	// strings.Split(s, sep string) []string
	// разбивает строку s на подстроки по разделителю sep
	// и возвращает их в виде среза строк.
	fmt.Println(strings.Split(s, " ")) // -> ["hello" "world"]

	s = "   hello world   "
	// strings.Trim(s string, cutset string) string
	// удаляет из начала и конца строки s все символы,
	// содержащиеся в cutset.
	fmt.Println(strings.Trim(s, " ")) // -> "hello world"

	s = "abc"
	// strings.Clone(s string) string
	// используется для явного создания копии строки.
	clone := strings.Clone(s)
	fmt.Println(s == clone) // true

	// strings.EqualFold(s1, s2)
	// функция выполняет сравнение без учета регистра
	// и возвращает true, если строки s1 и s2 совпадают
	fmt.Println(strings.EqualFold("Go", "go")) // true

	var str1 string
	fmt.Scan(&str1)
	str2 := strings.Clone(str1)
	runes := []rune(str2)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(strings.EqualFold(str1, string(runes)))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s1 := scanner.Text()
	count := 0
	maxCount := 0
	slice := strings.Fields(s1)
	for _, v := range slice {
		v = strings.Trim(v, ",.!")
		count = utf8.RuneCountInString(v)
		if count > maxCount {
			maxCount = count
		}
	}
	for _, v := range slice {
		v = strings.Trim(v, ",.!")
		if utf8.RuneCountInString(v) == maxCount {
			fmt.Println(v)
		}
	}
}
