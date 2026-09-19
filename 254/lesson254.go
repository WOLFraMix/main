package main

import (
	"fmt"
	"os"
)

// MyBuffer - кольцевой буфер фиксированного размера.
type MyBuffer struct {
	data     []int // содержимое
	capacity int   // вместимость
	head     int   // индекс самого старого элемента
	tail     int   // индекс, куда запишем следующий элемент
	size     int   // текущее количество элементов
}

// NewMyBuffer создает новый буфер заданной емкости.
// При capacity <= 0 паникует - т.к. это уже ошибка программиста.
func NewMyBuffer(capacity int) *MyBuffer {
	if capacity <= 0 {
		panic(fmt.Sprintf("capacity must be positive, got %d", capacity))
	}
	return &MyBuffer{
		data:     make([]int, capacity),
		capacity: capacity,
	}
}

// Push записывает значение.
// Если буфер полон - перезаписывает самый старый элемент.
func (mb *MyBuffer) Push(val int) {
	// Записываем значение в текущую позицию tail
	mb.data[mb.tail] = val

	// Двигаем tail вперёд по кругу
	mb.tail = (mb.tail + 1) % mb.capacity

	// Если буфер был полон, сдвигаем head
	// что само перезаписывает самый старый элемент
	if mb.size == mb.capacity {
		mb.head = (mb.head + 1) % mb.capacity
		// size остаётся равным capacity
		// (буфер по-прежнему полон)
		return
	}

	// Иначе просто увеличиваем размер
	mb.size++
}

// Pop забирает самый старый элемент.
// Если буфер пуст - возвращает (0, false).
func (mb *MyBuffer) Pop() (int, bool) {
	if mb.size == 0 {
		return 0, false
	}

	// Берем значение из head
	val := mb.data[mb.head]

	// Сдвигаем head вперёд по кругу
	// и теперь самый старый элемент - следующий
	mb.head = (mb.head + 1) % mb.capacity

	// Уменьшаем размер
	mb.size--

	return val, true
}

// Iterate проходит от самого старого к самому новому
// и вызывает callback для каждого.
// Физические индексы в слайсе «зациклены»,
// поэтому нужно выдать логические — от 0 до size-1 по очереди.
func (mb *MyBuffer) Iterate(callback func(int, int)) {
	for i := 0; i < mb.size; i++ {
		// Вычисляем физический индекс в слайсе с учётом кольца
		idx := (mb.head + i) % mb.capacity
		// Передаем логический индекс i и значение элемента
		callback(i, mb.data[idx])
	}
}

// Функция main и все тесты будут скрыты от вас при проверке на сайте.
func main() {
	if !test1() || !test2() || !test3() || !test4() || !test5() ||
		!test6() || !test7() || !test8() || !test9() || !test10() {
		os.Exit(1)
	}
	fmt.Println("Все тесты успешно пройдены!")
}

// sliceEq простая функция для сравнения слайсов
func sliceEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// collectIterate собирает значения и логические индексы из Iterate
func collectIterate(mb *MyBuffer) (vals []int, idxs []int) {
	mb.Iterate(func(i, v int) {
		idxs = append(idxs, i)
		vals = append(vals, v)
	})
	return
}

// expectPanic запускает f и возвращает true, если та паниковала
func expectPanic(f func()) bool {
	panicked := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		f()
	}()
	return panicked
}

// Базовый вариант, когда заполнили, обошли, вычитали все.
func test1() bool {
	buf := NewMyBuffer(3)
	buf.Push(10)
	buf.Push(20)
	buf.Push(30)

	vals, idxs := collectIterate(buf)
	if !sliceEq(vals, []int{10, 20, 30}) {
		fmt.Fprintf(os.Stderr, "Тест 1: значения после 3 Push: %v, ожидали [10 20 30]\n", vals)
		return false
	}
	if !sliceEq(idxs, []int{0, 1, 2}) {
		fmt.Fprintf(os.Stderr, "Тест 1: логические индексы: %v, ожидали [0 1 2]\n", idxs)
		return false
	}

	v, ok := buf.Pop()
	if !ok || v != 10 {
		fmt.Fprintf(os.Stderr, "Тест 1: первый Pop: (%d, %v), ожидали (10, true)\n", v, ok)
		return false
	}
	v, ok = buf.Pop()
	if !ok || v != 20 {
		fmt.Fprintf(os.Stderr, "Тест 1: второй Pop: (%d, %v), ожидали (20, true)\n", v, ok)
		return false
	}
	v, ok = buf.Pop()
	if !ok || v != 30 {
		fmt.Fprintf(os.Stderr, "Тест 1: третий Pop: (%d, %v), ожидали (30, true)\n", v, ok)
		return false
	}
	v, ok = buf.Pop()
	if ok || v != 0 {
		fmt.Fprintf(os.Stderr, "Тест 1: Pop из пустого: (%d, %v), ожидали (0, false)\n", v, ok)
		return false
	}
	return true
}

// Перезапись при переполнении - кейс из условия
func test2() bool {
	buf := NewMyBuffer(3)
	buf.Push(10)
	buf.Push(20)
	buf.Push(30)

	// Освобождаем одну ячейку
	v, _ := buf.Pop()
	if v != 10 {
		fmt.Fprintf(os.Stderr, "Тест 2: Pop вернул %d, ожидали 10\n", v)
		return false
	}

	buf.Push(40) // добавится в освободившуюся ячейку
	buf.Push(50) // буфер снова полон - 50 затирает 20

	vals, _ := collectIterate(buf)
	// Порядок от старого к новому: 30, 40, 50
	if !sliceEq(vals, []int{30, 40, 50}) {
		fmt.Fprintf(os.Stderr, "Тест 2: после Push(40), Push(50): %v, ожидали [30 40 50]\n", vals)
		return false
	}
	return true
}

// Pop из пустого буфера - должен вернуть zero-value и false
func test3() bool {
	buf := NewMyBuffer(5)
	v, ok := buf.Pop()
	if ok {
		fmt.Fprintf(os.Stderr, "Тест 3: Pop из пустого вернул ok=true, должен false\n")
		return false
	}
	if v != 0 {
		fmt.Fprintf(os.Stderr, "Тест 3: zero-value должен быть 0, получили %d\n", v)
		return false
	}

	// Вычитаем все до дна, потом еще раз пробуем
	buf.Push(100)
	buf.Push(200)
	buf.Pop()
	buf.Pop()
	v, ok = buf.Pop()
	if ok || v != 0 {
		fmt.Fprintf(os.Stderr, "Тест 3: после исчерпания элементов Pop: (%d, %v), ожидали (0, false)\n", v, ok)
		return false
	}
	return true
}

// Емкость 1
func test4() bool {
	buf := NewMyBuffer(1)

	v, ok := buf.Pop()
	if ok {
		fmt.Fprintf(os.Stderr, "Тест 4: Pop из пустого буфера емкости 1 должен вернуть false\n")
		return false
	}

	buf.Push(7)
	vals, _ := collectIterate(buf)
	if !sliceEq(vals, []int{7}) {
		fmt.Fprintf(os.Stderr, "Тест 4: после Push(7): %v, ожидали [7]\n", vals)
		return false
	}

	v, ok = buf.Pop()
	if !ok || v != 7 {
		fmt.Fprintf(os.Stderr, "Тест 4: Pop: (%d, %v), ожидали (7, true)\n", v, ok)
		return false
	}

	// Каждый новый Push должен затирать единственный слот
	buf.Push(1)
	buf.Push(2)
	buf.Push(3)
	v, ok = buf.Pop()
	if !ok || v != 3 {
		fmt.Fprintf(os.Stderr, "Тест 4: после тройного Push в емкость 1 Pop: (%d, %v), ожидали (3, true)\n", v, ok)
		return false
	}
	return true
}

// Конструктор паникует на невалидной емкости, проверим что не меняли поведение
func test5() bool {
	if !expectPanic(func() { _ = NewMyBuffer(0) }) {
		fmt.Fprintf(os.Stderr, "Тест 5: NewMyBuffer(0) должен паниковать\n")
		return false
	}
	if !expectPanic(func() { _ = NewMyBuffer(-1) }) {
		fmt.Fprintf(os.Stderr, "Тест 5: NewMyBuffer(-1) должен паниковать\n")
		return false
	}
	if !expectPanic(func() { _ = NewMyBuffer(-100) }) {
		fmt.Fprintf(os.Stderr, "Тест 5: NewMyBuffer(-100) должен паниковать\n")
		return false
	}
	return true
}

// Много кругов, должно работать стабильно
func test6() bool {
	buf := NewMyBuffer(3)
	for i := 1; i <= 10; i++ {
		buf.Push(i)
	}
	// После 10 Push в буфер 3 остаются последние три: 8, 9, 10
	vals, _ := collectIterate(buf)
	if !sliceEq(vals, []int{8, 9, 10}) {
		fmt.Fprintf(os.Stderr, "Тест 6: после 10 Push: %v, ожидали [8 9 10]\n", vals)
		return false
	}

	v, _ := buf.Pop()
	if v != 8 {
		fmt.Fprintf(os.Stderr, "Тест 6: Pop вернул %d, ожидали 8\n", v)
		return false
	}
	buf.Push(11)
	vals, _ = collectIterate(buf)
	if !sliceEq(vals, []int{9, 10, 11}) {
		fmt.Fprintf(os.Stderr, "Тест 6: после Pop(8)+Push(11): %v, ожидали [9 10 11]\n", vals)
		return false
	}
	return true
}

// Логический индекс не должен зависеть от физического положения
func test7() bool {
	buf := NewMyBuffer(4)
	buf.Push(100)
	buf.Push(200)
	buf.Push(300)
	// Имитируем сдвиг головы: два Pop, потом два Push
	buf.Pop()
	buf.Pop()
	buf.Push(400)
	buf.Push(500)

	// Порядок от старого к новому: 300, 400, 500, индексы 0, 1, 2
	vals, idxs := collectIterate(buf)
	if !sliceEq(vals, []int{300, 400, 500}) {
		fmt.Fprintf(os.Stderr, "Тест 7: значения: %v, ожидали [300 400 500]\n", vals)
		return false
	}
	if !sliceEq(idxs, []int{0, 1, 2}) {
		fmt.Fprintf(os.Stderr, "Тест 7: индексы: %v, ожидали [0 1 2] - логический индекс не должен зависеть от физического положения\n", idxs)
		return false
	}
	return true
}

// Много операций, проверяем консистентность состояния
func test8() bool {
	buf := NewMyBuffer(100)
	for i := range 100 {
		buf.Push(i)
	}
	vals, _ := collectIterate(buf)
	if len(vals) != 100 || vals[0] != 0 || vals[99] != 99 {
		fmt.Fprintf(os.Stderr, "Тест 8: после 100 Push в буфер 100: len=%d first=%d last=%d\n", len(vals), vals[0], vals[99])
		return false
	}

	// Перезаписываем все новыми значениями
	for i := 100; i < 200; i++ {
		buf.Push(i)
	}
	vals, _ = collectIterate(buf)
	if len(vals) != 100 || vals[0] != 100 || vals[99] != 199 {
		fmt.Fprintf(os.Stderr, "Тест 8: после перезаписи: len=%d first=%d last=%d, ожидали first=100 last=199\n", len(vals), vals[0], vals[99])
		return false
	}

	// Вычитаем все
	for i := range 100 {
		v, ok := buf.Pop()
		if !ok || v != 100+i {
			fmt.Fprintf(os.Stderr, "Тест 8: Pop #%d: (%d, %v), ожидали (%d, true)\n", i, v, ok, 100+i)
			return false
		}
	}
	_, ok := buf.Pop()
	if ok {
		fmt.Fprintf(os.Stderr, "Тест 8: после 100 Pop буфер должен быть пуст\n")
		return false
	}
	return true
}

// Iterate по пустому буферу не должен вызывать callback
func test9() bool {
	buf := NewMyBuffer(5)
	called := false
	buf.Iterate(func(i, v int) {
		called = true
	})
	if called {
		fmt.Fprintf(os.Stderr, "Тест 9: Iterate по пустому буферу не должен вызывать callback\n")
		return false
	}

	// И после полного опустошения тоже
	buf.Push(1)
	buf.Push(2)
	buf.Pop()
	buf.Pop()
	called = false
	buf.Iterate(func(i, v int) {
		called = true
	})
	if called {
		fmt.Fprintf(os.Stderr, "Тест 9: Iterate после исчерпания элементов не должен вызывать callback\n")
		return false
	}
	return true
}

// Чередуем Push/Pop, порядок должен оставаться FIFO
func test10() bool {
	buf := NewMyBuffer(5)
	expected := []int{}
	for i := 1; i <= 20; i++ {
		buf.Push(i)
		expected = append(expected, i)
		if len(expected) > 5 {
			expected = expected[1:]
		}
		// Каждые 3 итерации достаем один
		if i%3 == 0 {
			v, ok := buf.Pop()
			if !ok || v != expected[0] {
				fmt.Fprintf(os.Stderr, "Тест 10: на шаге %d Pop вернул (%d, %v), ожидали %d\n", i, v, ok, expected[0])
				return false
			}
			expected = expected[1:]
		}
		// Проверяем текущее состояние
		vals, _ := collectIterate(buf)
		if !sliceEq(vals, expected) {
			fmt.Fprintf(os.Stderr, "Тест 10: на шаге %d буфер %v, ожидали %v\n", i, vals, expected)
			return false
		}
	}
	return true
}
