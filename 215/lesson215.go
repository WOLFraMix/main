package main

import (
	"errors"
	"fmt"
)

// User — структура пользователя
type User struct {
	ID   string
	Name string
}

// userStore — хранилище пользователей в памяти (для примера)
var userStore = map[string]*User{
	"1": {ID: "1", Name: "Alice"},
	"2": {ID: "2", Name: "Bob"},
}

// GetUserByID возвращает пользователя по ID или ошибку
func GetUserByID(id string) (*User, error) {
	user, ok := userStore[id]
	if !ok {
		return nil, fmt.Errorf("user with ID %q not found", id)
	}
	// Возвращаем копию указателя,
	// чтобы избежать случайного изменения состояния хранилища снаружи
	return user, nil
}

// SaveUser сохраняет или обновляет пользователя в хранилище
func SaveUser(user *User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	if user.ID == "" {
		return errors.New("user ID cannot be empty")
	}

	// В реальной системе здесь была бы логика валидации, транзакций и т.д.
	userStore[user.ID] = user
	return nil
}

// UpdateUserName обновляет имя пользователя,
// используя GetUserByID и SaveUser.
func UpdateUserName(id, name string) error {
	user, err := GetUserByID(id)
	if err != nil {
		return fmt.Errorf("GetUserByID: %w", err)
	}

	user.Name = name

	if err := SaveUser(user); err != nil {
		return fmt.Errorf("SaveUser: %w", err)
	}

	return nil
}

func main() {
	// Пример использования
	err := UpdateUserName("1", "Alicia")
	if err != nil {
		fmt.Printf("Error updating user: %v\n", err)
		return
	}

	updatedUser, _ := GetUserByID("1")
	fmt.Printf("Updated user: ID=%s, Name=%s\n", updatedUser.ID, updatedUser.Name)
}
