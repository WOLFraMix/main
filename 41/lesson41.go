package main

import "fmt"

func main() {
	day := 5

	switch day {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница, после которой наконец-то будет...")
		fallthrough
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресенье")
	default:
		fmt.Println("Неизвестный день")
	}

	saturday := 6
	sunday := 7

	switch {
	case day >= 1 && day <= 5:
		fmt.Println("Будний день")
	case day == saturday || day == sunday:
		fmt.Println("Выходной день")
	default:
		fmt.Println("Некорректное значение")
	}

	var x any = "тут какой-то тип данных"
	// находим тип переменной через switch
	switch x.(type) {
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	case float64:
		fmt.Println("x is float64")
	case bool:
		fmt.Println("x is bool")
	default:
		fmt.Println("x is unknown type")
	}
}
