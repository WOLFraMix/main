package main

import "fmt"

func main() {
	// продукты - map прейскурант в рублях
	products := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	// деликатесы и их цены
	delicacies := make(map[string]int)
	for k, v := range products {
		if v > 500 {
			delicacies[k] = v
		}
	}
	fmt.Println(delicacies)

	// вычисляем стоимость заказа
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	sum := 0

	for k, v := range products {
		for i := range order {
			if k == order[i] {
				sum += v
			}
		}
	}
	fmt.Println(sum)
}
