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

	// Читаем N
	line, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	n, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil || n < 2 {
		return
	}

	// Читаем массив
	line, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	parts := strings.Fields(line)
	a := make([]int, 0, n)
	for _, p := range parts {
		val, err := strconv.Atoi(p)
		if err == nil {
			a = append(a, val)
		}
	}

	if len(a) < 2 {
		return
	}

	// Инициализируем первую пару соседей
	minDiff := a[1] - a[0]
	bestFirst := a[0]
	bestSecond := a[1]

	// Проходим по всем соседним парам
	for i := 1; i < len(a)-1; i++ {
		diff := a[i+1] - a[i]
		if diff < minDiff {
			minDiff = diff
			bestFirst = a[i]
			bestSecond = a[i+1]
		}
		// Если diff == minDiff, ничего не делаем — нужна первая пара
	}

	fmt.Printf("%d %d\n", bestFirst, bestSecond)
}
