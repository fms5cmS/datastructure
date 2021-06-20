package graph

import (
	"container/list"
	"log"
)

// 无向图
type Graph struct {
	v   int          // 顶点个数
	adj []*list.List // 邻接表
}

func NewGraph(v int) *Graph {
	graph := &Graph{v: v, adj: make([]*list.List, v)}
	for i := range graph.adj {
		graph.adj[i] = list.New()
	}
	return graph
}

func (graph *Graph) addEdge(s, t int) {
	graph.adj[s].PushBack(t)
	graph.adj[t].PushBack(s)
}

func (graph Graph) BFS(s, t int) {
	if s == t {
		return
	}
	queue := make([]int, 0)
	// 记录到达当前顶点(顶点值等于索引)的前一个顶点
	prev := make([]int, graph.v)
	// 记录当前顶点是否已经经过了
	visited := make([]bool, graph.v)
	// 是否已找到目标 t
	find := false
	queue = append(queue, s)
	for i := range prev {
		prev[i] = -1
	}
	visited[s] = true
	for len(queue) > 0 && !find {
		curVal := queue[0]
		queue = queue[1:]
		for next := graph.adj[curVal].Front(); next != nil; next = next.Next() {
			nextVal := next.Value.(int)
			if !visited[nextVal] {
				prev[nextVal] = curVal
				visited[nextVal] = true
				if t == nextVal {
					find = true
					break
				}
				queue = append(queue, nextVal)
			}
		}
	}
	if find {
		printPth(prev, s, t)
	} else {
		log.Printf("can not find path from %d to %d", s, t)
	}
}

func (graph Graph) DFS(s, t int) {
	find := false
	visited := make([]bool, graph.v)
	prev := make([]int, graph.v)
	for i := range prev {
		prev[i] = -1
	}
	graph.recurDFS(s, t, visited, prev, find)
	printPth(prev, s, t)
}

func (graph *Graph) recurDFS(s, t int, visited []bool, prev []int, find bool) {
	if find {
		return
	}
	visited[s] = true
	if s == t {
		find = true
		return
	}
	for next := graph.adj[s].Front(); next != nil; next = next.Next() {
		nextVal := next.Value.(int)
		if !visited[nextVal] {
			prev[nextVal] = s
			graph.recurDFS(nextVal, t, visited, prev, find)
		}
	}
}

func printPth(prev []int, s, t int) {
	if t == s || prev[t] == -1 {
		print(t)
	} else {
		printPth(prev, s, prev[t])
		print(" -> ", t)
	}
}
