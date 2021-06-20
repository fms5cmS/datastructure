package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 21.将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	ret := l
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}
	if l1 == nil {
		l.Next = l2
	}
	if l2 == nil {
		l.Next = l1
	}
	return ret.Next
}

// 19.给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	// 1. 快指针先走 n 步
	for i := 0; i < n; i++ {
		// 删除头节点 head，如链表共有 n 个元素，删除倒数第 n 个
		if fast.Next == nil && i == n-1 {
			head, head.Next = head.Next, nil
			return head
		}
		// 如果保证 n 是有效的，可以不做这一步判断
		// 要删除的节点不存在，如链表共 2 个元素，但 n==3
		if fast.Next == nil && i < n-1 {
			return nil
		}
		fast = fast.Next
	}
	// 2. 快慢指针同时走，直到快指针到达尾节点
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 这里借助 fast 暂时存储要删除的节点，防止删除后链表断裂
	fast = slow.Next
	slow.Next = fast.Next
	fast.Next = nil
	return head
}

// 206.反转一个单链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

// 206.递归反转一个单链表
func reverseListWithRecursion1(head *ListNode) *ListNode {
	return reverse(head, nil)
}
func reverse(head, pre *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	temp := head.Next
	head.Next = pre
	return reverse(temp, head)
}

// 206.递归反转一个单链表
// 1->2->3->4->...-> k -> k+1<-k+2<-...<-n
// 假设现在处于 head == k，而 k 之后的节点都已经反转了
func reverseListWithRecursion2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 将 head 之后的节点进行反转，并返回反转后链表的头节点
	p := reverseList(head.Next)
	// (k+1).Next = k
	head.Next.Next = head
	// 防止相邻节点循环指向，这里先将 k.Next 置为 nil
	head.Next = nil
	return p
}

// 876.给定一个带有 head 的非空单链表，返回链表的中间结点。如果有两个中间结点，则返回第二个中间结点。
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
