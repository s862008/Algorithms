// Алгоритмы вычисления факториала
package maths

import "fmt"

// Iterative returns the iteratively brute forced factorial of n
func Iterative(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Recursive This function recursively computes the factorial of a number
func Recursive(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Recursive(n-1)
	}
}

func main(){

  n:= 5

  fmt.Println(Iterative(n) )
  fmt.Println(Recursive(n))
}
