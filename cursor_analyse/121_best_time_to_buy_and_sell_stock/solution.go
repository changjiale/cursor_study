package main

import "fmt"

/*
题目：买卖股票的最佳时机
难度：简单
标签：数组、动态规划

题目描述：
给定一个数组 prices，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
你只能选择某一天买入这只股票，并选择在未来的某一个不同的日子卖出该股票。设计一个算法来计算你所能获取的最大利润。
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0。

要求：
- 时间复杂度：O(n)，其中 n 是数组长度
- 空间复杂度：O(1)

示例：
输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5。
注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
*/

// 解法一：一次遍历（推荐）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0] // 记录历史最低价格
	maxProfit := 0        // 记录最大利润

	// 遍历每一天的价格
	for i := 1; i < len(prices); i++ {
		// 更新最大利润：当前价格 - 历史最低价格
		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}

		// 更新历史最低价格
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return maxProfit
}

// 解法二：动态规划
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// dp[i] 表示第i天能获得的最大利润
	// dp[i] = max(dp[i-1], prices[i] - minPrice)
	dp := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		// 当前利润 = 当前价格 - 历史最低价格
		currentProfit := prices[i] - minPrice

		// 更新最大利润
		if currentProfit > dp {
			dp = currentProfit
		}

		// 更新历史最低价格
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	return dp
}

// 解法三：暴力法（不推荐，仅用于理解）
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
func maxProfit3(prices []int) int {
	maxProfit := 0

	// 遍历所有可能的买入和卖出组合
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

// 解法四：分治法（不推荐，仅用于理解）
// 时间复杂度：O(n log n)
// 空间复杂度：O(log n)，递归栈深度
func maxProfit4(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	return maxProfitHelper(prices, 0, len(prices)-1)
}

func maxProfitHelper(prices []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := (left + right) / 2

	// 递归求解左半部分和右半部分
	leftProfit := maxProfitHelper(prices, left, mid)
	rightProfit := maxProfitHelper(prices, mid+1, right)

	// 计算跨越中点的最大利润
	crossProfit := maxProfitCross(prices, left, mid, right)

	// 返回三者中的最大值
	return max(leftProfit, max(rightProfit, crossProfit))
}

func maxProfitCross(prices []int, left, mid, right int) int {
	// 在左半部分找最小值
	minPrice := prices[left]
	for i := left; i <= mid; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}

	// 在右半部分找最大值
	maxPrice := prices[mid+1]
	for i := mid + 1; i <= right; i++ {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}
	}

	return maxPrice - minPrice
}

// 解法五：单调栈思想（变种）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit5(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	stack := []int{prices[0]} // 维护一个单调递减栈
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// 如果当前价格比栈顶价格高，计算利润
		for len(stack) > 0 && prices[i] > stack[len(stack)-1] {
			profit := prices[i] - stack[len(stack)-1]
			if profit > maxProfit {
				maxProfit = profit
			}
			stack = stack[:len(stack)-1]
		}

		// 将当前价格入栈
		stack = append(stack, prices[i])
	}

	return maxProfit
}

// 解法六：Kadane算法思想
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxProfit6(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// 计算价格差数组
	diffs := make([]int, len(prices)-1)
	for i := 0; i < len(prices)-1; i++ {
		diffs[i] = prices[i+1] - prices[i]
	}

	// 使用Kadane算法找最大子数组和
	maxSoFar := 0
	maxEndingHere := 0

	for _, diff := range diffs {
		maxEndingHere = max(0, maxEndingHere+diff)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}

	return maxSoFar
}

// 辅助函数：返回两个数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 测试用例
	testCases := []struct {
		prices []int
		expect int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{5, 4, 3, 2, 1}, 0},
		{[]int{1}, 0},
		{[]int{}, 0},
		{[]int{2, 4, 1}, 2},
		{[]int{3, 2, 6, 5, 0, 3}, 4},
		{[]int{1, 2}, 1},
		{[]int{2, 1}, 0},
		{[]int{1, 1, 1, 1}, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9},
	}

	fmt.Println("=== 解法一：一次遍历（推荐）===")
	for i, tc := range testCases {
		result := maxProfit(tc.prices)
		fmt.Printf("测试用例 %d: prices=%v, 结果=%d, 期望=%d, 通过=%t\n",
			i+1, tc.prices, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法二：动态规划 ===")
	for i, tc := range testCases {
		result := maxProfit2(tc.prices)
		fmt.Printf("测试用例 %d: prices=%v, 结果=%d, 期望=%d, 通过=%t\n",
			i+1, tc.prices, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法三：暴力法 ===")
	for i, tc := range testCases {
		result := maxProfit3(tc.prices)
		fmt.Printf("测试用例 %d: prices=%v, 结果=%d, 期望=%d, 通过=%t\n",
			i+1, tc.prices, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法四：分治法 ===")
	for i, tc := range testCases {
		result := maxProfit4(tc.prices)
		fmt.Printf("测试用例 %d: prices=%v, 结果=%d, 期望=%d, 通过=%t\n",
			i+1, tc.prices, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法五：单调栈思想 ===")
	for i, tc := range testCases {
		result := maxProfit5(tc.prices)
		fmt.Printf("测试用例 %d: prices=%v, 结果=%d, 期望=%d, 通过=%t\n",
			i+1, tc.prices, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法六：Kadane算法思想 ===")
	for i, tc := range testCases {
		result := maxProfit6(tc.prices)
		fmt.Printf("测试用例 %d: prices=%v, 结果=%d, 期望=%d, 通过=%t\n",
			i+1, tc.prices, result, tc.expect, result == tc.expect)
	}
}

/*
解题思路：

1. 一次遍历（推荐）：
   - 维护一个历史最低价格变量
   - 遍历每一天，计算当前价格与历史最低价格的差值
   - 更新最大利润和历史最低价格
   - 时间复杂度O(n)，空间复杂度O(1)

2. 动态规划：
   - 定义dp[i]为第i天能获得的最大利润
   - 状态转移方程：dp[i] = max(dp[i-1], prices[i] - minPrice)
   - 优化空间复杂度，只使用一个变量

3. 暴力法：
   - 遍历所有可能的买入和卖出组合
   - 时间复杂度O(n²)，不推荐

4. 分治法：
   - 将问题分解为左半部分、右半部分和跨越中点的情况
   - 时间复杂度O(n log n)，不推荐

5. 单调栈思想：
   - 维护一个单调递减栈
   - 当遇到更高价格时，计算利润并更新最大值
   - 时间复杂度O(n)

6. Kadane算法思想：
   - 将问题转化为求最大子数组和
   - 计算相邻价格差，使用Kadane算法
   - 时间复杂度O(n)

时间复杂度分析：
- 解法一、二、五、六：O(n)，一次遍历
- 解法三：O(n²)，双重循环
- 解法四：O(n log n)，分治递归

空间复杂度分析：
- 解法一、二、三：O(1)，只使用常数额外空间
- 解法四：O(log n)，递归栈深度
- 解法五：O(n)，栈的大小
- 解法六：O(n)，价格差数组

关键点：
1. 一次遍历是最优解，维护历史最低价格
2. 动态规划可以优化空间复杂度
3. 理解不同解法的适用场景
4. 注意边界条件：空数组、单个元素等
5. 暴力法虽然简单，但时间复杂度较高
*/
