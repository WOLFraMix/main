package main

import "fmt"

// TreeNode представляет узел бинарного дерева
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree рекурсивно строит дерево из массива
func buildTree(arr []int, i int) *TreeNode {
	// Если индекс вышел за границы массива, возвращаем nil
	if i >= len(arr) {
		return nil
	}

	// Создаем текущий узел
	root := &TreeNode{Val: arr[i]}

	// Рекурсивно строим левое и правое поддеревья
	root.Left = buildTree(arr, 2*i+1)
	root.Right = buildTree(arr, 2*i+2)

	return root
}

// printInOrder выводит дерево в порядке in-order обхода (для наглядности структуры)
func printInOrder(node *TreeNode) {
	if node == nil {
		return
	}
	printInOrder(node.Left)
	fmt.Printf("%d ", node.Val)
	printInOrder(node.Right)
}

func main() {
	// Пример массива, представляющего полное бинарное дерево:
	//       1
	//      / \
	//     2   3
	//    / \ / \
	//   4  5 6  7
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("Исходный массив:", arr)

	// Строим дерево, начиная с индекса 0
	root := buildTree(arr, 0)

	fmt.Print("In-order обход построенного дерева: ")
	printInOrder(root)
	fmt.Println()

	// Пример с неполным деревом (пропущенные элементы обозначены как -1, если нужно сохранить структуру,
	// но в данной реализации nil ставится автоматически, если индекс выходит за границы)
	// Для демонстрации возьмем массив меньшей длины, чтобы некоторые узлы были nil
	arr2 := []int{10, 20, 30, 40}
	fmt.Println("\nИсходный массив:", arr2)
	root2 := buildTree(arr2, 0)
	fmt.Print("In-order обход второго дерева: ")
	printInOrder(root2)
	fmt.Println()
}
