package main

import (
	"fmt"
)

func main() {
	var time int
	fmt.Println("Введите текущее время (кол-во часов):")

	_, err := fmt.Scan(&time)
	if err != nil {
		fmt.Println("Ошибка ввода, неверные данные")
		return
	}

	switch {
	case time >= 0 && time <= 23:
		fmt.Println("Время:", time, "часов")
	default:
		fmt.Println("Это бобовая система счисления времени?")
	}

	switch {
	case time >= 8 && time < 12:
		fmt.Println("Сейчас утро!")
	case time >= 12 && time < 18:
		fmt.Println("Сейчас день.")
	case time >= 18 && time <= 23:
		fmt.Println("Сейчас вечер.")
	case time >= 0 && time < 8:
		fmt.Println("Сейчас ночь... Сладких снов.")
	default:
		fmt.Println("Неверно указанное число... Попробуйте снова.")
	}
}
