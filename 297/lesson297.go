package main

import (
	"fmt"
	"time"
)

func main() {
	findFriday13(2026, 2027)
}

// findFriday13 ищет Пятницу 13ое от года к году
func findFriday13(startYear, endYear int) {
	for year := startYear; year <= endYear; year++ {
		for month := time.January; month <= time.December; month++ {
			thirteen := time.Date(year, month, 13, 0, 0, 0, 0, time.UTC)
			if thirteen.Weekday() == time.Friday {
				fmt.Printf("%04d-%02d-%02d\n", year, month, thirteen.Day())
			}
		}
	}
}
