package main

import (
	"fmt"
	"math"
)

// s - строка с постоянным значением.
const s string = "constant"

// main демонстрирует использование констант в Go
// включая числовые операции и использование математических функций.
// Выводит на экран значение строковой константы
// результат деления
// преобразование типа
// и значение синуса.
func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
