package linandbin
import "fmt"
import "sort"

func linearSearch(arr []int, x int) bool{
	for _,a :=range arr{
		if a==x{
			return true
		}
	}
	return false
}
func binarySearch(arr []int, x int) bool{
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == x {
			return true
		}
		if arr[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func minmaxSearch(arr []int) (int, int) {

	max := arr[0]
	min := arr[0]
	for _, a := range arr {
		if a > max {
			max = a
		}
		if a < min {
			min = a
		}
	}
	return min, max
}


func main() {
  arr := []int{45, 55, 23, 4, 56, 90, 22}
  srch := 4
	// прямой поиск
	if linearSearch(arr, srch) {
		println("Искомое найдено!")
	} else {
		println("Искомое НЕ найдено!")
	}
	
	// бинарный поиск, где массив должен быть отсортирован
	sort.Ints(arr)
	
	if binarySearch(arr, srch) {
		println("Искомое найдено!")
	} else {
		println("Искомое НЕ найдено!")
	}
	
	// поиск имнимального и максимального значения
	
	arr2 := []int{45, 55, 203, 47, 56, 90, 56}
	mn, mx := minmaxSearch(arr2)
	fmt.Println("мин и макс:", mn, " ", mx)
  
}
