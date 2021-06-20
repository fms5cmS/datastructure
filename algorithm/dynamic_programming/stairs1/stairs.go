package stairs1

// 画出递归求解的递归树
// 节点数即为计算次数，
// 所以时间复杂度为 O(2^n)
func Stairs(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return Stairs(n-1) + Stairs(n-2)
}
