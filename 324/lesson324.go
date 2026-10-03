package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02"
	var value string
	fmt.Scan(&value)
	t, err := time.Parse(layout, value)
	if err != nil {
		fmt.Println("Введенное значение некорректно!")
		return
	}
	fmt.Println(t.Weekday().String())
}
