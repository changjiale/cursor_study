package main

import "fmt"

/*
题目：搜索旋转排序数组
难度：中等
标签：数组、二分查找

题目描述：
整数数组 nums 按升序排列，数组中的值互不相同。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了旋转，
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标从 0 开始计数）。

例如，[0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2]。
给你旋转后的数组 nums 和一个整数 target，如果 nums 中存在这个目标值 target，则返回它的下标，否则返回 -1。

要求：
- 时间复杂度：O(log n)
- 空间复杂度：O(1)

示例：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

输入：nums = [1], target = 0
输出：-1
*/

// 解法一：二分查找（推荐）
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// 判断左半部分是否有序
		if nums[left] <= nums[mid] {
			// 左半部分有序
			if nums[left] <= target && target < nums[mid] {
				// target在左半部分
				right = mid - 1
			} else {
				// target在右半部分
				left = mid + 1
			}
		} else {
			// 右半部分有序
			if nums[mid] < target && target <= nums[right] {
				// target在右半部分
				left = mid + 1
			} else {
				// target在左半部分
				right = mid - 1
			}
		}
	}

	return -1
}

// 解法二：先找旋转点，再二分查找
// 时间复杂度：O(log n)
// 空间复杂度：O(1)
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 找到旋转点（最小值的位置）
	pivot := findPivot(nums)

	// 根据target和nums[0]的关系决定在哪个区间搜索
	if pivot == 0 {
		// 数组没有旋转，直接二分查找
		return binarySearch(nums, target, 0, len(nums)-1)
	}

	if target >= nums[0] {
		// target在左半部分
		return binarySearch(nums, target, 0, pivot-1)
	} else {
		// target在右半部分
		return binarySearch(nums, target, pivot, len(nums)-1)
	}
}

// 找到旋转点（最小值的位置）
func findPivot(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// 旋转点在右半部分
			left = mid + 1
		} else {
			// 旋转点在左半部分或当前位置
			right = mid
		}
	}

	return left
}

// 标准二分查找
func binarySearch(nums []int, target, left, right int) int {
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// 解法三：暴力解法（不推荐，仅用于理解）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func search3(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	// 测试用例
	testCases := []struct {
		nums   []int
		target int
		expect int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{1, 3}, 3, 1},
		{[]int{3, 1}, 1, 1},
		{[]int{5, 1, 3}, 3, 2},
		{[]int{5, 1, 3}, 5, 0},
		{[]int{5, 1, 3}, 1, 1},
		{[]int{5, 1, 3}, 0, -1},
	}

	fmt.Println("=== 解法一：二分查找 ===")
	for i, tc := range testCases {
		result := search(tc.nums, tc.target)
		fmt.Printf("测试用例 %d: nums=%v, target=%d, 结果=%d, 期望=%d, 通过=%t\n",
			i+1, tc.nums, tc.target, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法二：先找旋转点再二分 ===")
	for i, tc := range testCases {
		result := search2(tc.nums, tc.target)
		fmt.Printf("测试用例 %d: nums=%v, target=%d, 结果=%d, 期望=%d, 通过=%t\n",
			i+1, tc.nums, tc.target, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法三：暴力解法 ===")
	for i, tc := range testCases {
		result := search3(tc.nums, tc.target)
		fmt.Printf("测试用例 %d: nums=%v, target=%d, 结果=%d, 期望=%d, 通过=%t\n",
			i+1, tc.nums, tc.target, result, tc.expect, result == tc.expect)
	}
}

/*
解题思路：

1. 二分查找法（推荐）：
   - 虽然数组被旋转了，但仍然可以使用二分查找
   - 关键是要判断哪一半是有序的
   - 如果 nums[left] <= nums[mid]，说明左半部分有序
   - 如果 nums[mid] <= nums[right]，说明右半部分有序
   - 根据target的值决定在哪个区间继续搜索

2. 先找旋转点再二分：
   - 先找到旋转点（最小值的位置）
   - 根据target和nums[0]的关系决定搜索区间
   - 在确定的区间内进行标准二分查找

3. 暴力解法：
   - 直接遍历数组查找目标值
   - 时间复杂度O(n)，不满足题目要求

时间复杂度分析：
- 解法一和二：O(log n)，每次都将搜索范围减半
- 解法三：O(n)，需要遍历整个数组

空间复杂度分析：
- 所有解法：O(1)，只使用了常数额外空间

关键点：
1. 理解旋转数组的特性：数组被分成两个有序部分
2. 二分查找的关键是判断哪一半是有序的
3. 根据target的值和有序部分的范围来决定搜索方向
*/
