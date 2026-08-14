package main

import "fmt"

type figures int

const (
	square figures = iota
	circle
	triangle
)

func area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		return func(x float64) float64 {
			return x * x
		}, true

	case circle:
		return func(x float64) float64 {
			return 3.142 * x * x // x здесь — радиус
		}, true

	case triangle:
		// Здесь подразумевается равносторонний треугольник со стороной x;
		return func(x float64) float64 {
			return 0.433 * x * x
		}, true

	default:
		return nil, false
	}
}

func main() {
	squareArea, ok := area(square)
	if ok {
		fmt.Println("Площадь квадрата со стороной 5:", squareArea(5))
	}

	circleArea, ok := area(circle)
	if ok {
		fmt.Println("Площадь круга с радиусом 3:", circleArea(3))
	}

	triangleArea, ok := area(triangle)
	if ok {
		fmt.Println("Площадь равностороннего треугольника со стороной 4:", triangleArea(4))
	}

	_, ok = area(10)                       // неизвестная фигура
	fmt.Println("Неизвестная фигура:", ok) // false
}
