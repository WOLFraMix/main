package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Рассмотрим последовательность, состоящую из:
круглых, квадратных и фигурных скобок.
Программа должна определить,
является ли данная скобочная последовательность правильной.
*/

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	line, _ := reader.ReadString('\n')
	n := strings.TrimSpace(line)

	// Инициализируем пустой "стек" как срез байтов
	stack := make([]byte, 0, len(n))

	// Карта соответствий закрывающих скобок к открытым
	m := map[byte]byte{
		']': '[',
		')': '(',
		'}': '{',
	}

	// Проходим по каждому символу строки
	for i := 0; i < len(n); i++ {
		c := n[i]

		// Открывающую скобку вносим в стек
		if c == '[' || c == '(' || c == '{' {
			stack = append(stack, c)
		} else { // Закрывающая скобка
			// Проверяем, есть ли что открывать в стеке
			if len(stack) == 0 {
				fmt.Fprintln(writer, "no")
				return
			}

			// Берем последнюю открытую скобку из стека
			top := stack[len(stack)-1]
			// Удаляем её из стека
			stack = stack[:len(stack)-1]

			// Сравниваем пары
			if m[c] != top {
				fmt.Fprintln(writer, "no")
				return
			}
		}
	}

	// Если после прохода стек пуст — всё хорошо
	if len(stack) == 0 {
		fmt.Fprintln(writer, "yes")
	} else {
		fmt.Fprintln(writer, "no")
	}
}
