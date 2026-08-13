package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

func main() {
	slice, err := CreateSlice(6)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(slice)

	fmt.Println(FilterSlice([]int{10, 3, 8, 4, 7, 18, 2}))

	fmt.Println(MaxSumWithNegative([]int{-1, 6, 2, 5, 8, -2}, 3))

	fmt.Println(SortByParity([]int{-1, -2, 0, 6, 3, 2, -5, 5, 8, -2}))
}

func CreateSlice(n int) ([]int, error) {
	if n < 0 {
		return nil, fmt.Errorf("slice cant be negative")
	}
	if n == 0 {
		return []int{}, nil
	}
	result := make([]int, n)
	for i := range result {
		result[i] = randomInRange(-10, 10)
	}
	return result, nil
}

// случайное число в диапазоне [min, max] включительно
func randomInRange(min, max int) int {
	return min + rand.IntN(max-min+1)
}

func FilterSlice(numbers []int) []int {
	if len(numbers) < 2 {
		return nil
	}

	result := make([]int, 0, len(numbers))

	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			if numbers[i]%2 == 0 || numbers[i]%5 == 0 || numbers[i]%6 == 0 || numbers[i]%9 == 0 {
				result = append(result, numbers[i])
			}
		}
	}

	return result
}

func MaxSumWithNegative(numbers []int, k int) []int {
	if k <= 0 || k > len(numbers) {
		return nil
	}

	// Считаем сумму первого окна и количество отрицательных в нём
	windowSum := 0
	negCount := 0 // сколько отрицательных чисел в текущем окне
	for i := 0; i < k; i++ {
		windowSum += numbers[i]
		if numbers[i] < 0 {
			negCount++
		}
	}

	maxSum := windowSum
	maxStart := 0
	foundValid := negCount > 0 // первое окно валидно, если в нём есть отрицательное

	// Скользим окном
	for i := k; i < len(numbers); i++ {
		outgoing := numbers[i-k]
		incoming := numbers[i]

		// Обновляем сумму
		windowSum += incoming - outgoing

		// Обновляем счётчик отрицательных
		if outgoing < 0 {
			negCount--
		}
		if incoming < 0 {
			negCount++
		}

		// Если в текущем окне есть отрицательное число
		if negCount > 0 {
			if !foundValid || windowSum > maxSum {
				maxSum = windowSum
				maxStart = i - k + 1
				foundValid = true
			}
		}
	}

	// Если ни одно окно не содержало отрицательных чисел — по условию можно вернуть nil
	if !foundValid {
		return nil
	}

	return numbers[maxStart : maxStart+k]
}

func SortByParity(numbers []int) []int {
	if len(numbers) < 1 {
		return nil
	}

	even := make([]int, 0, len(numbers)) // чётные
	odd := make([]int, 0, len(numbers))  // нечётные
	for _, v := range numbers {
		if v%2 == 0 {
			even = append(even, v)
		}
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}

	slices.Sort(even)
	slices.Reverse(even)

	slices.Sort(odd)

	result := make([]int, 0, len(numbers))
	result = append(even, odd...)
	return result
}
