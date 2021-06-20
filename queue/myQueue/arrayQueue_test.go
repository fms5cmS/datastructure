package myQueue

import "testing"

func TestArrayQueue(t *testing.T) {
	queue := NewArrayQueue()
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue(3)
	// Output: front<-[first second 3 ]; size:3, front: 0, tail: 3
	t.Log(queue)
	// Output: first
	t.Log(queue.Dequeue())
	// Output: front<-[second 3 ]; size: 2, front: 1, tail: 3
	t.Log(queue)
}
