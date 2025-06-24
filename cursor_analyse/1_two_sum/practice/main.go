package main

import (
	"fmt"
)

/*
题目：两数之和
难度：简单
标签：数组、哈希表

题目描述：
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

要求：
1. 返回两个数的下标
2. 同一个元素不能重复使用
3. 假设每种输入只会对应一个答案

示例：
输入: nums = [2,7,11,15], target = 9
输出: [0,1]
解释: 因为 nums[0] + nums[1] == 9，返回 [0, 1]

输入: nums = [3,2,4], target = 6
输出: [1,2]

输入: nums = [3,3], target = 6
输出: [0,1]

提示：
1. 可以使用哈希表来优化查找
2. 注意处理重复元素的情况
3. 考虑边界情况（空数组、无解等）
*/

// TODO: 在这里实现你的算法
func twoSum(nums []int, target int) []int {
	// 请实现你的代码
	return nil
}

func main() {
	// 测试用例
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9}, // 普通情况
		{[]int{3, 2, 4}, 6},      // 普通情况
		{[]int{3, 3}, 6},         // 重复元素
		{[]int{1, 2, 3, 4}, 10},  // 无解情况
		{[]int{1}, 1},            // 单个元素
		{[]int{}, 0},             // 空数组
	}

	for _, tc := range testCases {
		fmt.Printf("\n输入: nums = %v, target = %v\n", tc.nums, tc.target)
		result := twoSum(tc.nums, tc.target)
		fmt.Printf("输出: %v\n", result)
	}
}

/*
预期输出：
输入: nums = [2 7 11 15], target = 9
输出: [0 1]

输入: nums = [3 2 4], target = 6
输出: [1 2]

输入: nums = [3 3], target = 6
输出: [0 1]

输入: nums = [1 2 3 4], target = 10
输出: []

输入: nums = [1], target = 1
输出: []

输入: nums = [], target = 0
输出: []
*/
