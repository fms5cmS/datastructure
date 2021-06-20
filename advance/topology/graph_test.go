package topology

import "testing"

func TestGraph_TopologicalSortByKahn(t *testing.T) {
	graph := NewGraph(7)
	graph.addEdge(0, 1)
	graph.addEdge(1, 2)
	graph.addEdge(3, 2)
	graph.addEdge(2, 4)
	graph.addEdge(4, 5)
	graph.addEdge(5, 3)
	graph.addEdge(3, 6)
	// 输出的顶点个数小于图中顶点的个数，即图中还有入度不为 0 的顶点，所以图中存在环
	graph.TopologicalSortByKahn()
}

func TestGraph_TopologicalSortByDFS(t *testing.T) {
	graph := NewGraph(7)
	graph.addEdge(0, 1)
	graph.addEdge(1, 2)
	graph.addEdge(2, 3)
	graph.addEdge(2, 4)
	graph.addEdge(4, 5)
	graph.addEdge(5, 3)
	graph.addEdge(3, 6)
	graph.TopologicalSortByDFS()
}
