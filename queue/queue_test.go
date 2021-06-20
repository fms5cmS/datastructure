package queue

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	t.Log(queue.String())
}

func TestQueue_Dequeue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	t.Log(queue.Dequeue())
	t.Log(queue.String())
}

func TestQueue_Size(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	t.Log(queue.Size())
	t.Log(queue.Dequeue())
	t.Log(queue.Size())
}

func TestQueue_Empty(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	t.Log(queue.Empty())
	queue.Dequeue()
	t.Log(queue.Empty())
}
