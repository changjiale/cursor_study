package main

import "container/heap"

type MaxHeap1 []int

func (h MaxHeap1) Len() int {
	return len(h)
}

func (h MaxHeap1) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap1) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topK(nums []int, k int) int {
	h := &MaxHeap1{}
	heap.Init(h)

	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}

	// 弹出前k-1个最大元素
	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	return heap.Pop(h).(int)
}
