package shortest_path

import (
	"container/heap"
	"container/list"
)

type Graph struct {
	v   int
	adj []*list.List
}

type edge struct {
	sid, tid int // 边的起始/终止顶点编号
	w        int // 边的权重
}

func NewGraph(v int) *Graph {
	graph := &Graph{
		v:   v,
		adj: make([]*list.List, v),
	}
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return graph
}

func newEdge(sid, tid, w int) *edge {
	return &edge{sid, tid, w}
}

func (graph *Graph) AddEdge(s, t, w int) {
	graph.adj[s].PushBack(newEdge(s, t, w))
}

// s 到 t 的最短路径
func (graph *Graph) Dijkstra(s, t int) {
	// 用来还原最短路径
	predecessor := make([]int, graph.v)
	// 记录从起始顶点到每个顶点的距离 dist
	vertexes := make([]*Vertex, graph.v)
	for i := 0; i < graph.v; i++ {
		// 这里初始化距离为 1000
		vertexes[i] = NewVertex(i, 1000)
	}
	queue := make(PriorityQueue, 0) // 小顶堆
	heap.Init(&queue)
	// 标记是否进入过队列
	inqueue := make([]bool, graph.v)
	dist = 0
	heap.Push(&queue, vertexes[s])
	inqueue[s] = true
	for Len() > 0 {
		// 从优先队列中取出 dist 最小的顶点，考察该点可达的所有点
		minVertex := heap.Pop(&queue).(*Vertex)
		if id == t { // 最短路径已产生
			break
		}
		front := graph.adj[minVertex.id].Front()
		for ; front != nil; front = front.Next() {
			// 取出一条 minVertex 相连的边
			e := front.Value.(*edge)
			// minVertex -> nextVertex
			nextVertex := vertexes[e.tid]
			// 更新 next 的 dist
			if dist+e.w < dist {
				dist = dist + e.w
				predecessor[id] = id
				if inqueue[id] == true {
					// 对更新了 dist 的 nextVertex 进行堆化
					heap.Fix(&queue, id)
				} else {
					heap.Push(&queue, nextVertex)
					inqueue[id] = true
				}
			}
		}
	}
	// 输出最短路径
	print(s)
	graph.print(s, t, predecessor)
}

func (graph *Graph) print(s, t int, predecessor []int) {
	if s == t {
		return
	}
	graph.print(s, predecessor[t], predecessor)
	print("->", t)
}
