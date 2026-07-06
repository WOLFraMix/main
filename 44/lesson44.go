package main

import "fmt"

func main() {
	var imt float64
	var weight float64
	var height float64

	fmt.Println("Введите ваш вес (кг):")
	_, err := fmt.Scan(&weight)
	if err != nil || weight <= 0 || weight > 300 {
		fmt.Println("Ошибка ввода веса")
		return
	}

	fmt.Println("Введите ваш рост (м):")
	_, err = fmt.Scan(&height)
	if err != nil || height <= 0 || height > 3 {
		fmt.Println("Ошибка ввода роста")
		return
	}

	imt = weight / (height * height)
	fmt.Printf("Ваш ИМТ: %.2f\n", imt)

	switch {
	case imt < 18.5:
		fmt.Println("Недостаточный вес")
	case imt >= 18.5 && imt < 25:
		fmt.Println("Нормальный вес")
	case imt >= 25 && imt < 30:
		fmt.Println("Избыточный вес")
	case imt >= 30:
		fmt.Println("Ожирение")
	}
}
