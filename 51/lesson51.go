package main

import "fmt"

// divide - делить
// defer - отложенное выполнение
func main() {
	divide(10, 5)
	divide(10, 2)
	divide(10, 1)
	fn()
	fn1()
	fn2()
	fmt.Println(fn3())
}

func divide(n1, n2 int) {
	defer fmt.Println("Конец функции divide") // defer выполнится даже в случае ошибки в функции если был вызван заранее
	fmt.Println(n1 / n2)
	defer fmt.Println("Конец функции 2") // но не после ошибки (например, деление на 0)
}

func fn() {
	i := 0
	defer fmt.Println(i) // тут defer получил предыдущее значение и результат будет 0
	defer func() {       // тут функция берёт значение из предыдущей "ступеньки", т.е. из функции fn и результат будет 5
		fmt.Println(i) // но defer выводятся в обратном порядке, поэтому результат будет сначала 5 потом 0
	}()
	i = 5
}

func fn1() {
	i := 10
	defer func(i int) {
		fmt.Println(i) // 10
	}(i)
	i = 20
}

func fn2() {
	i := 10
	defer func() {
		fmt.Println(i) // 20
	}()
	i = 20
}

func fn3() (i int) {
	i = 5
	defer func() { // 6
		i = 6
	}()
	return 7
}
