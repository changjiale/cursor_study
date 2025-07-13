package main

import (
	"fmt"
)

/*
题目：102. 二叉树的层序遍历
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
- 考虑如何处理空节点
*/

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO: 在这里实现你的算法
// 要求：实现二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	// 请实现你的代码
	return [][]int{}
}

// 辅助函数：创建二叉树（用于测试）
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

// 测试用例结构
type TestCase struct {
	name   string
	input  []interface{}
	output [][]int
}

func main() {
	// 测试用例
	testCases := []TestCase{
		{
			name:   "标准二叉树",
			input:  []interface{}{3, 9, 20, nil, nil, 15, 7},
			output: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:   "单节点",
			input:  []interface{}{1},
			output: [][]int{{1}},
		},
		{
			name:   "空树",
			input:  []interface{}{},
			output: [][]int{},
		},
		{
			name:   "左斜树",
			input:  []interface{}{1, 2, nil, 3},
			output: [][]int{{1}, {2}, {3}},
		},
		{
			name:   "右斜树",
			input:  []interface{}{1, nil, 2, nil, nil, nil, 3},
			output: [][]int{{1}, {2}},
		},
		{
			name:   "完全二叉树",
			input:  []interface{}{1, 2, 3, 4, 5, 6, 7},
			output: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name:   "只有左子树",
			input:  []interface{}{1, 2, nil, 3, nil, 4},
			output: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:   "只有右子树",
			input:  []interface{}{1, nil, 2, nil, nil, nil, 3, nil, nil, nil, nil, nil, nil, nil, 4},
			output: [][]int{{1}, {2}, {3}, {4}},
		},
	}

	fmt.Println("=== 二叉树的层序遍历测试 ===")

	allPassed := true
	for i, tc := range testCases {
		fmt.Printf("\n测试用例 %d: %s\n", i+1, tc.name)
		fmt.Printf("输入: %v\n", tc.input)

		root := createTree(tc.input)
		result := levelOrder(root)

		fmt.Printf("输出: %v\n", result)
		fmt.Printf("期望: %v\n", tc.output)

		if compareResults(result, tc.output) {
			fmt.Println("✅ 通过")
		} else {
			fmt.Println("❌ 失败")
			allPassed = false
		}
	}

	fmt.Printf("\n=== 测试总结 ===\n")
	if allPassed {
		fmt.Println("🎉 所有测试用例都通过了！")
	} else {
		fmt.Println("⚠️  有测试用例未通过，请检查你的实现")
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

/*
预期输出：
=== 二叉树的层序遍历测试 ===

测试用例 1: 标准二叉树
输入: [3 9 20 <nil> <nil> 15 7]
输出: [[3] [9 20] [15 7]]
期望: [[3] [9 20] [15 7]]
✅ 通过

测试用例 2: 单节点
输入: [1]
输出: [[1]]
期望: [[1]]
✅ 通过

测试用例 3: 空树
输入: []
输出: []
期望: []
✅ 通过

测试用例 4: 左斜树
输入: [1 2 <nil> 3]
输出: [[1] [2] [3]]
期望: [[1] [2] [3]]
✅ 通过

测试用例 5: 右斜树
输入: [1 <nil> 2 <nil> <nil> <nil> 3]
输出: [[1] [2]]
期望: [[1] [2]]
✅ 通过

测试用例 6: 完全二叉树
输入: [1 2 3 4 5 6 7]
输出: [[1] [2 3] [4 5 6 7]]
期望: [[1] [2 3] [4 5 6 7]]
✅ 通过

测试用例 7: 只有左子树
输入: [1 2 <nil> 3 <nil> 4]
输出: [[1] [2] [3] [4]]
期望: [[1] [2] [3] [4]]
✅ 通过

测试用例 8: 只有右子树
输入: [1 <nil> 2 <nil> <nil> <nil> 3 <nil> <nil> <nil> <nil> <nil> <nil> <nil> 4]
输出: [[1] [2] [3] [4]]
期望: [[1] [2] [3] [4]]
✅ 通过

=== 测试总结 ===
🎉 所有测试用例都通过了！
*/
