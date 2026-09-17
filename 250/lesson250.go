package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Отсортируйте данный массив.
Используйте пирамидальную сортировку.
*/

// siftDown просеивает элемент вниз по куче,
// чтобы восстановить свойство max-heap
func siftDown(arr []int, start, end int) {
	root := start
	for root*2+1 <= end {
		child := root*2 + 1 // левый потомок
		swap := root        // индекс элемента для обмена

		// выбираем большего из потомков
		if arr[swap] < arr[child] {
			swap = child
		}
		// проверяем наличие правого потомка и сравниваем
		if child+1 <= end && arr[swap] < arr[child+1] {
			swap = child + 1
		}

		// если корень уже больше потомков, куча восстановлена
		if swap == root {
			return
		}

		// меняем местами и продолжаем просеивание
		arr[root], arr[swap] = arr[swap], arr[root]
		root = swap
	}
}

// heapSort выполняет пирамидальную сортировку
func heapSort(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}

	// Просеиваем все внутренние узлы снизу вверх
	for i := n/2 - 1; i >= 0; i-- {
		siftDown(arr, i, n-1)
	}

	// Меняем корень с последним элементом и уменьшаем размер кучи
	for end := n - 1; end > 0; end-- {
		// Меняем местами именно элементы, а не срез целиком
		arr[0], arr[end] = arr[end], arr[0]
		// Восстанавливаем свойство кучи для оставшейся части
		siftDown(arr, 0, end-1)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение N
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// Чтение массива
	arr := make([]int, n)
	if n > 0 {
		line, _ = reader.ReadString('\n')
		parts := strings.Fields(line)
		// Защита от ситуации, когда чисел меньше, чем ожидалось
		limit := n
		if len(parts) < n {
			limit = len(parts)
		}
		for i := 0; i < limit; i++ {
			val, convErr := strconv.Atoi(parts[i])
			if convErr != nil {
				// Если число не удалось распарсить, считаем его нулем
				val = 0
			}
			arr[i] = val
		}
	}

	// Сортировка
	heapSort(arr)

	// Вывод результата
	writer := bufio.NewWriter(os.Stdout)
	for i, v := range arr {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
