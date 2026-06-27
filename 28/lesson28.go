package main

import (
	"fmt"
)

func main() {
	fmt.Print("Добрый, ") // печать без переноса
	fmt.Print("день!\n")  // добавляем \n иначе дальше будет печатать на этой же строке

	name := "Stepan"
	fmt.Println("Моё имя на следующей строке:\n", name) // Println = Print + \n

	fmt.Println("Hi,", "my name is", name) // можно вставлять много значений в одну строку

	fmt.Printf("Привет, %s, %s", "мой", "друг\n") // %s - строка

	age := 27
	fmt.Printf("My name is %s, and I am %d years old.\n", name, age) // %d - число

	str := fmt.Sprint("My name is - ", name, " and I am ", age, " years old.\n")
	fmt.Print(str)

	str2 := fmt.Sprintf("My name is %s, and I am %d years old.", name, age)
	fmt.Println(str2)
}
