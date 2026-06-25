package main

import (
	"fmt"
	"math"
)

func main() {
	i := 42                    // тип int
	var f float64 = float64(i) // создаём float64 и просим представить число int в виде float64
	fmt.Println(f)

	d := 58.95         // тип float64
	var q int = int(d) // создаем int и просим представить число float64 в виде int
	fmt.Println(q)
	var g int = int(math.Round(d)) // сначала делаем округление через math.Round, а потом преобразуем в int
	fmt.Println(g)
}
