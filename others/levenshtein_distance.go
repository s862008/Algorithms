//This code calculates the Levenshtein distance between two strings, s1 and s2. 
//The Levenshtein distance is a measure of the minimum number of operations (insertions, deletions, or substitutions) needed to change one string into the other. 

//Вот пример реализации алгоритма Левенштейна на языке Go:

package main

import (
  "fmt"
)

func min(a, b, c int) int {
  if a < b && a < c {
    return a
  } else if b < c {
    return b
  } else {
    return c
  }
}

func levenshteinDistance(s1, s2 string) int {
  m := len(s1)
  n := len(s2)

  // Создаем матрицу размером (m+1)x(n+1)
  // Initialize the dynamic programming table (DP table)
  d := make([][]int, m+1)
  for i := 0; i <= m; i++ {
    d[i] = make([]int, n+1)
  }

  // Инициализируем первую строку и столбец матрицы
  // Initialize the first row and column of the DP table
  for i := 0; i <= m; i++ {
    d[i][0] = i
  }
  for j := 1; j <= n; j++ {
    d[0][j] = j
  }

  // Заполняем матрицу по правилу алгоритма
  // Fill in the DP table
  for j := 1; j <= n; j++ {
    for i := 1; i <= m; i++ {
      cost := 0
      if s1[i-1] != s2[j-1] {
        cost = 1
      }
      d[i][j] = min(d[i-1][j]+1, d[i][j-1]+1, d[i-1][j-1]+cost)
    }
  } 

  // Возвращаем нижний правый элемент матрицы, содержащий расстояние Левенштейна
  // Calculate the minimum distance
  return d[m][n]
}

func main() {
  s1 := "kitten"
  s2 := "sitting"
  distance := levenshteinDistance(s1, s2)
  fmt.Printf("The Levenshtein distance between \"%s\" and \"%s\" is %d.\n", s1, s2, distance)
}


/*
Этот код представляет собой функцию levenshteinDistance, которая принимает две строки s1 и s2 в качестве аргументов и возвращает расстояние Левенштейна между ними.
Функция min используется для определения минимального значения из трех чисел.
В основном блоке main создаются две строки для тестирования алгоритма (s1 и s2), и выполняется вызов функции levenshteinDistance для вычисления расстояния Левенштейна между ними. Результат выводится на консоль.

Пример вывода:
Расстояние Левенштейна между "kitten" и "sitting" равно 3

Обратите внимание, что в данном примере используется рекурсивный подход для вычисления расстояния Левенштейна. В реальных приложениях для оптимизации рекомендуется использовать
метод динамического программирования с использованием заполнения матрицы. 
*/
