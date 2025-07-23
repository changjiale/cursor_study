package main

func sum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {

		if pos, exist := numMap[target-num]; exist {
			return []int{i, pos}
		}
		numMap[num] = i
	}
	return nil
}
