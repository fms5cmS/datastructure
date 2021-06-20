package queueStack

import "testing"

func TestMyStack_Push(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Log(stack.String())
	t.Log(stack.a.Size())
}

func TestMyStack_Pop(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Log(stack.String())
	t.Log("pop ", stack.Pop())
	t.Log(stack.String())
}

func TestMyStack_Top(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	t.Log(stack.Top())
	t.Log(stack.Pop())
	t.Log(stack.Top())
}
