package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Дан неориентированный граф.
// Найдите длину минимального пути между двумя вершинами.

// readInput считывает входные данные из стандартного ввода.
func readInput() (int, [][]bool, int, int) {
	scanner := bufio.NewScanner(os.Stdin)

	// Сначала читаем количество вершин N.
	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	// Затем создаём матрицу смежности размером N*N.
	graph := make([][]bool, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}

	// Заполняем матрицу.
	for i := 0; i < n; i++ {
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		for j, v := range row {
			val, _ := strconv.Atoi(v)
			// наличие ребра = 1
			if val == 1 {
				graph[i][j] = true
			} else {
				graph[i][j] = false
			}
		}
	}

	// В конце читаем начальную и конечную вершины.
	// Между ними нужно найти путь.
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	// Переводим номера с 1-based на 0-based индексацию.
	return n, graph, start - 1, end - 1
}

// findShortestPath использует алгоритм поиска в ширину (BFS).
func findShortestPath(n int, graph [][]bool, start, end int) int {
	// Если нам нужно найти путь от вершины к самой себе.
	if start == end {
		return 0
	}

	// visited хранит информацию о том, были ли мы уже в этой вершине.
	visited := make([]bool, n)

	// distance хранит минимальное расстояние от начальной вершины до текущей.
	distance := make([]int, n)
	for i := range distance {
		// Инициализируем значением (-1), т.к. сначала пути нет.
		distance[i] = -1
	}

	// Очередь для BFS. Храним индексы вершин.
	queue := []int{start}
	visited[start] = true
	distance[start] = 0 // старт 0 шагов

	// Основной цикл BFS.
	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:] // достаём первую вершину из очереди

		// Проверка всех соседей текущей вершины.
		for neighbor := 0; neighbor < n; neighbor++ {
			// Условие: между вершинами есть ребро и мы не были там.
			if !graph[currentVertex][neighbor] || visited[neighbor] {
				continue // пропускаем эту итерацию цикла
			}

			// Если нашли новый путь к vertex:
			// то его длина равна длине пути до currentVertex + 1.
			distance[neighbor] = distance[currentVertex] + 1
			visited[neighbor] = true
			queue = append(queue, neighbor)

			// Если эта вершина оказалась нашей целью,
			// выводим ответ.
			if neighbor == end {
				return distance[end]
			}
		}
	}

	// Если мы вышли из цикла,
	// а цель не найдена, возвращаем -1.
	return -1
}

func main() {
	// Чтение входных данных.
	n, graph, start, end := readInput()

	// Поиск пути.
	result := findShortestPath(n, graph, start, end)
	fmt.Println(result)
}
