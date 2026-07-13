package main

import "fmt"

type Direction int

const (
	North Direction = iota + 1
	West
	South
	East
)

func main() {
	action(South)
}

func action(d Direction) {
	fmt.Println("Действие в направлении:", d)
}
