package main

import "fmt"

type account struct {
	balance float32
}

func main() {
	accounts := []account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
	for i := range accounts {
		accounts[i].balance += 1000
	}
	fmt.Println(accounts)

	johnPrice := computePrice(145.90, 3)
	fmt.Printf("TOTAL : %0.2f $\n", johnPrice)

	fmt.Println(factorial(3))

	fmt.Println(max(1, 2, 3, 4, 5))

	fmt.Println(unique([]string{"apple", "orange", "apple", "pear", "orange", "avocado"}))
}

func computePrice(rate float32, nights int) (price float32) {
	p := rate * float32(nights)
	p *= 2
	return p
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func max(n ...int) int {
	max := 0
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return max
}

func unique(s []string) []string {
	result := make([]string, 0, len(s))
	order := make(map[string]struct{})
	for _, str := range s {
		if _, ok := order[str]; !ok {
			result = append(result, str)
		}
		order[str] = struct{}{}
	}
	return result
}
