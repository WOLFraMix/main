package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	/*
		В первой строке вводится целое число m — наибольшее значение, которое может отображать счётчик.
		Во второй строке вводится целое число a, которое изначально отображал счётчик.
		В третьей строке вводится целое число b — число, которое должно отображаться на счётчике.
	*/

	// Чтение m
	line1, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(line1))

	// Чтение a
	line2, _ := reader.ReadString('\n')
	a, _ := strconv.Atoi(strings.TrimSpace(line2))

	// Чтение b
	line3, _ := reader.ReadString('\n')
	b, _ := strconv.Atoi(strings.TrimSpace(line3))

	// Циклический счётчик от 1 до m
	count := 0
	i := a
	// Двигаемся счётчик, пока не достигнем b
	for i != b {
		i++
		count++
		// Сбрасываем счётчик если он переполнился
		if i > m {
			i = 1
		}
	}

	writer.WriteString(strconv.Itoa(count))
	writer.WriteByte('\n')

}
