package queue

import (
	"fmt"
)

// 自定义 Queue
type Queue struct {
	data        []int
	front, tail int
}

func (q *Queue) String() string {
	if q.front == q.tail {
		return "nothing"
	}
	return fmt.Sprint("front: ", q.data, " tail")
}

func NewQueue() Queue {
	return Queue{
		data:  make([]int, 0),
		front: 0,
		tail:  0,
	}
}

func (q *Queue) Enqueue(x int) {
	q.data = append(q.data, x)
	q.tail++
}

func (q *Queue) Dequeue() int {
	ret := 0
	if q.front < q.tail {
		ret = q.data[q.front]
		q.front++
	}
	return ret
}

func (q *Queue) Size() int {
	return q.tail - q.front
}

func (q *Queue) Empty() bool {
	return q.front == q.tail
}
