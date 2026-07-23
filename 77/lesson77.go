package main

import "fmt"

func main() {
	printDiamond(3)
}

func printDiamond(n int) {
	fmt.Println("Мой бриллиант:")

	// Верхняя половина (включая среднюю строку)
	for i := 1; i <= n; i++ {
		spaces := n - i
		hashes := 2*i - 1

		// 1. Печатаем отступы слева
		for k := 0; k < spaces; k++ {
			fmt.Print(" ")
		}

		// 2. Печатаем сам ромб
		if hashes == 1 {
			// Верхняя вершина или нижняя вершина (если бы мы были в нижнем цикле, но тут только верх)
			fmt.Print("#")
		} else {
			// Если это средняя строка (i == n), часто её делают сплошной для красоты,
			// но по строгому требованию "без решёток внутри" она тоже должна быть полой,
			// если только n не равно 1.
			// Однако, если n=2, средняя строка имеет длину 3. Полая: "# #".

			fmt.Print("#") // Левый край

			// Внутренние пробелы: общая длина - 2 (два края)
			innerSpaces := hashes - 2
			for k := 0; k < innerSpaces; k++ {
				fmt.Print(" ")
			}

			fmt.Print("#") // Правый край
		}
		fmt.Println()
	}

	// Нижняя половина
	for i := n - 1; i >= 1; i-- {
		spaces := n - i
		hashes := 2*i - 1

		// 1. Печатаем отступы слева
		for k := 0; k < spaces; k++ {
			fmt.Print(" ")
		}

		// 2. Печатаем сам ромб
		if hashes == 1 {
			fmt.Print("#")
		} else {
			fmt.Print("#") // Левый край

			innerSpaces := hashes - 2
			for k := 0; k < innerSpaces; k++ {
				fmt.Print(" ")
			}

			fmt.Print("#") // Правый край
		}
		fmt.Println()
	}
}
