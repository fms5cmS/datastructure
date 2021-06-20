package myStack

import (
	"fmt"
)

type ArrayStack struct {
	data []interface{}
	// 栈顶指针
	top int
}

func (s ArrayStack) String() string {
	if len(s.data) == 0 {
		return "nothing"
	}
	return fmt.Sprint("ArrayStack: ", s.data, " top")
}

func NewStack() ArrayStack {
	return ArrayStack{
		data: make([]interface{}, 0),
		top:  -1,
	}
}

func (s *ArrayStack) Push(x interface{}) {
	s.data = append(s.data, x)
	s.top++
}

func (s *ArrayStack) Pop() interface{} {
	ret := s.data[s.top]
	s.data = s.data[0:s.top]
	s.top--
	return ret
}

func (s *ArrayStack) Peek() interface{} {
	return s.data[s.top]
}

func (s *ArrayStack) Size() interface{} {
	return len(s.data)
}
