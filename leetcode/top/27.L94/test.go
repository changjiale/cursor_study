package main

func test(root *TreeNode) []int {

	var result []int

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		result = append(result, root.Val)
		dfs(root.Right)

	}

	dfs(root)

	return result
}
