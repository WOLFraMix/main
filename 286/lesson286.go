package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{} // пустой интерфейс, изначально nil

	fmt.Println(i == nil) // true, так как i не содержит ни типа, ни значения

	i = 42         // теперь i хранит значение типа int
	fmt.Println(i) // вывод: 42

	i = "hello"    // теперь i хранит строку
	fmt.Println(i) // вывод: hello

	fmt.Println(CheckIfNil(i))
}

// CheckIfNil - функция для проверки значения на nil
// Функция проверяет, является ли переданное значение nil
// Работает с различными типами данных, включая указатели, срезы, мапы, каналы, функции и интерфейсы
// Возвращает true, если значение является nil, иначе возвращает false
func CheckIfNil(i interface{}) bool {
	// Первый уровень проверки - простая проверка на nil
	if i == nil {
		return true
	}

	// Получаем reflect.Value для более детальной проверки
	val := reflect.ValueOf(i)
	// Определяем тип значения с помощью метода Kind()
	// Kind() возвращает базовый тип значения
	switch val.Kind() {
	// Проверяем типы, которые могут быть nil
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface:
		// Метод IsNil() проверяет, является ли значение nil
		return val.IsNil()
	default:
		// Для всех остальных типов возвращаем false
		// Так как они не могут быть nil
		return false
	}
}
