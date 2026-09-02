package main

import "fmt"

// Структура (struct)
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод (method) — привязан к типу Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Метод с указателем может менять поля структуры
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Интерфейс (interface)
type Shape interface {
	Area() float64
}

// Rectangle реализует интерфейс Shape,
// потому что у него есть метод Area()

// Функция-конструктор (constructor function)
func NewRectangle(width, height float64) *Rectangle {
	if width <= 0 || height <= 0 {
		// можно вернуть nil или ошибку
		return nil
	}
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

func main() {
	// использование конструктора
	rect := NewRectangle(5, 3)
	if rect == nil {
		fmt.Println("Некорректные размеры")
		return
	}

	// вызов метода
	fmt.Printf("Площадь: %.2f\n", rect.Area())

	// изменение через метод с указателем
	rect.Scale(2)
	fmt.Printf("После масштабирования: %.2f\n", rect.Area())

	// полиморфизм через интерфейс
	var s Shape = rect
	fmt.Printf("Через интерфейс: %.2f\n", s.Area())
}
