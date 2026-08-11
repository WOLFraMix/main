package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
https://new.contest.yandex.ru/contests/80784/problems?id=149944%2F2025_08_30%2F0psuPDpNE3&contestId=80784
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	q, _ := strconv.Atoi(scanner.Text())

	var list []int
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 0; i < q; i++ {
		if !scanner.Scan() {
			break
		}
		parts := strings.Fields(scanner.Text())
		t, _ := strconv.Atoi(parts[0])

		switch t {
		case 1: // 1 x y
			x, _ := strconv.Atoi(parts[1])
			y, _ := strconv.Atoi(parts[2])
			// Позиции с 1, x=0 означает «в начало»
			idx := x // после x-ого элемента: индекс = x
			// Если x=0, то idx=0 — вставка в начало
			// Сдвигаем элементы справа от idx на одну позицию вправо
			list = append(list, 0) // увеличиваем ёмкость
			copy(list[idx+1:], list[idx:])
			list[idx] = y

		case 2: // 2 x
			x, _ := strconv.Atoi(parts[1])
			// позиции с 1
			fmt.Fprintln(out, list[x-1])

		case 3: // 3 x
			x, _ := strconv.Atoi(parts[1])
			idx := x - 1 // позиции с 1 -> индекс с 0
			// Удаляем элемент по индексу idx
			copy(list[idx:], list[idx+1:])
			list = list[:len(list)-1]
		}
	}
}
