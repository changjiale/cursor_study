package main

func test(nums []int) int {
	minPrice := nums[0]
	maxSum := 0

	for i := 1; i < len(nums); i++ {

		if nums[i]-minPrice > maxSum {
			maxSum = nums[i] - minPrice
		}

		if minPrice > nums[i] {
			minPrice = nums[i]
		}
	}

	return maxSum
}
