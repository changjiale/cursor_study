package main

import (
	"fmt"
	"strings"
)

/*
题目：合并两个有序数组
难度：简单
标签：数组、双指针、排序

题目描述：
给你两个按非递减顺序排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n，分别表示 nums1 和 nums2 中的元素数目。
请你合并 nums2 到 nums1 中，使合并后的数组同样按非递减顺序排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0，应忽略。nums2 的长度为 n。

要求：
- 时间复杂度：O(m + n)
- 空间复杂度：O(1)

示例：
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6]。
合并结果是 [1,2,2,3,5,6]，其中斜体加粗标注的为 nums1 中的元素。

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 []。
合并结果是 [1]。

输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并 [] 和 [1]。
合并结果是 [1]。

提示：
1. 利用nums1末尾有足够的空间来存储合并结果
2. 从后往前合并可以避免覆盖nums1中的有效元素
3. 使用双指针技术，一个指向nums1的有效元素，一个指向nums2
4. 注意边界条件的处理（空数组、单个元素等）
5. 合并完成后，nums1就是最终结果，不需要返回值
*/

// TODO: 在这里实现你的算法
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 请实现你的代码
	// 注意：直接修改nums1，不需要返回值
}

// 测试用例结构
type TestCase struct {
	nums1  []int
	m      int
	nums2  []int
	n      int
	expect []int
}

func main() {
	// 测试用例
	testCases := []TestCase{
		// 基本测试用例
		{
			[]int{1, 2, 3, 0, 0, 0}, 3,
			[]int{2, 5, 6}, 3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			[]int{1}, 1,
			[]int{}, 0,
			[]int{1},
		},
		{
			[]int{0}, 0,
			[]int{1}, 1,
			[]int{1},
		},

		// 边界测试用例
		{
			[]int{4, 5, 6, 0, 0, 0}, 3,
			[]int{1, 2, 3}, 3,
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 3, 5, 0, 0}, 3,
			[]int{2, 4}, 2,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{0, 0, 0}, 0,
			[]int{1, 2, 3}, 3,
			[]int{1, 2, 3},
		},

		// 特殊情况测试用例
		{
			[]int{1, 2, 3, 4, 5, 0, 0, 0}, 5,
			[]int{6, 7, 8}, 3,
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{6, 7, 8, 0, 0, 0}, 3,
			[]int{1, 2, 3}, 3,
			[]int{1, 2, 3, 6, 7, 8},
		},
		{
			[]int{1, 1, 1, 0, 0}, 3,
			[]int{1, 1}, 2,
			[]int{1, 1, 1, 1, 1},
		},
	}

	fmt.Println("开始测试合并两个有序数组...")
	fmt.Println(strings.Repeat("=", 50))

	passed := 0
	total := len(testCases)

	for i, tc := range testCases {
		fmt.Printf("\n测试用例 %d:\n", i+1)
		fmt.Printf("输入: nums1=%v, m=%d, nums2=%v, n=%d\n", tc.nums1, tc.m, tc.nums2, tc.n)

		// 创建nums1的副本，避免修改原数据
		nums1 := make([]int, len(tc.nums1))
		copy(nums1, tc.nums1)

		// 调用合并函数
		merge(nums1, tc.m, tc.nums2, tc.n)

		fmt.Printf("输出: %v\n", nums1)
		fmt.Printf("期望: %v\n", tc.expect)

		// 检查结果
		if compareSlices(nums1, tc.expect) {
			fmt.Printf("✅ 通过\n")
			passed++
		} else {
			fmt.Printf("❌ 失败\n")
		}
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("测试结果: %d/%d 通过\n", passed, total)

	if passed == total {
		fmt.Println("🎉 所有测试用例都通过了！")
		fmt.Println("\n💡 提示：")
		fmt.Println("- 从后往前合并是最优解，避免覆盖nums1中的有效元素")
		fmt.Println("- 利用nums1末尾的空间来存储合并结果")
		fmt.Println("- 注意边界条件的处理")
		fmt.Println("- 时间复杂度O(m+n)，空间复杂度O(1)")
	} else {
		fmt.Println("💡 还有一些测试用例没有通过，请检查你的实现。")
	}
}

// 比较两个切片是否相等
func compareSlices(a, b []int) bool {
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

/*
预期输出：
开始测试合并两个有序数组...
==================================================

测试用例 1:
输入: nums1=[1 2 3 0 0 0], m=3, nums2=[2 5 6], n=3
输出: [1 2 2 3 5 6]
期望: [1 2 2 3 5 6]
✅ 通过

测试用例 2:
输入: nums1=[1], m=1, nums2=[], n=0
输出: [1]
期望: [1]
✅ 通过

测试用例 3:
输入: nums1=[0], m=0, nums2=[1], n=1
输出: [1]
期望: [1]
✅ 通过

测试用例 4:
输入: nums1=[4 5 6 0 0 0], m=3, nums2=[1 2 3], n=3
输出: [1 2 3 4 5 6]
期望: [1 2 3 4 5 6]
✅ 通过

测试用例 5:
输入: nums1=[1 3 5 0 0], m=3, nums2=[2 4], n=2
输出: [1 2 3 4 5]
期望: [1 2 3 4 5]
✅ 通过

测试用例 6:
输入: nums1=[0 0 0], m=0, nums2=[1 2 3], n=3
输出: [1 2 3]
期望: [1 2 3]
✅ 通过

测试用例 7:
输入: nums1=[1 2 3 4 5 0 0 0], m=5, nums2=[6 7 8], n=3
输出: [1 2 3 4 5 6 7 8]
期望: [1 2 3 4 5 6 7 8]
✅ 通过

测试用例 8:
输入: nums1=[6 7 8 0 0 0], m=3, nums2=[1 2 3], n=3
输出: [1 2 3 6 7 8]
期望: [1 2 3 6 7 8]
✅ 通过

测试用例 9:
输入: nums1=[1 1 1 0 0], m=3, nums2=[1 1], n=2
输出: [1 1 1 1 1]
期望: [1 1 1 1 1]
✅ 通过

==================================================
测试结果: 9/9 通过
🎉 所有测试用例都通过了！

💡 提示：
- 从后往前合并是最优解，避免覆盖nums1中的有效元素
- 利用nums1末尾的空间来存储合并结果
- 注意边界条件的处理
- 时间复杂度O(m+n)，空间复杂度O(1)
*/
