package main

import (
	"fmt"
	"sort"
)

/*
题目：三数之和
难度：中等
标签：数组、双指针、排序

题目描述：
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

要求：
1. 返回所有和为0的三元组
2. 不能包含重复的三元组
3. 三元组中的元素顺序不重要

示例：
输入: nums = [-1,0,1,2,-1,-4]
输出: [[-1,-1,2],[-1,0,1]]
解释:
- nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0
- nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0
- nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。

输入: nums = [0,1,1]
输出: []

输入: nums = [0,0,0]
输出: [[0,0,0]]

提示：
- 使用排序 + 双指针
- 注意去重处理
- 可以提前剪枝优化
*/

// TODO: 在这里实现你的算法
func threeSum(nums []int) [][]int {
	// 请实现你的代码
	return nil
}

// 辅助函数：比较两个二维数组是否相等
func equalSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// 对每个三元组排序
	sortSlices := func(slices [][]int) {
		for i := range slices {
			sort.Ints(slices[i])
		}
		// 对三元组数组排序
		sort.Slice(slices, func(i, j int) bool {
			if len(slices[i]) != len(slices[j]) {
				return len(slices[i]) < len(slices[j])
			}
			for k := 0; k < len(slices[i]); k++ {
				if slices[i][k] != slices[j][k] {
					return slices[i][k] < slices[j][k]
				}
			}
			return false
		})
	}

	sortSlices(a)
	sortSlices(b)

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	// 测试用例
	testCases := []struct {
		input  []int
		output [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},                  // 普通情况
		{[]int{0, 1, 1}, [][]int{}},                                                     // 无解
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},                                            // 全零
		{[]int{}, [][]int{}},                                                            // 空数组
		{[]int{0}, [][]int{}},                                                           // 单个元素
		{[]int{1, 2}, [][]int{}},                                                        // 两个元素
		{[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},                        // 复杂情况
		{[]int{-1, 0, 1, 2, -1, -4, 0, 0}, [][]int{{-1, -1, 2}, {-1, 0, 1}, {0, 0, 0}}}, // 包含重复
	}

	for _, tc := range testCases {
		fmt.Printf("\n输入: %v\n", tc.input)
		result := threeSum(tc.input)
		fmt.Printf("输出: %v\n", result)
		fmt.Printf("期望: %v\n", tc.output)
		if equalSlices(result, tc.output) {
			fmt.Println("✓ 通过")
		} else {
			fmt.Println("✗ 失败")
		}
	}
}

/*
预期输出：
输入: [-1 0 1 2 -1 -4]
输出: [[-1 -1 2] [-1 0 1]]
期望: [[-1 -1 2] [-1 0 1]]
✓ 通过

输入: [0 1 1]
输出: []
期望: []
✓ 通过

输入: [0 0 0]
输出: [[0 0 0]]
期望: [[0 0 0]]
✓ 通过

输入: []
输出: []
期望: []
✓ 通过

输入: [0]
输出: []
期望: []
✓ 通过

输入: [1 2]
输出: []
期望: []
✓ 通过

输入: [-2 0 1 1 2]
输出: [[-2 0 2] [-2 1 1]]
期望: [[-2 0 2] [-2 1 1]]
✓ 通过

输入: [-1 0 1 2 -1 -4 0 0]
输出: [[-1 -1 2] [-1 0 1] [0 0 0]]
期望: [[-1 -1 2] [-1 0 1] [0 0 0]]
✓ 通过
*/
