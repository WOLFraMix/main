package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

var y int

// main — точка входа в программу. Генерирует случайное число, запрашивает ввод пользователя и запускает игровой цикл.
func main() {
	y = rand.IntN(100) + 1
	fmt.Println("Компьютер загадал случайное число от 1 до 100 включительно. Угадайте его!")

	num := inputNum()
	randomNumGame(num)
}

// inputNum обрабатывает ввод пользователя, проверяет его корректность и возвращает угаданное число.
// Также проверяет команду завершения игры ("выход").
func inputNum() (num int) {
	var input string
	fmt.Print("Ваше предположение (либо, для завершения, введите слово выход): ")
	fmt.Scanln(&input)
	if input == "выход" {
		os.Exit(1)
	}

	num, err := checkInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	return num
}

// checkInput валидирует строку ввода пользователя.
// Удаляет пробелы, преобразует в целое число и проверяет, находится ли оно в диапазоне [1, 100].
func checkInput(s string) (int, error) {
	s = strings.TrimSpace(s)

	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Не целое число: %w", err) // сюда попадут ошибки - "3.14", "abc" и т.д.
	}

	if v < 1 || v > 100 {
		return 0, fmt.Errorf("Некорректное значение: %d (должно быть от 1 до 100)", v)
	}

	return v, nil
}

// randomNumGame запускает основной игровой цикл (до 20 попыток).
// Сравнивает ввод пользователя с загаданным числом, выводит подсказки и обрабатывает условия победы или поражения.
func randomNumGame(num int) {
	for i := 1; i <= 20; i++ {
		if num < y {
			fmt.Println("Загаданное число больше.")
		}
		if num > y {
			fmt.Println("Загаданное число меньше.")
		}
		if num == y {
			fmt.Printf("Правильно! Вы угадали число с %d попытки.\n", i)
			fmt.Print("Хотите сыграть ещё раз? (да/нет): ")
			var repeat string
			fmt.Scanln(&repeat)
			if repeat == "да" {
				main() // Рекурсивный перезапуск игры
			} else {
				fmt.Println("Спасибо за игру! До свидания!")
				os.Exit(1)
			}
		}
		num = inputNum()
		if i == 20 {
			fmt.Println("Слишком много попыток:", i)
			fmt.Println("Загаданное число было:", y)
			os.Exit(1)
		}
	}
}
