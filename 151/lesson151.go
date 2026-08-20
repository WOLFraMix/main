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

	m, _ := readInt()

	/*
		нужные переменные
		a, _ := readInt()
		b, _ := readInt()
	*/

	/*
		логика задачи
	*/

	result := m // результат

	writer.WriteString(strconv.Itoa(result)) // вывод
	writer.WriteByte('\n')
}
