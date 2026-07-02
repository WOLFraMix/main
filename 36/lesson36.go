package main

import (
	"fmt"
	"strings"
)

// Конкатенация - это операция «склеивания» последовательностей, чаще всего строк.
func main() {
	str1 := "Hello,"
	str2 := " world!"
	fmt.Println(str1 + str2) // первый вариант
	result := str1 + str2
	fmt.Println(result) // второй вариант
	result2 := fmt.Sprintf("%s%s", str1, str2)
	fmt.Println(result2) // третий вариант

	var buffer strings.Builder // вариант через структуры и builder чтобы не расходовать память
	buffer.WriteString(str1)
	buffer.WriteString(str2)
	buffer.WriteString(" И это ещё сюда добавим")
	fmt.Println(buffer.String()) // выводим результат
	// или создаём переменную а потом выводим
	result3 := buffer.String()
	fmt.Println(result3)
}
