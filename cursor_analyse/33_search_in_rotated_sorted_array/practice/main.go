package main

import (
	"fmt"
	"strings"
)

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

提示：
1. 虽然数组被旋转了，但仍然可以使用二分查找
2. 关键是要判断哪一半是有序的
3. 如果 nums[left] <= nums[mid]，说明左半部分有序
4. 如果 nums[mid] <= nums[right]，说明右半部分有序
5. 根据target的值决定在哪个区间继续搜索
*/

// TODO: 在这里实现你的算法
func search(nums []int, target int) int {
	// 请实现你的代码
	return -1
}

// 测试用例结构
type TestCase struct {
	nums   []int
	target int
	expect int
}

func main() {
	// 测试用例
	testCases := []TestCase{
		// 基本测试用例
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},  // 目标值在右半部分
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1}, // 目标值不存在
		{[]int{1}, 0, -1},                   // 单个元素，目标值不存在
		{[]int{1}, 1, 0},                    // 单个元素，目标值存在

		// 边界测试用例
		{[]int{1, 3}, 3, 1},     // 两个元素，目标值在末尾
		{[]int{3, 1}, 1, 1},     // 两个元素，旋转后目标值在末尾
		{[]int{5, 1, 3}, 3, 2},  // 三个元素，目标值在末尾
		{[]int{5, 1, 3}, 5, 0},  // 三个元素，目标值在开头
		{[]int{5, 1, 3}, 1, 1},  // 三个元素，目标值在中间
		{[]int{5, 1, 3}, 0, -1}, // 三个元素，目标值不存在

		// 特殊情况测试用例
		{[]int{0, 1, 2, 4, 5, 6, 7}, 0, 0}, // 没有旋转的数组
		{[]int{0, 1, 2, 4, 5, 6, 7}, 7, 6}, // 没有旋转的数组，目标值在末尾
		{[]int{7, 0, 1, 2, 4, 5, 6}, 7, 0}, // 旋转一个位置
		{[]int{6, 7, 0, 1, 2, 4, 5}, 6, 0}, // 旋转两个位置
	}

	fmt.Println("开始测试搜索旋转排序数组...")
	fmt.Println(strings.Repeat("=", 50))

	passed := 0
	total := len(testCases)

	for i, tc := range testCases {
		fmt.Printf("\n测试用例 %d:\n", i+1)
		fmt.Printf("输入: nums = %v, target = %d\n", tc.nums, tc.target)

		result := search(tc.nums, tc.target)

		fmt.Printf("输出: %d\n", result)
		fmt.Printf("期望: %d\n", tc.expect)

		if result == tc.expect {
			fmt.Printf("✅ 通过\n")
			passed++
		} else {
			fmt.Printf("❌ 失败\n")
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("测试结果: %d/%d 通过\n", passed, total)

	if passed == total {
		fmt.Println("🎉 所有测试用例都通过了！")
	} else {
		fmt.Println("💡 还有一些测试用例没有通过，请检查你的实现。")
	}
}

/*
预期输出：
开始测试搜索旋转排序数组...
==================================================

测试用例 1:
输入: nums = [4 5 6 7 0 1 2], target = 0
输出: 4
期望: 4
✅ 通过

测试用例 2:
输入: nums = [4 5 6 7 0 1 2], target = 3
输出: -1
期望: -1
✅ 通过

测试用例 3:
输入: nums = [1], target = 0
输出: -1
期望: -1
✅ 通过

测试用例 4:
输入: nums = [1], target = 1
输出: 0
期望: 0
✅ 通过

测试用例 5:
输入: nums = [1 3], target = 3
输出: 1
期望: 1
✅ 通过

测试用例 6:
输入: nums = [3 1], target = 1
输出: 1
期望: 1
✅ 通过

测试用例 7:
输入: nums = [5 1 3], target = 3
输出: 2
期望: 2
✅ 通过

测试用例 8:
输入: nums = [5 1 3], target = 5
输出: 0
期望: 0
✅ 通过

测试用例 9:
输入: nums = [5 1 3], target = 1
输出: 1
期望: 1
✅ 通过

测试用例 10:
输入: nums = [5 1 3], target = 0
输出: -1
期望: -1
✅ 通过

测试用例 11:
输入: nums = [0 1 2 4 5 6 7], target = 0
输出: 0
期望: 0
✅ 通过

测试用例 12:
输入: nums = [0 1 2 4 5 6 7], target = 7
输出: 6
期望: 6
✅ 通过

测试用例 13:
输入: nums = [7 0 1 2 4 5 6], target = 7
输出: 0
期望: 0
✅ 通过

测试用例 14:
输入: nums = [6 7 0 1 2 4 5], target = 6
输出: 0
期望: 0
✅ 通过

==================================================
测试结果: 14/14 通过
🎉 所有测试用例都通过了！
*/
