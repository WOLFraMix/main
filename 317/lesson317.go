package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ProcessLogs фильтрует логи по minLevel и конвертирует время в UTC.
func ProcessLogs(logs []string, minLevel string, sourceTZ string) ([]string, error) {
	// Загружаем локацию один раз на весь пакет логов — это оптимизация
	loc, err := time.LoadLocation(sourceTZ)
	if err != nil {
		return nil, fmt.Errorf("failed to load location: %w", err)
	}

	result := make([]string, 0, len(logs))

	for _, log := range logs {
		parts, err := parseLogLine(log)
		if err != nil {
			return nil, err
		}

		levelValid, err := validateLogLevel(parts[1])
		if err != nil {
			return nil, err
		}
		if !levelValid {
			// unknown level уже обработан в validateLogLevel, но здесь просто пропускаем
			return nil, fmt.Errorf("unknown log level: %s", parts[1])
		}

		if !isLevelAtLeast(parts[1], minLevel) {
			continue // пропускаем логи ниже порога
		}

		t, err := parseTimeInLocation(parts[0], loc)
		if err != nil {
			return nil, fmt.Errorf("failed to parse time '%s': %w", parts[0], err)
		}

		utcTimeStr := t.UTC().Format("2006-01-02 15:04:05 MST")
		newLog := utcTimeStr + " | " + parts[1] + " | " + parts[2]
		result = append(result, newLog)
	}

	return result, nil
}

// parseLogLine проверяет формат строки и разбивает на части.
func parseLogLine(log string) ([]string, error) {
	if !strings.Contains(log, " | ") {
		return nil, fmt.Errorf("invalid log format: %s", log)
	}

	parts := strings.SplitN(log, " | ", 3) // ограничиваем до 3 частей, остальное не нужно
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid log format: %s", log)
	}
	return parts, nil
}

// validateLogLevel проверяет, что уровень логирования входит в разрешённый список.
func validateLogLevel(level string) (bool, error) {
	validLevels := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	for _, l := range validLevels {
		if level == l {
			return true, nil
		}
	}
	return false, fmt.Errorf("unknown log level: %s", level)
}

// isLevelAtLeast проверяет, что уровень лога не ниже minLevel.
func isLevelAtLeast(current, min string) bool {
	levelsOrder := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

	idxCurrent, idxMin := -1, -1
	for i, l := range levelsOrder {
		if l == current {
			idxCurrent = i
		}
		if l == min {
			idxMin = i
		}
	}

	// Если один из уровней не найден, это уже ошибка на предыдущем шаге,
	// но на всякий случай возвращаем false.
	if idxCurrent == -1 || idxMin == -1 {
		return false
	}

	return idxCurrent >= idxMin
}

// parseTimeInLocation парсит время в заданной локации.
func parseTimeInLocation(timeStr string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:00", timeStr, loc)
}

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
