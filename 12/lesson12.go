package main

import (
	"log"
	"net/http"
)

func HelloWeb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Web!"))
	// Функция HelloWeb принимает два параметра: http.ResponseWriter и *http.Request.
}

func main() {
	http.HandleFunc("/", HelloWeb)
	// Функция http.HandleFunc устанавливает функцию HelloWeb
	// по пути /
	// а http.ListenAndServe запускает сервер по указанному адресу.
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Println("listen and serve:", err)
		// Если сервер не может запуститься
		// будет выведено сообщение об ошибке.
	}
}
