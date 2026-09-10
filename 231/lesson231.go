package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
В постфиксной записи (или обратной польской записи)
операция записывается после двух операндов.
Например, сумма двух чисел A и B записывается как A B +.
*/

// Reader считывает строку
// с выражением в постфиксной записи.
func Reader() []string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	return strings.Fields(line)
}

// isOperator проверяет,
// является ли строка арифметическим оператором.
func isOperator(r string) bool {
	switch r {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}

func main() {
	s := Reader()                   // срез токенов
	stack := make([]int, 0, len(s)) // стек
	var top2 int                    // переменная для хранения второго операнда
	var top1 int                    // переменная для хранения первого операнда
	var result int

	for _, v := range s { // проходим по каждому токену
		if isOperator(v) { // если токен оператор
			top1 = stack[len(stack)-1]
			top2 = stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch v { // вычисляем
			case "+":
				result = top2 + top1
			case "-":
				result = top2 - top1
			case "*":
				result = top2 * top1
			case "/":
				result = top2 / top1
			}

			// результат возвращаем в стек
			stack = append(stack, result)

		} else { // если токен операнд
			// парсим и убираем в стек
			n, _ := strconv.Atoi(v)
			stack = append(stack, n)
		}
	}

	result = stack[len(stack)-1]
	fmt.Println(result)
}
