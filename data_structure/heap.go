//Вот пример реализации структуры данных "куча" (heap) на языке Go:

package main

import (
"fmt"
)

type Heap struct {
arr []int
}

func NewHeap() *Heap {
return &Heap{
arr: []int{},
}
}

func (h *Heap) Insert(val int) {
h.arr = append(h.arr, val)
h.heapifyUp(len(h.arr) - 1)
}

func (h *Heap) Delete() int {
if len(h.arr) == 0 {
return -1
}

root := h.arr[0]
lastIndex := len(h.arr) - 1
h.arr[0] = h.arr[lastIndex]
h.arr = h.arr[:lastIndex]

h.heapifyDown(0)

return root
}

func (h *Heap) heapifyUp(index int) {
parentIndex := (index - 1) / 2

if index > 0 && h.arr[parentIndex] < h.arr[index] {
h.arr[parentIndex], h.arr[index] = h.arr[index], h.arr[parentIndex]
h.heapifyUp(parentIndex)
}
}

func (h *Heap) heapifyDown(index int) {
leftChildIndex := 2*index + 1
rightChildIndex := 2*index + 2
largestIndex := index

if leftChildIndex < len(h.arr) && h.arr[leftChildIndex] > h.arr[largestIndex] {
largestIndex = leftChildIndex
}

if rightChildIndex < len(h.arr) && h.arr[rightChildIndex] > h.arr[largestIndex] {
largestIndex = rightChildIndex
}

if largestIndex != index {
h.arr[index], h.arr[largestIndex] = h.arr[largestIndex], h.arr[index]
h.heapifyDown(largestIndex)
}
}

func main() {
heap := NewHeap()

heap.Insert(10)
heap.Insert(30)
heap.Insert(20)
heap.Insert(50)
heap.Insert(40)

fmt.Println("Max Heap:")
for len(heap.arr) > 0 {
fmt.Println(heap.Delete())
}
}
/*
В данном примере представлена реализация структуры данных "куча" (heap) в виде бинарной кучи (max-heap).
Структура `Heap` содержит срез `arr`, который представляет собой массив элементов кучи.
Метод `NewHeap` создает новую кучу.
Метод `Insert` вставляет новый элемент в кучу и выполняет перестройку кучи снизу вверх.
Метод `Delete` удаляет и возвращает корневой элемент (максимальный элемент в случае max-heap) из кучи и выполняет перестройку кучи сверху вниз.
Методы `heapifyUp` и `heapifyDown` выполняют перестройку кучи в соответствии с ее основными свойствами.
В функции `main` создается новая куча, в нее вставляются элементы, а затем по очереди удаляются и выводятся на экран с помощью метода `Delete`.
В результате выполнения программы вы увидите вывод, представляющий собой элементы кучи в порядке убывания: `50 40 30 20 10`.
*/
