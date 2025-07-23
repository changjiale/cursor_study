package main

// 回溯法
func test(nums []int) [][]int {

	var result [][]int

	var back func(path []int, visit map[int]bool)
	back = func(path []int, visit map[int]bool) {
		if len(path) == len(nums) {
			//找到一个全排列
			//创建副本
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		for i, _ := range nums {
			if !visit[i] {
				visit[i] = true
				path = append(path, nums[i])

				//处理剩余的数字
				back(path, visit)

				//回退
				visit[i] = false
				path = path[:len(path)-1]
			}

		}
	}

	visit := make(map[int]bool, len(nums))
	back([]int{}, visit)

	return result

}
