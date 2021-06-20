package myStack

import "testing"

func TestListStack(t *testing.T) {
	stack := NewListStack()
	t.Logf("%s, Size: %d", stack, stack.Size())
	stack.Push("first")
	stack.Push("second")
	stack.Push(3)
	t.Logf("%s, Size: %d", stack, stack.Size())
	t.Log(stack.Pop())
	t.Logf("%s, Size: %d", stack, stack.Size())
	t.Log(stack.Peek())
}
