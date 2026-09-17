package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Реализуйте сортировку слиянием.
На каждом шаге делите массив на две части,
сортируйте их независимо и сливайте.
*/

// merge выполняет слияние двух отсортированных массивов
func merge(a []int, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы
	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}

	return result
}

// mergeSort рекурсивно сортирует массив методом слияния
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Читаем N
	if !scanner.Scan() {
		return
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil || n < 0 {
		return
	}

	// Читаем массив из N чисел
	arr := make([]int, 0, n)
	for i := 0; i < n && scanner.Scan(); i++ {
		val, err := strconv.Atoi(scanner.Text())
		if err == nil {
			arr = append(arr, val)
		}
	}

	// Сортируем
	sorted := mergeSort(arr)

	// Выводим результат
	out := make([]string, len(sorted))
	for i, v := range sorted {
		out[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(out, " "))
}
