/*
Дан массив строк s1, s2, ... sn.
Требуется вывести все строки,
которые встречаются наибольшее количество раз в этом массиве,
в лексикографическом порядке.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) // Сканируем по словам.

	// Читаем пока не достигнем конца.
	if !scanner.Scan() {
		return
	}

	var n int // Записываем.
	fmt.Sscan(scanner.Text(), &n)

	// Создаём мапу (словарь), где:
	// ключ — это слово из ввода;
	// значение — сколько раз слово встретилось.
	freq := make(map[string]int, n)

	// Читаем n строк и считаем частоту.
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		s := scanner.Text()
		freq[s]++
	}

	// Находим максимальную частоту.
	maxCount := 0
	for _, count := range freq {
		if count > maxCount {
			maxCount = count
		}
	}

	// Собираем все строки с максимальной частотой.
	result := make([]string, 0, len(freq))
	for s, count := range freq {
		if count == maxCount {
			result = append(result, s)
		}
	}

	// Сортируем в лексикографическом порядке.
	sort.Strings(result)

	// Выводим результат.
	for _, s := range result {
		fmt.Println(s)
	}
}
