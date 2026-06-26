package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 555
	str := strconv.Itoa(i) // преобразование числа в строку
	fmt.Println(str)
	str2 := strconv.FormatInt(int64(i), 2)  // преобразование числа в строку в двоичной системе
	str3 := strconv.FormatInt(int64(i), 8)  // преобразование числа в строку в восьмеричной системе
	str4 := strconv.FormatInt(int64(i), 16) // преобразование числа в строку в шестнадцатеричной системе
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)
	str5 := fmt.Sprintf("Сюда вставляем int: %d", i) // вставка значения int в готовую строку
	fmt.Println(str5)
	str6 := fmt.Sprintf("Сюда вставляем любое значение: %v", i) // вставка значения в готовую строку
	fmt.Println(str6)
	// %f для float64, %d для int, %v для любого типа

	str7 := "12345"
	num, err := strconv.Atoi(str7) // преобразование строки в число
	if err != nil {
		fmt.Println("Ошибка:", err) // вывод ошибки
	} else {
		fmt.Println("Число:", num) // вывод числа
	}
}
