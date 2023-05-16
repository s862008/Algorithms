package main

import "fmt"

func selectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		minindex := i
		for j := i + 1; j < n; j++ {
			if arr[minindex] > arr[j] {
				minindex = j
			}
		}
		arr[i], arr[minindex] = arr[minindex], arr[i]
  }
	return arr
}
func main() {
  //сортировка выбором
	arr = []int{45, 55, 23, 4, 56, 90, 22}
  fmt.Println("исходный массив:", arr)
	fmt.Println("отсортированный  выбором:", selectionSort(arr))
}
