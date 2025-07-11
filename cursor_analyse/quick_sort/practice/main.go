package main

import (
	"fmt"
)

/*
题目：快速排序
难度：中等
标签：排序、分治

题目描述：
实现快速排序算法，将给定的数组按照升序排序。

要求：
1. 实现基础的快速排序算法
2. 考虑处理重复元素的情况
3. 优化基准值的选择
4. 注意处理边界情况

示例：
输入: nums = [3,2,1,5,6,4]
输出: [1,2,3,4,5,6]

输入: nums = [1,1,1,1,1]
输出: [1,1,1,1,1]

输入: nums = [5,4,3,2,1]
输出: [1,2,3,4,5]

提示：
1. 可以使用最后一个元素作为基准值
2. 也可以随机选择基准值
3. 考虑使用三路快排处理重复元素
4. 注意递归的终止条件
*/

// TODO: 在这里实现你的快速排序算法
func quickSort(nums []int) {
	// 请实现你的代码
}

func main() {
	// 测试用例
	testCases := [][]int{
		{3, 2, 1, 5, 6, 4}, // 普通数组
		{1, 1, 1, 1, 1},    // 所有元素相同
		{5, 4, 3, 2, 1},    // 逆序数组
		{1},                // 单个元素
		{},                 // 空数组
	}

	for _, nums := range testCases {
		fmt.Printf("\n输入: nums = %v\n", nums)

		// 复制数组用于测试
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)

		// 调用排序函数
		quickSort(numsCopy)

		fmt.Printf("输出: %v\n", numsCopy)
	}
}

/*
预期输出：
输入: nums = [3 2 1 5 6 4]
输出: [1 2 3 4 5 6]

输入: nums = [1 1 1 1 1]
输出: [1 1 1 1 1]

输入: nums = [5 4 3 2 1]
输出: [1 2 3 4 5]

输入: nums = [1]
输出: [1]

输入: nums = []
输出: []
*/
