package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?") // когда суббота?
	today := time.Now().Weekday()

	// Вычисляем разницу с учётом цикличности недели (7 дней)
	daysUntilSaturday := (time.Saturday - today + 7) % 7
	// % - это оператор остатка от деления, который гарантирует
	// что результат будет в диапазоне от 0 до 6

	switch daysUntilSaturday {
	case 0:
		fmt.Println("Today.")
	case 1:
		fmt.Println("Tomorrow.")
	case 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
