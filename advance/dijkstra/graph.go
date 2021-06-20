package dijkstra

import "container/list"

// 边
type edge struct {
	sid, tid uint
	weight   uint
}

func newEdge(sid, tid, weight uint) *edge {
	return &edge{sid: sid, tid: tid, weight: weight}
}

// 无向有权图
type Graph struct {
	v        uint // 顶点个数
	i        uint
	vertexes map[string]uint // 顶点数据和 id 的映射关系
	adj      []*list.List    //邻接表
}

func NewGraph(v uint) *Graph {
	graph := &Graph{
		v:        v,
		i:        0,
		vertexes: make(map[string]uint),
		adj:      make([]*list.List, v),
	}
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return graph
}

func (graph *Graph) AddPlace(place string) bool {
	if graph.i >= graph.v {
		return false
	}
	graph.vertexes[place] = 0
	graph.i++
	return true
}

func (graph *Graph) AddEdge(sP, tP string, weight uint) {
	sid, tid := graph.vertexes[sP], graph.vertexes[tP]
	graph.adj[sid].PushBack(newEdge(sid, tid, weight))
	graph.adj[tid].PushBack(newEdge(tid, sid, weight))
}

func (graph *Graph) getId(place string) uint {
	return graph.vertexes[place]
}
