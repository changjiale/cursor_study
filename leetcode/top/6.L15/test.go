package main

import "sort"

/*
*
核心思想   先排序 然后区间查找
*/
func treeNumsSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return nil
	}

	n := len(nums)

	//排序
	sort.Ints(nums)
	for i := 0; i < n-2; i++ {

		if i > 0 && nums[i] == nums[i+1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		left, right := i+1, n-1
		for left < right {

			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
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
