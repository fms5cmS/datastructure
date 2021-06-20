package stackQueue

import (
	"awesomeProject/stack/myStack"
	"strings"
)

// 232. 用栈实现队列
type MyQueue struct {
	in, out myStack.ArrayStack
}

func (this *MyQueue) String() string {
	if this.Empty() {
		return "nothing"
	}
	inSize := this.in.Size()
	if this.out.Size() == 0 {
		for i := 0; i < inSize; i++ {
			this.out.Push(this.in.Pop())
		}
	}
	ret := this.out.String()
	ret = strings.Replace(ret, "ArrayStack", "tail", -1)
	ret = strings.Replace(ret, "top", "front", -1)
	return ret
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		in:  myStack.NewStack(),
		out: myStack.NewStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Enqueue(x int) {
	this.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Dequeue() int {
	inSize := this.in.Size()
	if this.out.Size() == 0 {
		for i := 0; i < inSize; i++ {
			this.out.Push(this.in.Pop())
		}
	}
	return this.out.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	inSize := this.in.Size()
	if this.out.Size() == 0 {
		for i := 0; i < inSize; i++ {
			this.out.Push(this.in.Pop())
		}
	}
	return this.out.Peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.out.Size() == 0 && this.in.Size() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
