package array_linkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	temp := head.Next
	for temp != nil {
		head.Next = pre
		pre = head
		head = temp
		temp = head.Next
	}
	head.Next = pre
	return head
}

func swapPairs_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := head.Next
	head.Next = swapPairs_1(head.Next.Next)
	result.Next = head
	return result
}

func swapPairs(head *ListNode) *ListNode {
	result := &ListNode{Next: head}
	pre := result
	for pre.Next != nil && pre.Next.Next != nil {
		a, b := pre.Next, pre.Next.Next
		pre.Next, b.Next, a.Next = b, a, b.Next
	}
	return result.Next
}
