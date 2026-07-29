package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{4, 5, 6}
	arr4 := [3]int{3, 2, 1}

	fmt.Println(arr1 == arr2)
	fmt.Println(arr2 == arr3)
	fmt.Println(arr1 == arr4)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(SumNeighbours(arr))
}

func SumNeighbours(array [10]int) (result [10]int) {
	// находим сумму соседей в массиве
	for i := range array { // проходимся по индексу
		switch i { // сравниваем его
		case 0: // у первого числа только правый сосед
			result[i] = array[1]
		case 9: // у последнего числа только левый сосед
			result[i] = array[8]
		default: // все остальные случаи левый+правый соседи
			result[i] = array[i-1] + array[i+1]
		}
	}
	return result
}
