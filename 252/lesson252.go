package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Определите, сколько обменов
сделает алгоритм пузырьковой сортировки
по возрастанию для данного массива.
Метод через сортировку слиянием.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Считываем размер массива
	var n int
	fmt.Fscan(reader, &n)

	// Заполняем массив элементами
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	var swaps int // Счётчик
	// Временный буфер для промежуточных результатов слияния
	temp := make([]int, n)

	// Рекурсивная функция сортировки слиянием
	var mergeSort func(l, r int)
	mergeSort = func(l, r int) {
		if l >= r {
			return
		}
		m := (l + r) / 2  // Находим середину текущего отрезка
		mergeSort(l, m)   // Сортируем левую половину
		mergeSort(m+1, r) // Сортируем правую половину

		// Слияние с подсчётом инверсий
		// i — указатель на левый срез, j — на правый, k — на temp
		i, j, k := l, m+1, l
		for i <= m && j <= r {
			if arr[i] <= arr[j] {
				temp[k] = arr[i]
				i++
			} else {
				temp[k] = arr[j]
				j++
				swaps += int(m - i + 1)
			}
			k++
		}

		// Копируем оставшиеся элементы,
		// если одна из половин закончилась раньше другой
		for i <= m {
			temp[k] = arr[i]
			i++
			k++
		}
		for j <= r {
			temp[k] = arr[j]
			j++
			k++
		}

		// Переносим отсортированный и слитый результат
		// обратно в основной массив
		for idx := l; idx <= r; idx++ {
			arr[idx] = temp[idx]
		}
	}

	mergeSort(0, n-1)
	fmt.Println(swaps)
}
