package main

import "fmt"

// Создадим структуру нашего дерева
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

// Метод всавки элемента
func (t *BinaryTree) Insert(value int) {
	if t.Root == nil {
		t.Root = &TreeNode{Val: value, Left: nil, Right: nil}
	} else {
		InsertRecursive(t.Root, value)
	}
}

// Функция вставка элемента
func InsertRecursive(t *TreeNode, value int) error {
	if t.Val > value {
		if t.Left == nil {
			t.Left = &TreeNode{Val: value, Left: nil, Right: nil}
			return nil
		}
		return InsertRecursive(t.Left, value)
	}

	if t.Val < value {
		if t.Right == nil {
			t.Right = &TreeNode{Val: value, Left: nil, Right: nil}
			return nil
		}
		return InsertRecursive(t.Right, value)
	}
	return nil
}

// Метод поиска элемента
func (t *BinaryTree) FindValue(value int) bool {
	if t.Root == nil {
		return false
	} else {
		return FindValueRecursive(t.Root, value)
	}
}

// Функция поиска элемента
func FindValueRecursive(t *TreeNode, value int) bool {
	if t.Val > value {
		if t.Left == nil {
			return false
		} else {
			return FindValueRecursive(t.Left, value)
		}
	}
	if t.Val < value {
		if t.Right == nil {
			return false
		} else {
			return FindValueRecursive(t.Right, value)
		}
	} else {
		return true
	}
}

// Метод удаления элемента
func (t *BinaryTree) Delete(value int) {
	t.Root = DeleteValue(t.Root, value)
}

// Функция удаления элемента
func DeleteValue(t *TreeNode, value int) *TreeNode {
	if t == nil {
		return t
	}

	if value < t.Val {
		t.Left = DeleteValue(t.Left, value)
	} else if value > t.Val {
		t.Right = DeleteValue(t.Right, value)
	} else {
		if t.Left == nil {
			return t.Right
		} else if t.Right == nil {
			return t.Left //1
		}

		t.Val = minVlaue(t.Right) //2
		t.Right = DeleteValue(t.Right, t.Val)

	}
	return t
}

// Функция нахождения самого минимального (самого левого элемента)
func minVlaue(t *TreeNode) int {
	minVlaue := t.Val
	for t.Left != nil {
		minVlaue = t.Left.Val
		t = t.Left
	}
	return minVlaue
}

// Обход дерева
func (r *BinaryTree) PrintInorderMethod() {
	PrintInOrder(r.Root)
}

func PrintInOrder(value *TreeNode) {
	if value != nil {
		PrintInOrder(value.Left)
		fmt.Print(value.Val, " ")
		PrintInOrder(value.Right)
	}
}
