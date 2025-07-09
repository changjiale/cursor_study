package main

import (
	"fmt"
)

/*
题目：53. 最大子数组和
难度：中等
标签：数组、分治、动态规划

题目描述：
给你一个整数数组 nums，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组是数组中的一个连续部分。

要求：
- 1 <= nums.length <= 105
- -104 <= nums[i] <= 104

示例：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6。

输入：nums = [1]
输出：1

输入：nums = [5,4,-1,7,8]
输出：23

提示：
- 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
- 考虑使用动态规划的思想
- 注意处理负数的情况
*/

// TODO: 在这里实现你的算法
// 要求：实现时间复杂度为 O(n) 的解法
func maxSubArray(nums []int) int {
	// 请实现你的代码
	return 0
}

// 测试用例结构
type TestCase struct {
	input  []int
	output int
}

func main() {
	// 测试用例
	testCases := []TestCase{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6}, // 普通情况
		{[]int{1}, 1},               // 单个元素
		{[]int{5, 4, -1, 7, 8}, 23}, // 全正数
		{[]int{-1}, -1},             // 单个负数
		{[]int{-2, -1}, -1},         // 全负数
		{[]int{1, 2, 3, 4, 5}, 15},  // 递增序列
		{[]int{-1, -2, -3, -4}, -1}, // 递减负数序列
		{[]int{0, 0, 0, 0}, 0},      // 全零
		{[]int{1, -2, 3, -4, 5}, 5}, // 交替正负
		{[]int{2, -1, 3, -2, 4}, 6}, // 复杂情况
	}

	fmt.Println("=== 最大子数组和测试 ===")

	allPassed := true
	for i, tc := range testCases {
		fmt.Printf("\n测试用例 %d:\n", i+1)
		fmt.Printf("输入: nums = %v\n", tc.input)

		// 复制数组用于测试
		nums := make([]int, len(tc.input))
		copy(nums, tc.input)

		result := maxSubArray(nums)
		fmt.Printf("输出: %d\n", result)
		fmt.Printf("期望: %d\n", tc.output)

		if result == tc.output {
			fmt.Println("✅ 通过")
		} else {
			fmt.Println("❌ 失败")
			allPassed = false
		}
	}

	fmt.Printf("\n=== 测试总结 ===\n")
	if allPassed {
		fmt.Println("🎉 所有测试用例都通过了！")
	} else {
		fmt.Println("⚠️  有测试用例未通过，请检查你的实现")
	}
}

/*
预期输出：
=== 最大子数组和测试 ===

测试用例 1:
输入: nums = [-2 1 -3 4 -1 2 1 -5 4]
输出: 6
期望: 6
✅ 通过

测试用例 2:
输入: nums = [1]
输出: 1
期望: 1
✅ 通过

测试用例 3:
输入: nums = [5 4 -1 7 8]
输出: 23
期望: 23
✅ 通过

测试用例 4:
输入: nums = [-1]
输出: -1
期望: -1
✅ 通过

测试用例 5:
输入: nums = [-2 -1]
输出: -1
期望: -1
✅ 通过

测试用例 6:
输入: nums = [1 2 3 4 5]
输出: 15
期望: 15
✅ 通过

测试用例 7:
输入: nums = [-1 -2 -3 -4]
输出: -1
期望: -1
✅ 通过

测试用例 8:
输入: nums = [0 0 0 0]
输出: 0
期望: 0
✅ 通过

测试用例 9:
输入: nums = [1 -2 3 -4 5]
输出: 5
期望: 5
✅ 通过

测试用例 10:
输入: nums = [2 -1 3 -2 4]
输出: 6
期望: 6
✅ 通过

=== 测试总结 ===
🎉 所有测试用例都通过了！
*/
