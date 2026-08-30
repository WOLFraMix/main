package main

import "fmt"

// структура Dog
type Dog struct {
}

// метод Speak для Dog
func (Dog) Speak() string {
	return "Гав!"
}

// структура Cat
type Cat struct {
}

// метод Speak для Cat
func (Cat) Speak() string {
	return "Мяу!"
}

// интерфейс с методом Speak()
type Speaker interface {
	Speak() string
}

// функция реализует интерфейс Speaker
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	makeSound(dog) // "Гав!"
	makeSound(cat) // "Мяу!"
}
