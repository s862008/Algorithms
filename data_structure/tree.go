package main

import (
	"fmt"
)
func main() {
// струкутра данных дерево
	var root *Node

	root = Insert(root, 50)
	root = Insert(root, 30)
	root = Insert(root, 20)
	root = Insert(root, 40)
	root = Insert(root, 70)
	root = Insert(root, 60)
	root = Insert(root, 80)
	
	fmt.Println("Обход дерева в порядке in-order:")
	InOrderTraversal(root)
	fmt.Println()

  
}
// Определяем структуру узла дерева
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Функция для создания нового узла
func NewNode(data int) *Node {
	node := &Node{}
	node.data = data
	node.left = nil
	node.right = nil
	return node
}

// Функция для вставки нового узла в дерево
func Insert(root *Node, data int) *Node {
	if root == nil {
		return NewNode(data)
	} else {
		if data <= root.data {
			root.left = Insert(root.left, data)
		} else {
			root.right = Insert(root.right, data)
		}
		return root
	}
}

// Функция для выполнения обхода дерева в порядке in-order (левый, корень, правый)
func InOrderTraversal(root *Node) {
	if root != nil {
		InOrderTraversal(root.left)
		fmt.Printf("%d ", root.data)
		InOrderTraversal(root.right)
	}
}

/*
В этом примере мы определяем структуру `Node`, которая представляет узел дерева. Каждый узел содержит значение `data` и указатели на левый (`left`) и правый (`right`) подузлы.
Мы также определяем функцию `NewNode`, которая создает новый узел и функцию `Insert`, которая вставляет новый узел в дерево. Функция `Insert` рекурсивно проходит по дереву, сравнивая значение нового узла с текущим узлом и вставляет его в соответствующее поддерево (левое или правое).
Для демонстрации функциональности, мы используем функцию `InOrderTraversal`, которая выполняет обход дерева в порядке "in-order" (левый, корень, правый) и выводит значения узлов на экран.
В функции `main` мы создаем дерево и вставляем в него несколько узлов. Затем мы вызываем функцию `InOrderTraversal` для обхода дерева в порядке in-order и выводим значения узлов на экран.
Это лишь один из множества способов реализации дерева на языке Go.
*/


