package backtracking

// 下标表示行, 值表示 queen 存储在哪一列
var result = [8]int{}

func Cal8Queens(row int) {
	if row == 8 {
		// 8 个棋子都放置好了，打印结果
		printQueens(result)
		return
	}
	// 每一行都有 8 种放法
	for column := 0; column < 8; column++ {
		// 有些放法不满足要求
		if isOk(row, column) {
			// 第 row 行的棋子放到了 column 列
			result[row] = column
			// 考察下一行
			Cal8Queens(row + 1)
		}
	}
}

// 判断坐标为 (row, column) 的位置放置棋子是否合适
// 由于是从上往下放置棋子，故只需判断当前坐标的上方(正上方、左上、右上对角线)
func isOk(row, column int) bool {
	leftup, rightup := column-1, column+1
	// 逐行往上考察每一行
	for i := row - 1; i >= 0; i-- {
		// 当前坐标正上方 (row-1, column) 是否有棋子
		if result[i] == column {
			return false
		}
		// 当前坐标左上对角线 (row-1, column-1) 是否有棋子
		if leftup >= 0 {
			if result[i] == leftup {
				return false
			}
		}
		// 当前坐标右上对角线 (row-1, column-1) 是否有棋子
		if rightup < 8 {
			if result[i] == rightup {
				return false
			}
		}
		// 继续向左、右上对角线移动判断
		leftup--
		rightup++
	}
	return true
}

// 打印出一个二维矩阵
func printQueens(result [8]int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				print("Q ")
			} else {
				print("* ")
			}
		}
		println()
	}
	println()
}
