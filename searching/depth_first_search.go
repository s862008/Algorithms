Вот пример реализации алгоритма поиска в глубину (Depth-First Search, DFS) на языке Go:

```go
package main

import "fmt"

// Структура для представления графа в виде списка смежности
type Graph struct {
vertices int
adjList map[int][]int
}

// Инициализация графа
func NewGraph(vertices int) *Graph {
adjList := make(map[int][]int)
return &Graph{vertices: vertices, adjList: adjList}
}

// Добавление ребра в граф
func (g *Graph) AddEdge(u, v int) {
g.adjList[u] = append(g.adjList[u], v)
g.adjList[v] = append(g.adjList[v], u)
}

// Вспомогательная функция для рекурсивного обхода в глубину
func dfsUtil(vertex int, visited []bool, g *Graph) {
visited[vertex] = true
fmt.Printf("%d ", vertex)

// Рекурсивно обходим все смежные вершины
for _, v := range g.adjList[vertex] {
if !visited[v] {
dfsUtil(v, visited, g)
}
}
}

// Обход графа в глубину
func DFS(g *Graph, startVertex int) {
visited := make([]bool, g.vertices)
dfsUtil(startVertex, visited, g)
}

func main() {
g := NewGraph(5)
g.AddEdge(0, 1)
g.AddEdge(0, 2)
g.AddEdge(1, 3)
g.AddEdge(1, 4)

startVertex := 0
fmt.Printf("DFS traversal starting from vertex %d: ", startVertex)
DFS(g, startVertex)
}
```

В этом примере реализован алгоритм поиска в глубину на основе рекурсии. В структуре `Graph` используется список смежности для представления графа. Метод `AddEdge` добавляет ребра в граф. Функция `DFS` и вспомогательная функция `dfsUtil` выполняют обход графа в глубину, начиная с заданной стартовой вершины.

В функции `main` создается граф с 5 вершинами и добавляются несколько ребер. Затем вызывается функция `DFS` для обхода графа в глубину, начиная с вершины 0. Результат обхода выводится на экран.

Результат выполнения программы будет следующим:

```
DFS traversal starting from vertex 0: 0 1 3 4 2
```

Это пример простого обхода графа в глубину на языке Go. Вы можете изменить граф и стартовую вершину в соответствии с вашими потребностями.
