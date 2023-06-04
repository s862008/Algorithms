/*
Этот код создает структуру Stack, которая содержит указатель на верхний элемент и размер стека.
Элементы стека реализованы в виде структуры Element, которая содержит значение элемента и указатель на следующий элемент.
Метод Push добавляет новый элемент на вершину стека,
метод Pop удаляет и возвращает последний элемент, метод Size возвращает текущее количество элементов в стеке, 
а метод Peek возвращает элемент на вершине стека без удаления.
*/

package main

import "fmt"

type Stack struct {
top *Element
size int
}

type Element struct {
value interface{}
next *Element
}

//Добавление элемента в стек.
func (s *Stack) Push(value interface{}) {
s.top = &Element{value, s.top}
s.size++
}

//Удаление и возвращение верхнего элемента стека.
func (s *Stack) Pop() interface{} {
if s.size == 0 {
return nil
}
value := s.top.value
s.top = s.top.next
s.size--
return value
}

//Возвращает текущее количество элементов в стеке.
func (s *Stack) Size() int {
return s.size
}

//Возвращает верхний элемент стека без удаления.
func (s *Stack) Peek() interface{} {
if s.Size() == 0 {
return nil
}
return s.top.value
}

func main(){
 stack := &Stack{}
    
    // Добавление элементов в стек.
    stack.Push("a")
    stack.Push("b")
    stack.Push("c")
    
    // Удаление верхнего элемента стека и печать его значения.
    top := stack.Pop()
    fmt.Println(top)

    // Получение элемента на вершине стека без его удаления.
    peek := stack.Peek()
    fmt.Println(peek)

    // Получение текущего размера стека.
    size := stack.Size()
    fmt.Println(size)


}
