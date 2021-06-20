package shortest_path

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := make(PriorityQueue, 6)
	pq[0] = &Vertex{0, 77}
	pq[1] = &Vertex{1, 100}
	pq[2] = &Vertex{2, 120}
	pq[3] = &Vertex{4, 10}
	pq[4] = &Vertex{5, 75}
	pq[5] = &Vertex{3, 99}
	heap.Init(&pq)
	v := &Vertex{5, 101}
	heap.Push(&pq, v)
	for _, v := range pq {
		t.Log(v)
	}
	for Len() > 0 {
		v := heap.Pop(&pq).(*Vertex)
		t.Logf("%+v", v)
	}
}
