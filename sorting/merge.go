// Вот пример кода на Go для сортировки слиянием:
package main

import (
"fmt"
)

func main() {
  arr := []int{9, 7, 5, 3, 1, 2, 4, 6, 8}
	sortedArr := mergeSort(arr)
	fmt.Println(sortedArr)

 }
// Функция сортировки слиянием
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

// Функция для слияния двух отсортированных массивов
func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// Добавляем оставшиеся элементы
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}
