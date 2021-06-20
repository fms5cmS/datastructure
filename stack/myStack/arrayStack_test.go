package myStack

import "testing"

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push("second")
	stack.Push(3)
	t.Log(stack)
	t.Log(stack.Pop())
	t.Log(stack)
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	stack.Push("one")
	stack.Push(2)
	stack.Push("third")
	t.Log(stack)
	t.Log(stack.Pop())
	t.Log(stack)
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Log(stack)
	t.Log(stack.Pop())
	t.Log(stack)
	t.Log(stack.Peek())
	t.Log(stack)
}

func TestStack_Size(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	t.Log(stack)
	t.Log(stack.Pop())
	t.Log(stack)
	t.Log(stack.Size())
}
