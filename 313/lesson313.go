package main

import (
	"fmt"
	"regexp"
	"strconv"
)

/*
На вход программы подается строка.
Найдите все числа в строке и выведите их сумму.
*/

func main() {
	var input string
	fmt.Scan(&input)

	// Компиляция регулярного выражения для поиска целых чисел
	// (включая отрицательные)
	re := regexp.MustCompile(`-?\d+`)

	// Поиск всех вхождений чисел в строке
	// -1 означает "найти все совпадения"
	matches := re.FindAllString(input, -1)
	sum := 0
	for _, m := range matches {
		// Преобразование строки в целое число
		num, _ := strconv.Atoi(m)
		sum += num
	}
	// Вывод итоговой суммы
	fmt.Println(sum)
}
