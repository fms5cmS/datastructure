package minstack

import "testing"

func TestMinStack_GetMin(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	t.Log(minStack.GetMin())
	t.Log(minStack.String())
	minStack.Pop()
	t.Log(minStack.Top())
	t.Log(minStack.GetMin())
}
