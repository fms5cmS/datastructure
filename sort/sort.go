package sort

// 冒泡排序
// 时间复杂度：O(n^2)
// 原地排序，空间复杂度为 O(1)，具有稳定性
func BubbleSort(nums []int, len int) {
	if nums == nil || len < 2 {
		return
	}
	for i := len - 1; i >= 1; i-- {
		// 假定未排序区间内元素有序
		sorted := true
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				// 发生了交换，则说明未排序区间内元素无序
				sorted = false
			}
		}
		if sorted {
			break // 未排序区间元素有序，所有元素已排序完成，退出循环
		}
	}
}

// 插入排序
// 时间复杂度 O(n^2)
// 原地排序，空间复杂度 O(1)。不具有稳定性
func InsertionSort(nums []int, len int) {
	if nums == nil || len < 2 {
		return
	}
	// [0,i)是有序区
	// 无序区的第一个元素依次与有序区的元素从后往前比较
	for i := 1; i < len; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 选择排序
// 时间复杂度为 O(n^2)
// 原地排序，空间复杂度为 O(1)，具有稳定性
func SelectionSort(nums []int, len int) {
	if nums == nil || len < 2 {
		return
	}
	for i := 0; i < len-1; i++ {
		// 使用一个第三变量保存未排序范围内最小值的索引，每次赋初值为 i
		minIndex := i
		// 查找 [i, len-1] 范围内最小值的索引
		for j := i + 1; j < len; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		// 每次移动i，最多仅发生一次交换操作
		nums[minIndex], nums[i] = nums[i], nums[minIndex]
	}
}

func MergeSort(nums []int, len int) {
	if nums == nil || len < 2 {
		return
	}
	mergeSort(nums, 0, len-1)
}

func mergeSort(nums []int, start, end int) {
	if start >= end { // 注意这里的结束条件
		return
	}
	middle := start + (end-start)>>1
	mergeSort(nums, start, middle)
	mergeSort(nums, middle+1, end)
	merge(nums, start, middle, end)
}

func merge(nums []int, start, middle, end int) {
	temp := make([]int, end-start+1)
	pointA, pointB := start, middle+1
	i := 0
	for pointA <= middle && pointB <= end {
		if nums[pointA] <= nums[pointB] {
			temp[i] = nums[pointA]
			pointA++
		} else {
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
	copy(nums[start:end+1], temp) // 由于是左闭右开区间，所以是 [start:end+1]
}

func QuickSort(nums []int, len int) {
	if nums == nil || len < 2 {
		return
	}
	quickSort(nums, 0, len-1)
}

func quickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	q := partition(nums, low, high)
	quickSort(nums, low, q-1)
	quickSort(nums, q+1, high)
}

// 获取分区点的索引
func partition(nums []int, low, high int) int {
	pivot := nums[high] // 以最后一个元素作为分区点
	var i = low
	// j 从左往右开始移动，当 nums[j] < pivot 时，交换索引 i 和 j 各自对应的值
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	// 此时索引 i 对应的值就是从左往右第一个大于 pivot 的值
	// 交换 i 和 high(即 pivot) 对应的值，则 i 就是分区点的索引
	nums[i], nums[high] = nums[high], nums[i]
	return i
}

func CountingSort(nums []int, len int) {
	if nums == nil || len < 2 {
		return
	}
	// 获取切片中的最大值
	max := nums[0]
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	temp := make([]int, max+1)
	for i := range nums {
		temp[nums[i]]++
	}
	for i := 1; i <= max; i++ {
		temp[i] += temp[i-1]
	}
	ret := make([]int, len)
	for i := len - 1; i >= 0; i-- {
		index := temp[nums[i]] - 1
		ret[index] = nums[i]
		temp[nums[i]]--
	}
	copy(nums, ret)
}
