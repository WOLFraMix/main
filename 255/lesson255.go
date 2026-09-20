package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Пещера представлена кубом,
разбитым на N частей по каждому измерению (куб).
Каждая клетка может быть или пустой,
или полностью заполненной камнем.
Исходя из положения спелеолога в пещере,
требуется найти минимальное количество перемещений по клеткам,
чтобы выбраться на поверхность.
*/

// Point — координата точки
type Point struct {
	z, y, x int
}

// QueueItem — элемент очереди для алгоритма BFS,
// хранит позицию и пройденное расстояние
type QueueItem struct {
	point    Point
	distance int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение N
	line, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(line))

	// Карта пещеры
	cave := make([][][]byte, n)
	for i := 0; i < n; i++ {
		cave[i] = make([][]byte, n)
		for j := 0; j < n; j++ {
			cave[i][j] = make([]byte, n)
		}
	}

	var start Point // Позиция спелеолога

	// Чтение данных пещеры
	for z := 0; z < n; z++ {
		for y := 0; y < n; y++ {
			rowLine, _ := reader.ReadString('\n')
			rowLine = strings.TrimRight(rowLine, "\r\n")

			// Пропуск пустых строк
			for len(rowLine) == 0 {
				rowLine, _ = reader.ReadString('\n')
				rowLine = strings.TrimRight(rowLine, "\r\n")
			}

			// Заполнение текущей строки слоя данными
			for x, char := range rowLine {
				cave[z][y][x] = byte(char)
				if char == 'S' { // Поиск стартовой позиции
					start = Point{z, y, x}
					// Считаем стартовую клетку свободной
					cave[z][y][x] = '.'
				}
			}
		}
	}

	// Возможные направления
	directions := []Point{
		{1, 0, 0}, {-1, 0, 0}, // Вверх/Вниз по уровням
		{0, -1, 0}, {0, 1, 0}, // Север/Юг (строки)
		{0, 0, -1}, {0, 0, 1}, // Запад/Восток (столбцы)
	}

	// Посещённые клетки
	visited := make([][][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			visited[i][j] = make([]bool, n)
		}
	}

	// Очередь для реализации поиска в ширину (BFS)
	queue := make([]QueueItem, 0, n*n*n)
	// Добавляем стартовую точку с расстоянием 0
	queue = append(queue, QueueItem{start, 0})
	// Помечаем её как посещенную
	visited[start.z][start.y][start.x] = true

	// Основной цикл обхода графа пещеры
	for len(queue) > 0 {
		// Извлекаем первый элемент
		current := queue[0]
		queue = queue[1:] // сдвигаем очередь

		currP := current.point
		currD := current.distance

		// Проверка если мы уже на верхнем уровне (z == 0)
		if currP.z == 0 {
			fmt.Println(currD)
			return
		}

		// Перебор всех соседних клеток
		for _, dir := range directions {
			nz := currP.z + dir.z
			ny := currP.y + dir.y
			nx := currP.x + dir.x

			// Проверка выхода за границы куба
			if nz < 0 || nz >= n || ny < 0 || ny >= n || nx < 0 || nx >= n {
				continue
			}

			// Проверка на препятствие (камень '#') и уже посещенные клетки
			if cave[nz][ny][nx] == '#' || visited[nz][ny][nx] {
				continue
			}

			visited[nz][ny][nx] = true // Помечаем новую клетку как посещенную
			// Добавляем соседа в конец очереди с увеличенным на 1 расстоянием
			queue = append(queue, QueueItem{Point{nz, ny, nx}, currD + 1})
		}
	}

	// Если очередь опустела и выход не найден
	fmt.Println(-1)
}
