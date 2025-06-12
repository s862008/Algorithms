// Проверка числа на простоту означает определение, делится ли данное число ровно на какое-то другое число кроме единицы и самого себя.
// Простые числа имеют только два делителя: единицу и само число.

package main

import (
 "fmt"
 "math"
)

// isPrime проверяет, является ли переданное число простым
func isPrime(num int) bool {
 if num <= 1 {
  return false // Числа меньше 2 не являются простыми
 }
 if num == 2 || num == 3 {
  return true // Первые два простых числа
 }
 if num%2 == 0 {
  return false // Все четные числа (кроме 2) составные
 }
 sqrtNum := int(math.Sqrt(float64(num))) + 1
 for i := 3; i <= sqrtNum; i += 2 {
  if num%i == 0 {
   return false // Если найдено хотя бы одно простое деление, значит число составное
  }
 }
 return true
}

func main() {
 var inputNumber int
 fmt.Print("Введите число для проверки на простоту: ")
 fmt.Scanf("%d", &inputNumber)

 if isPrime(inputNumber) {
  fmt.Println(inputNumber, "- простое число.")
 } else {
  fmt.Println(inputNumber, "- составное число.")
 }
}
