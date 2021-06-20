package linkedlist

import (
	"fmt"
	"strings"
)

type Node struct {
	Value interface{}
	next  *Node
}

func (node *Node) String() string {
	return fmt.Sprint(node.Value)
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) String() string {
	cur := l.head
	ret := strings.Builder{}
	for cur != nil {
		ret.WriteString(cur.String())
		if cur.next != nil {
			ret.WriteString(" -> ")
		}
		cur = cur.next
	}
	return ret.String()
}

// 链表的反转操作，pre 代表了反转后链表的尾节点。O(N) 复杂度
// 如果想直接反转链表，传入 pre 为 nil 即可
func (l *LinkedList) Reverse(pre *Node) {
	if l.len <= 1 {
		return
	}
	for l.head != nil {
		temp := l.head.next
		l.head.next = pre
		pre = l.head
		l.head = temp
	}
	l.head = pre
}

// 找到链表的中间节点，共计偶数个节点时返回前一个中间节点。O(N) 时间复杂度
// 1—>2—>3—>4—>5  返回值为 3 的节点
// 1—>2—>3—>4     返回值为 2 的节点
func (l *LinkedList) FindMiddleNode() *Node {
	if l.len < 2 {
		return nil
	}
	slow, fast := l.head, l.head
	// 共计偶数个节点时返回后一个中间节点时，修改条件为
	// for fast != nil && fast.next != nil{
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

// 已知链表的长度，删除链表的倒数第 n 个节点。O(N)复杂度
// 如果 n 大于链表的长度，返回 nil
func (l *LinkedList) DeleteBottomNWithLength(n int) *Node {
	index := l.len - n
	if index < 0 || index == l.len {
		return nil
	}
	// 删除头节点需要单独判断
	if index == 0 {
		temp := l.head
		l.head = temp.next
		l.len--
		return temp
	}
	// 1.获取要删除的节点的前驱结点
	i := 0
	pre := l.head
	for i+1 != index {
		pre = pre.next
		i++
	}
	// 2.删除节点
	temp := pre.next
	pre.next = temp.next
	temp.next = nil
	l.len--
	return temp
}

// 在不遍历链表就无法得到链表长度时，删除链表的倒数第 n 个节点。O(N)复杂度
// 如果 n 大于链表的长度，返回 nil
func (l *LinkedList) DeleteBottomN(n int) *Node {
	if n <= 0 || l.head == nil {
		return nil
	}
	fast, slow := l.head, l.head
	// 1. 快指针先走 n 步
	for i := 0; i < n; i++ {
		// 删除头节点 head
		if fast.next == nil && i == n-1 {
			temp := l.head
			l.head = temp.next
			temp.next = nil
			return temp
			// 要删除的节点不存在，如链表共 2 个元素，但 n==3
		} else if fast.next == nil && i < n-1 {
			return nil
		}
		fast = fast.next
	}
	// 2. 快慢指针同时走，直到快指针到达尾节点
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	// 这里借助 fast 暂时存储要删除的节点，防止删除后链表断裂
	fast = slow.next
	slow.next = fast.next
	fast.next = nil
	return fast
}

// 合并两个有序链表，假设这两个链表都按从小到大的顺序排列
func MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if 0 == l1.len {
		return l2
	}
	if 0 == l2.len {
		return l1
	}

	cur1, cur2 := l1.head, l2.head
	l := &LinkedList{&Node{}, 0}
	cur := l.head
	for cur1 != nil && cur2 != nil {
		if cur1.Value.(int) > cur2.Value.(int) {
			cur.next = cur2
			cur2 = cur2.next
			l.len++
		} else {
			cur.next = cur1
			cur1 = cur1.next
			l.len++
		}
		cur = cur.next
	}
	if cur1 == nil {
		cur.next = cur2
		for cur2 != nil {
			l.len++
			cur2 = cur2.next
		}
	}
	if cur2 == nil {
		cur.next = cur1
		for cur1 != nil {
			l.len++
			cur1 = cur1.next
		}
	}
	l.head = l.head.next
	return l
}
