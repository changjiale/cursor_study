package main

/*
*定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

	示例 1：

	输入：root = [1,null,2,3]
	输出：[1,3,2]
	示例 2：

	输入：root = []
*/
func main() {

}

// 解法一：递归（推荐）
// 核心思想：递归实现中序遍历，先遍历左子树，再访问根节点，最后遍历右子树
// 时间复杂度：O(n)
// 空间复杂度：O(h)，h为树的高度（递归调用栈的深度）
func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorderRecursive(root, &result)
	return result
}

func inorderRecursive(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	// 中序遍历：左 -> 根 -> 右
	inorderRecursive(node.Left, result)  // 遍历左子树
	*result = append(*result, node.Val)  // 访问根节点
	inorderRecursive(node.Right, result) // 遍历右子树
}
