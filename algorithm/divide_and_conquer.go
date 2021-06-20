package algorithm

var num int // 全局变量记录逆序度

// 借用分治思想来计算数组的逆序度
func ReverseOrderByDivideAndConquer(nums []int) int {
	num = 0
	mergeSortCounting(nums, 0, len(nums)-1)
	return num
}

func mergeSortCounting(nums []int, start, end int) {
	if start >= end {
		return
	}
	middle := start + (end-start)>>1
	mergeSortCounting(nums, start, middle)
	mergeSortCounting(nums, middle+1, end)
	merge(nums, start, middle, end)
}

func merge(nums []int, start, middle, end int) {
	pointA, pointB := start, middle+1
	i := 0
	temp := make([]int, end-start+1)
	for pointA <= middle && pointB <= end {
		if nums[pointA] <= nums[pointB] {
			temp[i] = nums[pointA]
			pointA++
		} else {
			// 计算分开的两部分之间的逆序度
			num += middle - i + 1
			temp[i] = nums[pointB]
			pointB++
		}
		i++
	}
	for pointA <= middle {
		temp[i] = nums[pointA]
		i++
		pointA++
	}
	for pointB <= end {
		temp[i] = nums[pointB]
		i++
		pointB++
	}
	copy(nums[start:end+1], temp)
}
