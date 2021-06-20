package stairs3

// f(n)=f(n-1)+f(n-2)
// 每次迭代的结果仅依赖前两次的状态
// 只要保留前两次的状态就可推导出新的状态
// 不需要想备忘录算法那样保留所有的子状态
// 时间复杂度 O(n),空间复杂度 O(1)
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
	// 用于保存前两个子状态
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
