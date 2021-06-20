package myQueue

import (
	"container/list"
	"fmt"
	"strings"
)

type ListQueue struct {
	data        *list.List
	front, tail *list.Element
}

func (queue ListQueue) String() string {
	if queue.front == nil {
		return "there is nothing"
	}
	ret := strings.Builder{}
	ret.WriteString("front<-[")
	cur, count := queue.front, 0
	for cur != nil {
		ret.WriteString(fmt.Sprintf("%v", cur.Value))
		cur = cur.Next()
		if cur != nil {
			ret.WriteString(" ")
		}
		count++
	}
	ret.WriteString("];")
	ret.WriteString(fmt.Sprintf("size: %d", count))
	return ret.String()
}

func NewListQueue() *ListQueue {
	return &ListQueue{data: list.New()}
}

func (queue *ListQueue) Enqueue(val interface{}) {
	queue.tail = queue.data.PushBack(val)
	if queue.front == nil {
		queue.front = queue.tail
	}
}

func (queue *ListQueue) Dequeue() (ret interface{}) {
	if queue.front == nil {
		return "there is nothing"
	}
	front := queue.data.Front()
	queue.front = front.Next()
	ret = queue.data.Remove(front)
	return
}
