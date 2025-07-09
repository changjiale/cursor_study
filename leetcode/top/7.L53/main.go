package main

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

// 核心思想：dp[i] = max(nums[i], dp[i-1] + nums[i])
// 最大和的连续子数组
func main() {

}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
