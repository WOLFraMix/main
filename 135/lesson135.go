package main

import "fmt"

// структура узла односвязного списка
type ListNode struct {
	Val  int
	Next *ListNode
}

// алгоритм разворота односвязного списка
func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		next := current.Next // 1. Сохраняем ссылку на следующий узел
		current.Next = prev  // 2. Разворачиваем указатель текущего узла
		prev = current       // 3. Сдвигаем prev на текущий узел
		current = next       // 4. Переходим к следующему узлу
	}

	return prev // новый head списка
}

// printList — вспомогательная функция для вывода списка
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func main() {
	// Создаем список: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("Исходный список:")
	printList(head)

	// Вызываем функцию разворота
	head = reverseLinkedList(head)

	fmt.Println("Развернутый список:")
	printList(head)
}
