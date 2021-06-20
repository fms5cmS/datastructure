package queueStack

import (
	"awesomeProject/queue"
	"strings"
)

// 225. 用队列实现栈
type MyStack struct {
	a, b queue.Queue
}

func (this *MyStack) String() string {
	if this.a.Size() == 0 {
		return "nothing"
	}
	ret := this.a.String()
	ret = strings.Replace(ret, "front", "Stack", -1)
	ret = strings.Replace(ret, "tail", "top", -1)
	return ret
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		a: queue.NewQueue(),
		b: queue.NewQueue(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.a.Enqueue(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.a.Size() == 1 {
		return this.a.Dequeue()
	}
	for i := 0; i < this.a.Size(); i++ {
		this.b.Enqueue(this.a.Dequeue())
	}
	ret := this.a.Dequeue()
	this.a, this.b = this.b, this.a
	return ret
}

/** Get the top element. */
func (this *MyStack) Top() int {
	top := this.Pop()
	this.Push(top)
	return top
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.a.Empty()
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
