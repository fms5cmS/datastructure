package tree

type BST struct {
	root *Node
}

func NewBST(data interface{}) *BST {
	return &BST{root: NewTreeNode(data)}
}

// 比较函数，这里假设二叉查找树中的值都可以转为 int 类型
// val 表示要比较的值，cur 表示当点节点的值
// 如果要比较的值大，返回 1，小返回 -1，相等返回 0
func (bst BST) compare(val, cur interface{}) int {
	v := val.(int)
	curV := cur.(int)
	if v > curV {
		return 1
	} else if v < curV {
		return -1
	}
	return 0
}

// 插入操作
func (bst *BST) Insert(v interface{}) {
	bst.root = bst.insert(bst.root, v)
}

func (bst *BST) insert(node *Node, v interface{}) *Node {
	if node == nil {
		return NewTreeNode(v)
	}
	compareResult := bst.compare(v, node.data)
	if compareResult > 0 {
		node.right = bst.insert(node.right, v)
	} else if compareResult < 0 {
		node.left = bst.insert(node.left, v)
	}
	return node
}

// 查找值为 v 的节点
func (bst *BST) Find(v interface{}) *Node {
	cur := bst.root
	for cur != nil {
		compareResult := bst.compare(v, cur.data)
		if compareResult < 0 {
			cur = cur.left
		} else if compareResult > 0 {
			cur = cur.right
		} else {
			return cur
		}
	}
	return nil
}

// 获取树中最小值对应的节点
func (bst *BST) Minimum() *Node {
	return bst.minimum(bst.root)
}

func (bst *BST) minimum(node *Node) *Node {
	for node.left != nil {
		node = node.left
	}
	return node
	// 递归写法
	// if node.left == nil{
	//	return node
	// }
	// return bst.minimum(node.left)
}

// 移除并返回树中的最小值对应的节点
func (bst *BST) RemoveMin() *Node {
	min := bst.Minimum()
	bst.root = bst.removeMin(bst.root)
	return min
}

func (bst *BST) removeMin(node *Node) *Node {
	if node.left == nil {
		right := node.right
		node.right = nil
		return right
	}
	node.left = bst.removeMin(node.left)
	return node
}

func (bst *BST) Remove(v interface{}) {
	bst.root = bst.remove(bst.root, v)
}

// 删除以 node 为根节点的树中值为 v 的节点，并返回删除后该树的根节点
func (bst *BST) remove(node *Node, v interface{}) *Node {
	if node == nil {
		return nil
	}
	compareResult := bst.compare(v, node.data)
	if compareResult < 0 {
		node.left = bst.remove(node.left, v)
		return node
	} else if compareResult > 0 {
		node.right = bst.remove(node.right, v)
		return node
	} else {
		// 如果待删除节点的左子树为 nil，直接删除即可
		if node.left == nil {
			right := node.right
			node.right = nil
			return right
		}
		// 如果待删除节点的右子树为 nil，直接删除即可
		if node.right == nil {
			left := node.left
			node.left = nil
			return left
		}
		// 待删除节点的左右子节点均不为 nil
		// 找到右子树的最小值所在节点
		successor := bst.minimum(node.right)
		successor.right = bst.removeMin(node.right)
		successor.left = node.left
		// 将待删除的节点脱离
		node.left = nil
		node.right = nil
		node = successor
		return node
	}
}

// 非递归删除
func (bst *BST) Delete(v interface{}) {
	cur := bst.root
	var curP *Node // cur 的父节点
	// 找到待删除的节点 cur
	compareResult := bst.compare(v, cur.data)
	for cur != nil && compareResult != 0 {
		curP = cur
		if compareResult > 0 {
			cur = cur.right
		} else {
			cur = cur.left
		}
		compareResult = bst.compare(v, cur.data)
	}
	if cur == nil {
		return
	}
	// 此时 cur 即为待删除的节点
	// 待删除的节点有两个子节点
	if cur.left != nil && cur.right != nil {
		// 找到右子树的最小节点
		min := cur.right
		minP := cur // min 的父节点
		for min.left != nil {
			minP = min
			min = min.left
		}
		// 注意，这里并没有交换节点，而仅替换了值
		cur.data = min.data
		// 此时待删除的节点就是之前右子树的最小节点
		cur = min
		curP = minP
	}

	// 待删除的节点是叶子节点或仅有一个子节点
	var child *Node
	if cur.left != nil {
		child = cur.left
	} else if cur.right != nil {
		child = cur.right
	} else {
		child = nil
	}

	if curP == nil {
		// 这种情况是根节点只有左子树，且删除的是根节点
		bst.root = child
	} else if curP.left == cur {
		curP.left = child
	} else {
		curP.right = child
	}
}
