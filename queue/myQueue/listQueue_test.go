package myQueue

import "testing"

func TestListQueue(t *testing.T) {
	queue := NewListQueue()
	t.Log(queue) // Output: there is nothing
	queue.Enqueue(1)
	queue.Enqueue("second")
	t.Log(queue)           // Output: front<-[1 second];size: 2
	t.Log(queue.Dequeue()) // Output: 1
	t.Log(queue)           // Output: front<-[second];size: 1
	t.Log(queue.Dequeue()) // second
	t.Log(queue)           // Output: there is nothing
}
