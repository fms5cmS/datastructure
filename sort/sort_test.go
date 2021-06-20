package sort

import (
	"github.com/google/go-cmp/cmp"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func generateRandomSlice(maxSize, maxValue int) ([]int, []int) {
	rand.Seed(time.Now().UnixNano())
	nums1 := make([]int, rand.Intn(maxSize+1))
	nums2 := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		// 表达式后面的注释是因为计数排序只能对非负整数排序
		nums1[i] = rand.Intn(maxValue + 1) // - rand.Intn(maxValue)
	}
	copy(nums2, nums1)
	return nums1, nums2
}

func TestGenerateRandomSlice(t *testing.T) {
	for i := 0; i < 10; i++ {
		nums1, nums2 := generateRandomSlice(10, 99)
		t.Log(nums1, "\n", nums2)
		time.Sleep(time.Nanosecond)
	}
}

func TestBubbleSort(t *testing.T) {
	for i := 0; i < 1000; i++ {
		nums1, nums2 := generateRandomSlice(15, 99)
		BubbleSort(nums1, len(nums1))
		sort.Ints(nums2)
		diff := cmp.Diff(nums1, nums2)
		if diff != "" {
			t.Fatalf(diff)
		}
		time.Sleep(time.Nanosecond)
	}
}

func TestInsertionSort(t *testing.T) {
	for i := 0; i < 10000; i++ {
		nums1, nums2 := generateRandomSlice(200, 99)
		InsertionSort(nums1, len(nums1))
		sort.Ints(nums2)
		diff := cmp.Diff(nums1, nums2)
		if diff != "" {
			t.Fatalf(diff)
		}
		time.Sleep(time.Nanosecond)
	}
}

func TestSelectionSort(t *testing.T) {
	for i := 0; i < 10000; i++ {
		nums1, nums2 := generateRandomSlice(15, 99)
		SelectionSort(nums1, len(nums1))
		sort.Ints(nums2)
		diff := cmp.Diff(nums1, nums2)
		if diff != "" {
			t.Fatalf(diff)
		}
		time.Sleep(time.Nanosecond)
	}
}

func TestMergeSort(t *testing.T) {
	for i := 0; i < 10000; i++ {
		nums1, nums2 := generateRandomSlice(15, 99)
		MergeSort(nums1, len(nums1))
		sort.Ints(nums2)
		diff := cmp.Diff(nums1, nums2)
		if diff != "" {
			t.Fatalf(diff)
		}
		time.Sleep(time.Nanosecond)
	}
}

func TestQuickSort(t *testing.T) {
	for i := 0; i < 10000; i++ {
		nums1, nums2 := generateRandomSlice(15, 99)
		QuickSort(nums1, len(nums1))
		sort.Ints(nums2)
		diff := cmp.Diff(nums1, nums2)
		if diff != "" {
			t.Fatalf(diff)
		}
		time.Sleep(time.Nanosecond)
	}
}

// 计数排序只能对非负整数进行排序
func TestCountingSort(t *testing.T) {
	for i := 0; i < 10000; i++ {
		nums1, nums2 := generateRandomSlice(15, 99)
		CountingSort(nums1, len(nums1))
		sort.Ints(nums2)
		diff := cmp.Diff(nums1, nums2)
		if diff != "" {
			t.Fatalf(diff)
		}
		time.Sleep(time.Nanosecond)
	}
}
