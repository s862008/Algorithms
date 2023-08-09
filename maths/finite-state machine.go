/* 
В этом примере на языке Go мы создаем конечный автомат с тремя состояниями (StateA, StateB, StateC) и определяем последовательность входных данных (input).
Алгоритм просто проверяет последовательность символов во входных данных и в зависимости от текущего состояния и входного символа переходит в следующее состояние.
В данном примере, при входных данных "a", "b", "c", "b", "a", конечное состояние будет StateC.
*/

package main

import "fmt"

// Состояния конечного автомата
const (
StateA = iota   // 0
StateB          // 1
StateC          // 2
)

func main() {
input := []string{"a", "b", "c", "b", "a"} // Пример входных данных

currentState := StateA // Начальное состояние

for _, inp := range input {
		switch currentState {
		case StateA:
			if inp == "a" {
				currentState = StateB
			}
		case StateB:
			if inp == "b" {
				currentState = StateC
			}
		case StateC:
			if inp == "c" {
				currentState = StateB
			}
		}
	}

fmt.Println("Конечное состояние:", currentState)
}
