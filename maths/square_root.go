package main

import (
 "fmt"
)

func squareRoot(num float64) float64 {
 guess := num / 2
 for i := 0; i < 10; i++ {
  guess = (guess + num/guess) / 2
 }
 return guess
}

func main() {
 num := 10.0
 result := squareRoot(num)
 fmt.Printf("Square root of %v is %v\n", num, result)
} 
