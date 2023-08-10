/*
В этом примере мы создаем структуру `Queue`, которая содержит срез элементов. 
Метод `Enqueue` добавляет элемент в конец очереди, а метод `Dequeue` извлекает и возвращает первый элемент из очереди.
В функции `main` мы создаем экземпляр очереди, добавляем несколько элементов и затем извлекаем их с помощью метода `Dequeue`. 
Если очередь пуста, метод `Dequeue` вернет -1 и выведет сообщение "Очередь пуста".
*/

package main

import "fmt"

type Queue struct {
items []int
}

func (q *Queue) Enqueue(item int) {
q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
if len(q.items) == 0 {
fmt.Println("Очередь пуста")
return -1
}

item := q.items[0]
q.items = q.items[1☺
return item
}

func main() {
queue := Queue{}

queue.Enqueue(1)
queue.Enqueue(2)
queue.Enqueue(3)

fmt.Println(queue.Dequeue())
fmt.Println(queue.Dequeue())
fmt.Println(queue.Dequeue())
fmt.Println(queue.Dequeue())
}
