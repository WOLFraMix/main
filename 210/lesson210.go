/*
В левом верхнем углу прямоугольной таблицы размером N×M находится черепашка.
В каждой клетке таблицы записано некоторое число.
Черепашка может перемещаться вправо или вниз,
при этом маршрут черепашки заканчивается в правом нижнем углу таблицы.
Подсчитаем сумму чисел, записанных в клетках,
через которую проползла черепашка (включая начальную и конечную клетку).
Найдите наибольшее возможное значение этой суммы и маршрут, на котором достигается эта сумма.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type cell struct {
	sum int
	dir byte // 'D' — пришли сверху, 'R' — слева; для (0,0) не используется
}

// readInput возвращает три значения:
// размеры таблицы n (строки), m (столбцы) и саму таблицу grid.
func readInput() (int, int, [][]int) {
	// Создаём сканер для чтения из стандартного ввода.
	scanner := bufio.NewScanner(os.Stdin)

	// Читаем по словам (пробелы и переносы строк — это разделители).
	scanner.Split(bufio.ScanWords)

	// Если ввода нет — паникует.
	if !scanner.Scan() {
		panic("no input")
	}
	// Читает первое слово и конвертирует в число строк.
	n, _ := strconv.Atoi(scanner.Text())

	if !scanner.Scan() {
		panic("missing M")
	}
	// Читает второе слово и конвертирует в число столбцов.
	m, _ := strconv.Atoi(scanner.Text())

	// Создаёт двумерный слайс (таблицу) из n строк.
	grid := make([][]int, n)

	// Заполняем таблицу.
	for i := 0; i < n; i++ {
		row := make([]int, m)
		for j := 0; j < m; j++ {
			if !scanner.Scan() {
				panic("unexpected end of input")
			}
			val, _ := strconv.Atoi(scanner.Text())
			row[j] = val
		}
		grid[i] = row
	}
	return n, m, grid // Возвращает размеры и таблицу.
}

// findMaxPath принимает размеры и таблицу.
// Возвращает: путь как слайс байтов (R D) и максимальную сумму.
func findMaxPath(n, m int, grid [][]int) ([]byte, int) {
	// Создаёт таблицу dp размером n × m.
	dp := make([][]cell, n)
	for i := range dp {
		dp[i] = make([]cell, m)
	}

	// Инициализируем стартовую клетку
	dp[0][0].sum = grid[0][0]

	// Первая строка: только справа
	for j := 1; j < m; j++ {
		dp[0][j].sum = dp[0][j-1].sum + grid[0][j]
		dp[0][j].dir = 'R'
	}

	// Первый столбец: только снизу
	for i := 1; i < n; i++ {
		dp[i][0].sum = dp[i-1][0].sum + grid[i][0]
		dp[i][0].dir = 'D'
	}

	// Основное ДП
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			fromTop := dp[i-1][j].sum
			fromLeft := dp[i][j-1].sum

			// Смотрим два варианта: прийти сверху или слева.
			// Выбираем тот, где сумма больше.
			if fromTop > fromLeft {
				dp[i][j].sum = fromTop + grid[i][j]
				dp[i][j].dir = 'D'
			} else {
				// Если равны — можно выбрать любое.
				dp[i][j].sum = fromLeft + grid[i][j]
				dp[i][j].dir = 'R'
			}
		}
	}

	// Максимальная сумма для всего пути.
	maxSum := dp[n-1][m-1].sum

	// Длина пути:
	// (n-1) шагов вниз + (m-1) шагов вправо = n+m-2
	pathLen := n + m - 2
	if pathLen == 0 {
		// Таблица 1x1:
		// если маршрута нет, возвращаем пустой путь
		return []byte{}, maxSum
	}

	// Создаём слайс для пути
	// и восстанавливаем с правого нижнего угла (x, y).
	path := make([]byte, pathLen)
	x, y := n-1, m-1

	// Восстанавливаем путь с конца до (0,0).
	for k := pathLen - 1; k >= 0; k-- {
		path[k] = dp[x][y].dir
		switch path[k] {
		case 'D':
			x--
		case 'R':
			y--
		default:
			// Это не должно происходить
			panic("invalid direction")
		}
	}

	return path, maxSum // Возвращает путь и сумму.
}

func main() {
	// Считываем входные данные.
	n, m, grid := readInput()
	path, sum := findMaxPath(n, m, grid)

	// Выводит сначала максимальную сумму.
	fmt.Println(sum)
	// Затем выводит путь: буквы R и D, c пробелами.
	for i, dir := range path {
		fmt.Printf("%c", dir)
		if i < len(path)-1 {
			fmt.Print(" ")
		}
	}
}
