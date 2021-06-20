package bplus

// B+ 树非叶子节点
// 假设 keywords = {3,5,8,10}
// 4 个键值将数据分为 5 个区间：(-INF, 3), (3, 5), (5, 8), (8, 10), (10, INF)
// 5 个区间分别对应 children[0]...children[4]
// m 是事先计算得到的，其依据是让所有信息的大小正好等于页的大小
// PAGE_SIZE=(m-1)*4[keywords 大小] + m*8[children 大小]
type TreeNode struct {
	m        int         // m 叉树
	keywords []int       // 键值，用于划分数据区间
	children []*TreeNode // 保存子节点指针
}

// 叶子节点存储的是值，而不是区间
// 这里每个叶子节点存储 3 个数据行的键值和地址信息
// k 是事先计算得到的，其依据是让所有信息的大小正好等于页的大小
// PAGE_SIZE=k*4[keywords 大小]+k*8[dataAddress 大小]+8[prev 大小]+8[next 大小]
type TreeLeafNode struct {
	k           int
	keywords    []int        // 数据键值
	dataAddress []int        // 数据地址
	prev, next  TreeLeafNode // 该节点在链表种的前驱、后继节点
}
