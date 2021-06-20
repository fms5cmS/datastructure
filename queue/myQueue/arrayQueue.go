package myQueue

import (
	"fmt"
	"strings"
)

type ArrayQueue struct {
	data        []interface{}
	front, tail int
}

func (queue ArrayQueue) String() string {
	if queue.front == queue.tail {
		return "there is nothing"
	}
	ret := strings.Builder{}
	ret.WriteString("front<-[")
	for i := queue.front; i < queue.tail; i++ {
		ret.WriteString(fmt.Sprintf("%v ", queue.data[i]))
	}
	ret.WriteString("];")
	info := fmt.Sprintf(" size: %d, front: %d, tail: %d", queue.tail-queue.front, queue.front, queue.tail)
	ret.WriteString(info)
	return ret.String()
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		data:  make([]interface{}, 0),
		front: 0,
		tail:  0,
	}
}

func (queue *ArrayQueue) Enqueue(ele interface{}) {
	queue.data = append(queue.data, ele)
	queue.tail++
}

func (queue *ArrayQueue) Dequeue() (ret interface{}) {
	if queue.front < queue.tail {
		ret = queue.data[queue.front]
		queue.front++
	}
	return
}
