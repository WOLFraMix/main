package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	now := time.Now()
	later2h := now.Add(2 * time.Hour)
	t := time.Date(1999, time.January, 30, 12, 30, 0, 0, time.UTC)

	fmt.Println(now)
	fmt.Println(later2h)
	fmt.Println(t)

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	nyTime := now.In(loc)
	fmt.Printf("%s\t%s\n", nyTime.Location(), nyTime)

	t1 := time.Now()
	t2 := t1.UTC()
	t3 := time.Now().UTC().Format(time.RFC3339)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)

	duration := time.Since(start)
	fmt.Println(duration)
}
