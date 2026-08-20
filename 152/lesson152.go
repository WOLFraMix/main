package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix[i][j])
		}
	}

	ans := 0

	// Обрабатываем группы из до 4 элементов в «верхнем левом квадранте»
	for i := 0; i < n/2; i++ {
		for j := 0; j < m/2; j++ {
			// Собираем уникальные позиции, которые должны быть равны
			coords := make(map[[2]int]struct{})
			coords[[2]int{i, j}] = struct{}{}
			coords[[2]int{i, m - 1 - j}] = struct{}{}
			coords[[2]int{n - 1 - i, j}] = struct{}{}
			coords[[2]int{n - 1 - i, m - 1 - j}] = struct{}{}

			// Собираем значения в этих позициях
			values := make(map[int]int)
			for coord := range coords {
				// coord[0] — строка, coord[1] — столбец
				r, c := coord[0], coord[1]
				values[matrix[r][c]]++
			}

			// Находим максимальную частоту (моду)
			maxFreq := 0
			for _, freq := range values {
				if freq > maxFreq {
					maxFreq = freq
				}
			}

			// Количество изменений = размер группы − частота моды
			ans += len(coords) - maxFreq
		}
	}

	// Если n нечётное, обрабатываем центральную строку
	if n%2 == 1 {
		midRow := n / 2
		for j := 0; j < m/2; j++ {
			if matrix[midRow][j] != matrix[midRow][m-1-j] {
				ans++
			}
		}
	}

	// Если m нечётное, обрабатываем центральный столбец
	if m%2 == 1 {
		midCol := m / 2
		for i := 0; i < n/2; i++ {
			if matrix[i][midCol] != matrix[n-1-i][midCol] {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
