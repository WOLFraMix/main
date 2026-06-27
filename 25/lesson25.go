package main

import (
	"fmt"
	"unicode/utf8"
)

// string
func main() {
	var str string = "first"
	fmt.Println(str)

	str2 := "second"
	fmt.Println(str2)

	var str3 string // empty string
	fmt.Println(str3)

	str4 := "" // пустая строка
	fmt.Println(str4)

	str5 := "hello \nworld" // перенос строки это \n
	fmt.Println(str5)

	str6 := `my 
name 
	is 
Stepan` // вывод на каждую строку (включая табуляцию)
	fmt.Println(str6)

	str7 := "VScode" // в английском языке 1 буква занимает 1 байт
	fmt.Println(str7)
	fmt.Println(len(str7)) // подсчёт символов строки через функцию len в байтах

	str8 := "русский" // русский язык по 2 байта
	fmt.Println(str8)
	fmt.Println(len(str8))         // подсчёт символов строки через функцию len в байтах
	fmt.Println(len([]rune(str8))) // подсчёт через руны

	str9 := "hello"
	fmt.Println("str9 это:", str9) // вывод сообщения и строки

	str10 := "привет"                                                            // через импортируемую функцию
	fmt.Println(str10, "- в этом тексте:", utf8.RuneCountInString(str10), "рун") // подсчёт через руны

	str11 := "\thello \n\tworld" // перенос строки \n табуляция \t
	fmt.Println(str11)

	str12 := "привет мой \"друг\" запиши \\это\\" // экранирование
	fmt.Println(str12)
}
