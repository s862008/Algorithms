package main

import (
	"fmt"
	"math"
)

// solveQuadratic решает квадратное уравнение ax^2 + bx + c = 0
// и выводит корни в консоль. Учитывает все возможные случаи:
// - два действительных корня
// - один корень (кратный)
// - комплексные корни
// - вырожденные случаи (не квадратное уравнение)
func solveQuadratic(a, b, c float64) {
	// Небольшой эпсилон для безопасного сравнения float64 с нулём.
	// В реальных системах (например, при валидации данных в МИС/СЭМД)
	// это помогает избежать ошибок из-за погрешностей вычислений.
	const epsilon = 1e-9

	// Случай 1: a ≈ 0 — уравнение не квадратное, а линейное (или вообще не уравнение)
	if math.Abs(a) < epsilon {
		if math.Abs(b) < epsilon {
			// Уравнение вида 0x^2 + 0x + c = 0 → c = 0
			if math.Abs(c) < epsilon {
				fmt.Println("Уравнение вырождено: 0 = 0. Имеет бесконечно много решений.")
			} else {
				fmt.Printf("Уравнение противоречиво: %.2f = 0. Решений нет.\n", c)
			}
		} else {
			// Линейное уравнение: bx + c = 0 → x = -c / b
			x := -c / b
			fmt.Printf("Линейное уравнение. Один корень: x = %.6f\n", x)
		}
		return
	}

	// Вычисляем дискриминант: D = b^2 - 4ac
	discriminant := b*b - 4*a*c

	if discriminant > epsilon {
		// Два различных действительных корня
		sqrtD := math.Sqrt(discriminant)
		x1 := (-b + sqrtD) / (2 * a)
		x2 := (-b - sqrtD) / (2 * a)
		fmt.Printf("Два действительных корня: x1 = %.6f, x2 = %.6f\n", x1, x2)
	} else if math.Abs(discriminant) <= epsilon {
		// Один действительный корень (дискриминант ≈ 0)
		x := -b / (2 * a)
		fmt.Printf("Один корень (кратный): x = %.6f\n", x)
	} else {
		// Комплексные корни: D < 0
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(math.Abs(discriminant)) / (2 * a)
		fmt.Printf("Комплексные корни:\n")
		fmt.Printf("x1 = %.6f + %.6fi\n", realPart, imaginaryPart)
		fmt.Printf("x2 = %.6f - %.6fi\n", realPart, imaginaryPart)
	}
}

func main() {
	// Примеры вызовов для разных случаев (можно менять коэффициенты)
	fmt.Println("Пример 1 (два корня): a=1, b=-5, c=6 → x=2, x=3")
	solveQuadratic(1, -5, 6)

	fmt.Println("\nПример 2 (один корень): a=1, b=2, c=1 → x=-1")
	solveQuadratic(1, 2, 1)

	fmt.Println("\nПример 3 (комплексные корни): a=1, b=0, c=1 → x=±i")
	solveQuadratic(1, 0, 1)

	fmt.Println("\nПример 4 (линейное уравнение): a=0, b=2, c=-4 → x=2")
	solveQuadratic(0, 2, -4)

	fmt.Println("\nПример 5 (противоречивое уравнение): a=0, b=0, c=5")
	solveQuadratic(0, 0, 5)
}
