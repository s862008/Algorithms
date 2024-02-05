package main
import "fmt"

// Структура, представляющая ребро графа
type Edge struct {
  source, destination, weight int
}

// Структура, представляющая граф
type Graph struct {
  edges []Edge
}

// Функция для поиска корня множества
func find(parent []int, i int) int {
  if parent[i] == -1 {
    return i
  }
  return find(parent, parent[i])
}

// Функция для объединения двух множеств
func union(parent []int, x, y int) {
  xRoot := find(parent, x)
  yRoot := find(parent, y)
  parent[xRoot] = yRoot
}

// Функция, реализующая алгоритм Крускала
func Kruskal(graph Graph, vertices int) []Edge {
  result := make([]Edge, 0)
  i, e := 0, 0
  sortedEdges := make([]Edge, len(graph.edges))
  copy(sortedEdges, graph.edges)

  // Сортировка ребер по возрастанию веса
  for i := 0; i < len(sortedEdges); i++ {
    for j := i + 1; j < len(sortedEdges); j++ {
      if sortedEdges[i].weight > sortedEdges[j].weight {
      sortedEdges[i], sortedEdges[j] = sortedEdges[j], sortedEdges[i]
      }
    }
  }

  parent := make([]int, vertices)
  for i := range parent {
    parent[i] = -1
  }

  // Выбираем ребра, пока не достигнем нужного количества вершин
  for e < vertices-1 {
    edge := sortedEdges[i]
    x := find(parent, edge.source)
    y := find(parent, edge.destination)

    if x != y {
      result = append(result, edge)
      union(parent, x, y)
      e++
    }
    i++
  }

  return result
}

func main() {
// Пример использования
  graph := Graph{
    edges: []Edge{
    {0, 1, 10},
    {0, 2, 6},
    {0, 3, 5},
    {1, 3, 15},
    {2, 3, 4},
    },
  }

  result := Kruskal(graph, 4)

  fmt.Println("Минимальное остовное дерево:")
  for _, edge := range result {
    fmt.Printf("%d - %d: %d\n", edge.source, edge.destination, edge.weight)
  }
}
