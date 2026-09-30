package main

import (
	"fmt"
	"time"
)

func main() {
	d := 5 * time.Minute
	fmt.Println(d)

	start := time.Date(2026, time.March, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, time.March, 2, 0, 0, 0, 0, time.UTC)
	diff := end.Sub(start)
	fmt.Println(diff)

	now := time.Now()
	elapsed := time.Since(now)
	fmt.Printf("Операция заняла %v\n", elapsed)

	deadline := time.Now().Add(10 * time.Hour)
	remaining := time.Until(deadline)
	fmt.Println(remaining)

	ti := time.Hour
	fmt.Println(ti / 4)
	fmt.Println(ti * 4)

	timeInfo(now)
	fmt.Println(FormatDuration(d))
}

func timeInfo(t time.Time) {
	leapYear := false
	if t.Year()%400 == 0 {
		leapYear = true
	} else if t.Year()%100 == 0 {
		leapYear = false
	} else if t.Year()%4 == 0 {
		leapYear = true
	}
	fmt.Println("Високосный:", leapYear)

	fmt.Print("День недели: ")
	switch t.Weekday() {
	case time.Monday:
		fmt.Print("Понедельник")
	case time.Tuesday:
		fmt.Print("Вторник")
	case time.Wednesday:
		fmt.Print("Среда")
	case time.Thursday:
		fmt.Print("Четверг")
	case time.Friday:
		fmt.Print("Пятница")
	case time.Saturday:
		fmt.Print("Суббота")
	case time.Sunday:
		fmt.Print("Воскресенье")
	}

	fmt.Println()
	fmt.Println("Unix:", t.Unix())
}

func FormatDuration(d time.Duration) string {
	// return d.String()

	seconds := int(d / time.Second)

	if d < 60*time.Second {
		return fmt.Sprintf("%02d", seconds)
	}
	if d < 3600*time.Second {
		minutes := seconds / 60
		seconds %= 60
		return fmt.Sprintf("%02d:%02d", minutes, seconds)
	}
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds %= 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
