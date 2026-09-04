package main

import (
	"bufio"
	"fmt"
	"os"
)

// Дана прямоугольная доска N×M (N строк и M столбцов).
// В левом верхнем углу находится шахматный конь,
// которого необходимо переместить в правый нижний угол доски.
// Необходимо определить, сколько существует различных маршрутов.

// readInput считывает входных данные:
// N строк и M столбцов.
func readInput() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n, m int
	fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)
	return n, m
}

// solveKnightPaths решает задачу
// методом динамического программирования.
func solveKnightPaths(n, m int) int64 {
	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, m+1)
	}

	// Начальная клетка
	dp[1][1] = 1

	// Заполняем таблицу
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// Ход "вниз-вправо" (2, 1)
			if i >= 2 && j >= 1 {
				dp[i][j] += dp[i-2][j-1]
			}
			// Ход "влево-вниз" (1, 2)
			if i >= 1 && j >= 2 {
				dp[i][j] += dp[i-1][j-2]
			}
		}
	}
	return dp[n][m]
}

func main() {
	n, m := readInput()
	result := solveKnightPaths(n, m)
	fmt.Println(result)
}
