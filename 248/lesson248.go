package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Базовый алгоритм для сортировки слиянием —
алгоритм слияния двух упорядоченных массивов
в один упорядоченный массив.
Эта операция выполняется за линейное время
с линейным потреблением памяти.
Реализуйте слияние двух массивов
в качестве первого шага для написания сортировки слиянием.
*/

// merge выполняет слияние двух отсортированных массивов
func merge(a []int, b []int) []int {
	result := make([]int, len(a)+len(b))
	// Индексы для a, b и result
	i, j, k := 0, 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result[k] = a[i]
			k++
			i++
		} else {
			result[k] = b[j]
			k++
			j++
		}
	}

	// Добавляем оставшиеся элементы
	for ; i < len(a); i++ {
		result[k] = a[i]
		k++
	}
	for ; j < len(b); j++ {
		result[k] = b[j]
		k++
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Читаем слова (числа) по одному
	scanner.Split(bufio.ScanWords)
	var n, m int

	// Чтение N и массива A
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	a := make([]int, n)
	for i := range a {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	// Чтение M и массива B
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())
	b := make([]int, m)
	for i := range b {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	// Слияние и вывод
	result := merge(a, b)
	output := make([]string, len(result))
	for i, v := range result {
		output[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(output, " "))
}
