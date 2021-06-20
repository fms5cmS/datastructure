package backtracking

import (
	"log"
)

type Knapsack struct {
	// 存储当前背包中物品总重量的最大值
	maxW uint
	// 当前已装进去的物品的总重量
	cw uint
	// 背包承受的最大重量
	w uint
	// 装入的物品各自重量
	weight []uint
}

// 对共有 n 个物品(重量分别为 items) 的第 i 个物品进行添加判断
func (k *Knapsack) add(i, n uint, items []uint) {
	if k.cw == k.w || i == n {
		if k.cw > k.maxW {
			k.maxW = k.cw
		}
		log.Printf("%+v", k)
		return
	}
	// 当前物品不装进背包
	k.add(i+1, n, items)
	// 是否超过了背包承受的重量
	if k.cw+items[i] <= k.w {
		// 没超过，将当前物品装进背包
		k.weight[i] = items[i]
		k.cw += items[i]
		k.add(i+1, n, items)
	}
}
