package shortest_path

import "testing"

func TestGraph_Dijkstra(t *testing.T) {
	graph := NewGraph(6)
	AddEdge(0, 1, 10)
	AddEdge(0, 4, 17)
	AddEdge(1, 2, 20)
	AddEdge(2, 3, 15)
	AddEdge(2, 4, 25)
	AddEdge(4, 5, 6)
	Dijkstra(0, 3)
}
