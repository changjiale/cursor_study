package main

/*
*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]

提示：

树中节点数目在范围 [0, 2000] 内
-1000 <= Node.val <= 1000
*/
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leverOrder(root *TreeNode) [][]int {
	arr := [][]int{}
	depth := 0
	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)
		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}
	order(root, depth)

	return arr
}

func levelOrderV2(root *TreeNode) [][]int {
	arr := [][]int{}

	depth := 0

	var order func(root *TreeNode, depth int)

	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)

		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}

	order(root, depth)

	return arr
}

// 前序遍历：根 -> 左 -> 右
func preorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 先访问根节点
		result = append(result, node.Val)
		// 再访问左子树
		dfs(node.Left)
		// 最后访问右子树
		dfs(node.Right)
	}

	dfs(root)
	return result
}

// 中序遍历：左 -> 根 -> 右
func inorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 先访问左子树
		dfs(node.Left)
		// 再访问根节点
		result = append(result, node.Val)
		// 最后访问右子树
		dfs(node.Right)
	}

	dfs(root)
	return result
}

// 后序遍历：左 -> 右 -> 根
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 先访问左子树
		dfs(node.Left)
		// 再访问右子树
		dfs(node.Right)
		// 最后访问根节点
		result = append(result, node.Val)
	}

	dfs(root)
	return result
}
