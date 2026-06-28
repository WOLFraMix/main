package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Имя - %s\n", "Stepan")

	str := fmt.Sprintf("Имя - %s\n", "Vlad")
	fmt.Print(str)

	fmt.Printf("%d\n", 70) // int

	fmt.Printf("%f\n", 60.345) // float

	fmt.Printf("%t\n", true) // bool

	fmt.Printf("Loading... %d%%\n", 99) // add percent

	fmt.Printf("%b\n", 50)  // binary
	fmt.Printf("%o\n", 50)  // octal
	fmt.Printf("%#x\n", 50) // hex

	fmt.Printf("%.2f\n", 65.3458) // float with 2 decimal places

	fmt.Printf("%v, %v, %v\n", 37, "Den", false) // displays any information
}
