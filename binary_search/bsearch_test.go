package binary_search

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func generateSortedNums(size, max int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, rand.Intn(size)+1)
	for i := range nums {
		nums[i] = rand.Intn(max+1) - rand.Intn(max)
	}
	sort.Ints(nums)
	return nums
}

func TestBinarySearch(t *testing.T) {
	for i := 0; i < 1000; i++ {
		nums := generateSortedNums(15, 100)
		target := nums[rand.Intn(len(nums))]
		index := BinarySearch(nums, target)
		if nums[index] != target {
			t.Errorf("%v, the target %d in %d", nums, target, index)

		}
		time.Sleep(time.Nanosecond)
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	for i := 0; i < 10000; i++ {
		nums := generateSortedNums(15, 100)
		target := nums[rand.Intn(len(nums))]
		index := BinarySearchRecursive(nums, target)
		if nums[index] != target {
			t.Errorf("%v, the target %d in %d", nums, target, index)

		}
		time.Sleep(time.Nanosecond)
	}
}

func TestElement(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 7, 7, 7, 7, 8, 9}
	t.Log(FirstElement(nums, 7))
	t.Log(LastElement(nums, 7))
	t.Log(FirstGreaterOrEqualElement(nums, 7))
	t.Log(LastLessOrEqualElement(nums, 7))
}
