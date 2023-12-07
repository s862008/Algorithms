//Пример реализации  чисел фибоначчи на языке GO (РЕКУРСИВНЫЙ АЛГОРИТМ)
package main

import (
"fmt"
)

func fibonacci(n int) int {
    if n<= 1 {
        return n
    }
    return fibonacci(n-1)+fibonacci(n-2)
}

func main() {
    num := 10 // Количество чисел,которые надо вычислить
    for i:=0; i<num;i++ {
    fmt.Println(fibonacci(i))
    }
}
