package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Обёртка
type JSON struct {
	ID       int       `json:"id"`
	Course   int       `json:"course"`
	Students []Student `json:"students"`
}

type Student struct {
	LastName  string `json:"LastName"`
	FirstName string `json:"FirstName"`
	Rating    []int  `json:"Rating"`
}

func main() {
	file, err := os.Open("stepik4-4-2.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var json JSON
	if err := decoder.Decode(&json); err != nil {
		log.Fatalf("ошибка декодирования JSON: %v", err)
	}

	sum := 0.0
	count := 0.0
	for _, student := range json.Students {
		for _, v := range student.Rating {
			sum += float64(v)
			count++
		}
	}
	avg := sum / count
	fmt.Printf("%0.2f", avg)
}
