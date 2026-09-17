package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Базовым алгоритмом для быстрой сортировки является алгоритм partition,
который разбивает набор элементов на две части относительно заданного предиката.
По сути элементы массива просто меняются местами так,
что левее некоторой точки в нем после этой операции лежат элементы,
удовлетворяющие заданному предикату,
а справа — не удовлетворяющие ему.
Например, при сортировке можно использовать предикат «меньше опорного».
*/

func main() {
	// Используем bufio для быстрого чтения
	reader := bufio.NewReader(os.Stdin)

	// Чтение n - количества элементов массива
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// Чтение массива элементов
	line, _ = reader.ReadString('\n')
	parts := strings.Fields(line)

	// Чтение опорного элемента x
	line, _ = reader.ReadString('\n')
	x, _ := strconv.Atoi(strings.TrimSpace(line))

	// Количество элементов меньших чем х
	countLess := 0

	for _, p := range parts {
		// Преобразуем строки в числа прямо в цикле
		val, _ := strconv.Atoi(p)
		if val < x {
			countLess++
		}
	}

	// Вывод результатов
	fmt.Println(countLess)     // Меньше х
	fmt.Println(n - countLess) // Остальные
}
