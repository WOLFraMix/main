package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

var (
	ErrInvalidRange = errors.New("invalid range")
	ErrSeatOccupied = errors.New("seat already occupied")
)

type Range struct {
	Start int
	End   int // включительно
}

type Cinema struct {
	seats []bool
}

func NewCinema(seats []bool) *Cinema {
	return &Cinema{seats: slices.Clone(seats)}
}

// FindAvailable ищет свободные диапазоны длиной >= k.
// Метод должен вернуть все найденные диапазоны.
// Если передается отрицательный аргумент или ноль - вернется пустой слайс.
func (c *Cinema) FindAvailable(k int) []Range {
	if k <= 0 {
		return []Range{}
	}

	var result []Range
	i := 0

	for i < len(c.seats) {
		if c.seats[i] == true {
			i++
			continue
		}
		start := i
		for i < len(c.seats) && c.seats[i] == false {
			i++
		}
		end := i - 1
		length := end - start + 1
		if length >= k {
			result = append(result, Range{
				Start: start,
				End:   end,
			})
		}
	}

	return result
}

// Book бронирует диапазон мест [start, end] включительно.
func (c *Cinema) Book(start, end int) error {
	if start > end || start < 0 || end >= len(c.seats) {
		return ErrInvalidRange
	}
	for i := start; i <= end; i++ {
		if c.seats[i] == true {
			return ErrSeatOccupied
		}
	}
	for i := start; i <= end; i++ {
		c.seats[i] = true
	}
	return nil
}

// Функция main и все тесты будут скрыты от вас на сайте во время проверки.
func main() {
	if !test1() || !test2() || !test3() || !test4() || !test5() || !test6() || !test7() || !test8() || !test9() {
		os.Exit(1)
	}
	fmt.Println("Все тесты успешно пройдены!")
}

// Проверка поиска свободных мест на краях, чекаем баг с индексом 0 и концом среза
func test1() bool {
	seats := []bool{false, false, false, true, false, false, false}
	cinema := NewCinema(seats)

	res := cinema.FindAvailable(3)
	if len(res) != 2 {
		fmt.Fprintf(os.Stderr, "Тест 1: Ожидали 2 диапазона по 3 места, получили: %v\n", res)
		return false
	}
	if res[0].Start != 0 || res[0].End != 2 || res[1].Start != 4 || res[1].End != 6 {
		fmt.Fprintf(os.Stderr, "Тест 1: Неверные границы диапазонов на краях зала: %v\n", res)
		return false
	}
	return true
}

// Проверка на корректность фильтрации k <= 0
func test2() bool {
	cinema := NewCinema([]bool{false, false, false})
	if len(cinema.FindAvailable(0)) != 0 || len(cinema.FindAvailable(-2)) != 0 {
		fmt.Fprintf(os.Stderr, "Тест 2: При k <= 0 метод FindAvailable должен возвращать пустой слайс\n")
		return false
	}
	return true
}

// Защита от дублирования внутренних состояний
func test3() bool {
	seats := []bool{true, false, false, true, false, true}
	cinema := NewCinema(seats)
	res := cinema.FindAvailable(2)
	if len(res) != 1 || res[0].Start != 1 || res[0].End != 2 {
		fmt.Fprintf(os.Stderr, "Тест 3: Неверная фильтрация диапазонов, ожидали только {1, 2}, получено: %v\n", res)
		return false
	}
	return true
}

// Валидация диапазонов и защита от паники Out of Range при бронировании
func test4() bool {
	cinema := NewCinema([]bool{false, false, false})

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Тест 4: Метод Book упал в панику: %v\n", r)
		}
	}()

	if err := cinema.Book(0, len(cinema.seats)); !errors.Is(err, ErrInvalidRange) {
		fmt.Fprintf(os.Stderr, "Тест 4: Book(0, %d) вышел за границы, ожидали ErrInvalidRange, получили: %v\n", len(cinema.seats), err)
		return false
	}
	if err := cinema.Book(-1, 2); !errors.Is(err, ErrInvalidRange) {
		fmt.Fprintf(os.Stderr, "Тест 4: Book(-1, 2) содержит отрицательный индекс, ожидали ErrInvalidRange, получили: %v\n", err)
		return false
	}
	if err := cinema.Book(2, 1); !errors.Is(err, ErrInvalidRange) {
		fmt.Fprintf(os.Stderr, "Тест 4: Book(2, 1) нарушен порядок start > end, ожидали ErrInvalidRange, получили: %v\n", err)
		return false
	}
	return true
}

// Проверяем "атомарность" операции бронирования
func test5() bool {
	seats := []bool{false, false, true}
	cinema := NewCinema(seats)

	err := cinema.Book(0, 2)
	if !errors.Is(err, ErrSeatOccupied) {
		fmt.Fprintf(os.Stderr, "Тест 5: Место занято, ожидали ErrSeatOccupied, получили: %v\n", err)
		return false
	}
	if cinema.seats[0] || cinema.seats[1] {
		fmt.Fprintf(os.Stderr, "Тест 5: Нарушена атомарность. Места успели частично забронироваться до падения с ошибкой\n")
		return false
	}

	// Тут стандартно все должно быть.
	if err := cinema.Book(0, 1); err != nil {
		fmt.Fprintf(os.Stderr, "Тест 5: Не удалось забронировать свободные места 0-1: %v\n", err)
		return false
	}
	if !cinema.seats[0] || !cinema.seats[1] {
		fmt.Fprintf(os.Stderr, "Тест 5: Места 0 и 1 должны быть заняты после успешного вызова Book\n")
		return false
	}
	return true
}

// Проверяем непрерывные ряды, длина которых больше k.
// Метод должен возвращать один целостный Range для каждого ряда, без дублирования и "наползания" индексов.
func test6() bool {
	// Кейс из условия задачи. Ряд из 4 свободных мест подряд (индексы 5-8) при k = 3.
	seats := []bool{true, false, false, false, true, false, false, false, false, true}
	cinema := NewCinema(seats)

	res := cinema.FindAvailable(3)

	// 2 диапазона должно быть
	if len(res) != 2 {
		fmt.Fprintf(os.Stderr, "Тест 6: Ожидали ровно 2 диапазона, получили %d: %v\n", len(res), res)
		return false
	}

	if res[0].Start != 1 || res[0].End != 3 || res[1].Start != 5 || res[1].End != 8 {
		fmt.Fprintf(os.Stderr, "Тест 6: Неверно определены границы сплошного ряда свободных мест: %v\n", res)
		return false
	}
	return true
}

// Проверка на лишнее бронирование и полноту бронирования диапазона
func test7() bool {
	cinema := NewCinema([]bool{false, false, false, false, false})

	// Бронируем индексы 1, 2 и 3. Индексы 0 и 4 должны остаться false.
	if err := cinema.Book(1, 3); err != nil {
		fmt.Fprintf(os.Stderr, "Тест 7: Не удалось забронировать места 1-3: %v\n", err)
		return false
	}

	// Проверяем, что забронировались все места внутри диапазона
	for _, idx := range []int{1, 2, 3} {
		if !cinema.seats[idx] {
			fmt.Fprintf(os.Stderr, "Тест 7: Метод Book не забронировал место внутри диапазона (индекс %d)\n", idx)
			return false
		}
	}

	// Проверяем, что не забронировалось ничего лишнего за пределами
	for _, idx := range []int{0, 4} {
		if cinema.seats[idx] {
			fmt.Fprintf(os.Stderr, "Тест 7: Метод Book забронировал лишнее место за пределами диапазона (индекс %d)\n", idx)
			return false
		}
	}

	return true
}

// Проверка фильтрации короткого диапазона в конце слайса
func test8() bool {
	// В конце слайса 2 свободных места, а ищем мы 3, этот хвост не должен попасть в результат.
	seats := []bool{false, false, false, true, false, false}
	cinema := NewCinema(seats)

	res := cinema.FindAvailable(3)

	if len(res) != 1 {
		fmt.Fprintf(os.Stderr, "Тест 8: Ожидали 1 диапазон, короткий хвост в конце должен отсеяться. Получено: %v\n", res)
		return false
	}

	if res[0].Start != 0 || res[0].End != 2 {
		fmt.Fprintf(os.Stderr, "Тест 8: Неверные границы диапазона: %v\n", res)
		return false
	}

	return true
}

// Проверка, когда k больше любого свободного диапазона в зале
func test9() bool {
	// Свободные диапазоны длиной 2, 3 и 1, а ищем 4.
	seats := []bool{false, false, true, false, false, false, true, false}
	cinema := NewCinema(seats)

	res := cinema.FindAvailable(4)

	if len(res) != 0 {
		fmt.Fprintf(os.Stderr, "Тест 9: Ожидали пустой слайс, так как нет диапазонов длиной >= 4. Получено: %v\n", res)
		return false
	}

	return true
}
