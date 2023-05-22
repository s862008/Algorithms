/*
Здесь функция `heapSort` принимает слайс (срез) целых чисел `arr` и сортирует его в порядке возрастания с помощью сортировки кучей. 
Функция `heapify` служит для построения кучи и превращения поддерева в кучу максимального элемента. 
Функция `main` вызывает функцию `heapSort` на примере неотсортированного слайса и выводит исходный и 
отсортированный массивы в консоль.
*/

package main
import "fmt"

func heapSort(arr []int) {
    n := len(arr)
    
    // Построение кучи (Heapify)
    for i := n/2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }
    
    // Извлечение элементов по одному из кучи
    for i := n - 1; i > 0; i-- {
        // Перемещаем текущий корень в конец
        arr[0], arr[i] = arr[i], arr[0]
        
        // Вызываем heapify на уменьшенной куче
        heapify(arr, i, 0)
    }
}

// Функция для превращения поддерева в кучу максимального элемента (Heapify)
func heapify(arr []int, n int, i int) {
    largest := i  // Инициализируем наибольший элемент как корень поддерева
    left := 2*i + 1  // Левый потомок
    right := 2*i + 2  // Правый потомок

    // Если левый дочерний элемент больше корня
    if left < n && arr[left] > arr[largest] {
        largest = left
    }
    
    // Если правый дочерний элемент больше, чем наибольший элемент на данный момент
    if right < n && arr[right] > arr[largest] {
        largest = right
    }
    
    // Если наибольший элемент не корень
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]  // Обмен местами
        
        // Применяем Heapify к поддереву.
        heapify(arr, n, largest)
    }
}

func main() {
 
    arr := []int{12, 11, 13, 5, 6, 7}
    fmt.Println("Исходный массив: ")
    fmt.Println(arr)

    heapSort(arr)
    fmt.Println("Отсортированный массив: ")
    fmt.Println(arr)

}
