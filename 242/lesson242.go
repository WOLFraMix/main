package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
В этой задаче вам необходимо самостоятельно
(не используя соответствующие классы и функции стандартной библиотеки)
организовать структуру данных Heap (кучу) для хранения целых чисел.
*/

// MaxHeap — структура для хранения элементов в виде максимальной кучи
type MaxHeap struct {
	heap []int // массив, представляющий кучу
}

// Insert добавляет элемент в кучу и восстанавливает порядок
func (h *MaxHeap) Insert(k int) {
	h.heap = append(h.heap, k)
	i := len(h.heap) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if h.heap[i] <= h.heap[parent] {
			break
		}
		// Поднимаем элемент вверх
		h.heap[i], h.heap[parent] = h.heap[parent], h.heap[i]
		i = parent
	}
}

// Extract достает максимальный элемент из кучи и удаляет его
func (h *MaxHeap) Extract() int {
	n := len(h.heap)
	if n == 0 {
		return 0 // или любое другое значение по умолчанию
	}
	maxVal := h.heap[0]
	last := h.heap[n-1]
	h.heap = h.heap[:n-1]
	if n > 1 {
		h.heap[0] = last
		h.siftDown(0)
	}
	return maxVal
}

// siftDown опускает элемент вниз до нужной позиции
func (h *MaxHeap) siftDown(i int) {
	n := len(h.heap)
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < n && h.heap[left] > h.heap[largest] {
			largest = left
		}
		if right < n && h.heap[right] > h.heap[largest] {
			largest = right
		}
		if largest != i {
			h.heap[i], h.heap[largest] = h.heap[largest], h.heap[i]
			i = largest
		} else {
			break
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем количество команд
	scanner.Scan()
	var N int
	fmt.Sscan(scanner.Text(), &N)

	// Создаем пустую кучу
	heap := MaxHeap{}

	// Обрабатываем команды
	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()

		// Операция извлечения
		if line == "1" {
			fmt.Println(heap.Extract())
		} else { // Операция вставки
			_, numStr := line[0], line[2:]
			var num int
			fmt.Sscanf(numStr, "%d", &num)
			heap.Insert(num)
		}
	}
}
