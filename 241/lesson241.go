package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Рассмотрим последовательность целых чисел длины n.
По ней с шагом 1 двигается «окно» длины k,
то есть сначала в «окне» видны первые k чисел,
на следующем шаге в «окне» уже будут находиться k чисел,
начиная со второго, и так далее до конца последовательности.
Требуется для каждого положения «окна» определить минимум в нём.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	// Создаём слайс для хранения последовательности чисел
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		arr[i] = num
	}

	// minQueue будет хранить индексы элементов массива
	minQueue := make([]int, 0, k)

	// Инициализация первого окна
	for i := 0; i < k; i++ {
		for len(minQueue) > 0 && arr[i] <= arr[minQueue[len(minQueue)-1]] {
			minQueue = minQueue[:len(minQueue)-1]
		}
		minQueue = append(minQueue, i)
	}

	// Первый минимум
	fmt.Println(arr[minQueue[0]])

	// Обработка остальных окон
	for i := k; i < n; i++ {
		// Удаляем индекс, который выходит из окна
		if minQueue[0] <= i-k {
			minQueue = minQueue[1:]
		}

		// Поддерживаем монотонность очереди
		for len(minQueue) > 0 && arr[i] <= arr[minQueue[len(minQueue)-1]] {
			minQueue = minQueue[:len(minQueue)-1]
		}
		// Добавляем индекс текущего элемента в очередь
		minQueue = append(minQueue, i)

		// Минимум в текущем окне
		fmt.Println(arr[minQueue[0]])
	}
}
