package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
Реализуйте быструю сортировку, используя алгоритм partition.
На каждом шаге выбирайте опорный элемент
и выполняйте partition относительно него.
Затем рекурсивно запуститесь от двух частей,
на которые разбился исходный массив.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Читаем количество элементов
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	if n == 0 {
		// Если массив пуст, завершаем программу
		return
	}

	// Читаем массив
	line, _ = reader.ReadString('\n')
	parts := strings.Fields(line)

	// Создаём слайс
	arr := make([]int, 0, n)

	// Заполняем слайс
	for _, p := range parts {
		// Преобразуем строки в числа прямо в цикле
		val, _ := strconv.Atoi(p)
		arr = append(arr, val)
	}

	result := quickSort(arr)
	for i, v := range result {
		fmt.Print(v)
		if i < len(result)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// quickSort сортирует слайс разделяя его на 3 части
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Выбираем три кандидата на роль опорного элемента
	l := arr[0]
	m := arr[(len(arr)-1)/2]
	r := arr[len(arr)-1]

	// Помещаем кандидатов в маленький слайс
	// и сортируем, чтобы найти медиану
	medianArr := []int{l, m, r}
	slices.Sort(medianArr)
	median := medianArr[1]

	// Трёхстороннее разбиение:
	// < median, == median, > median
	left := make([]int, 0, len(arr))
	middle := make([]int, 0, len(arr))
	right := make([]int, 0, len(arr))

	for _, v := range arr {
		if v < median {
			left = append(left, v)
		} else if v == median {
			middle = append(middle, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	result := make([]int, 0, len(arr))
	result = append(result, left...)
	result = append(result, middle...)
	result = append(result, right...)
	return result
}
