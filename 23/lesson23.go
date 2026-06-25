package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	randomNum := rand.IntN(100) // генерируем случайное число от [0 до 100) т.е. вкл 0 но не вкл 100
	fmt.Println(randomNum)

	min := 10
	max := 30
	randomNum2 := rand.IntN(max-min) + min // генерируем случайное число от [10 до 30)
	fmt.Println(randomNum2)

	min = 20
	max = 40
	randomNum3 := rand.IntN(max-min+1) + min // генерируем случайное число от [20 до 40]
	fmt.Println(randomNum3)
}
