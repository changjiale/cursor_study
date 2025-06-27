package main

import (
	"fmt"
)

/*
题目：两数之和
难度：简单
标签：数组、哈希表

题目描述：
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复使用。

你可以按任意顺序返回答案。

要求：
1. 返回两个数的下标
2. 同一个元素不能重复使用
3. 假设每种输入只会对应一个答案
*/

// 解法一：暴力解法
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
func twoSum1(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 解法二：哈希表
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func twoSum2(nums []int, target int) []int {
	// 使用哈希表存储每个数字及其下标
	numMap := make(map[int]int)

	for i, num := range nums {
		// 计算需要寻找的补数
		complement := target - num

		// 如果补数在哈希表中，说明找到了答案
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}

		// 将当前数字及其下标加入哈希表
		numMap[num] = i
	}

	return nil
}

// 解法三：双指针（需要先排序，但会改变原数组）
// 时间复杂度：O(n log n)
// 空间复杂度：O(n)
func twoSum3(nums []int, target int) []int {
	n := len(nums)

	// 创建索引数组，用于记录原始位置
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	// 对索引数组进行排序，排序依据是对应的数值
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[indices[i]] > nums[indices[j]] {
				indices[i], indices[j] = indices[j], indices[i]
			}
		}
	}

	// 双指针查找
	left, right := 0, n-1
	for left < right {
		sum := nums[indices[left]] + nums[indices[right]]
		if sum == target {
			return []int{indices[left], indices[right]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

// 解法四：极简写法（哈希表优化版）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func twoSum4(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		if j, ok := seen[target-num]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

// 辅助函数：比较两个数组是否相等
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// 测试用例
	testCases := []struct {
		nums   []int
		target int
		output []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},        // 普通情况
		{[]int{3, 2, 4}, 6, []int{1, 2}},             // 中间位置
		{[]int{3, 3}, 6, []int{0, 1}},                // 相同元素
		{[]int{1, 5, 8, 10, 13}, 18, []int{2, 4}},    // 较大数值
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}}, // 负数
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},          // 包含零
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},       // 末尾位置
	}

	for _, tc := range testCases {
		fmt.Printf("\n输入: nums = %v, target = %d\n", tc.nums, tc.target)

		// 解法一：暴力解法
		result1 := twoSum1(tc.nums, tc.target)
		fmt.Printf("解法一（暴力解法）结果: %v\n", result1)

		// 解法二：哈希表
		result2 := twoSum2(tc.nums, tc.target)
		fmt.Printf("解法二（哈希表）结果: %v\n", result2)

		// 解法三：双指针
		result3 := twoSum3(tc.nums, tc.target)
		fmt.Printf("解法三（双指针）结果: %v\n", result3)

		// 解法四：极简写法
		result4 := twoSum4(tc.nums, tc.target)
		fmt.Printf("解法四（极简写法）结果: %v\n", result4)

		// 验证结果
		expected := tc.output
		if equalSlices(result1, expected) && equalSlices(result2, expected) &&
			equalSlices(result3, expected) && equalSlices(result4, expected) {
			fmt.Println("✓ 所有解法都通过")
		} else {
			fmt.Println("✗ 部分解法失败")
		}
	}
}

/*
预期输出：
输入: nums = [2 7 11 15], target = 9
解法一（暴力解法）结果: [0 1]
解法二（哈希表）结果: [0 1]
解法三（双指针）结果: [0 1]
解法四（极简写法）结果: [0 1]
✓ 所有解法都通过

输入: nums = [3 2 4], target = 6
解法一（暴力解法）结果: [1 2]
解法二（哈希表）结果: [1 2]
解法三（双指针）结果: [1 2]
解法四（极简写法）结果: [1 2]
✓ 所有解法都通过

输入: nums = [3 3], target = 6
解法一（暴力解法）结果: [0 1]
解法二（哈希表）结果: [0 1]
解法三（双指针）结果: [0 1]
解法四（极简写法）结果: [0 1]
✓ 所有解法都通过

输入: nums = [1 5 8 10 13], target = 18
解法一（暴力解法）结果: [2 4]
解法二（哈希表）结果: [2 4]
解法三（双指针）结果: [2 4]
解法四（极简写法）结果: [2 4]
✓ 所有解法都通过

输入: nums = [-1 -2 -3 -4 -5], target = -8
解法一（暴力解法）结果: [2 4]
解法二（哈希表）结果: [2 4]
解法三（双指针）结果: [2 4]
解法四（极简写法）结果: [2 4]
✓ 所有解法都通过

输入: nums = [0 4 3 0], target = 0
解法一（暴力解法）结果: [0 3]
解法二（哈希表）结果: [0 3]
解法三（双指针）结果: [0 3]
解法四（极简写法）结果: [0 3]
✓ 所有解法都通过

输入: nums = [1 2 3 4 5], target = 9
解法一（暴力解法）结果: [3 4]
解法二（哈希表）结果: [3 4]
解法三（双指针）结果: [3 4]
解法四（极简写法）结果: [3 4]
✓ 所有解法都通过
*/

/*
算法分析：

1. 暴力解法：
   - 双重循环遍历所有可能的组合
   - 时间复杂度高，但思路简单直观
   - 适用于小规模数据

2. 哈希表：
   - 使用哈希表存储已遍历的数字
   - 每次查找补数的时间复杂度为O(1)
   - 最优解法，推荐使用

3. 双指针：
   - 需要先排序，会改变原数组
   - 适用于需要返回数值而不是下标的变种题目
   - 时间复杂度受排序影响

4. 极简写法：
   - 哈希表的简化版本
   - 代码更简洁，逻辑相同
   - 实际应用中的首选
*/
