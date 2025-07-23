package main

import "container/heap"

/*
*
再来一遍
*/
type MaxHead1 []int

func (h MaxHead1) Len() int {
	return len(h)
}
func (h MaxHead1) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHead1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHead1) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHead1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topK1(nums []int, k int) int {

	h := &MaxHeap1{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
	}

	for i := 0; i < k-1; k++ {
		heap.Pop(h)
	}

	return heap.Pop(h).(int)

}
