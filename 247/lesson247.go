package main

import (
	"fmt"
	"sort"
)

/*
Игровое поле представляет собой квадрат из N×N клеток,
на котором расположено N кораблей
(каждый корабль занимает одну клетку).
Определите минимальное количество ходов,
необходимых для построения кораблей в одном столбце.
*/

func main() {
	var n int
	fmt.Scan(&n)

	rows := make([]int, n) // Строки
	cols := make([]int, n) // Столбцы

	// Считывание исходных координат всех кораблей
	for i := 0; i < n; i++ {
		fmt.Scan(&rows[i], &cols[i])
	}

	// Считаем минимальное количество ходов по вертикали
	sort.Ints(rows) // Сортируем координаты
	movesVertical := 0
	for i := 0; i < n; i++ {
		// Целевая строка для i-го корабля — это i + 1
		targetRow := i + 1
		// Разница между текущей и целевой позицией
		diff := rows[i] - targetRow
		// Берем модуль разницы
		if diff < 0 {
			diff = -diff
		}
		movesVertical += diff
	}

	// Считаем минимальное количество ходов по горизонтали
	sort.Ints(cols)        // Сортируем координаты
	medianCol := cols[n/2] // Находим центральный элемент (медиану)
	movesHorizontal := 0
	for i := 0; i < n; i++ {
		diff := cols[i] - medianCol
		if diff < 0 {
			diff = -diff
		}
		movesHorizontal += diff
	}

	// Выводим итоговое минимальное количество ходов
	fmt.Println(movesVertical + movesHorizontal)
}
