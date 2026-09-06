package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
На клеточном поле, размером N*M сидит Q блох.
«Прием пищи» блохами возможен только в кормушке.
Блохи перемещаются как шахматный конь.
Определить минимальное значение суммы длин путей блох до кормушки.
Вывод −1, если сбор невозможен.
В первой строке входного файла находится 5 чисел,
разделенных пробелом: N, M, S, T, Q.
N, M — размеры доски (отсчет начинается с 1).
S, T — координаты клетки — кормушки.
Q — количество блох на доске.
Далее Q строк по два числа — координаты каждой блохи.
*/

// Структура для хранения координат клетки при поиске в ширину.
type Point struct {
	r, c int // Координаты строки и столбца
	dist int // Расстояние до кормушки
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Читаем размеры поля и параметры задачи.
	var N, M, S, T, Q int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &M)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &S)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &T)
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &Q)

	// Переводим координаты из 1-based в 0-based.
	startR, startC := S-1, T-1

	// Матрица расстояний: -1 — клетка ещё не посещена.
	dist := make([][]int, N)
	for i := range dist {
		dist[i] = make([]int, M)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	// Возможные ходы шахматного коня (8 направлений).
	moves := [][]int{
		{2, 1}, {2, -1},
		{-2, 1}, {-2, -1},
		{1, 2}, {1, -2},
		{-1, 2}, {-1, -2},
	}

	// BFS (Поиск в ширину).
	// Находим кратчайшие пути от кормушки ко всем клеткам.
	queue := []Point{{startR, startC, 0}}
	dist[startR][startC] = 0

	head := 0
	for head < len(queue) {
		current := queue[head]
		head++

		// Пробуем сделать все возможные прыжки.
		for _, move := range moves {
			nr := current.r + move[0]
			nc := current.c + move[1]

			if nr >= 0 && nc >= 0 && nr < N && nc < M && dist[nr][nc] == -1 {
				// Если клетка доступна и мы её ещё не посетили.
				dist[nr][nc] = current.dist + 1
				queue = append(queue, Point{nr, nc, current.dist + 1})
			}
		}
	}

	// Подсчёт суммы путей всех блох.
	totalSum := 0
	for i := 0; i < Q*2; i += 2 {
		scanner.Scan()
		bugR := scanner.Text()
		scanner.Scan()
		bugC := scanner.Text()

		// Преобразование к 0-based прямо во время чтения.
		r := toZeroBased(bugR)
		c := toZeroBased(bugC)

		if dist[r][c] == -1 {
			fmt.Println(-1)
			return
		}

		totalSum += dist[r][c]
	}

	fmt.Println(totalSum)
}

// Вспомогательная функция для перевода из текстового вида в число,
// вычитая единицу для перехода к индексации с нуля.
func toZeroBased(s string) int {
	num, _ := strconv.Atoi(s)
	return num - 1
}
