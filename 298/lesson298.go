package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"time"
)

type LoginRecord struct {
	Username  string
	LoginTime time.Time
}

// FindNightOwls находит пользователей, которые входили в систему в ночное время.
// nightStart и nightEnd — часы начала и конца ночного интервала (0-23).
// Интервал может пересекать полночь (например, 22:00 - 06:00).
func FindNightOwls(records []LoginRecord, nightStart, nightEnd int) []string {
	// nameCounts хранит количество ночных входов для каждого пользователя.
	nameCounts := make(map[string]int)
	// result — итоговый слайс с именами пользователей.
	result := []string{}
	// orderIndex запоминает индекс самого первого входа для каждого пользователя
	// (нужен для стабильной сортировки при равном количестве ночных входов).
	orderIndex := make(map[string]int)

	// Проходим по всем записям о входах.
	for i, login := range records {
		// Если пользователь встречается впервые,
		// запоминаем индекс его первого появления.
		if _, ok := orderIndex[login.Username]; !ok {
			orderIndex[login.Username] = i
		}

		// Извлекаем час из времени входа.
		h := login.LoginTime.Hour()
		var isNight bool

		// Проверяем, попадает ли час входа в заданный ночной интервал.
		if nightStart < nightEnd {
			// Обычный интервал в пределах одних суток
			// (например, 9:00 - 17:00).
			isNight = h >= nightStart && h < nightEnd
		} else {
			// Интервал, пересекающий полночь
			// (например, 22:00 - 06:00).
			// Это либо время после начала, либо время до конца.
			isNight = h >= nightStart || h < nightEnd
		}

		// Если вход был ночным:
		if isNight {
			// Если это первый ночной вход пользователя,
			// добавляем его в итоговый список.
			if _, ok := nameCounts[login.Username]; !ok {
				result = append(result, login.Username)
			}
			// Увеличиваем счетчик ночных входов для этого пользователя.
			nameCounts[login.Username]++
		}
	}

	// Сортируем итоговый список пользователей.
	// sort.SliceStable сохраняет относительный порядок элементов,
	// если функция сравнения возвращает false.
	sort.SliceStable(result, func(i, j int) bool {
		// Сначала сортируем по убыванию количества ночных входов.
		if nameCounts[result[i]] != nameCounts[result[j]] {
			return nameCounts[result[i]] > nameCounts[result[j]]
		}
		// Если количество входов одинаковое,
		// сортируем по возрастанию индекса первого появления.
		return orderIndex[result[i]] < orderIndex[result[j]]
	})

	return result
}

// Функция main и все тесты будут скрыты от вас при проверке на сайте.
func main() {
	if !test1() ||
		!test2() ||
		!test3() ||
		!test4() ||
		!test5() ||
		!test6() ||
		!test7() ||
		!test8() ||
		!test9() ||
		!test10() ||
		!test11() {
		os.Exit(1)
	}

	fmt.Println("Все тесты успешно пройдены!")
}

func mustParse(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}
	return t
}

// Обычный пример, проверка сортировки по количеству ночных входов.
func test1() bool {
	records := []LoginRecord{
		{"Иван Петров", mustParse("2026-01-15T23:45:00Z")},
		{"Мария Сидорова", mustParse("2026-01-15T14:30:00Z")},
		{"Иван Петров", mustParse("2026-01-16T02:10:00Z")},
		{"Алексей Смирнов", mustParse("2026-01-15T03:20:00Z")},
		{"Мария Сидорова", mustParse("2026-01-16T01:05:00Z")},
		{"Иван Петров", mustParse("2026-01-16T23:55:00Z")},
	}

	expected := []string{
		"Иван Петров",
		"Мария Сидорова",
		"Алексей Смирнов",
	}

	got := FindNightOwls(records, 23, 6)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 1:\nОжидали: %v\nПолучили: %v\n", expected, got)
		return false
	}

	return true
}

// Никто ночью не входил, должен вернуться пустой слайс.
func test2() bool {
	records := []LoginRecord{
		{"Олег Морозов", mustParse("2026-03-10T10:00:00Z")},
		{"Анна Лебедева", mustParse("2026-03-10T15:45:00Z")},
	}

	expected := []string{}

	got := FindNightOwls(records, 23, 6)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 2:\nОжидали пустой слайс\nПолучили: %v\n", got)
		return false
	}

	return true
}

// Проверяем сортировку по количеству входов.
func test3() bool {
	records := []LoginRecord{
		{"Павел Тарасов", mustParse("2026-04-05T01:00:00Z")},
		{"Игорь Кубиков", mustParse("2026-04-05T02:00:00Z")},
		{"Павел Тарасов", mustParse("2026-04-06T03:00:00Z")},
		{"Игорь Кубиков", mustParse("2026-04-06T04:00:00Z")},
		{"Павел Тарасов", mustParse("2026-04-07T05:00:00Z")},
	}

	expected := []string{
		"Павел Тарасов",
		"Игорь Кубиков",
	}

	got := FindNightOwls(records, 0, 6)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 3:\nОжидали: %v\nПолучили: %v\n", expected, got)
		return false
	}

	return true
}

// Пустой список логов.
func test4() bool {
	records := []LoginRecord{}

	expected := []string{}

	got := FindNightOwls(records, 23, 6)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 4:\nОжидали пустой слайс\nПолучили: %v\n", got)
		return false
	}

	return true
}

// Проверяем подсчет большого количества ночных входов.
func test5() bool {
	records := []LoginRecord{
		{"Ночной хулиган", mustParse("2026-05-01T23:00:00Z")},
		{"Ночной хулиган", mustParse("2026-05-01T23:30:00Z")},
		{"Ночной хулиган", mustParse("2026-05-02T00:00:00Z")},
		{"Ночной хулиган", mustParse("2026-05-02T05:59:00Z")},
		{"Ранняя птичка", mustParse("2026-05-01T23:00:00Z")},
		{"Ранняя птичка", mustParse("2026-05-02T00:00:00Z")},
	}

	expected := []string{
		"Ночной хулиган",
		"Ранняя птичка",
	}

	got := FindNightOwls(records, 23, 6)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 5:\nОжидали: %v\nПолучили: %v\n", expected, got)
		return false
	}

	return true
}

// Проверяем границы ночного интервала.
// 21:00 входит в ночь.
// 05:59 входит.
// 06:00 уже не входит.
func test6() bool {
	records := []LoginRecord{
		{"Совенок", mustParse("2026-07-01T21:30:00Z")}, // 1 валидный
		{"Совенок", mustParse("2026-07-02T06:00:00Z")}, // 06:00 (nightEnd), не входит
		{"Сова", mustParse("2026-07-01T21:00:00Z")},    // 1 валидный
		{"Сова", mustParse("2026-07-01T22:00:00Z")},    // 1 валидный
	}

	expected := []string{
		"Сова",    // 2 входа
		"Совенок", // 1 вход
	}

	got := FindNightOwls(records, 21, 6)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 6:\nОжидали: %v\nПолучили: %v\n", expected, got)
		return false
	}

	return true
}

// При одинаковом количестве пользователи должны идти в порядке первого появления.
func test7() bool {
	records := []LoginRecord{
		{"Пользователь Б", mustParse("2026-01-01T23:00:00Z")},
		{"Пользователь А", mustParse("2026-01-01T23:10:00Z")},
		{"Пользователь Б", mustParse("2026-01-02T00:00:00Z")},
		{"Пользователь А", mustParse("2026-01-02T00:10:00Z")},
	}

	expected := []string{
		"Пользователь Б",
		"Пользователь А",
	}

	got := FindNightOwls(records, 23, 6)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 7:\nОжидали: %v\nПолучили: %v\n", expected, got)
		return false
	}

	return true
}

// Проверяем случай, когда ночной период не переходит через полночь.
func test8() bool {
	records := []LoginRecord{
		{"Ранний", mustParse("2026-08-01T02:00:00Z")},
		{"Дневной", mustParse("2026-08-01T12:00:00Z")},
		{"Вечерний", mustParse("2026-08-01T18:00:00Z")},
		{"Ранний", mustParse("2026-08-01T05:59:00Z")},
	}

	expected := []string{
		"Ранний",
	}

	got := FindNightOwls(records, 2, 6)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 8:\nОжидали: %v\nПолучили: %v\n", expected, got)
		return false
	}

	return true
}

// Проверяем логи до начала непересекающего полночь интервала, т.е. к примеру
// логин в 01:00 не должен засчитываться для интервала с 2 до 6.
func test9() bool {
	records := []LoginRecord{
		{"Ранний птах", mustParse("2026-08-01T01:00:00Z")}, // Не входит
		{"Ночная сова", mustParse("2026-08-01T02:30:00Z")}, // Входит
		{"Ранний птах", mustParse("2026-08-01T07:00:00Z")}, // Не входит
	}

	expected := []string{
		"Ночная сова",
	}

	got := FindNightOwls(records, 2, 6)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 9:\nОжидали: %v\nПолучили: %v\n", expected, got)
		return false
	}

	return true
}

// Проверка работы функции с дневным интервалом.
func test10() bool {
	records := []LoginRecord{
		{"Сотрудник Б", mustParse("2026-09-01T12:00:00Z")}, // 1 валидный
		{"Сотрудник Б", mustParse("2026-09-01T14:00:00Z")}, // 14:00 (nightEnd), не входит
		{"Сотрудник А", mustParse("2026-09-01T11:00:00Z")}, // Не входит
		{"Сотрудник В", mustParse("2026-09-01T12:30:00Z")}, // 1 валидный
		{"Сотрудник В", mustParse("2026-09-01T13:30:00Z")}, // 1 валидный
	}

	expected := []string{
		"Сотрудник В", // 2 входа
		"Сотрудник Б", // 1 вход
	}

	got := FindNightOwls(records, 12, 14)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 10:\nОжидали: %v\nПолучили: %v\n", expected, got)
		return false
	}

	return true
}

// Проверяем, что сортировка по количеству входов работает.
func test11() bool {
	// Пользователь А будет первым, но у пользователя Б будет больше ночных входов.
	records := []LoginRecord{
		{"Пользователь А", mustParse("2026-01-01T23:00:00Z")}, // 1 ночной вход
		{"Пользователь Б", mustParse("2026-01-01T23:10:00Z")}, // 1-й ночной вход
		{"Пользователь А", mustParse("2026-01-02T10:00:00Z")}, // Дневной вход (не считается)
		{"Пользователь Б", mustParse("2026-01-02T00:10:00Z")}, // 2-й ночной вход
	}

	expected := []string{
		"Пользователь Б", // У него 2 входа, должен быть первым
		"Пользователь А", // У этого 1 вход, должен быть вторым
	}

	got := FindNightOwls(records, 23, 6)

	if !reflect.DeepEqual(got, expected) {
		fmt.Fprintf(os.Stderr, "Тест 11:\nОжидали: %v\nПолучили: %v\n", expected, got)
		return false
	}

	return true
}
