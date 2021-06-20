package binary_search

func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	// 1.这里的循环退出条件必须是 low<=high
	for low <= high {
		// 2.这里不写成 (low+high)/2 是为了防止 low+high 溢出，使用位运算符可提高性能
		// 注意，Go 语言中位运算符的优先级和乘除运算符是同级的！！
		mid := low + (high-low)>>1
		// 3.注意 high、low 更新的值，并不是等于 mid
		// 否则，当 high=low=mid 时，就会陷入死循环
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func BinarySearchRecursive(nums []int, target int) int {
	return bSearch(nums, 0, len(nums)-1, target)
}

func bSearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)>>1
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return bSearch(nums, low, mid-1, target)
	} else {
		return bSearch(nums, mid+1, high, target)
	}
}

// 查找第一个等于给定值的元素
func FirstElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			high = mid - 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找最后一个等于给定值的元素
func LastElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func FirstGreaterOrEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 查找最后一个小于等于给定值的元素
func LastLessOrEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] <= target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
