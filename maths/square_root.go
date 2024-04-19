package main

import (
 "fmt"
)

func squareRoot(num float64) float64 {
 guess :=num / 2
 for i := 0; i < 10; i++ {
  t:= guess
  guess = (t + (num/t)) / 2
    if(t == guess){
        return guess
    }
}
 return guess
}

func main() {
 num := 10.0
 result := squareRoot(num)
 fmt.Printf("Square root of %v is %v\n", num, result)
} 
