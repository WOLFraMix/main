package main

import (
	"fmt"
	"reflect"
)

// Базовый интерфейс
type Animal interface {
	Speak() string
	GetName() string
}

// Конкретные типы
type Dog struct {
	Name string
}

func (d Dog) Speak() string   { return "Woof!" }
func (d Dog) GetName() string { return d.Name }

type Cat struct {
	Name string
}

func (c Cat) Speak() string   { return "Meow!" }
func (c Cat) GetName() string { return c.Name }

type Bird struct {
	Name string
}

func (b Bird) Speak() string   { return "Tweet!" }
func (b Bird) GetName() string { return b.Name }

// Прямой перебор через слайс интерфейсов
func traverseAnimals() {
	fmt.Println("=== Прямой перебор ===")
	animals := []Animal{
		Dog{Name: "Rex"},
		Cat{Name: "Whiskers"},
		Bird{Name: "Tweety"},
	}

	for _, animal := range animals {
		fmt.Printf("%s (%s): %s\n",
			animal.GetName(),
			reflect.TypeOf(animal).Name(),
			animal.Speak())
	}
}

func main() {
	// Демонстрация всех способов
	traverseAnimals()
}
