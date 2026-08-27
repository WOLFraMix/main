package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Printf("%s издаёт звук.\n", a.Name)
}

type Dog struct {
	Animal // встраивание
	Breed  string
}

func (d Dog) Bark() {
	fmt.Println("ГАВ!")
}

func main() {
	dog := Dog{
		Animal: Animal{
			Name: "Шарик",
		},
		Breed: "Далматин",
	}
	fmt.Printf("%+v\n", dog)
	dog.Speak()
	dog.Bark()
}
