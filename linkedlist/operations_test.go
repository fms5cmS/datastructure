package linkedlist

import (
	"testing"
)

func TestLinkedList_String(t *testing.T) {
	var l LinkedList
	l.head = &Node{"first",
		&Node{"second",
			&Node{"third",
				&Node{"fourth", nil}}}}
	t.Log(l.String())
}

func TestLinkedList_Reverse(t *testing.T) {
	var l LinkedList
	l.head = &Node{"first",
		&Node{"second",
			&Node{"third",
				&Node{"fourth", nil}}}}
	l.len = 4 // 这里手动设置了链表的长度
	t.Log("反转前：", l.String())
	l.Reverse(nil)
	t.Log("反转后：", l.String())
}

func TestLinkedList_FindMiddleNode(t *testing.T) {
	var l LinkedList
	t.Log("偶数个元素")
	l.head = &Node{"first",
		&Node{"second",
			&Node{"third",
				&Node{"fourth", nil}}}}
	l.len = 4
	t.Log(l.String())
	t.Log("中间节点为：", l.FindMiddleNode())
	t.Log("奇数个元素")
	l.head = &Node{"first",
		&Node{"second",
			&Node{"third",
				&Node{"fourth",
					&Node{"fifth", nil}}}}}
	l.len = 5
	t.Log(l.String())
	t.Log("中间节点为：", l.FindMiddleNode())
}

func TestLinkedList_DDeleteBottomNWithLength(t *testing.T) {
	var l LinkedList
	l.head = &Node{"first",
		&Node{"second",
			&Node{"third",
				&Node{"fourth", nil}}}}
	l.len = 4
	t.Log(l.String())
	t.Log("the value of deleted node is:", l.DeleteBottomNWithLength(5))
	t.Log(l.String())
	l.head = &Node{"1", nil}
	l.len = 1
	t.Log("the value of deleted node is:", l.DeleteBottomNWithLength(1))
	t.Log(l.String())
}

func TestLinkedList_DeleteBottomN(t *testing.T) {
	var l LinkedList
	l.head = &Node{"first",
		&Node{"second",
			&Node{"third",
				&Node{"fourth", nil}}}}
	l.len = 4
	t.Log(l.String())
	t.Log("the value of deleted node is:", l.DeleteBottomN(3))
	t.Log(l.String())

	l.head = &Node{"first", nil}
	t.Log(l.String())
	t.Log("the value of deleted node is:", l.DeleteBottomN(1))
	t.Log(l.String())
}

func TestMergeSortedList(t *testing.T) {
	var l1 LinkedList
	l1.head = &Node{1,
		&Node{2,
			&Node{4,
				&Node{8, nil}}}}
	l1.len = 4
	var l2 LinkedList
	l2.head = &Node{3,
		&Node{6,
			&Node{7,
				&Node{9, nil}}}}
	l2.len = 4
	l := MergeSortedList(&l1, &l2)
	t.Log("l1:", l1.String())
	t.Log("l2:", l2.String())
	t.Log("result after merging:", l.String())
	t.Log("length of result:", l.len)
}
