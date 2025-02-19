/*
пример реализации алгоритма K-средних для кластеризации данных на языке Go. Этот пример включает в себя генерацию случайных данных,
реализацию алгоритма K-средних и вывод результатов.
*/
package main

import (
 "fmt"
 "math"
 "math/rand"
 "time"
)

// Point представляет точку данных в двумерном пространстве
type Point struct {
 X, Y float64
}

// Cluster представляет кластер с его центром и точками данных
type Cluster struct {
 Centroid Point
 Points   []Point
}

// euclideanDistance вычисляет евклидово расстояние между двумя точками
func euclideanDistance(p1, p2 Point) float64 {
 return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

// assignToClusters назначает каждую точку данных ближайшему кластеру
func assignToClusters(points []Point, clusters []*Cluster) {
 for i := range clusters {
  clusters[i].Points = []Point{} // Очистить существующие точки кластера
 }

 for _, point := range points {
  closestClusterIndex := 0
  minDistance := math.MaxFloat64

  for i, cluster := range clusters {
   distance := euclideanDistance(point, cluster.Centroid)
   if distance < minDistance {
    minDistance = distance
    closestClusterIndex = i
   }
  }

  clusters[closestClusterIndex].Points = append(clusters[closestClusterIndex].Points, point)
 }
}

// updateCentroids вычисляет новые центроиды для каждого кластера
func updateCentroids(clusters []*Cluster) bool {
 moved := false
 for i, cluster := range clusters {
  if len(cluster.Points) == 0 {
   continue // Пропустить кластеры без точек
  }

  sumX := 0.0
  sumY := 0.0
  for _, point := range cluster.Points {
   sumX += point.X
   sumY += point.Y
  }

  newCentroidX := sumX / float64(len(cluster.Points))
  newCentroidY := sumY / float64(len(cluster.Points))

  if newCentroidX != cluster.Centroid.X || newCentroidY != cluster.Centroid.Y {
   moved = true
  }

  clusters[i].Centroid = Point{X: newCentroidX, Y: newCentroidY}
 }

 return moved
}

// kMeans выполняет алгоритм K-средних
func kMeans(points []Point, k int, maxIterations int) []*Cluster {
 // 1. Инициализация:  Выбор k случайных точек в качестве начальных центроидов
 rand.Seed(time.Now().UnixNano())
 clusters := make([]*Cluster, k)
 for i := 0; i < k; i++ {
  randomIndex := rand.Intn(len(points))
  clusters[i] = &Cluster{Centroid: points[randomIndex]}
 }

 // 2. Итерация:  Повторять до сходимости или достижения максимального количества итераций
 for i := 0; i < maxIterations; i++ {
  // a.  Назначение: Назначить каждую точку ближайшему кластеру
  assignToClusters(points, clusters)

  // b. Обновление:  Пересчитать центроиды каждого кластера
  moved := updateCentroids(clusters)

  // c. Проверка сходимости:  Если центроиды не переместились, алгоритм сошелся
  if !moved {
   fmt.Printf("Converged after %d iterations\n", i+1)
   break
  }
 }

 return clusters
}

func main() {
 // 1. Генерация случайных данных
 numPoints := 100
 points := make([]Point, numPoints)
 for i := 0; i < numPoints; i++ {
  points[i] = Point{X: rand.Float64() * 100, Y: rand.Float64() * 100} // случайные точки от 0 до 100
 }

 // 2. Запуск алгоритма K-средних
 k := 3 // Количество кластеров
 maxIterations := 100
 clusters := kMeans(points, k, maxIterations)

 // 3. Вывод результатов
 fmt.Println("Clusters:")
 for i, cluster := range clusters {
  fmt.Printf("Cluster %d: Centroid = (%f, %f), Number of points = %d\n",
   i+1, cluster.Centroid.X, cluster.Centroid.Y, len(cluster.Points))
  // Можно также вывести координаты точек в каждом кластере, если необходимо
  // fmt.Printf("  Points:\n")
  // for _, point := range cluster.Points {
  //  fmt.Printf("    (%f, %f)\n", point.X, point.Y)
  // }

 }
}
