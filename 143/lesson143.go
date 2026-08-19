package main

import (
	"fmt"
	"time"
)

// Stopwatch представляет собой секундомер, фиксирующий промежуточные результаты.
type Stopwatch struct {
	startTime time.Time       // Момент запуска секундомера
	splits    []time.Duration // Список промежуточных длительностей относительно старта
	running   bool            // Флаг, показывающий, запущен ли сейчас отсчет
}

// Start запускает новый отсчет времени или сбрасывает существующий.
// Если секундомер уже запущен, он будет перезапущен с новым временем старта.
func (s *Stopwatch) Start() {
	s.startTime = time.Now()
	s.splits = s.splits[:0] // Очищаем слайс результатов (сброс)
	s.running = true
}

// SaveSplit сохраняет текущее прошедшее время относительно момента старта.
// Если секундомер не запущен, метод ничего не делает (или можно добавить панику/ошибку).
func (s *Stopwatch) SaveSplit() {
	if !s.running {
		return
	}
	elapsed := time.Since(s.startTime)
	s.splits = append(s.splits, elapsed)
}

// GetResults возвращает срез длительностей, где каждый элемент — это время
// от общего старта до момента вызова SaveSplit.
func (s *Stopwatch) GetResults() []time.Duration {
	// Возвращаем копию слайса, чтобы внешний код не мог модифицировать внутреннее состояние
	result := make([]time.Duration, len(s.splits))
	copy(result, s.splits)
	return result
}

func main() {
	sw := Stopwatch{}

	// 1. Запуск секундомера (сброс и старт)
	sw.Start()

	// Имитация работы
	time.Sleep(100 * time.Millisecond)
	sw.SaveSplit() // Первый сплит: ~100 мс

	time.Sleep(250 * time.Millisecond)
	sw.SaveSplit() // Второй сплит: ~350 мс

	time.Sleep(50 * time.Millisecond)
	sw.SaveSplit() // Третий сплит: ~400 мс

	// Получение и вывод результатов
	results := sw.GetResults()
	fmt.Println("Результаты (относительно старта):")
	for i, res := range results {
		fmt.Printf("Сплит %d: %v\n", i+1, res)
	}

	// Пример повторного запуска (сброс)
	sw.Start()
	fmt.Println("\nПосле сброса количество результатов:", len(sw.GetResults()))
}
