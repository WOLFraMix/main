package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Напишите программу, содержащую описание очереди и моделирующую работу очереди.
Возможные команды для программы:
push n
Добавить в очередь число n. Программа должна вывести ok.
pop
Удалить из очереди первый элемент. Программа должна вывести его значение.
front
Программа должна вывести значение первого элемента, не удаляя его из очереди.
size
Программа должна вывести количество элементов в очереди.
clear
Программа должна очистить очередь и вывести ok.
exit
Программа должна вывести bye и завершить работу.
*/

// Queue — структура очереди на базе слайса
type Queue struct {
	data []int
}

// Push добавляет элемент в конец очереди
func (q *Queue) Push(n int) {
	q.data = append(q.data, n)
}

// Pop удаляет и возвращает первый элемент
func (q *Queue) Pop() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("error")
	}
	val := q.data
	q.data = q.data[1:]
	return val[0], nil
}

// Front возвращает первый элемент без удаления
func (q *Queue) Front() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("error")
	}
	return q.data[0], nil
}

// Size возвращает количество элементов в очереди
func (q *Queue) Size() int {
	return len(q.data)
}

// Clear очищает очередь
func (q *Queue) Clear() {
	q.data = q.data[:0]
}

func main() {
	var q Queue
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		cmd := parts[0] // Первое значение это команда

		switch cmd {
		case "push":
			// Ожидаем второй аргумент — число
			if len(parts) < 2 {
				fmt.Println("error")
				continue
			}
			n, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("error")
				continue
			}
			q.Push(n)
			fmt.Println("ok")

		case "pop":
			val, err := q.Pop()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(val)
			}

		case "front":
			val, err := q.Front()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(val)
			}

		case "size":
			fmt.Println(q.Size())

		case "clear":
			q.Clear()
			fmt.Println("ok")

		case "exit":
			fmt.Println("bye")
			return

		default:
			// Неизвестная команда
			fmt.Println("error")
		}
	}
}
