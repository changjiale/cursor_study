package main

/*
*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例 3：

输入：nums = [1]
输出：[[1]]

提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/
func main() {

}

// 回溯法
func permute(nums []int) [][]int {
	var result [][]int
	var backtrack func(path []int, used []bool)

	backtrack = func(path []int, used []bool) {
		// 如果路径长度等于数组长度，说明找到一个排列
		if len(path) == len(nums) {
			// 创建path的副本，避免后续修改影响结果
			perm := make([]int, len(path))
			copy(perm, path)
			result = append(result, perm)
			return
		}

		// 尝试将每个未使用的数字加入当前路径
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				// 标记当前数字为已使用
				used[i] = true
				// 将当前数字加入路径
				path = append(path, nums[i])

				// 递归处理剩余数字
				backtrack(path, used)

				// 回溯：恢复状态
				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}

	// 初始化used数组，记录每个数字是否被使用
	used := make([]bool, len(nums))
	backtrack([]int{}, used)

	return result
}
