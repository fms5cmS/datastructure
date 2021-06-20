package stackQueue

import "testing"

func TestMyQueue_Push(t *testing.T) {
	queue := Constructor()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	t.Log(queue.String())
}

func TestMyQueue_Empty(t *testing.T) {
	queue := Constructor()
	t.Log(queue.Empty())
	queue.Enqueue(1)
	t.Log(queue.Empty())
}

func TestMyQueue_Pop(t *testing.T) {
	queue := Constructor()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	t.Log(queue.String())
	t.Log(queue.Dequeue())
	t.Log(queue.String())
}

func TestMyQueue_Peek(t *testing.T) {
	queue := Constructor()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	t.Log(queue.Peek())
}
