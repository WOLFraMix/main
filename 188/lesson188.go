package main

import (
	"fmt"
)

type color int

const (
	red color = iota
	black
)

type node struct {
	key                 int
	color               color
	left, right, parent *node
}

type RedBlackTree struct {
	root *node
	size int
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

// --- Вспомогательные методы ---

func (t *RedBlackTree) insert(key int) {
	newNode := &node{key: key, color: red}
	var y *node
	x := t.root

	// Обычная вставка в BST
	for x != nil {
		y = x
		if newNode.key < x.key {
			x = x.left
		} else if newNode.key > x.key {
			x = x.right
		} else {
			// Дубликаты не добавляем (можно изменить логику при необходимости)
			return
		}
	}

	newNode.parent = y
	if y == nil {
		t.root = newNode
	} else if newNode.key < y.key {
		y.left = newNode
	} else {
		y.right = newNode
	}

	t.size++
	t.fixInsert(newNode)
}

// --- Вращения ---

func (t *RedBlackTree) rotateLeft(x *node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RedBlackTree) rotateRight(y *node) {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.right = y
	y.parent = x
}

// --- Восстановление после вставки (fixInsert) ---

func (t *RedBlackTree) fixInsert(z *node) {
	for z.parent != nil && z.parent.color == red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right // дядя
			if y != nil && y.color == red {
				// Случай 1: дядя красный — перекрашиваем
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					// Случай 2: z — правый ребёнок
					z = z.parent
					t.rotateLeft(z)
				}
				// Случай 3: z — левый ребёнок, делаем правое вращение
				z.parent.color = black
				z.parent.parent.color = red
				t.rotateRight(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left // дядя (симметричный случай)
			if y != nil && y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rotateRight(z)
				}
				z.parent.color = black
				z.parent.parent.color = red
				t.rotateLeft(z.parent.parent)
			}
		}
	}
	t.root.color = black // Корень всегда чёрный
}

// --- Поиск и обход ---

func (t *RedBlackTree) find(key int) bool {
	x := t.root
	for x != nil {
		if key < x.key {
			x = x.left
		} else if key > x.key {
			x = x.right
		} else {
			return true
		}
	}
	return false
}

func (t *RedBlackTree) inorder(n *node, out *[]int) {
	if n == nil {
		return
	}
	t.inorder(n.left, out)
	*out = append(*out, n.key)
	t.inorder(n.right, out)
}

func (t *RedBlackTree) Inorder() []int {
	var res []int
	t.inorder(t.root, &res)
	return res
}

func (t *RedBlackTree) Size() int {
	return t.size
}

func main() {
	tree := NewRedBlackTree()

	// Вставка нескольких значений
	vals := []int{10, 20, 5, 6, 12, 30, 7}
	for _, v := range vals {
		tree.insert(v)
	}

	fmt.Println("In order:", tree.Inorder()) // Отсортированный обход
	fmt.Println("Size:", tree.Size())        // Количество элементов
	fmt.Println("Find 12:", tree.find(12))   // true
	fmt.Println("Find 99:", tree.find(99))   // false
}
