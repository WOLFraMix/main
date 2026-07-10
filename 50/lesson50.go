package main

import "fmt"

func main() {
	val := 5

	fn(val)
	fmt.Println(val) // 5

	fnPointer(&val)
	fmt.Println(val) // 25
}

func fn(a int) {
	a = a + 10     // значение меняется только внутри функции
	fmt.Println(a) // 15
}

func fnPointer(a *int) {
	*a = *a + 20    // значение меняется по указателю, т.е. в основной переменной
	fmt.Println(*a) // 25
}
