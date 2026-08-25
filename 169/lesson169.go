package main

import (
	"fmt"
	"strings"
)

// User представляет пользователя
type User struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address Address
	Cart    []CartItem
}

// Address представляет адрес пользователя
type Address struct {
	Street     string
	City       string
	PostalCode string
}

// CartItem представляет элемент в корзине
type CartItem struct {
	Product  Product
	Quantity int
}

// Product представляет продукт в корзине
type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Category    string
	Brand       string
	Rating      float64
	Reviews     int
}

func main() {
	user := User{
		ID:    1,
		Name:  "Иван Петров",
		Email: "ivan.petrov@example.com",
		Phone: "+7 999 123-45-67",
		Address: Address{
			Street:     "Улица Ленина",
			City:       "Москва",
			PostalCode: "101000",
		},
		Cart: []CartItem{
			{
				Product: Product{
					ID:          1,
					Name:        "Ноутбук",
					Description: "Мощный ноутбук для работы и игр",
					Price:       59990,
					Category:    "Электроника",
					Brand:       "Brand A",
					Rating:      4.5,
					Reviews:     120,
				},
				Quantity: 1,
			},
			{
				Product: Product{
					ID:          2,
					Name:        "Смартфон",
					Description: "Современный смартфон с отличной камерой",
					Price:       29990,
					Category:    "Электроника",
					Brand:       "Brand B",
					Rating:      4.7,
					Reviews:     200,
				},
				Quantity: 2,
			},
			{
				Product: Product{
					ID:          3,
					Name:        "Наушники",
					Description: "Беспроводные наушники с шумоподавлением",
					Price:       7990,
					Category:    "Аудио",
					Brand:       "Brand C",
					Rating:      4.3,
					Reviews:     80,
				},
				Quantity: 1,
			},
		},
	}
	printInfo(user)
}

// printInfo вывод информации о пользователе
func printInfo(user User) {
	// Покупатель [ИМЯ]. Телефон: [ТЕЛЕФОН]. Адрес: г. [ГОРОД], [УЛИЦА].
	fmt.Printf("Покупатель %s. Телефон: %s. Адрес: г. %s, %s.\n", user.Name, user.Phone, user.Address.City, user.Address.Street)

	// Пользователь [является/не является] покупателем электроники.
	buyer := false
	for _, value := range user.Cart {
		if value.Product.Category == "Электроника" {
			buyer = true
			break
		}
	}
	if buyer {
		fmt.Printf("Пользователь является покупателем электроники.\n")
	} else {
		fmt.Printf("Пользователь не является покупателем электроники.\n")
	}

	// Товары в корзине, где цена 10000 и более: [НАЗВАНИЯ_ТОВАРОВ].
	// Товары в корзине, где цена 10000 и более: отсутствуют.
	price := 10000
	products := []string{}
	for _, value := range user.Cart {
		if value.Product.Price >= price {
			products = append(products, value.Product.Name)
		}
	}
	if len(products) == 0 {
		fmt.Printf("Товары в корзине, где цена %d и более: отсутствуют.\n", price)
	} else {
		fmt.Printf("Товары в корзине, где цена %d и более: %s.\n", price, strings.Join(products, ", "))
	}

	// Общая сумма покупки: [СУММА] руб.
	// Price * Quantity
	cartPrice := 0
	for _, value := range user.Cart {
		cartPrice += value.Product.Price * value.Quantity
	}
	fmt.Printf("Общая сумма покупки: %d руб.\n", cartPrice)
}
