// Алгоритм «Решето Эратосфена» — эффективный алгоритм для нахождения всех простых чисел до заданного предела n. 
// Работает по принципу «отсеивания» составных чисел.

package main

import (
	"fmt"
)

// SieveOfEratosthenes находит все простые числа до n (включительно)
func SieveOfEratosthenes(n int) []int {
	if n < 2 {
		return []int{} // Нет простых чисел меньше 2
	}

	// Создаём булевый массив, изначально все числа считаются простыми
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Решето Эратосфена
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			// Отмечаем все кратные i как составные
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Собираем простые числа в срез
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	n := 30
	primes := SieveOfEratosthenes(n)
	fmt.Printf("Простые числа до %d:\n", n)
	fmt.Println(primes)
}
