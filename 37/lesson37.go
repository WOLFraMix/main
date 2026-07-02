package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 5
	c := 10
	// операторы сравнения автоматически дают bool тип данных
	fmt.Println(a == b) // оператор равенства
	fmt.Println(a == c)
	fmt.Println(a != b) // оператор неравенства
	fmt.Println(a != c)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)

	fmt.Println(a > 3 && c > 3) // оператор логического И
	fmt.Println(a != 3 && c < 3)
	fmt.Println(a > 3 || c > 3) // оператор логического ИЛИ
	fmt.Println(b <= 3 || c >= 12)
	fmt.Println(!(b <= 3 || c >= 12)) // оператор логического НЕ !(меняет значение на противоположное)
}
