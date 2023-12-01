// Вот пример алгоритма Беллмана-Форда на языке Go:

package main

import (
"fmt"
"math"
)

type Edge struct {
source int
destination int
weight int
}

type Graph struct {
vertices int
edges []Edge
}

func BellmanFord(graph Graph, startVertex int) ([]int, []int) {
dist := make([]int, graph.vertices)
prev := make([]int, graph.vertices)

for i := 0; i < graph.vertices; i++ {
dist[i] = math.MaxInt32
prev[i] = -1
}

dist[startVertex] = 0

for i := 1; i < graph.vertices; i++ {
  for j := 0; j < len(graph.edges); j++ {
    u := graph.edges[j].source
    v := graph.edges[j].destination
    weight := graph.edges[j].weight
    if dist[u]+weight < dist[v] {
      dist[v] = dist[u] + weight
      prev[v] = u
    }
  }
}

for j := 0; j < len(graph.edges); j++ {
  u := graph.edges[j].source
  v := graph.edges[j].destination
  weight := graph.edges[j].weight
  if dist[u]+weight < dist[v] {
    fmt.Println("Граф содержит цикл с отрицательным весом.")
    return nil, nil
  }
}

return dist, prev
}

func main() {
  graph := Graph{
    vertices: 5,
    edges: []Edge{
    {source: 0, destination: 1, weight: -1},
    {source: 0, destination: 2, weight: 4},
    {source: 1, destination: 2, weight: 3},
    {source: 1, destination: 3, weight: 2},
    {source: 1, destination: 4, weight: 2},
    {source: 3, destination: 2, weight: 5},
    {source: 3, destination: 1, weight: 1},
    {source: 4, destination: 3, weight: -3},
    },
  }

  startVertex := 0
  dist, prev := BellmanFord(graph, startVertex)

  fmt.Println("Расстояния от начальной вершины:")
  for i := 0; i < graph.vertices; i++ {
      fmt.Printf("Вершина %d: %d\n", i, dist[i])
  }

  fmt.Println("\nКратчайшие пути:")
  for i := 0; i < graph.vertices; i++ {
    path := []int{}
    for j := i; j != -1; j = prev[j] {
    path = append([]int{j}, path...)
    }
    fmt.Printf("Кратчайший путь до вершины %d: %v\n", i, path)
  }
}
/*
В этом примере алгоритм Беллмана-Форда реализован для нахождения кратчайших путей от одной вершины до всех остальных вершин в графе. 
Заметьте, что в приведенном примере используется взвешенный граф с отрицательными весами ребер. 
Алгоритм Беллмана-Форда может обрабатывать такие графы и обнаруживать наличие циклов с отрицательным весом.
*/
