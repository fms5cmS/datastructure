package tree

// 最大堆
type MaxHeap struct {
	size int // 堆中元素的数量
	data []int
}

func NewHeap() *MaxHeap {
	heap := &MaxHeap{size: 0, data: make([]int, 0)}
	return heap
}

func Heapify(arr []int) *MaxHeap {
	if len(arr) > 1 {
		heap := &MaxHeap{size: len(arr), data: arr}
		for i := heap.parent(len(arr) - 1); i >= 0; i-- {
			heap.siftDown(i)
		}
	}
	heap := &MaxHeap{size: len(arr), data: arr}
	return heap
}

func (heap *MaxHeap) parent(index int) int {
	if index == 0 {
		panic("index 0 does not have parent")
	}
	return (index - 1) / 2
}

func (heap *MaxHeap) leftChild(index int) int {
	return 2*index + 1
}

func (heap *MaxHeap) rightChild(index int) int {
	return 2*index + 2
}

func (heap *MaxHeap) swap(i, j int) {
	if i < 0 || i >= heap.size || j < 0 || j >= heap.size {
		panic("index is illegal")
	}
	heap.data[i], heap.data[j] = heap.data[j], heap.data[i]
}

// 对索引为 index 的节点进行上浮操作
func (heap *MaxHeap) siftUp(index int) {
	// index>0 说明该节点的父节点存在，且 index 所对应节点的值比其父节点的值大就进行上浮
	for index > 0 && heap.data[index] > heap.data[heap.parent(index)] {
		// 交换 index 和 indexP 的值
		heap.swap(index, heap.parent(index))
		// 更新 index 和 indexP 的值
		index = heap.parent(index)
	}
}

// 对索引为 index 的节点进行下沉操作
func (heap *MaxHeap) siftDown(index int) {
	// index 存在左节点
	for heap.leftChild(index) < heap.size {
		// max 用于保存 index 两个子节点中的最大值的索引
		max := heap.leftChild(index)
		// index 存在右子节点，且右子节点的值 > 左子节点的值
		if max+1 < heap.size && heap.data[max+1] > heap.data[max] {
			max = heap.rightChild(index)
		}
		// 如果 index 的值比其两个子节点的值都大，停止下沉
		if heap.data[index] >= heap.data[max] {
			break
		}
		// 将子节点的最大值与 index 的值交换
		heap.swap(index, max)
		// 更新 index，便于继续下沉
		index = max
	}
}

func (heap *MaxHeap) Insert(data int) {
	heap.data = append(heap.data, data)
	heap.size++
	heap.siftUp(heap.size - 1)
}

func (heap *MaxHeap) RemoveMax() int {
	if heap.size == 0 {
		panic("can't find max when heap is empty")
	}
	max := heap.data[0]
	heap.swap(0, heap.size-1)
	heap.data = heap.data[:heap.size-1]
	heap.size--
	heap.siftDown(0)
	return max
}

func HeapSort(arr []int) {
	heap := Heapify(arr)
	k := len(arr) - 1
	for k > 0 {
		heap.swap(0, k)
		heap = Heapify(arr[:k])
		k--
	}
}
