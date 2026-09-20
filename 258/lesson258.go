package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Напишите программу, которая принимает на вход число n
и последовательность из n целых чисел
и выводит на экран количество раз,
которое каждое число встречается в этой последовательности.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n, _ := strconv.Atoi(line)

	data := make(map[int]int, n)

	line, _ = reader.ReadString('\n')
	parts := strings.Fields(line)

	for i := 0; i < n; i++ {
		key, _ := strconv.Atoi(parts[i])
		data[key]++
	}

	fmt.Println(data)
}
