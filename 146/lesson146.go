package main

import (
	"fmt"
	"net/http"
)

// handler — HTTP-обработчик запросов:
// принимает ответный писатель и запрос.
func handler(w http.ResponseWriter, r *http.Request) {
	// Отправляем клиенту строку «Hello, World!»
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Регистрируем обработчик для корневого пути "/"
	http.HandleFunc("/", handler)

	// Запускаем HTTP-сервер на порту 8080;
	// nil означает использование стандартного мультиплексора.
	http.ListenAndServe(":8080", nil)
}
