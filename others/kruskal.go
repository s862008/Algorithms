package main
import (
"fmt"
"sort"
)
type Edge struct {
src int
dest int
weight int
}
type Graph struct {
V, E int
edges []Edge
}
type Subset struct {
parent, rank int
}
func NewGraph(V, E int) *Graph {
return &Graph{V: V, E: E, edges: make([]Edge, E)}
}
func (g *Graph) AddEdge(src, dest, weight int) {
g.edges = append(g.edges, Edge{src: src, dest: dest, weight: weight})
}
// FindSubset finds the subset of an element i
func (subsets []*Subset) FindSubset(i int) int {
if subsets[i].parent != i {
subsets[i].parent = subsets.FindSubset(subsets[i].parent)
}
return subsets[i].parent
}
// Union subsets with x and y
func (subsets []*Subset) Union(x, y int) {
rootX := subsets.FindSubset(x)
rootY := subsets.FindSubset(y)
if subsets[rootX].rank < subsets[rootY].rank {
subsets[rootX].parent = rootY
} else if subsets[rootX].rank > subsets[rootY].rank {
subsets[rootY].parent = rootX
} else {
subsets[rootY].parent = rootX
subsets[rootX].rank++
}
}
// Kruskal algorithm implementation
func Kruskal(graph *Graph) {
V := graph.V
res := make([]Edge, 0)
i, e := 0, 0
sort.Slice(graph.edges, func(i, j int) bool {
return graph.edges[i].weight < graph.edges[j].weight
})
subsets := make([]*Subset, V)
for i := 0; i < V; i++ {
subsets[i] = &Subset{parent: i, rank: 0}
}
for e < V-1 && i < len(graph.edges) {
edge := graph.edges[i]
i++
x := subsets.FindSubset(edge.src)
y := subsets.FindSubset(edge.dest)
if x != y {
res = append(res, edge)
e++
subsets.Union(x, y)
}
}
fmt.Println("Minimum Spanning Tree using Kruskal's algorithm:")
for i := 0; i < e; i++ {
fmt.Printf("%d — %d == %d\n", res[i].src, res[i].dest, res[i].weight)
}
}
func main() {
V, E := 4, 5
graph := NewGraph(V, E)
graph.AddEdge(0, 1, 10)
graph.AddEdge(0, 2, 6)
graph.AddEdge(0, 3, 5)
graph.AddEdge(1, 3, 15)
graph.AddEdge(2, 3, 4)
Kruskal(graph)
}
