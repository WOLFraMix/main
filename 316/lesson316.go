package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ProcessLogs фильтрует логи по minLevel и конвертирует время в UTC.
func ProcessLogs(logs []string, minLevel string, sourceTZ string) ([]string, error) {
	result := make([]string, 0, len(logs))

	for _, log := range logs {
		hasSeparator := strings.Contains(log, " | ")
		if !hasSeparator {
			return []string{}, fmt.Errorf("invalid log format: %s", log)
		}

		logParts := strings.Split(log, " | ")
		if len(logParts) < 3 {
			return []string{}, fmt.Errorf("invalid log format: %s", log)
		}

		logLevels := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
		found := false
		for _, level := range logLevels {
			if logParts[1] == level {
				found = true
				break
			}
		}
		if !found {
			return []string{}, fmt.Errorf("unknown log level: %s", logParts[1])
		}

		level := false
		switch minLevel {
		case "FATAL":
			if logParts[1] == "FATAL" {
				level = true
			}
		case "ERROR":
			if logParts[1] == "FATAL" || logParts[1] == "ERROR" {
				level = true
			}
		case "WARN":
			if logParts[1] == "FATAL" || logParts[1] == "ERROR" || logParts[1] == "WARN" {
				level = true
			}
		case "INFO":
			if logParts[1] == "FATAL" || logParts[1] == "ERROR" || logParts[1] == "WARN" || logParts[1] == "INFO" {
				level = true
			}
		case "DEBUG":
			level = true
		default:
			return []string{}, fmt.Errorf("unknown minLevel: %s\n", minLevel)
		}

		if level {
			loc, err := time.LoadLocation(sourceTZ)
			if err != nil {
				return []string{}, fmt.Errorf("failed to load location: %w\n", err)
			}

			t, err := time.ParseInLocation("2006-01-02 15:04:00", logParts[0], loc)
			if err != nil {
				return []string{}, fmt.Errorf("failed to parse time '%s': %w\n", logParts[0], err)
			}

			utcTime := t.UTC()
			timeResult := utcTime.Format("2006-01-02 15:04:05 MST")

			s := timeResult + " | " + logParts[1] + " | " + logParts[2]
			result = append(result, s)
		}
	}
	return result, nil
}

// Функция main будет скрыта от вас
func main() {
	// Логи
	inputLogs := []string{
		"2026-10-25 14:30:00 | INFO | Application started",
		"2026-10-25 14:35:00 | ERROR | Database connection failed",
		"2026-10-25 14:40:00 | WARN | High memory usage",
		"2026-10-25 15:00:00 | FATAL | System shutdown initiated",
	}
	// Минимальный level
	minLevel := "WARN"
	// Часовой пояс логов
	sourceTZ := "Europe/Moscow"

	result, err := ProcessLogs(inputLogs, minLevel, sourceTZ)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %s\n", err)
		os.Exit(1)
	}

	// Просто вывод в консоль
	for _, line := range result {
		fmt.Println(line)
	}
}
