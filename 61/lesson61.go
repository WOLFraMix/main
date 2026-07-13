package main

import "fmt"

var config string

// Функция init используется для инициализации глобальных переменных и
// выполнения необходимой логики перед запуском функции main.
func init() {
	config = "Настройки загружены"
	fmt.Println("Инициализация пакета")
}

func main() {
	fmt.Println(config)
}
