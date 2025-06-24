package main

import (
	"fmt"
	"sort"
)

/*
图例说明：

1. 快速选择算法过程示意：
   原始数组：[3,2,1,5,6,4], k=2

   第一次分区：
   pivot=3
   [2,1,3,5,6,4]
    ^     ^
   left  pivot

   第二次分区：
   pivot=5
   [2,1,3,4,5,6]
          ^  ^
       left pivot

2. 堆排序过程示意：
   原始数组：[3,2,1,5,6,4], k=2

   构建最大堆：
       6
      / \
     5   4
    / \
   2   3
  /
 1

   第一次弹出：6
   第二次弹出：5 (第2大的元素)

3. 排序过程示意：
   原始：[3,2,1,5,6,4]
   排序：[1,2,3,4,5,6]
   第k大：5 (k=2)
*/

// 解法一：快速选择算法（基于快速排序的分区思想）
// 时间复杂度：平均 O(n)，最坏 O(n²)
// 空间复杂度：O(1)
func findKthLargest1(nums []int, k int) int {
	// 将k转换为第k小的元素（因为我们要找第k大的）
	k = len(nums) - k
	left, right := 0, len(nums)-1

	for left <= right {
		// 获取分区点
		pivot := partition(nums, left, right)

		if pivot == k {
			// 找到第k小的元素
			return nums[pivot]
		} else if pivot < k {
			// 在右半部分继续查找
			left = pivot + 1
		} else {
			// 在左半部分继续查找
			right = pivot - 1
		}
	}
	return -1
}

// 分区函数
func partition(nums []int, left, right int) int {
	// 选择最右边的元素作为基准
	pivot := nums[right]
	// i表示小于基准值的区域边界
	i := left

	// 遍历数组，将小于基准值的元素移到左边
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	// 将基准值放到正确的位置
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// 解法二：堆排序
// 时间复杂度：O(nlogk)
// 空间复杂度：O(k)
func findKthLargest2(nums []int, k int) int {
	// 构建小顶堆
	heap := make([]int, k)
	copy(heap, nums[:k])

	// 调整堆
	for i := k/2 - 1; i >= 0; i-- {
		heapify(heap, i, k)
	}

	// 遍历剩余元素
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			heapify(heap, 0, k)
		}
	}

	return heap[0]
}

// 堆化函数
func heapify(nums []int, i, size int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < size && nums[left] < nums[smallest] {
		smallest = left
	}
	if right < size && nums[right] < nums[smallest] {
		smallest = right
	}

	if smallest != i {
		nums[i], nums[smallest] = nums[smallest], nums[i]
		heapify(nums, smallest, size)
	}
}

// 解法三：排序后直接取（最简单但效率较低）
// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)
func findKthLargest3(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 解法四：极简写法（使用Go的sort包）
// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)
func findKthLargestSimple(nums []int, k int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums[k-1]
}

func main() {
	// 测试用例
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2},          // 预期输出：5
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4}, // 预期输出：4
		{[]int{1}, 1},                         // 预期输出：1
	}

	for _, tc := range testCases {
		fmt.Printf("\n输入: nums = %v, k = %d\n", tc.nums, tc.k)

		// 解法一
		nums1 := make([]int, len(tc.nums))
		copy(nums1, tc.nums)
		fmt.Printf("解法一（快速选择）结果: %d\n", findKthLargest1(nums1, tc.k))

		// 解法二
		nums2 := make([]int, len(tc.nums))
		copy(nums2, tc.nums)
		fmt.Printf("解法二（堆排序）结果: %d\n", findKthLargest2(nums2, tc.k))

		// 解法三
		nums3 := make([]int, len(tc.nums))
		copy(nums3, tc.nums)
		fmt.Printf("解法三（排序）结果: %d\n", findKthLargest3(nums3, tc.k))

		// 解法四
		nums4 := make([]int, len(tc.nums))
		copy(nums4, tc.nums)
		fmt.Printf("解法四（极简）结果: %d\n", findKthLargestSimple(nums4, tc.k))
	}
}
