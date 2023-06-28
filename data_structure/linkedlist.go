/*
Вот пример реализации простого связного списка на языке программирования Go:
В этом примере определены две структуры: Node (узел списка) и LinkedList (сам связный список). Метод Append добавляет новый элемент в конец списка, 
а метод Display выводит все элементы списка.
В функции main создается экземпляр связного списка и добавляются несколько элементов. Затем вызывается метод Display для вывода элементов списка. 
Результатом выполнения программы будет вывод списка: `1 2 3 4 5`.
Вы можете добавлять и удалять элементы из списка, обращаться к элементам по индексу и выполнять другие операции в зависимости от ваших потребностей.
*/
package main

import "fmt"

// Определение структуры узла списка
type Node struct {
data interface{}
next *Node
}

// Определение структуры связного списка
type LinkedList struct {
head *Node
}

// Добавление элемента в конец списка
func (list *LinkedList) Append(data interface{}) {
newNode := &Node{data: data, next: nil}

  if list.head == nil {
  list.head = newNode
  } else {
  current := list.head
    for current.next != nil {
    current = current.next
    }
  current.next = newNode
  }
}

// Вывод элементов списка
func (list *LinkedList) Display() {
current := list.head
  for current != nil {
  fmt.Printf("%v ", current.data)
  current = current.next
  }
fmt.Println()
}

func main() {
list := LinkedList{}

// Добавление элементов в список
list.Append(1)
list.Append(2)
list.Append(3)
list.Append(4)
list.Append(5)

// Вывод элементов списка
list.Display()
}
