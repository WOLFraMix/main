package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "15:04"
	var v string
	fmt.Scan(&v)
	var h int
	fmt.Scan(&h)
	t, err := time.Parse(layout, v)
	if err != nil {
		fmt.Println("Введенное значение некорректно!")
		return
	}
	result := t.Add(time.Duration(h) * time.Hour)
	fmt.Println(result.Format(layout))
}
