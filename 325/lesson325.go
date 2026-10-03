package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	layout := "2006-01-02"
	var v1 string
	fmt.Scan(&v1)
	var v2 string
	fmt.Scan(&v2)
	t1, err := time.Parse(layout, v1)
	if err != nil {
		fmt.Println("Введенные значения некорректны!")
		return
	}
	t2, err := time.Parse(layout, v2)
	if err != nil {
		fmt.Println("Введенные значения некорректны!")
		return
	}
	diff := t2.Sub(t1)
	days := int(math.Abs(diff.Hours()) / 24)
	fmt.Println(days)
}
