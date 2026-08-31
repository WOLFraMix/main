package main

import "fmt"

type Direction int

const (
	North Direction = iota + 1
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "not a direction"
	}
}

func main() {
	fmt.Println(South.String())
	fmt.Println(South)
}
