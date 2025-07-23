package main

/*
*给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例 1：

输入：root = [1,2,3,null,5,null,4]

输出：[1,3,4]

解释：

示例 2：

输入：root = [1,2,3,4,null,null,null,5]

输出：[1,3,4,5]

解释：

示例 3：

输入：root = [1,null,3]

输出：[1,3]

示例 4：

输入：root = []

输出：[]

提示:

二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100
*/
func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 解法一：BFS（层次遍历）
// 核心思想：使用层次遍历，记录每一层的最后一个节点
// 时间复杂度：O(n)
// 空间复杂度：O(w)，w为树的最大宽度
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		// 遍历当前层的所有节点
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 如果是当前层的最后一个节点，加入结果
			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			// 添加子节点到队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}
