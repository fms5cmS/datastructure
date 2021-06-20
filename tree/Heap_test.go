package tree

import (
	"github.com/google/go-cmp/cmp"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMaxHeap_Insert(t *testing.T) {
	heap := NewHeap()
	arr := []int{53, 62, 11, 71, -24, 39, -4, 5, -63, 51}
	for _, v := range arr {
		heap.Insert(v)
	}
	t.Logf("%+v", heap)
	heap.RemoveMax()
	t.Logf("%+v", heap)
}

func generateArr(maxSize, maxValue int) (arr []int) {
	rand.Seed(time.Now().UnixNano())
	arr = make([]int, rand.Intn(maxSize)+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue)
	}
	return
}

func TestMaxHeap_Heapify(t *testing.T) {
	arr := generateArr(10, 100)
	t.Log(arr)
	t.Logf("%+v", Heapify(arr))
}

func TestHeapSort(t *testing.T) {
	arr := generateArr(100, 100)
	compare := make([]int, len(arr))
	copy(compare, arr)
	HeapSort(arr)
	sort.Ints(compare)
	if diff := cmp.Diff(arr, compare); diff != "" {
		t.Fatalf(diff)
	}
}
