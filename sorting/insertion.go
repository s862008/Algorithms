package main
import 	"fmt"

func insrtSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}
func main () {
  
  arr = []int{45, 55, 23, 4, 56, 90, 22}
  
	fmt.Println("исходный массив:", arr)
	fmt.Println("отсортированный массив:", insrtSort(arr))
  
}
