package main

import (
	"fmt"
	"strings"
)

// Product represents an item with a name and price
type Product struct {
	Name  string
	Price int
}

func main() {
	// Initialize a slice of sample products
	products := []Product{
		{Name: "Клавиатура JZ9", Price: 19200},
		{Name: "Наушники N45", Price: 9600},
		{Name: "Смартфон S10", Price: 55000},
	}

	var find string
	fmt.Println("Введите название товара:")
	_, err := fmt.Scan(&find)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Convert search term to lowercase for case-insensitive matching
	findLower := strings.ToLower(find)
	found := false

	// Iterate through products to find a matching name
	for _, product := range products {
		if strings.Contains(strings.ToLower(product.Name), findLower) {
			fmt.Printf("%s: %d руб.\n", product.Name, product.Price)
			found = true
		}
	}

	// Output result if no match was found
	if !found {
		fmt.Printf("Товар %s не найден.\n", find)
	}
}
