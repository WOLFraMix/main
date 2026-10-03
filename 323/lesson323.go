package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006.01.02 15-04-05"
	value := "2022.02.02  13-43-05"
	t, err := time.Parse(layout, value)
	if err != nil {
		fmt.Println("Ошибка при парсинге строки:", err)
	}
	fmt.Println(t) // -> 2022-02-02 13:43:05 +0000 UTC

	strTime2 := "2023-03-20"
	layout2 := "2006-01-02"
	t2, err := time.Parse(layout2, strTime2)
	if err != nil {
		fmt.Println("Введенное значение некорректно!")
	}
	t2 = t2.AddDate(-23, -2, -19)
	fmt.Println(t2.Format(layout2))
}
