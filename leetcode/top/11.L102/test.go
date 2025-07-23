package main

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func levePrint(node *TreeNode1) {
	arr := [][]int{}
	depth := 0
	var dfs func(node *TreeNode1, depth int)
	dfs = func(node *TreeNode1, depth int) {

		if node == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}

		arr[depth] = append(arr[depth], node.Val)

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(node, depth)
}
