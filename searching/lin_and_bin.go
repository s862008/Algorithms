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

func main() {
  arr := []int{45, 55, 23, 4, 56, 90, 22}
  srch := 4

	if linearSearch(arr, srch) {
		println("Искомое найдено!")
	} else {
		println("Искомое НЕ найдено!")
	}
	
	
	sort.Ints(arr)
	
	if binarySearch(arr, srch) {
		println("Искомое найдено!")
	} else {
		println("Искомое НЕ найдено!")
	}
  
}
