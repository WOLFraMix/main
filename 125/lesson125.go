package main

import "fmt"

func main() {
	str := "Зашифруй меня!"                // шифр
	encodedStr := CaesarCode(str, 5, true) // зашифровка
	fmt.Println(encodedStr)

	decodedStr := CaesarCode(encodedStr, 5, false) // дешифровка
	fmt.Println(decodedStr)
}

// функция которая шифрует и дешифрует текст по алгоритму шифра Цезаря
func CaesarCode(text string, shift int, encode bool) string {
	// text - текст, который нужно зашифровать или расшифровать
	// shift - число, на которое нужно сдвинуть символы
	// encode - нужно зашифровать (true) или дешифровать (false) сообщение
	var result string

	if encode == true {
		for _, v := range text {
			v += rune(shift)
			result += string(v)
		}
	}

	if encode == false {
		for _, v := range text {
			v -= rune(shift)
			result += string(v)
		}
	}

	return result
}
