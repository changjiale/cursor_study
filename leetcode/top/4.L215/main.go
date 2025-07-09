package main

import (
	"container/heap"
	"fmt"
)

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1:

输入: [3,2,1,5,6,4], k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4


提示：

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
**/

func main() {
	// 测试堆方法
	testHeapMethods()

	// 演示快速选择算法
	demonstrateQuickSelect()

	// 测试用例
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{[]int{1}, 1, 1},
		{[]int{2, 1}, 1, 2},
		{[]int{2, 1}, 2, 1},
	}

	fmt.Println("\n=== 测试结果 ===")
	for i, tc := range testCases {
		fmt.Printf("测试用例 %d:\n", i+1)
		fmt.Printf("  输入: nums = %v, k = %d\n", tc.nums, tc.k)

		// 复制数组用于测试
		nums1 := make([]int, len(tc.nums))
		nums2 := make([]int, len(tc.nums))
		copy(nums1, tc.nums)
		copy(nums2, tc.nums)

		// 快速选择解法
		result1 := findKthLargestQuickSelect(nums1, tc.k)
		fmt.Printf("  快速选择解法结果: %d\n", result1)

		// 堆解法
		result2 := findKthLargestHeap(nums2, tc.k)
		fmt.Printf("  堆解法结果: %d\n", result2)

		fmt.Printf("  期望结果: %d\n", tc.want)
		fmt.Printf("  快速选择解法: %s\n", checkResult(result1, tc.want))
		fmt.Printf("  堆解法: %s\n", checkResult(result2, tc.want))
		fmt.Println()
	}
}

func findQSTest1(nums []int, k int) int {

	//todo
	target := len(nums) - k
	left := 0
	right := len(nums) - 1
	for left <= right {
		baseIndex := moveBase(nums, left, right)
		if baseIndex == target {
			return nums[baseIndex]
		} else if baseIndex < target {
			left = baseIndex + 1
		} else {
			right = baseIndex - 1
		}
	}
	return -1
}

func moveBase(nums []int, left, right int) int {
	base := nums[right]
	i := left - 1
	for j := i; j < right; j++ {
		if nums[j] <= base {
			//移动
			i++
			nums[i], nums[j] = nums[j], nums[i] // 交换元素
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

// 解法1：基于快排思想的快速选择算法
// 时间复杂度：平均 O(n)，最坏 O(n²)
// 空间复杂度：O(1)
func findKthLargestQuickSelect(nums []int, k int) int {
	// 关键理解：快排分区后，基准元素的位置表示"第几个最小元素"
	// 例如：数组 [3,2,1,5,6,4] 分区后变成 [1,2,3,4,5,6]，基准元素4在位置3
	// 这表示4是第4个最小元素（从0开始数）

	// 我们要找第k个最大元素，需要转换为"第几个最小元素"
	// 第k个最大元素 = 第(n-k+1)个最小元素
	// 例如：n=6, k=2时，第2个最大元素 = 第(6-2+1)=第5个最小元素
	// 注意：数组索引从0开始，所以位置 = (n-k+1) - 1 = n-k
	target := len(nums) - k
	left, right := 0, len(nums)-1

	for left <= right {
		// 对当前区间进行分区，返回基准元素的最终位置
		pivotIndex := partition(nums, left, right)

		// 如果基准元素的位置正好是我们要找的目标位置
		if pivotIndex == target {
			// 找到了！返回这个位置的元素
			return nums[pivotIndex]
		} else if pivotIndex < target {
			// 基准元素位置偏左，说明目标元素在右半部分
			// 例如：基准在位置2，目标在位置5，需要在右半部分继续找
			left = pivotIndex + 1
		} else {
			// 基准元素位置偏右，说明目标元素在左半部分
			// 例如：基准在位置5，目标在位置2，需要在左半部分继续找
			right = pivotIndex - 1
		}
	}
	return -1
}

// 分区函数，返回基准元素的最终位置
// 这个函数会将数组分为三部分：[小于基准] [基准] [大于基准]
func partition(nums []int, left, right int) int {
	// 选择最右边的元素作为基准（也可以选择其他位置的元素）
	pivot := nums[right]
	i := left - 1 // i指向小于基准区域的最后一个位置

	// 遍历[left, right-1]区间，将小于等于基准的元素移到左边
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			i++                                 // 扩展小于基准的区域
			nums[i], nums[j] = nums[j], nums[i] // 交换元素
		}
	}

	// 将基准元素放到正确位置（小于基准区域的右边）
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1 // 返回基准元素的最终位置
}

// 解法2：基于大根堆的解法
// 时间复杂度：O(n + k*log n)
// 空间复杂度：O(n)
func findKthLargestHeap(nums []int, k int) int {
	// 构建大根堆
	h := &MaxHeap{}
	heap.Init(h)

	// 将所有元素加入堆
	for _, num := range nums {
		heap.Push(h, num)
	}

	// 弹出前k-1个最大元素
	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	// 返回第k个最大元素
	return heap.Pop(h).(int)
}

// 大根堆实现
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 大根堆
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 演示快速选择算法的执行过程
func demonstrateQuickSelect() {
	fmt.Println("\n=== 快速选择算法演示 ===")
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Printf("原数组: %v, 找第%d个最大元素\n", nums, k)

	// 复制数组用于演示
	demoNums := make([]int, len(nums))
	copy(demoNums, nums)

	// 计算目标位置
	n := len(demoNums)
	target := n - k
	fmt.Printf("第%d个最大元素 = 第%d个最小元素 (位置%d)\n", k, target+1, target)
	fmt.Printf("计算过程: 第%d个最大元素 = 第(%d-%d+1)=第%d个最小元素 = 位置%d\n", k, n, k, target+1, target)

	// 演示分区过程
	fmt.Println("\n分区过程演示:")
	left, right := 0, len(demoNums)-1
	step := 1

	for left <= right {
		fmt.Printf("步骤%d: 区间[%d,%d], 数组: %v\n", step, left, right, demoNums)

		pivotIndex := partition(demoNums, left, right)
		fmt.Printf("  分区后基准位置: %d, 数组: %v\n", pivotIndex, demoNums)

		if pivotIndex == target {
			fmt.Printf("  ✅ 找到目标! 位置%d的元素是%d\n", pivotIndex, demoNums[pivotIndex])
			break
		} else if pivotIndex < target {
			fmt.Printf("  基准位置%d < 目标位置%d, 在右半部分继续查找\n", pivotIndex, target)
			left = pivotIndex + 1
		} else {
			fmt.Printf("  基准位置%d > 目标位置%d, 在左半部分继续查找\n", pivotIndex, target)
			right = pivotIndex - 1
		}
		step++
	}
}

// 测试堆方法是否被调用
func testHeapMethods() {
	fmt.Println("\n=== 测试堆方法调用 ===")

	// 创建一个测试堆
	h := &MaxHeap{}

	fmt.Println("1. 初始化堆...")
	heap.Init(h)

	fmt.Println("2. 添加元素...")
	heap.Push(h, 5)
	heap.Push(h, 3)
	heap.Push(h, 7)
	heap.Push(h, 1)

	fmt.Println("3. 弹出元素...")
	for h.Len() > 0 {
		val := heap.Pop(h).(int)
		fmt.Printf("   弹出: %d\n", val)
	}
}

func checkResult(result, want int) string {
	if result == want {
		return "✅ 正确"
	}
	return "❌ 错误"
}
