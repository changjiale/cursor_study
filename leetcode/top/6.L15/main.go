package main

import "sort"

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
func main() {

}

// 核心思想排序+双指针
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)

	var result [][]int
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i+1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 跳过重复元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result

}
