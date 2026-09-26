package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Дана последовательность,
требуется найти её наибольшую возрастающую подпоследовательность.
Наибольшая возрастающая подпоследовательность -
это строго возрастающая подпоследовательность наибольшей длины.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение N
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	// Чтение последовательности
	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)

	arr := make([]int, n) // Исходный массив чисел
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(parts[i])
	}

	// dp[i] — длина НВП
	// (наибольшей возрастающей подпоследовательности),
	// которая заканчивается строго в индексе i
	dp := make([]int, n)
	// prev[i] — индекс предыдущего элемента в этой НВП
	// для восстановления пути
	prev := make([]int, n)

	// Инициализация
	for i := 0; i < n; i++ {
		// Каждый элемент сам по себе
		// является подпоследовательностью длины 1
		dp[i] = 1
		// -1 означает, что у текущего элемента нет предшественника
		prev[i] = -1
	}

	// Заполнение таблиц динамики
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// Если предыдущий элемент меньше текущего,
			// мы можем нарастить подпоследовательность
			if arr[j] < arr[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					prev[i] = j
				}
			}
		}
	}

	// Поиск индекса конца максимальной подпоследовательности
	maxLen := 0
	maxIdx := 0
	for i := 0; i < n; i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
			maxIdx = i
		}
	}

	// Восстановление подпоследовательности
	result := make([]int, 0, maxLen)
	curr := maxIdx
	for curr != -1 {
		// Добавляем текущее число в результат
		result = append(result, arr[curr])
		// Переходим к индексу предыдущего числа
		curr = prev[curr]
	}

	// Разворот результата (так как шли с конца)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// Вывод результата
	for i, val := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}
