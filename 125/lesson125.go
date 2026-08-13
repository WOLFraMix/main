package main

import "fmt"

func main() {
	str := "Зашифруй меня!"
	encodedStr := CaesarCode(str, 5, true)
	fmt.Println(encodedStr)

	decodedStr := CaesarCode(encodedStr, 5, false)
	fmt.Println(decodedStr)
}

func CaesarCode(text string, shift int, encode bool) string {

}
