package tree

import (
	"awesomeProject/stack/myStack"
	"fmt"
)

type Node struct {
	data        interface{}
	left, right *Node
}

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.data)
}

func NewTreeNode(data interface{}) *Node {
	return &Node{data: data}
}

func PreOrderRecursion(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root, " ")
	PreOrderRecursion(root.left)
	PreOrderRecursion(root.right)
}

// 前序遍历的非递归实现

// 相当于用自己的栈模拟递归写法中系统栈的作用
func PreOrderNotRecursion(root *Node) {
	stack := myStack.NewStack()
	stack.Push(root)
	for stack.Size() != 0 {
		cur := stack.Pop().(*Node)
		fmt.Print(cur, " ")
		// 由于栈是 LIFO，所以先将右子树压栈，再将左子树压栈
		if cur.right != nil {
			stack.Push(cur.right)
		}
		if cur.left != nil {
			stack.Push(cur.left)
		}
	}
}

// 使用数组模拟栈
func PreOrderNonRecursion(root *Node) {
	stack := make([]*Node, 0)
	stack = append(stack, root) // 根节点入栈
	for len(stack) > 0 {
		length := len(stack) - 1
		cur := stack[length]   // 栈顶元素出栈
		stack = stack[:length] // 更新栈的长度
		fmt.Print(cur, " ")
		if cur.right != nil {
			stack = append(stack, cur.right)
		}
		if cur.left != nil {
			stack = append(stack, cur.left)
		}
	}
}

func InOrderRecursion(root *Node) {
	if root == nil {
		return
	}
	InOrderRecursion(root.left)
	fmt.Print(root, " ")
	InOrderRecursion(root.right)
}

func PostOrderRecursion(root *Node) {
	if root == nil {
		return
	}
	PostOrderRecursion(root.left)
	PostOrderRecursion(root.right)
	fmt.Print(root, " ")
}

// 层序遍历
func LevelOrder(root *Node) {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Print(cur, " ")
		if cur.left != nil {
			queue = append(queue, cur.left)
		}
		if cur.right != nil {
			queue = append(queue, cur.right)
		}
	}
}
