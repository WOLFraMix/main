package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
)

const (
	MinPasswordLength = 4
	MinPasswordsCount = 1
	MaxPasswordsCount = 50
)

var (
	ErrPasswordLengthTooLow = errors.New("password length too low")
	ErrPasswordsCountTooLow = errors.New("too low passwords count")
	ErrPasswordsCountTooBig = errors.New("too big passwords count")
)

var (
	upperChars   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerChars   = []rune("abcdefghijklmnopqrstuvwxyz")
	digitChars   = []rune("0123456789")
	specialChars = []rune("!@#$%^&*")
)

// generatePassword генерирует count паролей длиной length.
func generatePassword(length int, count int) ([]string, error) {
	result := make([]string, 0, count)
	if length < MinPasswordLength {
		return nil, ErrPasswordLengthTooLow
	}
	if count < MinPasswordsCount {
		return nil, ErrPasswordsCountTooLow
	}
	if count > MaxPasswordsCount {
		return nil, ErrPasswordsCountTooBig
	}
	for j := 0; j < count; j++ {
		pass := make([]rune, length)
		pass[0], _ = randomRune(upperChars)
		pass[1], _ = randomRune(lowerChars)
		pass[2], _ = randomRune(digitChars)
		pass[3], _ = randomRune(specialChars)
		for i := 4; i < length; i++ {
			// Выбираем категорию случайно, чтобы не было жёсткого паттерна.
			category := randInt(5)
			switch category {
			case 0:
				pass[i], _ = randomRune(upperChars)
			case 1:
				pass[i], _ = randomRune(lowerChars)
			case 2:
				pass[i], _ = randomRune(digitChars)
			default:
				pass[i], _ = randomRune(specialChars)
			}
		}
		shuffleRunes(pass)
		result = append(result, string(pass))
		if !UniqueSlice(result) {
			result = result[:len(result)-1]
			shuffleRunes(pass)
			result = append(result, string(pass))
		}
	}
	return result, nil
}

// UniqueSlice проверяет все ли значения в слайсе уникальные.
func UniqueSlice(s []string) bool {
	m := make(map[string]struct{}, len(s))

	for i := 0; i < len(s); i++ {
		val := s[i]
		if _, ok := m[val]; ok {
			return false
		}
		m[val] = struct{}{}
	}
	return true
}

// randomRune возвращает рандомную руну.
func randomRune(set []rune) (rune, error) {
	n := big.NewInt(int64(len(set)))
	index, err := rand.Int(rand.Reader, n)
	if err != nil {
		return 0, err
	}
	return set[index.Int64()], nil
}

// randInt возвращает криптостойкое случайное число в диапазоне [0, n).
func randInt(n int) int {
	rnd, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic(err)
	}
	return int(rnd.Int64())
}

// shuffleRunes перемешивает срез рун.
func shuffleRunes(runes []rune) error {
	for i := len(runes) - 1; i > 0; i-- {
		j := randInt(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
	return nil
}

func main() {
	passwords, err := generatePassword(12, 5)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(passwords)
}
