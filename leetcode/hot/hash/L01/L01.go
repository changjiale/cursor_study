package main

import "leetcode/util"

// 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
func main() {
	ints := find([]int{1, 2, 3, 4, 5, 6}, 9)
	println(util.MustMarshalToString(ints))
}

func find(list []int, target int) []int {

	itemMap := make(map[int]int, 0)
	for i, item := range list {
		diff := target - item
		if _, has := itemMap[diff]; has {
			return []int{itemMap[diff], i}
		}
		itemMap[item] = i
	}

	return nil
}
