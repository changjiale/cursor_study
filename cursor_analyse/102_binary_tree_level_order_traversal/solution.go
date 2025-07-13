package main

import (
	"fmt"
)

/**
102. 二叉树的层序遍历
难度：中等
标签：树、广度优先搜索、二叉树

题目描述：
给你二叉树的根节点 root，返回其节点值的层序遍历。（即逐层地，从左到右访问所有节点）。

要求：
- 树中节点数目在范围 [0, 2000] 内
- -1000 <= Node.val <= 1000

示例：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]

输入：root = [1]
输出：[[1]]

输入：root = []
输出：[]

提示：
- 使用队列进行广度优先搜索
- 需要记录每一层的节点数量
*/

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解法1：队列 + 广度优先搜索（标准解法）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 核心思想：使用队列存储每一层的节点，逐层处理
func levelOrderBFS(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue) // 当前层的节点数量
		var currentLevel []int  // 当前层的节点值

		// 处理当前层的所有节点
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // 出队

			currentLevel = append(currentLevel, node.Val)

			// 将子节点加入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

// 解法2：递归 + 深度优先搜索
// 时间复杂度：O(n)
// 空间复杂度：O(h) - h为树的高度
// 核心思想：使用递归遍历，记录每个节点的层数
func levelOrderDFS(root *TreeNode) [][]int {
	var result [][]int
	dfsHelper(root, 0, &result)
	return result
}

func dfsHelper(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	// 如果是新的一层，创建新的切片
	if level >= len(*result) {
		*result = append(*result, []int{})
	}

	// 将当前节点值加入对应层
	(*result)[level] = append((*result)[level], node.Val)

	// 递归处理左右子树
	dfsHelper(node.Left, level+1, result)
	dfsHelper(node.Right, level+1, result)
}

// 递归版本V2
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

// 解法3：双队列法（优化空间使用）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 核心思想：使用两个队列交替存储不同层的节点
func levelOrderTwoQueues(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	currentQueue := []*TreeNode{root}
	nextQueue := []*TreeNode{}

	for len(currentQueue) > 0 {
		var currentLevel []int

		// 处理当前队列中的所有节点
		for len(currentQueue) > 0 {
			node := currentQueue[0]
			currentQueue = currentQueue[1:]

			currentLevel = append(currentLevel, node.Val)

			// 将子节点加入下一层队列
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}

		result = append(result, currentLevel)

		// 交换队列
		currentQueue, nextQueue = nextQueue, currentQueue
		nextQueue = nextQueue[:0] // 清空下一层队列
	}

	return result
}

// 解法4：单队列 + 标记法
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 核心思想：在队列中插入层标记，区分不同层
func levelOrderMarker(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []interface{}{root, nil} // 使用interface{}存储不同类型的值
	var currentLevel []int

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item == nil {
			// 遇到标记，表示当前层结束
			if len(currentLevel) > 0 {
				result = append(result, currentLevel)
				currentLevel = []int{}
				queue = append(queue, nil) // 为下一层添加标记
			}
		} else {
			// 处理节点
			node := item.(*TreeNode)
			currentLevel = append(currentLevel, node.Val)

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

// 辅助函数：创建二叉树
func createTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// 右子节点
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// 辅助函数：打印二叉树（用于调试）
func printTree(root *TreeNode, prefix string, isLeft bool) {
	if root == nil {
		return
	}

	fmt.Printf("%s", prefix)
	if isLeft {
		fmt.Printf("├── ")
	} else {
		fmt.Printf("└── ")
	}
	fmt.Printf("%d\n", root.Val)

	printTree(root.Left, prefix+"│   ", true)
	printTree(root.Right, prefix+"    ", false)
}

// 演示算法执行过程
func demonstrateAlgorithm() {
	fmt.Println("=== 二叉树的层序遍历演示 ===")

	// 创建测试树：[3,9,20,null,null,15,7]
	values := []interface{}{3, 9, 20, nil, nil, 15, 7}
	root := createTree(values)

	fmt.Println("二叉树结构:")
	printTree(root, "", false)

	fmt.Println("\n层序遍历过程:")
	result := levelOrderBFS(root)

	for i, level := range result {
		fmt.Printf("第%d层: %v\n", i+1, level)
	}

	fmt.Printf("\n最终结果: %v\n", result)
}

func main() {
	// 演示算法
	demonstrateAlgorithm()

	// 测试用例
	testCases := []struct {
		name   string
		values []interface{}
		want   [][]int
	}{
		{
			name:   "标准二叉树",
			values: []interface{}{3, 9, 20, nil, nil, 15, 7},
			want:   [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:   "单节点",
			values: []interface{}{1},
			want:   [][]int{{1}},
		},
		{
			name:   "空树",
			values: []interface{}{},
			want:   [][]int{},
		},
		{
			name:   "左斜树",
			values: []interface{}{1, 2, nil, 3},
			want:   [][]int{{1}, {2}, {3}},
		},
		{
			name:   "右斜树",
			values: []interface{}{1, nil, 2, nil, nil, nil, 3},
			want:   [][]int{{1}, {2}},
		},
		{
			name:   "完全二叉树",
			values: []interface{}{1, 2, 3, 4, 5, 6, 7},
			want:   [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}

	fmt.Println("\n=== 测试结果 ===")
	for i, tc := range testCases {
		fmt.Printf("\n测试用例 %d: %s\n", i+1, tc.name)
		fmt.Printf("输入: %v\n", tc.values)

		root := createTree(tc.values)

		// 测试各种解法
		result1 := levelOrderBFS(root)
		result2 := levelOrderDFS(root)
		result3 := levelOrderTwoQueues(root)
		result4 := levelOrderMarker(root)

		fmt.Printf("BFS结果: %v\n", result1)
		fmt.Printf("DFS结果: %v\n", result2)
		fmt.Printf("双队列结果: %v\n", result3)
		fmt.Printf("标记法结果: %v\n", result4)
		fmt.Printf("期望结果: %v\n", tc.want)

		// 验证结果一致性
		if compareResults(result1, tc.want) &&
			compareResults(result2, tc.want) &&
			compareResults(result3, tc.want) &&
			compareResults(result4, tc.want) {
			fmt.Println("✅ 所有解法结果正确")
		} else {
			fmt.Println("❌ 有解法结果错误")
		}
	}
}

// 辅助函数：比较两个二维切片是否相等
func compareResults(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
