package main

// fallthrough - это ключевое слово в языке Go
// которое принудительно передаёт выполнение к следующему блоку case
// в конструкции switch
// даже если значение этого блока не совпадает с проверяемым
// но передача выполнения происходит только на один следующий блок case
// и fallthrough не выполняется если не выполнено условие текущего блока case

import "fmt"

func main() {
	x := 2
	switch x {
	case 0:
		fmt.Println("Ноль")
		fallthrough
	case 1:
		fmt.Println("Один")
	case 2:
		fmt.Println("Два")
		fallthrough
	case 3:
		fmt.Println("Три")
	default:
		fmt.Println("Другое")
	}
}
