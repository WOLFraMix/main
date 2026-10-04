package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	startStr := "2024-10-04 09:15:00"
	endStr := "2024-10-04 11:45:30"

	t1, err := time.Parse("2006-01-02 15:04:05", startStr)
	if err != nil {
		log.Fatal(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05", endStr)
	if err != nil {
		log.Fatal(err)
	}

	// Длительность или разница:
	diff := t2.Sub(t1)
	fmt.Printf("Длительность: %v (%.2f часов)\n", diff, diff.Hours())

	// Округление начала до 15 минут
	step := 15 * time.Minute
	t1Rounded := t1.Round(step)
	fmt.Println("Начало (округлённое до 15 мин):", t1Rounded.Format("15:04"))
}
