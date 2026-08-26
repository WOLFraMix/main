package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	FirstName         string
	LastName          string
	BirthYear         int
	FavoriteLanguages []string
}

// SecretIdentity() string - генерирует секретное имя.
// Составляется из первых букв имени и фамилии + случайное число.
func (u *User) SecretIdentity() string {
	runesA := []rune(u.FirstName)
	a := string(runesA[0])
	runesB := []rune(u.LastName)
	b := string(runesB[0])
	c := strconv.Itoa(random(1, 100))
	result := a + b + c
	return result
}

// random генерирует случайное число от min до max.
func random(min, max int) int {
	return min + rand.IntN(max-min+1)
}

// Age() int - возвращает текущий возраст пользователя.
func (u *User) Age() int {
	age := time.Now().Year() - u.BirthYear
	return age
}

// AddFavoriteLanguage(language string) error - добавляет язык в FavoriteLanguages.
func (u *User) AddFavoriteLanguage(language string) error {
	if language == "" || language == " " {
		return fmt.Errorf("empty language name")
	}

	for _, v := range u.FavoriteLanguages {
		if v == language {
			return fmt.Errorf("duplicate")
		}
	}

	u.FavoriteLanguages = append(u.FavoriteLanguages, language)
	return nil
}

// RemoveFavoriteLanguage(language string) error - удаляет язык из FavoriteLanguages.
func (u *User) RemoveFavoriteLanguage(language string) error {
	for i, v := range u.FavoriteLanguages {
		if v == language {
			u.FavoriteLanguages = append(u.FavoriteLanguages[:i], u.FavoriteLanguages[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("not found")
}

// IsProgrammingLanguageFavorite(language string) bool - проверяет, есть ли язык в списке.
func (u *User) IsProgrammingLanguageFavorite(language string) bool {
	for _, v := range u.FavoriteLanguages {
		if v == language {
			return true
		}
	}
	return false
}

// RandomFavoriteLanguage() (string, error) - возвращает случайный язык.
func (u *User) RandomFavoriteLanguage() (string, error) {
	if len(u.FavoriteLanguages) == 0 {
		return "", fmt.Errorf("no options")
	}
	r := random(0, len(u.FavoriteLanguages)-1)
	return u.FavoriteLanguages[r], nil
}

// GenerateProfile() string - возвращает строку с полным профилем пользователя.
func (u *User) GenerateProfile() string {
	fav := "[" + strings.Join(u.FavoriteLanguages, ", ") + "]"
	return fmt.Sprintf("Имя: %s.\nФамилия: %s.\nВозраст: %d.\nСписок любимых языков программирования: %s.", u.FirstName, u.LastName, u.Age(), fav)
}

// UpdateName(firstName, lastName string) error - обновляет имя и фамилию пользователя.
func (u *User) UpdateName(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return fmt.Errorf("empty data")
	}
	if firstName == strings.ToLower(u.FirstName) || lastName == strings.ToLower(u.LastName) {
		return fmt.Errorf("invalid data")
	}
	u.FirstName = firstName
	u.LastName = lastName
	return nil
}

// Функция main и все тесты будут скрыты при проверке на сайте.
func main() {
	if !test1() || !test2() || !test3() || !test4() || !test5() ||
		!test6() || !test7() || !test8() || !test9() || !test10() || !test11() {
		os.Exit(1)
	}
	fmt.Println("Все тесты успешно пройдены!")
}

// SecretIdentity - префикс из первых букв + число 1..100
func test1() bool {
	u := &User{FirstName: "Алексей", LastName: "Смирнов"}
	id := u.SecretIdentity()

	runes := []rune(id)
	if len(runes) < 3 {
		fmt.Fprintf(os.Stderr, "Тест 1: Слишком короткое имя получили: %q\n", id)
		return false
	}

	prefix := string(runes[:2])
	if prefix != "АС" {
		fmt.Fprintf(os.Stderr, "Тест 1: Ожидали префикс 'АС', а получили %q\n", prefix)
		return false
	}

	num, err := strconv.Atoi(string(runes[2:]))
	if err != nil || num < 1 || num > 100 {
		fmt.Fprintf(os.Stderr, "Тест 1: Число должно быть в [1,100], получили: %q\n", string(runes[2:]))
		return false
	}

	return true
}

// SecretIdentity - проверим английский язык отдельно
func test2() bool {
	u := &User{FirstName: "Pavel", LastName: "Tarasov"}
	id := u.SecretIdentity()

	runes := []rune(id)
	if len(runes) < 3 {
		fmt.Fprintf(os.Stderr, "Тест 2: Слишком короткое имя получили: %q\n", id)
		return false
	}

	prefix := string(runes[:2])
	if prefix != "PT" {
		fmt.Fprintf(os.Stderr, "Тест 2: Ожидали префикс 'PT', а получили %q\n", prefix)
		return false
	}

	num, err := strconv.Atoi(string(runes[2:]))
	if err != nil || num < 1 || num > 100 {
		fmt.Fprintf(os.Stderr, "Тест 2: Число должно быть в [1,100], получили: %q\n", string(runes[2:]))
		return false
	}

	return true
}

// Age - текущий год минус год рождения
func test3() bool {
	u := &User{BirthYear: 1990}
	expected := time.Now().Year() - 1990
	got := u.Age()
	if got != expected {
		fmt.Fprintf(os.Stderr, "Тест 3: Ожидали %d, получили %d\n", expected, got)
		return false
	}
	return true
}

// AddFavoriteLanguage - добавляем, ловим дубликат и пустое имя
func test4() bool {
	u := &User{}

	if err := u.AddFavoriteLanguage("Go"); err != nil {
		fmt.Fprintf(os.Stderr, "Тест 4: Не ждали ошибку при добавлении языка 'Go': %v. Ты что, не любишь этот прекрасный язык? :)\n", err)
		return false
	}
	if err := u.AddFavoriteLanguage("Python"); err != nil {
		fmt.Fprintf(os.Stderr, "Тест 4: Не ждали ошибку при добавлении языка 'Python': %v\n", err)
		return false
	}

	// Проверяем добавление дубля
	if err := u.AddFavoriteLanguage("Go"); err == nil || err.Error() != "duplicate" {
		fmt.Fprintf(os.Stderr, "Тест 4: Ждали ошибку с текстом 'duplicate', получили: %v\n", err)
		return false
	}

	// Проверяем что будет с пустым именем
	if err := u.AddFavoriteLanguage(""); err == nil || err.Error() != "empty language name" {
		fmt.Fprintf(os.Stderr, "Тест 4: Ждали ошибку с текстом 'empty language name', получили: %v\n", err)
		return false
	}

	// Дубликат не должен был добавиться ранее
	if len(u.FavoriteLanguages) != 2 {
		fmt.Fprintf(os.Stderr, "Тест 4: Ожидалось что у нас будет 2 языка, по факту их %d\n", len(u.FavoriteLanguages))
		return false
	}

	return true
}

// RemoveFavoriteLanguage - удаляем существующий, ловим not found
func test5() bool {
	u := &User{FavoriteLanguages: []string{"Go", "Python", "Rust"}}

	if err := u.RemoveFavoriteLanguage("Python"); err != nil {
		fmt.Fprintf(os.Stderr, "Тест 5: Не ждали ошибку при удалении 'Python': %v\n", err)
		return false
	}
	if len(u.FavoriteLanguages) != 2 {
		fmt.Fprintf(os.Stderr, "Тест 5: Ждали 2 языка после удаления, получили %d\n", len(u.FavoriteLanguages))
		return false
	}

	// Пытаемся удалить несуществующий язык
	if err := u.RemoveFavoriteLanguage("Java"); err == nil || err.Error() != "not found" {
		fmt.Fprintf(os.Stderr, "Тест 5: Ждали ошибку с текстом 'not found', получили: %v\n", err)
		return false
	}

	// При пустом списке тоже должно быть not found
	empty := &User{}
	if err := empty.RemoveFavoriteLanguage("Go"); err == nil || err.Error() != "not found" {
		fmt.Fprintf(os.Stderr, "Тест 5: Ждали ошибку с текстом 'not found' для пустого списка, получили: %v\n", err)
		return false
	}

	return true
}

// IsProgrammingLanguageFavorite - ищем в списке, проверяем что будет если не найдем
func test6() bool {
	u := &User{FavoriteLanguages: []string{"Go", "Rust"}}

	if !u.IsProgrammingLanguageFavorite("Go") {
		fmt.Fprintf(os.Stderr, "Тест 6: Как так? Язык 'Go' должен быть любимым! ;)\n")
		return false
	}
	if u.IsProgrammingLanguageFavorite("Java") {
		fmt.Fprintf(os.Stderr, "Тест 6: Язык 'Java' не должен быть любимым\n")
		return false
	}

	// Пустой список
	empty := &User{}
	if empty.IsProgrammingLanguageFavorite("Go") {
		fmt.Fprintf(os.Stderr, "Тест 6: Пустой список, значит мы ничего не должны любить. Хотя мне нравится ход ваших мыслей :)\n")
		return false
	}

	return true
}

// RandomFavoriteLanguage - случайный язык, проверяем на "no options"
func test7() bool {
	// Пустой список
	u := &User{}
	_, err := u.RandomFavoriteLanguage()
	if err == nil || err.Error() != "no options" {
		fmt.Fprintf(os.Stderr, "Тест 7: Ждали ошибку с текстом 'no options' при пустом списке, получили: %v\n", err)
		return false
	}

	// Непустой список - должны получить язык из него
	u.FavoriteLanguages = []string{"Go", "Python", "Rust"}
	lang, err := u.RandomFavoriteLanguage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Тест 7: Не ждали ошибку: %v\n", err)
		return false
	}

	found := false
	for _, l := range u.FavoriteLanguages {
		if l == lang {
			found = true
			break
		}
	}
	if !found {
		fmt.Fprintf(os.Stderr, "Тест 7: Получили язык %q, которого не было в списке.\n", lang)
		return false
	}

	return true
}

// GenerateProfile - точный формат с языками
func test8() bool {
	u := &User{
		FirstName:         "Павел",
		LastName:          "Тарасов",
		BirthYear:         1990,
		FavoriteLanguages: []string{"Go", "JavaScript"},
	}

	expected := fmt.Sprintf(
		"Имя: Павел.\nФамилия: Тарасов.\nВозраст: %d.\nСписок любимых языков программирования: [Go, JavaScript].",
		time.Now().Year()-1990,
	)

	got := u.GenerateProfile()
	if got != expected {
		fmt.Fprintf(os.Stderr, "Тест 8:\nОжидали:\n%s\nПолучили:\n%s\n", expected, got)
		return false
	}

	return true
}

// UpdateName - валидные данные, пустые, с маленькой буквы, в общем все проверяем
func test9() bool {
	u := &User{}

	// Нормальное обновление
	if err := u.UpdateName("Иван", "Иванов"); err != nil {
		fmt.Fprintf(os.Stderr, "Тест 9: Не ожидали ошибку для 'Иван Иванов': %v\n", err)
		return false
	}
	if u.FirstName != "Иван" || u.LastName != "Иванов" {
		fmt.Fprintf(os.Stderr, "Тест 9: Имя не обновилось, должно быть 'Иван Иванов', по факту лежит: %s %s\n", u.FirstName, u.LastName)
		return false
	}

	// Пустые данные
	if err := u.UpdateName("", "Иванов"); err == nil || err.Error() != "empty data" {
		fmt.Fprintf(os.Stderr, "Тест 9: Ждали 'empty data' для пустого имени, получили: %v\n", err)
		return false
	}
	if err := u.UpdateName("Иван", ""); err == nil || err.Error() != "empty data" {
		fmt.Fprintf(os.Stderr, "Тест 9: Ждали 'empty data' для пустой фамилии, получили: %v\n", err)
		return false
	}

	// С маленькой буквы
	if err := u.UpdateName("иван", "Иванов"); err == nil || err.Error() != "invalid data" {
		fmt.Fprintf(os.Stderr, "Тест 9: Ждали 'invalid data' для 'иван', получили: %v\n", err)
		return false
	}
	if err := u.UpdateName("Иван", "иванов"); err == nil || err.Error() != "invalid data" {
		fmt.Fprintf(os.Stderr, "Тест 9: Ждали 'invalid data' для 'иванов', получили: %v\n", err)
		return false
	}

	return true
}

// UpdateName - проверим обновление имени и фамилии на английском языке
func test10() bool {
	u := &User{}

	// Нормальное обновление
	if err := u.UpdateName("Pavel", "Tarasov"); err != nil {
		fmt.Fprintf(os.Stderr, "Тест 10: Не ожидали ошибку для 'Pavel Tarasov': %v\n", err)
		return false
	}
	if u.FirstName != "Pavel" || u.LastName != "Tarasov" {
		fmt.Fprintf(os.Stderr, "Тест 10: Имя не обновилось, должно быть 'Pavel Tarasov', по факту лежит: %s %s\n", u.FirstName, u.LastName)
		return false
	}

	// Пустые данные
	if err := u.UpdateName("", "Tarasov"); err == nil || err.Error() != "empty data" {
		fmt.Fprintf(os.Stderr, "Тест 10: Ждали 'empty data' для пустого имени, получили: %v\n", err)
		return false
	}
	if err := u.UpdateName("Pavel", ""); err == nil || err.Error() != "empty data" {
		fmt.Fprintf(os.Stderr, "Тест 10: Ждали 'empty data' для пустой фамилии, получили: %v\n", err)
		return false
	}

	// С маленькой буквы
	if err := u.UpdateName("pavel", "Tarasov"); err == nil || err.Error() != "invalid data" {
		fmt.Fprintf(os.Stderr, "Тест 10: Ждали 'invalid data' для 'pavel', получили: %v\n", err)
		return false
	}
	if err := u.UpdateName("Pavel", "tarasov"); err == nil || err.Error() != "invalid data" {
		fmt.Fprintf(os.Stderr, "Тест 10: Ждали 'invalid data' для 'tarasov', получили: %v\n", err)
		return false
	}

	return true
}

// GenerateProfile - проверяем nil-слайс языков
func test11() bool {
	u := &User{
		FirstName: "Анна",
		LastName:  "Ковалева",
		BirthYear: 2000,
		// FavoriteLanguages не инициализирован, т.е. nil
	}

	expected := fmt.Sprintf(
		"Имя: Анна.\nФамилия: Ковалева.\nВозраст: %d.\nСписок любимых языков программирования: [].",
		time.Now().Year()-2000,
	)

	got := u.GenerateProfile()
	if got != expected {
		fmt.Fprintf(os.Stderr, "Тест 11:\nОжидали:\n%s\nПолучили:\n%s\n", expected, got)
		return false
	}

	return true
}
