package stairs2

// 从之前的递归树可以发现有很多重复
// 的计算，借助备忘录算法优化
// 从 1 到 n 共计 n 次输入
// 哈希表中存了 n-2 个数据
// 所以时间复杂度和空间复杂度都是 O(n)
func Stairs(n int, mem map[int]int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if v, ok := mem[n]; ok {
		return v
	} else {
		v = Stairs(n-1, mem) + Stairs(n-2, mem)
		mem[n] = v
		return v
	}
}
