package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Лайнландия представляет из себя одномерный мир, являющийся прямой,
на котором располагаются N городов, последовательно пронумерованных от 0 до N - 1.
Направление в сторону от нулевого города названо восточным.
Когда в Лайнландии неожиданно начался кризис,
все жители мира стали испытывать глубокое смятение.
По всей Лайнландии стали ходить слухи, что на востоке живётся лучше, чем на западе.
Так и началось Великое Лайнландское переселение.
Обитатели мира целыми городами отправились на восток,
и двигались до тех пор, пока не приходили в город,
в котором средняя цена проживания была меньше, чем в родном.
*/

func readInput() (int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Читаем количество городов
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	// Читаем цены для всех городов
	prices := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		prices[i], _ = strconv.Atoi(scanner.Text())
	}

	return N, prices
}

func solveLandlineMigration(n int, prices []int) []int {
	result := make([]int, n)
	stack := make([]int, 0, n)

	// Идём справа налево — ищем ближайший меньший справа
	for i := n - 1; i >= 0; i-- {
		// Удаляем из стека элементы, которые >= чем текущий
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= prices[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			// Если не нашли подходящий город,
			// отправляем в Восточное Бесконечное Ничто (-1)
			result[i] = -1
		} else {
			result[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return result
}

func main() {
	// Считываем входные данные
	N, prices := readInput()

	// Защита на случай, если ввод 0 или некорректный
	if N == 0 || len(prices) == 0 {
		return
	}

	// Находим ответ
	result := solveLandlineMigration(N, prices)

	// Выводим результат
	for i, v := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
