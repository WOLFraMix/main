package main

import (
	"fmt"
	"runtime/debug"
)

// LIFO - Last In First Out
// callstack - стек вызовов
func main() {
	fn1()
}

func fn1() {
	fn2()
	fmt.Println("fn1")
}

func fn2() {
	fmt.Println("fn2")
	debug.PrintStack()
}
