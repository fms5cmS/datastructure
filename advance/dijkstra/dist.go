package dijkstra

type VertexMinDist struct {
	id   uint // 该点的 id
	dist uint // 起点到该点的最短距离
}

func newVertexMinDist(id, dist uint) *VertexMinDist {
	return &VertexMinDist{id, dist}
}

type Dist []*VertexMinDist

func (d Dist) Len() int {
	return len(d)
}

func (d Dist) Less(i, j int) bool {
	return d[i].dist < d[j].dist
}

func (d Dist) Swap(i, j int) {
	panic("implement me")
}

func (d *Dist) Push(x interface{}) {
	panic("implement me")
}

func (d *Dist) Pop() interface{} {
	panic("implement me")
}
