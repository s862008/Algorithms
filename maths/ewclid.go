package main

import (
	"fmt"
)

// GCD находит наибольший общий делитель (НОД) двух чисел с помощью алгоритма Евклида (итеративный подход)
func GCD(a, b int) int {
	// Обеспечиваем работу с положительными числами
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// Алгоритм Евклида: пока b не равно 0, заменяем (a, b) на (b, a % b)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// GCDRecursive находит НОД с помощью рекурсивной версии алгоритма Евклида
func GCDRecursive(a, b int) int {
	// Обеспечиваем работу с положительными числами
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// Базовый случай: если b равно 0, НОД — это a
	if b == 0 {
		return a
	}
	// Рекурсивный вызов: НОД(a, b) = НОД(b, a % b)
	return GCDRecursive(b, a%b)
}

func main() {
	// Примеры для тестирования
	testCases := []struct {
		a, b, expected int
	}{
		{48, 18, 6},
		{100, 25, 25},
		{35, 15, 5},
		{0, 5, 5},
		{7, 13, 1},  // взаимно простые числа
		{-48, 18, 6}, // проверка отрицательных чисел
	}

	fmt.Println("Алгоритм Евклида — итеративная реализация:")
	for _, tc := range testCases {
		result := GCD(tc.a, tc.b)
		fmt.Printf("НОД(%d, %d) = %d", tc.a, tc.b, result)
		if result == tc.expected {
			fmt.Println(" ✓")
		} else {
			fmt.Printf(" ✗ (ожидалось %d)\n", tc.expected)
		}
	}

	fmt.Println("\nАлгоритм Евклида — рекурсивная реализация:")
	for _, tc := range testCases {
		result := GCDRecursive(tc.a, tc.b)
		fmt.Printf("НОД(%d, %d) = %d", tc.a, tc.b, result)
		if result == tc.expected {
			fmt.Println(" ✓")
		} else {
			fmt.Printf(" ✗ (ожидалось %d)\n", tc.expected)
		}
	}
}
