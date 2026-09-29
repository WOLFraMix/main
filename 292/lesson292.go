package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Создаем текущее, локальное время, которое на компьютере
	now := time.Now()
	// Добавляем два часа к времени в переменной now.
	later := now.Add(2 * time.Hour) // now не изменится, а later - новый объект time.Time

	// Создаем определенную дату и время
	t := time.Date(2024, time.April, 5, 14, 30, 0, 0, time.UTC)
	fmt.Println(t.Format(time.RFC3339)) // Вывод в определенном формате

	// Загружаем локаль Нью-Йорка
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	// Получаем время Нью-Йорка
	nyTime := t.In(loc)
	fmt.Println(nyTime)

	// Сравнение выполняем с методами .Equal, .Before, .After
	if later.After(now) {
		fmt.Println("later позже now")
	}

	// Проверка на нулевое время
	var zero time.Time
	if zero.IsZero() {
		fmt.Println("Время не установлено")
	}
}
