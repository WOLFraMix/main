package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	// добавляем текст в буфер
	builder.WriteString("hello")
	builder.WriteString(" ")
	builder.WriteString("world")

	// получаем итоговую строку из буфера
	result := builder.String()

	// выводим информацию о буфере
	fmt.Println("Length:", builder.Len())   // -> "Length: 11"
	fmt.Println("Capacity:", builder.Cap()) // -> "Capacity: 64"
	fmt.Println(result)                     // -> hello world

	// очищаем буфер
	builder.Reset()

	// добавляем новый текст в буфер
	builder.WriteString("foo")
	builder.WriteString(" ")
	builder.WriteString("bar")

	// получаем итоговую строку из буфера
	result = builder.String()

	// выводим информацию о буфере
	fmt.Println("Length:", builder.Len())   // -> "Length: 7"
	fmt.Println("Capacity:", builder.Cap()) // -> "Capacity: 64"

	fmt.Println(result) // -> "foo bar"
}
