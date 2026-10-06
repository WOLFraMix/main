package main

import (
	"fmt"
	"net/http"
)

// handler — функция-обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	// Получаем все query parameters как map
	queryParams := r.URL.Query()

	// Получаем конкретный параметр
	name := queryParams.Get("name")
	age := queryParams.Get("age")

	// Если параметр может встречаться несколько раз
	ids := queryParams["id"] // возвращает []string

	fmt.Fprintf(w, "Name: %s, Age: %s, IDs: %v", name, age, ids)
}

func main() {
	// http.HandleFunc регистрирует функцию-обработчик для всех запросов,
	// путь которых начинается с "/" (корневой путь).
	http.HandleFunc("/", handler)
	// Запускает веб-сервер на порту 8080.
	http.ListenAndServe(":8080", nil)
}

// http://localhost:8080/?name=Ivan&age=25&id=10&id=20
