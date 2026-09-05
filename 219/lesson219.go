package main

import (
	"errors"
	"fmt"
)

// Данные
type Message struct {
	From    string
	Message string
}

type Chat struct {
	ID       int
	Messages []Message
}

// Ошибка базы данных
type DatabaseError struct {
	Message string
	Code    int
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error with code %d: %s", e.Code, e.Message)
}

// Сервис для работы
type Service struct {
	db interface {
		GetChatByIDWithMessages(id int) (*Chat, error)
	}
}

// вывод разных ошибок
func (w Service) PrintChat(id int) {
	chat, err := w.db.GetChatByIDWithMessages(id)
	if err != nil {
		if DatabaseError, ok := errors.AsType[*DatabaseError](err); ok {
			if DatabaseError.Code == 24 {
				fmt.Printf("Ошибка запроса: %s", err)
			} else {
				fmt.Printf("Инфраструктурная ошибка: %s", err)
			}
		} else {
			fmt.Printf("Неизвестная ошибка: %s", err)
		}
	}
	for _, message := range chat.Messages {
		fmt.Printf("%s: %s\n", message.From, message.Message)
	}
}
