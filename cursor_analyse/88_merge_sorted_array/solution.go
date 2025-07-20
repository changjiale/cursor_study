package main

import "fmt"

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
注意：因为 m = 0，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
*/

// 解法一：从后往前合并（推荐）
// 时间复杂度：O(m + n)
// 空间复杂度：O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 从后往前合并，避免覆盖nums1中的有效元素
	i := m - 1     // nums1的有效元素指针
	j := n - 1     // nums2的指针
	k := m + n - 1 // 合并后数组的指针

	// 从后往前比较，将较大的元素放到nums1的末尾
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// 如果nums2还有剩余元素，需要复制到nums1中
	// 如果nums1还有剩余元素，它们已经在正确的位置上了
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// 解法二：从前往后合并（需要额外空间）
// 时间复杂度：O(m + n)
// 空间复杂度：O(m)
func merge2(nums1 []int, m int, nums2 []int, n int) {
	// 先保存nums1的有效元素
	temp := make([]int, m)
	copy(temp, nums1[:m])

	i := 0 // temp的指针
	j := 0 // nums2的指针
	k := 0 // nums1的指针

	// 从前往后合并
	for i < m && j < n {
		if temp[i] <= nums2[j] {
			nums1[k] = temp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	// 复制剩余元素
	for i < m {
		nums1[k] = temp[i]
		i++
		k++
	}

	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}

// 解法三：使用sort包（不推荐，仅用于理解）
// 时间复杂度：O((m+n) * log(m+n))
// 空间复杂度：O(1)
func merge3(nums1 []int, m int, nums2 []int, n int) {
	// 将nums2的元素复制到nums1的末尾
	copy(nums1[m:], nums2)

	// 对整个nums1进行排序
	// 注意：这里只是演示，实际应该使用双指针方法
	// sort.Ints(nums1[:m+n])

	// 手动实现简单的排序（冒泡排序，仅用于演示）
	for i := 0; i < m+n-1; i++ {
		for j := 0; j < m+n-1-i; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}
}

// 解法四：双指针优化版本
// 时间复杂度：O(m + n)
// 空间复杂度：O(1)
func merge4(nums1 []int, m int, nums2 []int, n int) {
	// 如果nums2为空，nums1已经是正确结果
	if n == 0 {
		return
	}

	// 如果nums1为空，直接复制nums2
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	// 从后往前合并
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// 只需要处理nums2的剩余元素
	// nums1的剩余元素已经在正确位置
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}

func main() {
	// 测试用例
	testCases := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
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
	}

	fmt.Println("=== 解法一：从后往前合并（推荐）===")
	for i, tc := range testCases {
		nums1 := make([]int, len(tc.nums1))
		copy(nums1, tc.nums1)

		merge(nums1, tc.m, tc.nums2, tc.n)

		fmt.Printf("测试用例 %d:\n", i+1)
		fmt.Printf("输入: nums1=%v, m=%d, nums2=%v, n=%d\n", tc.nums1, tc.m, tc.nums2, tc.n)
		fmt.Printf("输出: %v\n", nums1)
		fmt.Printf("期望: %v\n", tc.expect)
		fmt.Printf("通过: %t\n\n", compareSlices(nums1, tc.expect))
	}

	fmt.Println("=== 解法二：从前往后合并（需要额外空间）===")
	for i, tc := range testCases {
		nums1 := make([]int, len(tc.nums1))
		copy(nums1, tc.nums1)

		merge2(nums1, tc.m, tc.nums2, tc.n)

		fmt.Printf("测试用例 %d:\n", i+1)
		fmt.Printf("输入: nums1=%v, m=%d, nums2=%v, n=%d\n", tc.nums1, tc.m, tc.nums2, tc.n)
		fmt.Printf("输出: %v\n", nums1)
		fmt.Printf("期望: %v\n", tc.expect)
		fmt.Printf("通过: %t\n\n", compareSlices(nums1, tc.expect))
	}

	fmt.Println("=== 解法三：使用排序（仅用于理解）===")
	for i, tc := range testCases {
		nums1 := make([]int, len(tc.nums1))
		copy(nums1, tc.nums1)

		merge3(nums1, tc.m, tc.nums2, tc.n)

		fmt.Printf("测试用例 %d:\n", i+1)
		fmt.Printf("输入: nums1=%v, m=%d, nums2=%v, n=%d\n", tc.nums1, tc.m, tc.nums2, tc.n)
		fmt.Printf("输出: %v\n", nums1)
		fmt.Printf("期望: %v\n", tc.expect)
		fmt.Printf("通过: %t\n\n", compareSlices(nums1, tc.expect))
	}

	fmt.Println("=== 解法四：双指针优化版本 ===")
	for i, tc := range testCases {
		nums1 := make([]int, len(tc.nums1))
		copy(nums1, tc.nums1)

		merge4(nums1, tc.m, tc.nums2, tc.n)

		fmt.Printf("测试用例 %d:\n", i+1)
		fmt.Printf("输入: nums1=%v, m=%d, nums2=%v, n=%d\n", tc.nums1, tc.m, tc.nums2, tc.n)
		fmt.Printf("输出: %v\n", nums1)
		fmt.Printf("期望: %v\n", tc.expect)
		fmt.Printf("通过: %t\n\n", compareSlices(nums1, tc.expect))
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
解题思路：

1. 从后往前合并（推荐）：
   - 利用nums1末尾有足够的空间
   - 从后往前比较，避免覆盖nums1中的有效元素
   - 将较大的元素放到nums1的末尾
   - 时间复杂度O(m+n)，空间复杂度O(1)

2. 从前往后合并（需要额外空间）：
   - 先保存nums1的有效元素到临时数组
   - 从前往后合并，需要额外空间
   - 时间复杂度O(m+n)，空间复杂度O(m)

3. 使用排序（不推荐）：
   - 将nums2复制到nums1末尾，然后排序
   - 时间复杂度O((m+n)*log(m+n))
   - 仅用于理解，实际不推荐

4. 双指针优化版本：
   - 优化边界条件处理
   - 减少不必要的操作
   - 代码更简洁高效

时间复杂度分析：
- 解法一、二、四：O(m+n)，每个元素最多访问一次
- 解法三：O((m+n)*log(m+n))，排序的时间复杂度

空间复杂度分析：
- 解法一、三、四：O(1)，只使用常数额外空间
- 解法二：O(m)，需要临时数组保存nums1的有效元素

关键点：
1. 利用nums1末尾的空间，避免覆盖有效元素
2. 从后往前合并是最优解
3. 注意边界条件的处理
4. 理解不同解法的适用场景
*/
