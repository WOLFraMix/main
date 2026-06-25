package main

import (
	"fmt"
)

const (
	a = iota // 0
	b        // 1
	c = 5    // 5
	d        // предыдущая строка (5)
	e = iota // отсчёт с нуля до текущей строки (4)
)

func main() {
	fmt.Println(a, b, c, d, e)
}
