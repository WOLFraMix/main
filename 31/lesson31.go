package main

import (
	"fmt"
)

func main() {
	var p *int     // указатель на int
	fmt.Println(p) // выводим nil
	// fmt.Println(*p) - в памяти нету значения а значит и указатель нам не вывести
	if p != nil {
		fmt.Println("Значение по указателю:", p)
	} else {
		fmt.Println("Указатель равен nil")
	}

	x := 10
	p = &x          // даём указателю адрес переменной x
	fmt.Println(p)  // т.к. p - указатель, то выводим адрес
	fmt.Println(*p) // выводим значение по указателю
	if p != nil {
		fmt.Println("Значение по указателю:", p)
	} else {
		fmt.Println("Указатель равен nil")
	}

	k := new(int) // new сразу создаёт zerovalue значение даже если мы ничего не передали переменной k
	fmt.Println(*k)
	fmt.Println(k)
}
