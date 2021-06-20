package operations

// 反转链表
// 1 -> 2 -> 3 -> 4 反转后得到 4 -> 3 -> 2 -> 1
func Reverse(head *Node) *Node {
	var pre *Node
	for head != nil {
		temp := head.next
		head.next = pre
		pre, head = head, temp
	}
	return pre
}

// 找到链表的中间节点，共计偶数个节点时返回前一个中间节点。
// 1—>2—>3—>4—>5  返回值为 3 的节点
// 1—>2—>3—>4     返回值为 2 的节点
func FindMiddleNode(head *Node) *Node {
	fast, slow := head, head
	// 共计偶数个节点时返回后一个中间节点时，可修改条件为
	// for fast != nil && fast.next != nil{
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

func RemoveNthFromEnd(head *Node, n int) *Node {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for i := 0; i < n; i++ {
		// fast 走到末尾且 fast 共计走了 n-1（从头节点到尾节点所需步数）步，即要删除的为链表头节点
		if fast.next == nil && i == n-1 {
			// 删除头节点
			head, head.next = head.next, nil
			return head
		}
		// 如果保证 n 是有效的，可以不做这一步判断
		// 要删除的节点不存在，如链表共 2 个元素，但 n==3
		if fast.next == nil && i < n-1 {
			return nil
		}
		fast = fast.next
	}
	// 快慢指针同时移动
	for fast.next != nil {
		fast, slow = fast.next, slow.next
	}
	// 删除指定节点
	fast = slow.next
	slow.next, fast.next = fast.next, nil
	return head
}
