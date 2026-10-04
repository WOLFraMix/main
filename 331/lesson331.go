package main

import (
	"fmt"
	"os"
	"time"
)

// utcDate возвращает полночь в UTC для заданного времени.
func utcDate(t time.Time) time.Time {
	tUTC := t.UTC()
	return time.Date(tUTC.Year(), tUTC.Month(), tUTC.Day(), 0, 0, 0, 0, time.UTC)
}

// localDate возвращает полночь (нормализованную в UTC) для календарной даты
// в исходном часовом поясе времени t.
func localDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func absDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}

func minutesForm(minutes int) string {
	if minutes >= 11 && minutes <= 14 {
		return "минут"
	}
	switch minutes % 10 {
	case 1:
		return "минута"
	case 2, 3, 4:
		return "минуты"
	default:
		return "минут"
	}
}

func daysForm(days int) string {
	if days%100 >= 11 && days%100 <= 14 {
		return "дней"
	}
	switch days % 10 {
	case 1:
		return "день"
	case 2, 3, 4:
		return "дня"
	default:
		return "дней"
	}
}

// formatMinutesAgo форматирует прошедшее время в минутах.
func formatMinutesAgo(diff time.Duration) string {
	n := int(-diff.Minutes())
	return fmt.Sprintf("%d %s назад", n, minutesForm(n))
}

// formatMinutesAhead форматирует будущее время в минутах.
func formatMinutesAhead(diff time.Duration) string {
	n := int(diff.Minutes())
	s := "минуту"
	if n != 1 {
		s = minutesForm(n)
	}
	return fmt.Sprintf("через %d %s", n, s)
}

// formatPast форматирует время target, находящееся в прошлом относительно now.
func formatPast(target, now time.Time, diff time.Duration) string {
	// Меньше часа назад
	if diff > -time.Hour {
		return formatMinutesAgo(diff)
	}

	targetDate := utcDate(target)
	nowDate := utcDate(now)
	targetLocalDate := localDate(target)
	nowLocalDate := localDate(now)
	timeStr := target.Format("15:04")

	// Тот же день (UTC)
	if targetDate.Equal(nowDate) {
		return fmt.Sprintf("в %s", timeStr)
	}

	// Вчера (UTC или по локальным календарным датам)
	if targetDate.Equal(nowDate.AddDate(0, 0, -1)) ||
		targetLocalDate.Equal(nowLocalDate.AddDate(0, 0, -1)) {
		return fmt.Sprintf("вчера в %s", timeStr)
	}

	// Дни: UTC-дата now минус локальная дата target
	days := int(nowDate.Sub(targetLocalDate).Hours() / 24)
	return fmt.Sprintf("%d %s назад", days, daysForm(days))
}

// formatFuture форматирует время target, находящееся в будущем относительно now.
func formatFuture(target, now time.Time, diff time.Duration) string {
	// Менее чем через час
	if diff < time.Hour {
		return formatMinutesAhead(diff)
	}

	targetDate := utcDate(target)
	nowDate := utcDate(now)
	targetLocalDate := localDate(target)
	nowLocalDate := localDate(now)
	timeStr := target.Format("15:04")

	// Сегодня (UTC)
	if targetDate.Equal(nowDate) {
		return fmt.Sprintf("в %s", timeStr)
	}

	// Завтра (UTC или по локальным календарным датам)
	if targetDate.Equal(nowDate.AddDate(0, 0, 1)) ||
		targetLocalDate.Equal(nowLocalDate.AddDate(0, 0, 1)) {
		return fmt.Sprintf("завтра в %s", timeStr)
	}

	// Дни: локальная дата target минус UTC-дата now
	days := int(targetLocalDate.Sub(nowDate).Hours() / 24)
	return fmt.Sprintf("через %d %s", days, daysForm(days))
}

// FormatRelativeTime принимает форматируемую дату target и текущую дату now
// и возвращает строковое представление их разницы на русском языке.
func FormatRelativeTime(target, now time.Time) string {
	cmp := target.Compare(now)
	diff := target.Sub(now)

	// Если разница меньше минуты
	if cmp == 0 || absDuration(diff) < time.Minute {
		return "только что"
	}

	if cmp < 0 {
		return formatPast(target, now, diff)
	}

	if cmp > 0 {
		return formatFuture(target, now, diff)
	}

	return "только что"
}

func main() {
	testCases := []testCase{
		{1, "2026-06-07T12:00:45Z", "2026-06-07T12:01:15Z", "только что"},
		{2, "2026-06-07T12:01:15Z", "2026-06-07T12:00:45Z", "только что"},
		{3, "2026-06-07T11:59:00Z", "2026-06-07T12:00:00Z", "1 минута назад"},
		{4, "2026-06-07T12:01:00Z", "2026-06-07T12:00:00Z", "через 1 минуту"},
		{5, "2026-06-07T11:45:00Z", "2026-06-07T12:00:00Z", "15 минут назад"},
		{6, "2026-06-07T12:22:00Z", "2026-06-07T12:00:00Z", "через 22 минуты"},
		{7, "2026-06-07T10:30:00Z", "2026-06-07T14:00:00Z", "в 10:30"},
		{8, "2026-06-07T16:45:00Z", "2026-06-07T14:00:00Z", "в 16:45"},
		{9, "2026-06-06T15:30:00Z", "2026-06-07T10:00:00Z", "вчера в 15:30"},
		{10, "2026-06-08T18:00:00Z", "2026-06-07T23:00:00Z", "завтра в 18:00"},
		{11, "2026-06-04T12:00:00Z", "2026-06-07T12:00:00Z", "3 дня назад"},
		{12, "2026-06-12T12:00:00Z", "2026-06-07T12:00:00Z", "через 5 дней"},
		{13, "2026-06-06T23:30:00Z", "2026-06-07T00:45:00Z", "вчера в 23:30"},
		{14, "2026-06-08T00:30:00Z", "2026-06-07T23:15:00Z", "завтра в 00:30"},
		{15, "2026-06-19T12:00:00Z", "2026-06-07T12:00:00Z", "через 12 дней"},
		{16, "2026-06-28T12:00:00Z", "2026-06-07T12:00:00Z", "через 21 день"},
		{17, "2026-06-06T23:55:00Z", "2026-06-07T00:05:00Z", "10 минут назад"},
		{18, "2026-03-28T12:00:00+01:00", "2026-03-30T12:00:00+02:00", "2 дня назад"},
		{19, "2026-10-24T12:00:00+02:00", "2026-10-26T12:00:00+01:00", "2 дня назад"},
		{20, "2026-10-25T01:30:00+02:00", "2026-10-26T03:30:00+01:00", "вчера в 01:30"},
		{21, "2026-03-28T12:00:00+01:00", "2026-03-30T12:00:00+02:00", "2 дня назад"},
		{22, "2026-10-25T12:00:00+01:00", "2026-10-26T12:00:00+01:00", "вчера в 12:00"},
		{23, "2026-06-07T11:00:00Z", "2026-06-07T12:00:00Z", "в 11:00"},
		{24, "2026-06-07T13:00:00Z", "2026-06-07T12:00:00Z", "в 13:00"},
		{25, "2026-05-15T10:30:00Z", "2026-06-15T14:00:00Z", "31 день назад"},
		{26, "2026-05-01T15:30:00Z", "2026-06-02T10:00:00Z", "32 дня назад"},
		{27, "2026-06-01T23:00:00Z", "2026-06-03T00:00:00Z", "2 дня назад"},
		{28, "2025-12-30T23:00:00Z", "2026-01-01T01:00:00Z", "2 дня назад"},
		{29, "2026-02-28T22:00:00Z", "2026-03-01T02:00:00Z", "вчера в 22:00"},
		{30, "2026-06-07T11:00:30Z", "2026-06-07T12:00:00Z", "59 минут назад"},
		{31, "2026-06-07T12:00:00Z", "2026-06-07T11:00:30Z", "через 59 минут"},
		{32, "2026-06-01T12:00:00+14:00", "2026-06-03T12:00:00-12:00", "3 дня назад"},
		{33, "2026-06-08T01:00:00+05:00", "2026-06-07T22:00:00+00:00", "в 01:00"},
		{34, "2026-06-08T01:00:00+03:00", "2026-06-08T01:00:00+03:00", "только что"},
		{35, "2026-06-02T01:00:00+05:00", "2026-06-02T20:00:00+00:00", "вчера в 01:00"},
		{36, "2026-06-04T01:00:00+05:00", "2026-06-02T22:00:00+00:00", "завтра в 01:00"},
		{37, "2026-06-10T13:00:00+14:00", "2026-06-11T01:00:00Z", "вчера в 13:00"},
	}

	errs := make([]string, 0)
	for _, testCase := range testCases {
		res := FormatRelativeTime(parseTime(testCase.target), parseTime(testCase.now))
		if res != testCase.expected {
			err := fmt.Sprintf(
				"#%d: c target=%s и now=%s вернулось %q вместо %q",
				testCase.number,
				testCase.target,
				testCase.now,
				res,
				testCase.expected,
			)
			errs = append(errs, err)
		}
	}

	if len(errs) == 0 {
		fmt.Println("Все тесты успешно пройдены")
		return
	}

	for _, err := range errs {
		fmt.Println(err)
	}
	os.Exit(1)
}

type testCase struct {
	number   int
	target   string
	now      string
	expected string
}

func parseTime(input string) time.Time {
	if t, err := time.Parse(time.RFC3339, input); err == nil {
		return t
	}

	return time.Time{}
}
