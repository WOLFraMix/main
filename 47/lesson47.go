package main

import (
	"fmt"
	"math"
)

func main() {
	// Приветствие по умолчанию
	hello("Катя", "ведущий")

	var name string
	var role string

	// Запрос и чтение имени пользователя
	fmt.Println("Введите ваш ник: ")
	fmt.Scanln(&name)
	// Запрос и чтение роли пользователя
	fmt.Println("Введите вашу роль: ")
	fmt.Scanln(&role)
	hello(name, role)

	// Тестирование функции сравнения котиков и собачек
	PetBattle(3, 2)
	PetBattle(1, 4)
	PetBattle(5, 5)

	// Тестирование функции анализа чисел
	printNumberInfo(-4)
	printNumberInfo(5)
	printNumberInfo(16)
	printNumberInfo(0)
}

// Функция приветствия пользователя
func hello(userName string, userRole string) {
	fmt.Printf("Привет, %s! Вы %s.\n", userName, userRole)
}

// Функция определения победителя между котиками и собачками
func PetBattle(cats int, dogs int) {
	if cats > dogs {
		fmt.Printf("Котики победили со счетом %d:%d!\n", cats, dogs)
	} else if cats < dogs {
		fmt.Printf("Собачки победили со счетом %d:%d!\n", dogs, cats)
	} else {
		fmt.Println("Ничья! Все дружат!")
	}
}

// Функция вывода информации о числе (знак, четность, квадратный корень)
func printNumberInfo(num int) {
	// Проверка на ноль, отрицательное или положительное число
	switch {
	case num == 0:
		fmt.Println("Число равно 0.")
	case num < 0:
		fmt.Printf("Число %d отрицательное.\n", num)
	case num > 0:
		fmt.Printf("Число %d положительное.\n", num)
	}
	// Проверка на четность или нечетность
	switch {
	case num%2 == 0:
		fmt.Printf("Число %d четное.\n", num)
	case num%2 != 0:
		fmt.Printf("Число %d нечетное.\n", num)
	}
	// Вычисление и анализ квадратного корня
	switch {
	case num > 0:
		sqrtNum := math.Sqrt(float64(num))
		if sqrtNum == math.Floor(sqrtNum) {
			fmt.Printf("Квадратный корень числа %d является целым числом и равен %.0f.\n", num, sqrtNum)
		} else if sqrtNum != math.Floor(sqrtNum) {
			fmt.Printf("Квадратный корень числа %d не является целым числом и равен %.5f.\n", num, sqrtNum)
		}
	}
}
