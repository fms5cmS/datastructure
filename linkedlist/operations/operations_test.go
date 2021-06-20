package operations

import (
	"testing"
)

var (
	evenList *Node
	oddList  *Node
)

func TestMain(m *testing.M) {
	evenList = &Node{Val: "first",
		next: &Node{Val: 2,
			next: &Node{Val: "third",
				next: &Node{Val: 4}}}}
	oddList = &Node{Val: 1,
		next: &Node{Val: "second",
			next: &Node{Val: 3,
				next: &Node{Val: "fourth",
					next: &Node{Val: 5}}}}}
	m.Run()
}

func TestReverse(t *testing.T) {
	t.Log(evenList)
	evenList = Reverse(evenList)
	t.Log(evenList)
}

func TestFindMiddleNode(t *testing.T) {
	t.Logf("奇数个节点链表: %s\n 中间节点值: %v", oddList, FindMiddleNode(oddList).Val)
	t.Logf("偶数个节点链表: %s\n 中间节点值: %v", evenList, FindMiddleNode(evenList).Val)
}

func TestRemoveNthFromEnd(t *testing.T) {
	// 删除之前：1 -> second -> 3 -> fourth -> 5
	t.Logf("删除之前：%s", oddList)
	// 删除倒数第二个节点后：1 -> second -> 3 -> 5
	oddList = RemoveNthFromEnd(oddList, 2)
	t.Logf("删除倒数第二个节点后：%s", oddList)
	// 删除倒数第三个节点后：1 -> 3 -> 5
	oddList = RemoveNthFromEnd(oddList, 3)
	t.Logf("删除倒数第三个节点后：%s", oddList)
}
