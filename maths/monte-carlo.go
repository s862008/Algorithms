//Пример кода на Go для вычисления числа Пи методом Монте-Карло

package main

import (
 "fmt"
 "math/rand"
 "time"
)

func estimatePi(numPoints int) float64 {
 rand.Seed(time.Now().UnixNano())
 hits := 0

 for i := 0; i < numPoints; i++ {
  // Генерируем случайные x и y координаты в пределах [-1, 1]
  x := rand.Float64()*2 - 1
  y := rand.Float64()*2 - 1

  // Проверяем, попадает ли точка в круг
  if x*x+y*y <= 1 {
   hits++
  }
 }

 // Оценка числа π
 return 4 * float64(hits) / float64(numPoints)
}

func main() {
 numPoints := 1000000 // Количество случайных точек
 estimatedPi := estimatePi(numPoints)

 fmt.Printf("Estimated value of Pi: %f\n", estimatedPi)
}