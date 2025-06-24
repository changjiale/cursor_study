package main

import (
	"fmt"
	"math/rand"
)

// 解法一：基础快速排序
/*
图例说明：
原始数组：[3,2,1,5,6,4]

第一次分区(pivot=4)：
[3,2,1,4,6,5]
     ^  ^
    i  j

第二次分区(左半部分，pivot=1)：
[1,2,3,4,6,5]
 ^  ^
i  j

第二次分区(右半部分，pivot=5)：
[1,2,3,4,5,6]
        ^  ^
        i  j

分区过程示意：
1. 选择基准值
2. 将小于基准值的元素移到左边
3. 将大于基准值的元素移到右边
4. 递归处理左右两部分
*/
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

// 解法二：随机化快速排序
/*
图例说明：
原始数组：[3,2,1,5,6,4]

随机选择基准值：
第一次：pivot = 2
[1,2,3,5,6,4]
 ^  ^
i  j

第二次：pivot = 5
[1,2,3,4,5,6]
        ^  ^
        i  j

随机化优势：
1. 避免最坏情况
2. 提高平均性能
3. 减少对输入数据的依赖
*/
// 时间复杂度：平均 O(nlogn)
// 空间复杂度：O(logn)
func quickSort2(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 随机选择基准值
	pivotIndex := rand.Intn(len(nums))
	nums[pivotIndex], nums[len(nums)-1] = nums[len(nums)-1], nums[pivotIndex]
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
	quickSort2(nums[:i])
	quickSort2(nums[i+1:])
}

// 解法三：三路快速排序
/*
图例说明：
原始数组：[3,2,1,5,6,4]

三路分区过程：
lt: 小于基准值的右边界
gt: 大于基准值的左边界
i: 当前扫描位置

初始状态：
[3,2,1,5,6,4]
 ^
lt,i
     ^
    gt

扫描过程：
[1,2,3,5,6,4]
   ^
  lt,i
     ^
    gt

最终状态：
[1,2,3,4,5,6]
     ^
    lt
     ^
    gt
*/
// 时间复杂度：平均 O(nlogn)
// 空间复杂度：O(logn)
func quickSort3(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 三路分区
	lt, gt := 0, len(nums)-1
	pivot := nums[0]
	i := 1
	for i <= gt {
		if nums[i] < pivot {
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		} else if nums[i] > pivot {
			nums[gt], nums[i] = nums[i], nums[gt]
			gt--
		} else {
			i++
		}
	}
	// 递归排序左右两部分
	quickSort3(nums[:lt])
	quickSort3(nums[gt+1:])
}

// 解法四：极简写法
/*
图例说明：
原始数组：[3,2,1,5,6,4]

极简写法特点：
1. 使用切片操作
2. 代码更简洁
3. 逻辑更清晰

执行过程：
[3,2,1,5,6,4] -> [1,2,3,4,5,6]
 ^     ^
pivot  i
*/
// 时间复杂度：平均 O(nlogn)
// 空间复杂度：O(logn)
func quickSortSimple(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[0]
	var left, right []int
	for _, num := range nums[1:] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	left = quickSortSimple(left)
	right = quickSortSimple(right)
	return append(append(left, pivot), right...)
}

func main() {
	// 测试用例
	testCases := [][]int{
		{3, 2, 1, 5, 6, 4},
		{1, 1, 1, 1, 1},
		{5, 4, 3, 2, 1},
		{1},
		{},
	}

	for _, nums := range testCases {
		fmt.Printf("\n输入: nums = %v\n", nums)

		// 解法一
		nums1 := make([]int, len(nums))
		copy(nums1, nums)
		quickSort1(nums1)
		fmt.Printf("解法一（基础快速排序）结果: %v\n", nums1)

		// 解法二
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		quickSort2(nums2)
		fmt.Printf("解法二（随机化快速排序）结果: %v\n", nums2)

		// 解法三
		nums3 := make([]int, len(nums))
		copy(nums3, nums)
		quickSort3(nums3)
		fmt.Printf("解法三（三路快速排序）结果: %v\n", nums3)

		// 解法四
		nums4 := make([]int, len(nums))
		copy(nums4, nums)
		result := quickSortSimple(nums4)
		fmt.Printf("解法四（极简写法）结果: %v\n", result)
	}
}
