package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Напишите программу,
которая принимает на вход число n
и последовательность из n пар
ключ (тип int) - значение (тип int),
сохраняет их в map и выводит содержимое.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Считываем n
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n, _ := strconv.Atoi(line)

	// Создаём map для хранения пар ключ-значение
	data := make(map[int]int, n)

	// Считываем n пар
	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		parts := strings.Fields(line)
		key, _ := strconv.Atoi(parts[0])
		value, _ := strconv.Atoi(parts[1])
		data[key] = value
	}

	// Выводим содержимое map
	fmt.Println("Содержимое map:")
	for k, v := range data {
		fmt.Printf("key: %d, value: %d\n", k, v)
	}
}
