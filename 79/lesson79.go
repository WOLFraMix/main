package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Считывает ввод пользователя, подсчитывает типы символов и выводит результаты.
func main() {
	// Получаем текст от пользователя
	text, err := GetInput()
	if err != nil {
		fmt.Println("Ошибка")
		os.Exit(1)
	}
	// Подсчитываем количество букв, цифр, пробелов и знаков препинания
	letters, digits, spaces, punctuation := CountCharacters(text)
	// Выводим результаты подсчёта
	DisplayResults(letters, digits, spaces, punctuation)
}

// GetInput считывает непустую строку со стандартного ввода.
// Повторяет запрос до тех пор, пока пользователь не введёт корректный текст.
func GetInput() (string, error) {
	fmt.Println("Ввод сообщения:")
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err // например, EOF или другая ошибка ввода
		}

		// Удаляем символы перевода строки и лишние пробелы по краям
		line = strings.TrimSpace(line)
		// Если после этого осталась только пустая строка, просим ввести текст заново
		if line == "" {
			fmt.Println("Строка не может быть пустой. Попробуйте ещё раз:")
			continue
		}

		return line, nil
	}
}

// CountCharacters анализирует строку и возвращает количество букв, цифр, пробелов и знаков препинания.
func CountCharacters(text string) (letters, digits, spaces, punctuation int) {
	// Считаем общее количество символов Юникода в строке
	letters = utf8.RuneCountInString(text)
	digits = 0
	// Считаем цифры, перебирая каждый символ
	for _, r := range text {
		if unicode.IsDigit(r) {
			digits++
		}
	}
	// Считаем пробелы с помощью стандартной функции
	spaces = strings.Count(text, " ")
	punctuation = 0
	// Считаем знаки препинания, перебирая каждый символ
	for _, r := range text {
		if unicode.IsPunct(r) {
			punctuation++
		}
	}
	return letters, digits, spaces, punctuation
}

// DisplayResults выводит результаты подсчёта символов на экран.
func DisplayResults(letters, digits, spaces, punctuation int) {
	fmt.Printf("Количество букв: %d\nКоличество цифр: %d\nКоличество пробелов: %d\nКоличество знаков препинания: %d\n", letters, digits, spaces, punctuation)
}
