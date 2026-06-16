package main

import "testing"

func TestIsValidEmail(t *testing.T) { // функция для тестирования функции IsValidEmail
	data := "email@example.com" // тестовые данные - валидный email адрес
	if !IsValidEmail(data) {    // если функция возвращает false, то выводим сообщение об ошибке
		t.Errorf("IsValidEmail(%v)=false, want true", data) // t.Errorf() - это функция для вывода сообщения об ошибке в тестах
		// %v - это формат для вывода значения переменной data
	}
}
