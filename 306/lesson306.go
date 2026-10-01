package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(parseUserDate("2010-04-05"))
}

// parseUserDate парсит строку с датой в одном из поддерживаемых форматов.
// Поддерживаемые форматы:
//   - YYYY-MM-DD, например: "2012-04-22"
//   - DD.MM.YYYY, например: "05.02.2016"
//   - DD/MM/YYYY, например: "09/07/2005"
//
// Возвращает время в UTC (полночь) и nil в случае успеха.
// Возвращает time.Time{} и ошибку, если строка не соответствует ни одному формату
// или представляет собой невалидную дату.
func parseUserDate(input string) (time.Time, error) {
	if input == "" || input == " " {
		return time.Time{}, fmt.Errorf("empty input")
	}

	t, err := time.Parse("2006-01-02", input)
	if err != nil {
		t, err = time.Parse("02.01.2006", input)
		if err != nil {
			t, err = time.Parse("02/01/2006", input)
			if err != nil {
				return time.Time{}, fmt.Errorf("unsupported or invalid date format: %s", input)
			}
		}
	}
	return t, nil
}
