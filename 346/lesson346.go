package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) Greet() string {
	return fmt.Sprintf("Привет, меня зовут %s, мне %d лет, я из %s", p.Name, p.Age, p.City)
}

func main() {
	p1 := Person{Name: "Ivan", Age: 25, City: "Moscow"}
	p2 := Person{Name: "Anna", Age: 30, City: "London"}
	p3 := Person{Name: "Stepan", Age: 27, City: "Astrakhan"}
	people := []Person{p1, p2, p3}

	n := 1
	fmt.Print("Введите количество человек (1-3): ")
	fmt.Scan(&n)

	if n < 1 {
		n = 1
	}
	if n > len(people) {
		n = len(people)
	}
	count := 1

	for _, p := range people {
		if count > n {
			break
		}
		count++

		fmt.Println("Имя:", p.Name)
		fmt.Println("Возраст:", p.Age)
		fmt.Println("Город:", p.City)
		fmt.Println(p.Greet())
		fmt.Println()
	}
}
