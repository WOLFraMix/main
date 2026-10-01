package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Форматирование даты и времени
	// Mon Jan 2 15:04:05 MST 2006
	// "2006-01-02"
	// "02/01/2006 15:04"

	// RFC3339

	now := time.Now()
	t := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond(),
		now.Location(),
	)
	fmt.Println(t)
	s := t.Format("02/01/2006 15:04")
	fmt.Println(s)

	tt, err := time.Parse("02/01/2006 15:04", s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tt)
}
