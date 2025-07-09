package main

import (
	"fmt"
	"math"
)

/**
53. 最大子数组和
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
*/

// 解法1：动态规划（Kadane算法）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 核心思想：dp[i] = max(nums[i], dp[i-1] + nums[i])
func maxSubArrayDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]     // 全局最大和
	currentSum := nums[0] // 当前子数组和

	// 从第二个元素开始遍历
	for i := 1; i < len(nums); i++ {
		// 关键：选择继续当前子数组，还是重新开始
		// 如果当前和小于0，重新开始；否则继续累加
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}

		// 更新全局最大和
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

// 解法2：分治法
// 时间复杂度：O(n log n)
// 空间复杂度：O(log n) - 递归调用栈
// 核心思想：将数组分成两半，最大子数组和可能在左半、右半或跨越中点
func maxSubArrayDivide(nums []int) int {
	return divideAndConquer(nums, 0, len(nums)-1)
}

func divideAndConquer(nums []int, left, right int) int {
	// 基础情况：只有一个元素
	if left == right {
		return nums[left]
	}

	// 计算中点
	mid := left + (right-left)/2

	// 递归求解左半部分和右半部分
	leftMax := divideAndConquer(nums, left, mid)
	rightMax := divideAndConquer(nums, mid+1, right)

	// 计算跨越中点的最大子数组和
	crossMax := maxCrossingSum(nums, left, mid, right)

	// 返回三者中的最大值
	return max(leftMax, max(rightMax, crossMax))
}

func maxCrossingSum(nums []int, left, mid, right int) int {
	// 计算左半部分的最大后缀和
	leftSum := 0
	leftMax := math.MinInt32
	for i := mid; i >= left; i-- {
		leftSum += nums[i]
		if leftSum > leftMax {
			leftMax = leftSum
		}
	}

	// 计算右半部分的最大前缀和
	rightSum := 0
	rightMax := math.MinInt32
	for i := mid + 1; i <= right; i++ {
		rightSum += nums[i]
		if rightSum > rightMax {
			rightMax = rightSum
		}
	}

	// 返回跨越中点的最大和
	return leftMax + rightMax
}

// 解法3：暴力解法（不推荐，仅用于理解）
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
func maxSubArrayBruteForce(nums []int) int {
	maxSum := math.MinInt32

	// 枚举所有可能的子数组
	for i := 0; i < len(nums); i++ {
		currentSum := 0
		for j := i; j < len(nums); j++ {
			currentSum += nums[j]
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}

	return maxSum
}

// 解法4：优化的动态规划（更清晰的实现）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func maxSubArrayOptimized(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// 状态转移方程：dp[i] = max(nums[i], dp[i-1] + nums[i])
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

// 辅助函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 演示算法执行过程
func demonstrateAlgorithm() {
	fmt.Println("=== 最大子数组和算法演示 ===")

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Printf("数组: %v\n", nums)

	fmt.Println("\n动态规划执行过程:")
	currentSum := nums[0]
	maxSum := nums[0]
	fmt.Printf("初始: currentSum = %d, maxSum = %d\n", currentSum, maxSum)

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}

		fmt.Printf("i=%d: nums[%d]=%d, currentSum=%d, maxSum=%d\n",
			i, i, nums[i], currentSum, maxSum)
	}

	fmt.Printf("\n最终结果: %d\n", maxSum)
}

// 解法5：双指针 + 滑动窗口（直观易懂）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 核心思想：维护一个滑动窗口，当窗口和为负数时收缩窗口
func maxSubArrayTwoPointers(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := 0
	left := 0

	// 右指针遍历数组
	for right := 0; right < len(nums); right++ {
		// 将当前元素加入窗口
		currentSum += nums[right]

		// 更新最大和
		if currentSum > maxSum {
			maxSum = currentSum
		}

		// 关键：如果当前窗口和为负数，收缩左边界
		// 因为负数会拖累后续的子数组和
		for left <= right && currentSum < 0 {
			currentSum -= nums[left]
			left++
		}
	}

	return maxSum
}

// 解法6：前缀和 + 贪心（另一种思路）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 核心思想：前缀和的思想，维护最小前缀和
func maxSubArrayPrefixSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := 0
	minPrefixSum := 0 // 最小前缀和

	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]

		// 当前前缀和减去最小前缀和，得到以i结尾的最大子数组和
		if currentSum-minPrefixSum > maxSum {
			maxSum = currentSum - minPrefixSum
		}

		// 更新最小前缀和
		if currentSum < minPrefixSum {
			minPrefixSum = currentSum
		}
	}

	return maxSum
}

// 解法7：模拟人类思维（最直观）
// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 核心思想：模拟人类如何找最大子数组和
func maxSubArrayHuman(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// 人类思维：如果当前累加和已经小于0，还不如重新开始
		// 因为负数会拖累后续的累加
		if currentSum < 0 {
			// 重新开始一个新的子数组
			currentSum = nums[i]
		} else {
			// 继续累加当前子数组
			currentSum += nums[i]
		}

		// 更新最大和
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

// 解法8：分治法的简化版本（更容易理解）
// 时间复杂度：O(n log n)
// 空间复杂度：O(log n)
func maxSubArrayDivideSimple(nums []int) int {
	return divideAndConquerSimple(nums, 0, len(nums)-1)
}

func divideAndConquerSimple(nums []int, left, right int) int {
	// 基础情况
	if left == right {
		return nums[left]
	}
	if left > right {
		return -10000 // 返回一个很小的数
	}

	// 计算中点
	mid := (left + right) / 2

	// 情况1：最大子数组完全在左半部分
	leftMax := divideAndConquerSimple(nums, left, mid-1)

	// 情况2：最大子数组完全在右半部分
	rightMax := divideAndConquerSimple(nums, mid+1, right)

	// 情况3：最大子数组跨越中点
	// 从中点向左找最大后缀和
	leftSum := 0
	leftMaxSum := 0
	for i := mid - 1; i >= left; i-- {
		leftSum += nums[i]
		if leftSum > leftMaxSum {
			leftMaxSum = leftSum
		}
	}

	// 从中点向右找最大前缀和
	rightSum := 0
	rightMaxSum := 0
	for i := mid + 1; i <= right; i++ {
		rightSum += nums[i]
		if rightSum > rightMaxSum {
			rightMaxSum = rightSum
		}
	}

	// 跨越中点的最大和 = 左半部分最大后缀 + 中点 + 右半部分最大前缀
	crossMax := leftMaxSum + nums[mid] + rightMaxSum

	// 返回三种情况中的最大值
	return max(max(leftMax, rightMax), crossMax)
}

// 解法9：暴力优化（剪枝）
// 时间复杂度：O(n²) 但实际运行更快
// 空间复杂度：O(1)
func maxSubArrayBruteForceOptimized(nums []int) int {
	maxSum := nums[0]

	for i := 0; i < len(nums); i++ {
		currentSum := 0
		// 剪枝：如果当前元素已经大于最大和，更新最大和
		if nums[i] > maxSum {
			maxSum = nums[i]
		}

		for j := i; j < len(nums); j++ {
			currentSum += nums[j]
			if currentSum > maxSum {
				maxSum = currentSum
			}
			// 剪枝：如果当前和为负数，后面的累加只会更小
			if currentSum < 0 {
				break
			}
		}
	}

	return maxSum
}

// 演示非DP解法的执行过程
func demonstrateNonDPAlgorithms() {
	fmt.Println("\n=== 非DP解法演示 ===")

	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Printf("数组: %v\n", nums)

	fmt.Println("\n1. 双指针 + 滑动窗口执行过程:")
	fmt.Println("窗口变化过程:")
	left := 0
	currentSum := 0
	maxSum := nums[0]

	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]
		fmt.Printf("  右指针到位置%d: 加入%d, 当前和=%d\n", right, nums[right], currentSum)

		if currentSum > maxSum {
			maxSum = currentSum
			fmt.Printf("    更新最大和: %d\n", maxSum)
		}

		// 收缩窗口
		oldLeft := left
		for left <= right && currentSum < 0 {
			currentSum -= nums[left]
			left++
		}
		if left > oldLeft {
			fmt.Printf("    窗口收缩: 移除位置%d-%d, 当前和=%d\n", oldLeft, left-1, currentSum)
		}
	}

	fmt.Printf("最终结果: %d\n", maxSum)

	fmt.Println("\n2. 人类思维模拟:")
	fmt.Println("思考过程:")
	currentSum = nums[0]
	maxSum = nums[0]
	fmt.Printf("  初始: 当前和=%d, 最大和=%d\n", currentSum, maxSum)

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			fmt.Printf("  位置%d: 当前和为负数，重新开始，选择%d\n", i, nums[i])
			currentSum = nums[i]
		} else {
			fmt.Printf("  位置%d: 继续累加，%d + %d = %d\n", i, currentSum, nums[i], currentSum+nums[i])
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
			fmt.Printf("    更新最大和: %d\n", maxSum)
		}
	}

	fmt.Printf("最终结果: %d\n", maxSum)
}

// 详细演示：为什么选择最新元素而不是重新计算最小的
func demonstrateWhyChooseLatest() {
	fmt.Println("\n=== 详细演示：为什么选择最新元素 ===")

	nums := []int{2, -3, 4, -1, 5}
	fmt.Printf("数组: %v\n", nums)

	fmt.Println("\n方法1: 遇到负数就重新开始（我们的方法）")
	currentSum := nums[0]
	maxSum := nums[0]
	fmt.Printf("初始: currentSum = %d, maxSum = %d\n", currentSum, maxSum)

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			fmt.Printf("位置%d: 当前和为负数(%d)，重新开始，选择%d\n", i, currentSum, nums[i])
			currentSum = nums[i]
		} else {
			fmt.Printf("位置%d: 继续累加，%d + %d = %d\n", i, currentSum, nums[i], currentSum+nums[i])
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
			fmt.Printf("  更新最大和: %d\n", maxSum)
		}
	}
	fmt.Printf("最终结果: %d\n", maxSum)

	fmt.Println("\n方法2: 遇到负数就找最小的重新开始（假设的方法）")
	currentSum = nums[0]
	maxSum = nums[0]
	fmt.Printf("初始: currentSum = %d, maxSum = %d\n", currentSum, maxSum)

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			// 假设我们找最小的重新开始
			minInRemaining := nums[i]
			for j := i + 1; j < len(nums); j++ {
				if nums[j] < minInRemaining {
					minInRemaining = nums[j]
				}
			}
			fmt.Printf("位置%d: 当前和为负数(%d)，找剩余最小元素%d，重新开始\n", i, currentSum, minInRemaining)
			currentSum = minInRemaining
		} else {
			fmt.Printf("位置%d: 继续累加，%d + %d = %d\n", i, currentSum, nums[i], currentSum+nums[i])
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
			fmt.Printf("  更新最大和: %d\n", maxSum)
		}
	}
	fmt.Printf("最终结果: %d\n", maxSum)

	fmt.Println("\n分析:")
	fmt.Println("1. 我们的方法：遇到负数就选择当前位置的元素")
	fmt.Println("2. 假设的方法：遇到负数就找剩余元素中的最小值")
	fmt.Println("3. 问题：找最小值需要额外遍历，时间复杂度变成O(n²)")
	fmt.Println("4. 而且：选择最小值不一定比选择当前位置的元素更好")

	fmt.Println("\n举例说明为什么选择当前位置更好:")
	fmt.Println("数组: [2, -3, 4, -1, 5]")
	fmt.Println("位置1: 当前和=2, 遇到-3, 和变成-1")
	fmt.Println("位置2: 当前和=-1(负数), 重新开始")
	fmt.Println("  选择位置2的4: 当前和=4 ✅")
	fmt.Println("  选择剩余最小-1: 当前和=-1 ❌")
	fmt.Println("  选择剩余最小-1: 当前和=-1 ❌")
	fmt.Println("结论：选择当前位置的元素4比选择最小值-1更好！")
}

// 演示不同策略的对比
func demonstrateDifferentStrategies() {
	fmt.Println("\n=== 不同策略对比 ===")

	testCases := [][]int{
		{2, -3, 4, -1, 5},     // 测试用例1
		{1, -2, 3, -4, 5},     // 测试用例2
		{-1, 2, -3, 4, -5, 6}, // 测试用例3
	}

	for i, nums := range testCases {
		fmt.Printf("\n测试用例 %d: %v\n", i+1, nums)

		// 策略1: 遇到负数选择当前位置
		result1 := strategyChooseCurrent(nums)

		// 策略2: 遇到负数选择剩余最小
		result2 := strategyChooseMin(nums)

		fmt.Printf("策略1(选择当前位置): %d\n", result1)
		fmt.Printf("策略2(选择剩余最小): %d\n", result2)

		if result1 >= result2 {
			fmt.Println("✅ 策略1更好或相等")
		} else {
			fmt.Println("❌ 策略2更好")
		}
	}
}

// 策略1: 遇到负数选择当前位置
func strategyChooseCurrent(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i] // 选择当前位置
		} else {
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

// 策略2: 遇到负数选择剩余最小（仅用于演示）
func strategyChooseMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			// 找剩余元素中的最小值
			minInRemaining := nums[i]
			for j := i + 1; j < len(nums); j++ {
				if nums[j] < minInRemaining {
					minInRemaining = nums[j]
				}
			}
			currentSum = minInRemaining
		} else {
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	// 演示算法
	demonstrateAlgorithm()

	// 测试用例
	testCases := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4}, // 预期输出：6
		{1},                             // 预期输出：1
		{5, 4, -1, 7, 8},                // 预期输出：23
		{-1},                            // 预期输出：-1
		{-2, -1},                        // 预期输出：-1
		{1, 2, 3, 4, 5},                 // 预期输出：15
	}

	fmt.Println("\n=== 测试结果 ===")
	for i, nums := range testCases {
		fmt.Printf("\n测试用例 %d:\n", i+1)
		fmt.Printf("输入: nums = %v\n", nums)

		// 复制数组用于测试
		nums1 := make([]int, len(nums))
		nums2 := make([]int, len(nums))
		nums3 := make([]int, len(nums))
		nums4 := make([]int, len(nums))
		nums5 := make([]int, len(nums))
		nums6 := make([]int, len(nums))
		nums7 := make([]int, len(nums))
		nums8 := make([]int, len(nums))
		nums9 := make([]int, len(nums))
		copy(nums1, nums)
		copy(nums2, nums)
		copy(nums3, nums)
		copy(nums4, nums)
		copy(nums5, nums)
		copy(nums6, nums)
		copy(nums7, nums)
		copy(nums8, nums)
		copy(nums9, nums)

		// 测试各种解法
		result1 := maxSubArrayDP(nums1)
		result2 := maxSubArrayDivide(nums2)
		result3 := maxSubArrayBruteForce(nums3)
		result4 := maxSubArrayOptimized(nums4)
		result5 := maxSubArrayTwoPointers(nums5)
		result6 := maxSubArrayPrefixSum(nums6)
		result7 := maxSubArrayHuman(nums7)
		result8 := maxSubArrayDivideSimple(nums8)
		result9 := maxSubArrayBruteForceOptimized(nums9)

		fmt.Printf("动态规划结果: %d\n", result1)
		fmt.Printf("分治法结果: %d\n", result2)
		fmt.Printf("暴力解法结果: %d\n", result3)
		fmt.Printf("优化DP结果: %d\n", result4)
		fmt.Printf("双指针结果: %d\n", result5)
		fmt.Printf("前缀和结果: %d\n", result6)
		fmt.Printf("人类思维结果: %d\n", result7)
		fmt.Printf("分治法简化结果: %d\n", result8)
		fmt.Printf("暴力优化结果: %d\n", result9)

		// 验证结果一致性
		if result1 == result2 && result2 == result3 && result3 == result4 && result4 == result5 && result5 == result6 && result6 == result7 && result7 == result8 && result8 == result9 {
			fmt.Println("✅ 所有解法结果一致")
		} else {
			fmt.Println("❌ 解法结果不一致")
		}
	}

	// 演示非DP解法
	demonstrateNonDPAlgorithms()

	// 演示详细演示
	demonstrateWhyChooseLatest()

	// 演示不同策略对比
	demonstrateDifferentStrategies()
}
