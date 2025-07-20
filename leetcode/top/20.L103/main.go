package main

/*
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]

提示：

树中节点数目在范围 [0, 2000] 内
-100 <= Node.val <= 100
*/
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解法一：BFS + 双端队列（推荐）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		levelNodes := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 根据层级决定插入位置
			if level%2 == 0 {
				// 偶数层：从左到右
				levelNodes[i] = node.Val
			} else {
				// 奇数层：从右到左
				levelNodes[levelSize-1-i] = node.Val
			}

			// 添加子节点到队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelNodes)
		level++
	}

	return result
}
