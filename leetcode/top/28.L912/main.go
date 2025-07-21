package main

/**
给你一个整数数组 nums，请你将该数组升序排列。

你必须在 不使用任何内置函数 的情况下解决问题，时间复杂度为 O(nlog(n))，并且空间复杂度尽可能小。



示例 1：

输入：nums = [5,2,3,1]
输出：[1,2,3,5]
解释：数组排序后，某些数字的位置没有改变（例如，2 和 3），而其他数字的位置发生了改变（例如，1 和 5）。
示例 2：

输入：nums = [5,1,1,2,0,0]
输出：[0,0,1,1,2,5]
解释：请注意，nums 的值不一定唯一。
*/

func main() {

}

//分区过程示意：
//1. 选择基准值
//2. 将小于基准值的元素移到左边
//3. 将大于基准值的元素移到右边
//4. 递归处理左右两部分

// 时间复杂度：平均 O(nlogn)，最坏 O(n²)
// 空间复杂度：O(logn)
func quickSort1(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 选择基准值
	pivot := nums[len(nums)-1]
	// 分区
	i := 0
	for j := 0; j < len(nums)-1; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	// 将基准值放到正确位置
	nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
	// 递归排序左右两部分
	quickSort1(nums[:i])
	quickSort1(nums[i+1:])
}
