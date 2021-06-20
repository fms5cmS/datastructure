package shortest_path

// 为实现 Dijkstra 算法而使用的
type Vertex struct {
	id   int // 顶点编号
	dist int // 起始顶点到该顶点的距离
}

func NewVertex(id, dist int) *Vertex {
	return &Vertex{id, dist}
}

type PriorityQueue []*Vertex

func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, j int) bool {
	return p[i].dist < p[j].dist
}

func (p PriorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].id = i
	p[j].id = j
}
func (p *PriorityQueue) Push(x interface{}) {
	n := len(*p)
	vertex := x.(*Vertex)
	vertex.id = n
	*p = append(*p, vertex)
}
func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	//item.id = -1
	*p = old[0 : n-1]
	return item
}
