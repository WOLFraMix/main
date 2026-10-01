package main

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	a := 10
	b := "hello"
	c := true

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	fmt.Fprint(file, a, b, c)
	file.Close()

	// Проверяем, является ли символ цифрой
	fmt.Println(unicode.IsDigit('5')) // -> true
	fmt.Println(unicode.IsDigit('a')) // -> false

	// Проверяем, является ли символ буквой
	fmt.Println(unicode.IsLetter('A')) // -> true
	fmt.Println(unicode.IsLetter('5')) // -> false

	// Проверяем, является ли символ пробельным
	fmt.Println(unicode.IsSpace(' '))  // -> true
	fmt.Println(unicode.IsSpace('\n')) // -> true
	fmt.Println(unicode.IsSpace('a'))  // -> false

	// Преобразуем символы в нижний и верхний регистр
	fmt.Println(string(unicode.ToLower('A'))) // -> "a"
	fmt.Println(string(unicode.ToUpper('a'))) // -> "A"

	// Проверяем, что буква в нижнем регистре
	fmt.Println(unicode.IsLower('ы')) // -> true

	// Проверяем, что буква в вверхнем регистре
	fmt.Println(unicode.IsUpper('Ы')) // -> true

	// Считаем количество символов Unicode в строке
	s := "Привет, мир!"
	// len() считает кол-во байт в строке
	fmt.Println(len(s)) // -> 21
	// Правильный способ посчитать кол-во символов в строке
	fmt.Println(utf8.RuneCountInString(s)) // -> 12

	var str string
	fmt.Scan(&str)
	fmt.Println(utf8.RuneCountInString(str))

	sum := 0
	for _, v := range str {
		if unicode.IsDigit(v) {
			sum++
		}
	}
	fmt.Println("Количество цифр в строке:", sum)
}
