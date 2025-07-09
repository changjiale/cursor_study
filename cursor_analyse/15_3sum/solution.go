package main

import (
	"fmt"
	"sort"
)

/*
图例说明：

1. 解法一：排序 + 双指针
   原始数组：[-1,0,1,2,-1,-4]
   排序后：  [-4,-1,-1,0,1,2]

   第一次遍历(i=0)：
   -4 + (-1) + 2 = -3 < 0，右指针左移
   -4 + (-1) + 1 = -4 < 0，右指针左移
   -4 + (-1) + 0 = -5 < 0，右指针左移
   -4 + (-1) + (-1) = -6 < 0，i++

   第二次遍历(i=1)：
   -1 + (-1) + 2 = 0，找到一组解
   -1 + 0 + 1 = 0，找到一组解

   指针移动示意：
   i=0: [-4,-1,-1,0,1,2]
        i  l        r
   i=1: [-4,-1,-1,0,1,2]
           i  l     r
   i=1: [-4,-1,-1,0,1,2]
           i     l  r

2. 解法二：哈希表
   原始数组：[-1,0,1,2,-1,-4]
   排序后：  [-4,-1,-1,0,1,2]

   第一次遍历(i=0)：
   target = 4
   hash表：{-1:true, -1:true, 0:true, 1:true, 2:true}
   检查：4-(-1)=5, 4-(-1)=5, 4-0=4, 4-1=3, 4-2=2
   无解

   第二次遍历(i=1)：
   target = 1
   hash表：{-1:true, 0:true, 1:true, 2:true}
   检查：1-(-1)=2, 1-0=1, 1-1=0, 1-2=-1
   找到解：[-1,0,1]

3. 解法三：极简写法（双指针优化版）
   原始数组：[-1,0,1,2,-1,-4]
   排序后：  [-4,-1,-1,0,1,2]

   优化点：
   1. 提前判断：if nums[i] > 0 { break }
   2. 跳过重复：if i > 0 && nums[i] == nums[i-1] { continue }
   3. 合并指针移动：left++, right--

   执行过程：
   i=0: [-4,-1,-1,0,1,2]
        i  l        r
   i=1: [-4,-1,-1,0,1,2]
           i  l     r
   i=1: [-4,-1,-1,0,1,2]
           i     l  r
*/

// 解法一：排序 + 双指针
/*
图例说明：
原始数组：[-1,0,1,2,-1,-4]
排序后：  [-4,-1,-1,0,1,2]

第一次遍历(i=0)：
-4 + (-1) + 2 = -3 < 0，右指针左移
-4 + (-1) + 1 = -4 < 0，右指针左移
-4 + (-1) + 0 = -5 < 0，右指针左移
-4 + (-1) + (-1) = -6 < 0，i++

第二次遍历(i=1)：
-1 + (-1) + 2 = 0，找到一组解
-1 + 0 + 1 = 0，找到一组解

指针移动示意：
i=0: [-4,-1,-1,0,1,2]
     i  l        r
i=1: [-4,-1,-1,0,1,2]
        i  l     r
i=1: [-4,-1,-1,0,1,2]
        i     l  r
*/
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
func threeSum1(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	// 排序
	sort.Ints(nums)
	var result [][]int

	// 遍历数组
	/**
	我们需要找三个数：nums[i] + nums[left] + nums[right] = 0
	i 是第一个数
	left 和 right 是另外两个数
	如果 i 遍历到 n-1，那么 left 和 right 就没有位置了
	所以 i 最多只能到 n-3，即 i < n-2
		**/
	for i := 0; i < n-2; i++ {
		// 跳过重复元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 如果当前元素大于0，后面的元素都大于0，不可能和为0
		if nums[i] > 0 {
			break
		}

		// 双指针
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				// 找到一组解
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 跳过重复元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

// 解法二：哈希表（不推荐，因为需要处理重复元素）
/*
图例说明：
原始数组：[-1,0,1,2,-1,-4]
排序后：  [-4,-1,-1,0,1,2]

第一次遍历(i=0)：
target = 4
hash表：{-1:true, -1:true, 0:true, 1:true, 2:true}
检查：4-(-1)=5, 4-(-1)=5, 4-0=4, 4-1=3, 4-2=2
无解

第二次遍历(i=1)：
target = 1
hash表：{-1:true, 0:true, 1:true, 2:true}
检查：1-(-1)=2, 1-0=1, 1-1=0, 1-2=-1
找到解：[-1,0,1]

哈希表变化示意：
i=0: {2:true, 1:true, 0:true, -1:true, -4:true}
i=1: {2:true, 1:true, 0:true, -1:true}
*/
// 时间复杂度：O(n²)
// 空间复杂度：O(n)
func threeSum2(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}

	// 排序
	sort.Ints(nums)
	var result [][]int
	seen := make(map[string]bool)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		hash := make(map[int]bool)

		for j := i + 1; j < n; j++ {
			complement := target - nums[j]
			if hash[complement] {
				// 生成唯一key
				key := fmt.Sprintf("%d,%d,%d", nums[i], complement, nums[j])
				if !seen[key] {
					result = append(result, []int{nums[i], complement, nums[j]})
					seen[key] = true
				}
			}
			hash[nums[j]] = true
		}
	}

	return result
}

// 解法三：极简写法（双指针）
/*
图例说明：
原始数组：[-1,0,1,2,-1,-4]
排序后：  [-4,-1,-1,0,1,2]

优化点：
1. 提前判断：if nums[i] > 0 { break }
2. 跳过重复：if i > 0 && nums[i] == nums[i-1] { continue }
3. 合并指针移动：left++, right--

执行过程：
i=0: [-4,-1,-1,0,1,2]
     i  l        r
i=1: [-4,-1,-1,0,1,2]
        i  l     r
i=1: [-4,-1,-1,0,1,2]
        i     l  r

优化效果：
1. 减少不必要的遍历
2. 简化代码结构
3. 提高代码可读性
*/
// 时间复杂度：O(n²)
// 空间复杂度：O(1)
func threeSumSimple(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func main() {
	// 测试用例
	testCases := [][]int{
		{-1, 0, 1, 2, -1, -4}, // 预期输出：[[-1,-1,2],[-1,0,1]]
		{0, 1, 1},             // 预期输出：[]
		{0, 0, 0},             // 预期输出：[[0,0,0]]
	}

	for _, nums := range testCases {
		fmt.Printf("\n输入: nums = %v\n", nums)

		// 解法一
		nums1 := make([]int, len(nums))
		copy(nums1, nums)
		fmt.Printf("解法一（双指针）结果: %v\n", threeSum1(nums1))

		// 解法二
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		fmt.Printf("解法二（哈希表）结果: %v\n", threeSum2(nums2))

		// 解法三
		nums3 := make([]int, len(nums))
		copy(nums3, nums)
		fmt.Printf("解法三（极简）结果: %v\n", threeSumSimple(nums3))
	}
}
