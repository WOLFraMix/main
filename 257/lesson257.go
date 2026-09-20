package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	data[8] = n
	data[1] += n
	delete(data, 3)
	var sum int
	for _, v := range data {
		sum += v
	}
	fmt.Println(sum)
}
