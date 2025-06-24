package main

import "leetcode/util"

func main() {
	println(util.MustMarshalToString(longestConsecutive([]int{1, 3, 5, 7, 9})))
}

func longestConsecutive(nums []int) (ans int) {

	baseMap := map[int]bool{}
	for _, num := range nums {
		baseMap[num] = true
	}
	for num, _ := range baseMap {
		//避免重复计算
		if baseMap[num-1] {
			continue
		}

		y := num + 1
		for baseMap[y] {
			y++
		}

		ans = max(ans, y-num)

	}
	return ans

}
