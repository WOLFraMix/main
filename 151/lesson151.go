package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Ввод
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	// Проверка правильности ввода
	readInt := func() (int, error) {
		line, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}
		line = strings.TrimSpace(line)
		v, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		return v, nil
	}

	// Создаем переменные
	m, _ := readInt()
	/*
		n, _ := readInt()
		a, _ := readInt()
		b, _ := readInt()
	*/

	/*
		Решение задачи
	*/

	result := m // Результат

	writer.WriteString(strconv.Itoa(result)) // Вывод
	writer.WriteByte('\n')
}
