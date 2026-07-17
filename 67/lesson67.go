package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите вашу оценку от 0 до 100:")
	var grade int
	fmt.Scanln(&grade)
	result, err := gradeToLetter(grade)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func gradeToLetter(grade int) (letter string, err error) {
	if grade < 0 || grade > 100 {
		return "", fmt.Errorf("Некорректное значение: %d.\n", grade)
	}
	if grade >= 90 && grade <= 100 {
		return "A", nil
	}
	if grade >= 80 {
		return "B", nil
	}
	if grade >= 70 {
		return "C", nil
	}
	if grade >= 60 {
		return "D", nil
	}
	return "F", nil
}
