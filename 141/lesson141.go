package main

import "fmt"

// объявление типа
type MyType int

// объявление метода
func (m MyType) String() string {
	return fmt.Sprintf("MyType: %d", m)
}

func main() {
	var m MyType = 5

	// вызов метода
	s := m.String()
	fmt.Println(s)
}

// Синтаксис метода типа похож на синтаксис обычной функции,
// но добавляется получатель (receiver) после ключевого слова func.
// Можно сказать, что получатель — это ещё один аргумент функции.
