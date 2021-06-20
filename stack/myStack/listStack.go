package myStack

import (
	"container/list"
	"fmt"
	"strings"
)

type ListStack struct {
	list *list.List
}

func NewListStack() *ListStack {
	return &ListStack{
		list: list.New(),
	}
}

func (stack ListStack) String() string {
	if stack.list.Len() == 0 {
		return "the stack has nothing"
	}
	ret := strings.Builder{}
	cur := stack.list.Front()
	ret.WriteString("ListStack: [ ")
	for cur != nil {
		ret.WriteString(fmt.Sprintf("%v ", cur.Value))
		cur = cur.Next()
	}
	ret.WriteString("] top")
	return ret.String()
}

func (stack *ListStack) Size() int {
	return stack.list.Len()
}

func (stack *ListStack) Push(ele interface{}) {
	stack.list.PushBack(ele)
}

func (stack *ListStack) Pop() interface{} {
	last := stack.list.Back()
	stack.list.Remove(last)
	return last.Value
}

func (stack *ListStack) Peek() interface{} {
	return stack.list.Back().Value
}
