package graph

import "testing"

func generateBFSTestCase() *Graph {
	graph := NewGraph(8)
	graph.addEdge(0, 1)
	graph.addEdge(0, 3)
	graph.addEdge(1, 2)
	graph.addEdge(1, 4)
	graph.addEdge(2, 5)
	graph.addEdge(3, 4)
	graph.addEdge(4, 5)
	graph.addEdge(4, 6)
	graph.addEdge(5, 7)
	graph.addEdge(6, 7)
	return graph
}

func TestGraph_BFS(t *testing.T) {
	graph := generateBFSTestCase()
	graph.BFS(0, 5)
}

func generateDFSTestCase() *Graph {
	graph := NewGraph(9)
	graph.addEdge(1, 2)
	graph.addEdge(1, 4)
	graph.addEdge(2, 3)
	graph.addEdge(2, 5)
	graph.addEdge(3, 6)
	graph.addEdge(4, 5)
	graph.addEdge(5, 6)
	graph.addEdge(5, 7)
	graph.addEdge(6, 8)
	graph.addEdge(7, 8)
	return graph
}

func TestGraph_DFS(t *testing.T) {
	graph := generateDFSTestCase()
	graph.DFS(1, 7)
}
