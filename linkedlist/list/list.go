package list

import (
	"errors"
	"fmt"
	"strings"
)

type List struct {
	dummyHead *Node // 虚拟头节点
	len       int
}

func NewList() *List {
	return &List{dummyHead: &Node{}}
}

type Node struct {
	Val  interface{}
	next *Node
}

func (node Node) String() string {
	return fmt.Sprintf("%v", node.Val)
}

func (list List) String() string {
	cur := list.dummyHead
	ret := strings.Builder{}
	for cur.next != nil {
		ret.WriteString(cur.next.String())
		ret.WriteString(" -> ")
		cur = cur.next
	}
	ret.WriteString("nil")
	return ret.String()
}

func (list *List) Size() int {
	return list.len
}

func (list *List) Insert(val interface{}) {
	new := &Node{Val: val}
	cur := list.dummyHead
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = new
	list.len++
}

func (list *List) InsertAfter(before, val interface{}) error {
	if prev := list.findPrev(before); prev != nil {
		new := &Node{Val: val}
		prev.next, new.next = new, prev.next
		list.len++
	}
	return errors.New(fmt.Sprintf("can't find the place, because there is not %v.", before))
}

func (list *List) findPrev(val interface{}) *Node {
	prev := list.dummyHead
	for prev.next != nil && prev.next.Val != val {
		prev = prev.next
	}
	if prev.next == nil {
		return nil
	}
	return prev
}

func (list *List) Remove(val interface{}) error {
	if prev := list.findPrev(val); prev != nil {
		prev.next, prev.next.next = prev.next.next, nil
		list.len--
		return nil
	}
	return errors.New("there is not this value in list")
}

func (list *List) Update(old, new interface{}) error {
	if prev := list.findPrev(old); prev != nil {
		prev.next.Val = new
	}
	return errors.New(fmt.Sprintf("there is not the old value: %v", old))
}
