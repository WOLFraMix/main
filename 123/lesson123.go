package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Объявляем структуру Person. Это шаблон для хранения данных о человеке.
type Person struct {
	// Поле Name типа string. Тег `json:"Имя"` говорит пакету json: «Используй ключ "Имя" вместо "Name"».
	Name string `json:"Имя"`
	// Поле Email типа string. Тег `json:"Почта"` задаёт ключ "Почта" в итоговом JSON.
	Email string `json:"Почта"`
	// Поле DateOfBirth типа time.Time. Тег `json:"-"` означает «Исключить это поле из JSON-результата».
	DateOfBirth time.Time `json:"-"`
}

func main() {
	// Создаём экземпляр структуры Person с нужными данными.
	// Поле DateOfBirth тоже нужно заполнить (тип time.Time обязателен), но в JSON оно не попадёт.
	p := Person{
		Name:        "Алекс",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(), // Используем текущее время как пример значения.
	}

	// Вызываем json.Marshal(p) — эта функция превращает структуру p в байтовый срез ([]byte) в формате JSON.
	data, err := json.Marshal(p)

	// Проверяем, не произошла ли ошибка при сериализации.
	if err != nil {
		// Если ошибка есть, выводим сообщение и завершаем выполнение функции.
		fmt.Println("Ошибка при сериализации:", err)
		return
	}

	// Преобразуем байтовый срез data в строку и выводим её.
	// Именно здесь мы увидим JSON вида {"Имя":"Алекс","Почта":"alex@yandex.ru"}.
	fmt.Println(string(data))
}
