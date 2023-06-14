/*
Реализация множеств на языке Go может быть выполнена несколькими способами, 
но одним из наиболее распространенных способов является использование карты (map) 
для хранения элементов множества.
*/

package main

import "fmt"

type Set struct {
data map[interface{}]bool
}

//Создание нового множества.
func NewSet() *Set {
return &Set{make(map[interface{}]bool)}
}

//Добавление элемента в множество.
func (s *Set) Add(value interface{}) {
s.data[value] = true
}

//Удаление элемента из множества.
func (s *Set) Remove(value interface{}) {
delete(s.data, value)
}

//Проверка наличия элемента в множестве.
func (s *Set) Contains(value interface{}) bool {
_, inSet := s.data[value]
return inSet
}

//Возвращает все элементы множества.
func (s *Set) GetAll() []interface{} {
result := make([]interface{}, len(s.data))
i := 0
for value := range s.data {
result[i] = value
i++
}
return result
}

//Возвращает количество элементов в множестве.
func (s *Set) Size() int {
return len(s.data)
}

//Проверка, является ли множество пустым.
func (s *Set) IsEmpty() bool {
return s.Size() == 0
}

func main() {
set := NewSet()

// Добавление элементов в множество.
set.Add("a")
set.Add("b")
set.Add("c")

// Проверка наличия элемента в множестве.
containsA := set.Contains("a")
fmt.Println(containsA)

// Удаление элемента и проверка его отсутствия в множестве.
set.Remove("b")
containsB := set.Contains("b")
fmt.Println(containsB)

// Получение всех элементов множества и их количество.
allValues := set.GetAll()
fmt.Println(allValues)
size := set.Size()
fmt.Println(size)
}
