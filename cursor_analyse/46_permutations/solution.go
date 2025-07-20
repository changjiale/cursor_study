package main

import "fmt"

/*
题目：全排列
难度：中等
标签：数组、回溯算法

题目描述：
给定一个不含重复数字的数组 nums，返回其所有可能的全排列。你可以按任意顺序返回答案。

要求：
- 时间复杂度：O(n!)，其中 n 是数组长度
- 空间复杂度：O(n)，递归栈深度

示例：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

输入：nums = [0,1]
输出：[[0,1],[1,0]]

输入：nums = [1]
输出：[[1]]
*/

// 解法一：回溯算法（推荐）
// 时间复杂度：O(n!)
// 空间复杂度：O(n)
func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(path []int, used []bool)

	backtrack = func(path []int, used []bool) {
		// 如果路径长度等于数组长度，说明找到一个排列
		if len(path) == len(nums) {
			// 创建path的副本，避免后续修改影响结果
			perm := make([]int, len(path))
			copy(perm, path)
			result = append(result, perm)
			return
		}

		// 尝试将每个未使用的数字加入当前路径
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				// 标记当前数字为已使用
				used[i] = true
				// 将当前数字加入路径
				path = append(path, nums[i])

				// 递归处理剩余数字
				backtrack(path, used)

				// 回溯：恢复状态
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}

	// 初始化used数组，记录每个数字是否被使用
	used := make([]bool, len(nums))
	backtrack([]int{}, used)

	return result
}

// 解法二：交换法
// 时间复杂度：O(n!)
// 空间复杂度：O(n)
func permute2(nums []int) [][]int {
	var result [][]int
	var swap func(index int)

	swap = func(index int) {
		// 如果到达数组末尾，说明找到一个排列
		if index == len(nums) {
			// 创建nums的副本
			perm := make([]int, len(nums))
			copy(perm, nums)
			result = append(result, perm)
			return
		}

		// 将当前位置的数字与后面的每个数字交换
		for i := index; i < len(nums); i++ {
			// 交换位置
			nums[index], nums[i] = nums[i], nums[index]

			// 递归处理下一个位置
			swap(index + 1)

			// 回溯：恢复交换
			nums[index], nums[i] = nums[i], nums[index]
		}
	}

	swap(0)
	return result
}

// 解法三：插入法
// 时间复杂度：O(n!)
// 空间复杂度：O(n!)
func permute3(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	var result [][]int

	// 获取前n-1个数字的全排列
	prevPerms := permute3(nums[:len(nums)-1])
	lastNum := nums[len(nums)-1]

	// 对于每个前n-1个数字的排列，将最后一个数字插入到所有可能的位置
	for _, perm := range prevPerms {
		// 在排列的每个位置插入最后一个数字
		for i := 0; i <= len(perm); i++ {
			newPerm := make([]int, len(perm)+1)
			copy(newPerm[:i], perm[:i])
			newPerm[i] = lastNum
			copy(newPerm[i+1:], perm[i:])
			result = append(result, newPerm)
		}
	}

	return result
}

// 解法四：字典序法（迭代）
// 时间复杂度：O(n!)
// 空间复杂度：O(n)
func permute4(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var result [][]int

	// 创建初始排列（升序）
	perm := make([]int, len(nums))
	copy(perm, nums)

	// 添加初始排列
	result = append(result, append([]int{}, perm...))

	// 生成下一个排列，直到没有下一个排列
	for nextPermutation(perm) {
		result = append(result, append([]int{}, perm...))
	}

	return result
}

// 生成下一个排列（字典序）
func nextPermutation(nums []int) bool {
	n := len(nums)

	// 从右向左找到第一个递减的位置
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 如果没有找到递减位置，说明已经是最后一个排列
	if i < 0 {
		return false
	}

	// 从右向左找到第一个大于nums[i]的数字
	j := n - 1
	for j > i && nums[j] <= nums[i] {
		j--
	}

	// 交换nums[i]和nums[j]
	nums[i], nums[j] = nums[j], nums[i]

	// 反转i+1到末尾的部分
	reverse(nums, i+1, n-1)

	return true
}

// 反转数组的指定范围
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	// 测试用例
	testCases := []struct {
		nums   []int
		expect [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			[]int{0, 1},
			[][]int{{0, 1}, {1, 0}},
		},
		{
			[]int{1},
			[][]int{{1}},
		},
		{
			[]int{},
			[][]int{},
		},
		{
			[]int{1, 2},
			[][]int{{1, 2}, {2, 1}},
		},
	}

	fmt.Println("=== 解法一：回溯算法 ===")
	for i, tc := range testCases {
		result := permute(tc.nums)
		fmt.Printf("测试用例 %d: nums=%v\n", i+1, tc.nums)
		fmt.Printf("结果: %v\n", result)
		fmt.Printf("期望: %v\n", tc.expect)
		fmt.Printf("通过: %t\n\n", len(result) == len(tc.expect))
	}

	fmt.Println("=== 解法二：交换法 ===")
	for i, tc := range testCases {
		nums := make([]int, len(tc.nums))
		copy(nums, tc.nums)
		result := permute2(nums)
		fmt.Printf("测试用例 %d: nums=%v\n", i+1, tc.nums)
		fmt.Printf("结果: %v\n", result)
		fmt.Printf("期望: %v\n", tc.expect)
		fmt.Printf("通过: %t\n\n", len(result) == len(tc.expect))
	}

	fmt.Println("=== 解法三：插入法 ===")
	for i, tc := range testCases {
		result := permute3(tc.nums)
		fmt.Printf("测试用例 %d: nums=%v\n", i+1, tc.nums)
		fmt.Printf("结果: %v\n", result)
		fmt.Printf("期望: %v\n", tc.expect)
		fmt.Printf("通过: %t\n\n", len(result) == len(tc.expect))
	}

	fmt.Println("=== 解法四：字典序法 ===")
	for i, tc := range testCases {
		result := permute4(tc.nums)
		fmt.Printf("测试用例 %d: nums=%v\n", i+1, tc.nums)
		fmt.Printf("结果: %v\n", result)
		fmt.Printf("期望: %v\n", tc.expect)
		fmt.Printf("通过: %t\n\n", len(result) == len(tc.expect))
	}
}

/*
解题思路：

1. 回溯算法（推荐）：
   - 使用递归和回溯的思想
   - 维护一个路径和已使用标记数组
   - 每次选择一个未使用的数字加入路径
   - 当路径长度等于数组长度时，找到一个排列
   - 回溯时恢复状态，尝试其他选择

2. 交换法：
   - 通过交换数组元素来生成排列
   - 将当前位置的数字与后面的每个数字交换
   - 递归处理下一个位置
   - 回溯时恢复交换

3. 插入法：
   - 基于递归的思想
   - 先求前n-1个数字的全排列
   - 将第n个数字插入到每个排列的所有可能位置
   - 适合理解全排列的生成过程

4. 字典序法：
   - 按照字典序生成所有排列
   - 从初始排列开始，不断生成下一个排列
   - 需要实现nextPermutation函数
   - 适合需要按顺序生成排列的场景

时间复杂度分析：
- 所有解法：O(n!)，因为n个数字的全排列数量是n!

空间复杂度分析：
- 回溯法：O(n)，递归栈深度
- 交换法：O(n)，递归栈深度
- 插入法：O(n!)，需要存储所有排列
- 字典序法：O(n)，递归栈深度

关键点：
1. 回溯算法的核心是状态恢复
2. 避免重复使用同一个数字
3. 注意数组的深拷贝，避免修改原数组
4. 理解不同解法的适用场景
*/
