package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "42"

	// Функция Atoi преобразует строку в целое число
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return // выходим из функции
	}

	fmt.Printf("значение %d типа %T\n", num, num) // -> значение 42 типа int

	// Функция Itoa преобразует целое число в строку
	str = strconv.Itoa(num)

	fmt.Printf("значение %s типа %T\n", str, str) // -> значение 42 типа string

	num = 42

	// Функция FormatInt преобразует целые числа в строки:
	// Мы передаем первым аргументом число,
	// а вторым - основание системы счисления
	// (10 в данном случае).
	str = strconv.FormatInt(int64(num), 10)

	fmt.Printf("значение %s типа %T\n", str, str) // -> значение 42 типа string

	var n int64
	fmt.Scan(&n)
	s := strconv.FormatInt(n, 2)
	fmt.Printf("Число %d в двоичной системе счисления: %s\n", n, s)
}
