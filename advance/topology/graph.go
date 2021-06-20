package topology

import (
	"container/list"
)

type Graph struct {
	v   int          // 顶点个数
	adj []*list.List // 邻接表
}

// s 指向 t
func (graph *Graph) addEdge(s, t int) {
	graph.adj[s].PushBack(t)
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

// 基于 Kahn 算法实现拓扑排序
func (graph Graph) TopologicalSortByKahn() {
	inDegree := make([]int, graph.v)
	// 统计所有顶点的入度
	for i := 0; i < graph.v; i++ {
		front := graph.adj[i].Front()
		for ; front != nil; front = front.Next() {
			v := front.Value.(int)
			inDegree[v]++
		}
	}
	// 向 queue 中加入所有入度为 0 的顶点
	queue := make([]int, 0)
	for i := 0; i < graph.v; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		print("->", cur) // 打印 cur
		// 将 cur 可到达的顶点入度减一
		front := graph.adj[cur].Front()
		for ; front != nil; front = front.Next() {
			v := front.Value.(int)
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
}

// 基于 DFS 算法实现拓扑排序
func (graph Graph) TopologicalSortByDFS() {
	// 通过邻接表构建逆邻接表
	inverseAdj := make([]*list.List, graph.v)
	for i := range inverseAdj {
		inverseAdj[i] = list.New()
	}
	for i := 0; i < graph.v; i++ {
		front := graph.adj[i].Front()
		for ; front != nil; front = front.Next() {
			v := front.Value.(int)    // i->v
			inverseAdj[v].PushBack(i) // v->i
		}
	}
	// 递归处理每个顶点
	visited := make([]bool, graph.v)
	for i := 0; i < graph.v; i++ {
		if visited[i] == false {
			visited[i] = true
			graph.dfs(i, inverseAdj, visited)
		}
	}
}

// 输出逆邻接表中 vertex 可到达的所有顶点，再输出 vertex
// 也即输出邻接表中所有可到达 vertex 的顶点，再输出 vertex
func (graph *Graph) dfs(vertex int, inverseAdj []*list.List, visited []bool) {
	curVertex := inverseAdj[vertex].Front()
	for ; curVertex != nil; curVertex = curVertex.Next() {
		curV := curVertex.Value.(int)
		if visited[curV] == true {
			continue
		}
		visited[curV] = true
		graph.dfs(curV, inverseAdj, visited)
	}
	print("->", vertex)
}
