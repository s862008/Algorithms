/*
В этом примере реализуется хэш-таблица на основе метода цепочек для разрешения коллизий. 
Структура `ListNode` представляет собой узел связанного списка, где каждый узел является элементом цепочки в массиве хэш-таблицы. 
Структура `HashTable` содержит массив указателей на узлы связанного списка и размер хэш
*/
package main

import "fmt"

// Узел списка цепочек
type ListNode struct {
Key string
Value interface{}
Next *ListNode
}

// Хеш-таблица
type HashTable struct {
Table []*ListNode
Size int
}

// Создание новой хэш-таблицы
func NewHashTable(size int) *HashTable {
return &HashTable{
Table: make([]*ListNode, size),
Size: size,
}
}

// Хэш-функция
func (ht *HashTable) hashFunction(key string) int {
hash := 0
for _, c := range key {
hash += int(c)
}
return hash % ht.Size
}

// Добавление значения в хэш-таблицу
func (ht *HashTable) Put(key string, value interface{}) {
index := ht.hashFunction(key)
newNode := &ListNode{
Key: key,
Value: value,
}
if ht.Table[index] == nil {
ht.Table[index] = newNode
} else {
current := ht.Table[index]
for current.Next != nil {
if current.Key == key {
// Если ключ уже существует, обновляем значение
current.Value = value
return
}
current = current.Next
}
if current.Key == key {
current.Value = value
} else {
current.Next = newNode
}
}
}

// Получение значения из хэш-таблицы по ключу
func (ht *HashTable) Get(key string) (interface{}, bool) {
index := ht.hashFunction(key)
current := ht.Table[index]
for current != nil {
if current.Key == key {
return current.Value, true
}
current = current.Next
}
return nil, false
}

func main() {
ht := NewHashTable(10)

ht.Put("key1", 10)
ht.Put("key2", "Hello")
ht.Put("key3", true)

val, found := ht.Get("key1")
if found {
fmt.Println("Value:", val)
} else {
fmt.Println("Key not found")
}
}
