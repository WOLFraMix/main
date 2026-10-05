package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type JSONResponse struct {
	People []People `json:"people"`
}

type People struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IP        string `json:"ip_address"`
}

func main() {
	// Открытие файла для чтения
	file, err := os.Open("stepik4-4.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Создание нового декодера JSON
	decoder := json.NewDecoder(file)

	// Декодирование данных JSON
	var response JSONResponse
	if err := decoder.Decode(&response); err != nil {
		log.Fatalf("ошибка декодирования JSON: %v", err)
	}

	for _, person := range response.People {
		if person.Gender == "Female" {
			fmt.Println(person.FirstName, person.LastName)
		}
	}
}
