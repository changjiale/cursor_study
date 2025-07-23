package main

func sort1(nums []int) {
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
	sort1(nums[:i])
	sort1(nums[i+1:])
}
