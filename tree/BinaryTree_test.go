package tree

import (
	"testing"
)

//		   1
//		2	  3
//    4  5  6
func generateBinaryTree() *Node {
	root := NewTreeNode(1)
	root.left = NewTreeNode(2)
	root.right = NewTreeNode(3)
	root.left.left = NewTreeNode(4)
	root.left.right = NewTreeNode(5)
	root.right.left = NewTreeNode(6)
	return root
}

// 前序遍历，输出应为 1 2 4 5 3 6
func TestPreOrderRecursion(t *testing.T) {
	PreOrderRecursion(generateBinaryTree())
}

func TestPreOrderNotRecursion(t *testing.T) {
	PreOrderNotRecursion(generateBinaryTree())
	println()
	PreOrderNonRecursion(generateBinaryTree())
}

// 中序遍历，输出应为 4 2 5 1 6 3
func TestInOrderRecursion(t *testing.T) {
	InOrderRecursion(generateBinaryTree())
}

// 后序遍历，输出应为 4 5 2 6 3 1
func TestPostOrderRecursion(t *testing.T) {
	PostOrderRecursion(generateBinaryTree())
}

// 层序遍历，输出应为 1 2 3 4 5 6
func TestLevelOrder(t *testing.T) {
	LevelOrder(generateBinaryTree())
}
