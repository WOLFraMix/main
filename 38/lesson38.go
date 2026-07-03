package main

import "fmt"

const (
	execute = 0b001 // 1 - право выполнения
	write   = 0b010 // 2 - право записи
	read    = 0b100 // 4 - право чтения
)

// ПОБИТОВЫЕ ОПЕРАЦИИ И ОПЕРАТОРЫ
func main() {
	permissions := 5                                                    // 0b101 - право выполнения и чтения
	fmt.Printf("Наши права: %b\n", permissions)                         // проверяем можем ли мы:
	fmt.Printf("Можем выполнять: %t\n", permissions&execute == execute) // выполнять
	fmt.Printf("Можем писать: %t\n", permissions&write == write)        // писать
	fmt.Printf("Можем читать: %t\n", permissions&read == read)          // читать
	// %b выводим бинарный формат для понимания прав. %t - true or false
	adminPermissions := 7 // 0b111 - право выполнения, записи и чтения
	fmt.Printf("Права администратора: %b\n", adminPermissions)
	fmt.Printf("Можем выполнять: %t\n", adminPermissions&execute == execute) // выполнять
	fmt.Printf("Можем писать: %t\n", adminPermissions&write == write)        // писать
	fmt.Printf("Можем читать: %t\n", adminPermissions&read == read)          // читать
}
