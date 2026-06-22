package main

import "fmt"

var (
	goodsCount = 13
	numberPair = 14
	classRoom  int
)

const (
	li = 3.0
	si = 8.1
	di
)

func main() {
	a := 5
	b := 10
	fmt.Println("lets GO", a)
	a = 20
	fmt.Println("say it", b, a)
	var age int = 30
	fmt.Println(age)
	ace := 50
	fmt.Println(ace)
	var x, y int = 1, 2
	fmt.Println(x, y)
	z, d := 3, 4
	fmt.Println(z, d)
	fmt.Println(goodsCount, numberPair, classRoom)
	age, ace = 31, 88
	fmt.Println(age, ace)
	const pi = 3.14
	fmt.Println(pi)
	fmt.Println(li, si, di)
}
