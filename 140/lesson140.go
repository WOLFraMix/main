package main

import "fmt"

// ListNode представляет узел односвязного списка.
type ListNode struct {
	Val  int       // Значение узла (int)
	Next *ListNode // Ссылка на следующий узел
}

// removeElements удаляет из списка все узлы, значение которых равно val.
func removeElements(head *ListNode, val int) *ListNode {
	// Создаем фиктивный узел, чтобы упростить удаление головы списка
	dummy := &ListNode{Next: head}
	cur := dummy

	// Итерируемся, проверяя следующий узел (cur.Next)
	for cur.Next != nil {
		if cur.Next.Val == val {
			// Пропускаем узел со значением val
			cur.Next = cur.Next.Next
			// Не двигаем cur, чтобы проверить новый cur.Next на следующей итерации
		} else {
			// Если значение не совпадает, двигаем cur вперед
			cur = cur.Next
		}
	}

	return dummy.Next
}

// NewList создает односвязный список из среза целых чисел.
func NewList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	// берем первый элемент среза (vals), а не весь срез
	head := &ListNode{Val: vals[0]}
	cur := head

	// Проходим по остальным элементам среза
	for _, v := range vals[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return head
}

// Slice конвертирует односвязный список в срез для удобной печати.
func Slice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func main() {
	// Тест 1: удаление элементов из середины и конца
	head := NewList([]int{1, 2, 6, 3, 4, 5, 6})
	fmt.Println("До:", Slice(head))

	head = removeElements(head, 6)
	fmt.Println("После удаления 6:", Slice(head))

	// Тест 2: удаление всех элементов
	head = NewList([]int{7, 7, 7, 7})
	head = removeElements(head, 7)
	fmt.Println("После удаления всех 7:", Slice(head))

	// Тест 3: пустой список
	head = NewList([]int{})
	head = removeElements(head, 1)
	fmt.Println("Пустой список после удаления:", Slice(head))
}
