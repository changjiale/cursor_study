package main

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
func main() {

}

func sort(nums []int) {
	if len(nums) < 2 {
		return
	}
	base := nums[len(nums)-1]

	i := 0
	for j := 0; j < len(nums)-1; j++ {
		if nums[j] < base {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
	sort(nums[:i])
	sort(nums[i+1:])
}
