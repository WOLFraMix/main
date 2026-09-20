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
принимающую на вход целое положительное число n
и последовательность из n целых положительных чисел,
которая выводит из этой последовательности только уникальные элементы
(уникальные элементы должны сохранять свой исходный порядок в последовательности).
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n, _ := strconv.Atoi(line)

	data := make(map[int]struct{}, n)
	var numbers []int

	line, _ = reader.ReadString('\n')
	parts := strings.Fields(line)

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(parts[i])
		numbers = append(numbers, num)
	}

	for _, num := range numbers {
		_, ok := data[num]
		if !ok {
			fmt.Print(num)
			fmt.Print(" ")
			data[num] = struct{}{}
		}
	}
	fmt.Println()
}
