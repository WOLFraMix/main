package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	rollDice(7) // пример: ищем сумму 7
}

func rollDice(target int) {
	// Проверка на невозможную сумму двух кубиков (от 2 до 12)
	if target < 2 || target > 12 {
		fmt.Printf("Сумма %d невозможна для двух шестигранных кубиков.\n", target)
		return
	}

	var z, y, attempts int

	for {
		z = rand.IntN(6) + 1
		y = rand.IntN(6) + 1
		attempts++

		if z+y == target {
			wordVerb := pluralVerb(attempts)
			wordCount := pluralCount(attempts)
			fmt.Printf("Выпало %d и %d, в сумме %d. На это %s %d %s.\n", z, y, target, wordVerb, attempts, wordCount)
			break // если выпал нужный результат, то завершаем функцию. иначе пропускаем if
		}

		fmt.Printf("Выпало %d и %d, в сумме %d, бросаем ещё раз.\n", z, y, z+y)
	}
}

// pluralVerb возвращает правильную форму глагола «потребовался/потребовалось»
func pluralVerb(n int) string {
	lastTwo := n % 100
	lastOne := n % 10

	if lastTwo >= 11 && lastTwo <= 19 {
		return "потребовалось"
	}
	if lastOne == 1 {
		return "потребовался"
	}
	return "потребовалось"
}

// pluralCount возвращает правильную форму слова «бросок/броска/бросков»
func pluralCount(n int) string {
	lastTwo := n % 100
	lastOne := n % 10

	if lastTwo >= 11 && lastTwo <= 19 {
		return "бросков"
	}
	switch lastOne {
	case 1:
		return "бросок"
	case 2, 3, 4:
		return "броска"
	default:
		return "бросков"
	}
}
