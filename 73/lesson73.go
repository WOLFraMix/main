package main

import "fmt"

func main() {
	str := "Хай. Ку-ку!" // Исходная строка

	// Итерация по строке как по рунам (Unicode-символам)
	for i, r := range str {
		fmt.Println(i, r, string(r)) // Вывод индекса, значения руны и её строкового представления
	}
	fmt.Println(str)

	// Итерация по строке как по байтам (UTF-8)
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i]) // Вывод индекса и значения байта
	}
	fmt.Println(str)

	// Преобразование в слайс байтов и итерация по ним
	for i, b := range []byte(str) {
		fmt.Println(i, b) // Вывод индекса и значения байта
	}
	fmt.Println(str)

	PrintReplaced("Суша!")          // Замена символа с формированием новой строки
	PrintReplacedDifferent("Лужа.") // Замена символа с прямым выводом
}

// Функция заменяет 'у' на 'а' и выводит результат
func PrintReplaced(x string) {
	var result []rune // Слайс для накопления результата
	for _, r := range x {
		if r == 'у' {
			result = append(result, 'а') // Замена символа
		} else {
			result = append(result, r) // Сохранение символа
		}
	}
	fmt.Println(string(result)) // Вывод итоговой строки
}

// Функция заменяет 'у' на 'а' и выводит результат посимвольно
func PrintReplacedDifferent(str string) {
	for _, r := range str {
		if r == 'у' {
			fmt.Print("а") // Вывод замены
		} else {
			fmt.Print(string(r)) // Вывод исходного символа
		}
	}
	fmt.Println() // Перенос строки
}
