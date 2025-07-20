package main

import "fmt"

/*
题目：二叉树的锯齿形层次遍历
难度：中等
标签：树、广度优先搜索、二叉树

题目描述：
给你二叉树的根节点 root，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

要求：
- 时间复杂度：O(n)，其中 n 是二叉树中节点的个数
- 空间复杂度：O(n)

示例：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]
解释：
第0层：从左到右 [3]
第1层：从右到左 [20,9]
第2层：从左到右 [15,7]

输入：root = [1]
输出：[[1]]

输入：root = []
输出：[]
*/

// 二叉树节点定义
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

// 解法二：BFS + 反转数组
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func zigzagLevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		levelNodes := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levelNodes = append(levelNodes, node.Val)

			// 添加子节点到队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 奇数层反转数组
		if level%2 == 1 {
			for i, j := 0, len(levelNodes)-1; i < j; i, j = i+1, j-1 {
				levelNodes[i], levelNodes[j] = levelNodes[j], levelNodes[i]
			}
		}

		result = append(result, levelNodes)
		level++
	}

	return result
}

// 解法三：BFS + 双端队列（使用切片模拟）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func zigzagLevelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		levelNodes := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 根据层级决定插入方式
			if level%2 == 0 {
				// 偶数层：从左到右，正常添加
				levelNodes = append(levelNodes, node.Val)
			} else {
				// 奇数层：从右到左，在开头插入
				levelNodes = append([]int{node.Val}, levelNodes...)
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

// 解法四：DFS + 递归
// 时间复杂度：O(n)
// 空间复杂度：O(h)，h为树的高度
func zigzagLevelOrder4(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	dfs(root, 0, &result)
	return result
}

func dfs(node *TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	// 确保result有足够的层级
	if level >= len(*result) {
		*result = append(*result, []int{})
	}

	// 根据层级决定插入位置
	if level%2 == 0 {
		// 偶数层：从左到右，正常添加
		(*result)[level] = append((*result)[level], node.Val)
	} else {
		// 奇数层：从右到左，在开头插入
		(*result)[level] = append([]int{node.Val}, (*result)[level]...)
	}

	// 递归处理子节点
	dfs(node.Left, level+1, result)
	dfs(node.Right, level+1, result)
}

// 解法五：BFS + 两个栈
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func zigzagLevelOrder5(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	stack1 := []*TreeNode{root} // 当前层
	stack2 := []*TreeNode{}     // 下一层
	level := 0

	for len(stack1) > 0 {
		levelNodes := make([]int, 0)

		// 处理当前层的所有节点
		for len(stack1) > 0 {
			node := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]

			levelNodes = append(levelNodes, node.Val)

			// 根据层级决定子节点的添加顺序
			if level%2 == 0 {
				// 偶数层：先左后右
				if node.Left != nil {
					stack2 = append(stack2, node.Left)
				}
				if node.Right != nil {
					stack2 = append(stack2, node.Right)
				}
			} else {
				// 奇数层：先右后左
				if node.Right != nil {
					stack2 = append(stack2, node.Right)
				}
				if node.Left != nil {
					stack2 = append(stack2, node.Left)
				}
			}
		}

		result = append(result, levelNodes)

		// 交换栈
		stack1, stack2 = stack2, stack1
		level++
	}

	return result
}

// 解法六：BFS + 链表（使用切片模拟）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func zigzagLevelOrder6(root *TreeNode) [][]int {
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

			// 根据层级决定索引
			index := i
			if level%2 == 1 {
				index = levelSize - 1 - i
			}
			levelNodes[index] = node.Val

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

// 创建测试用的二叉树
func createTestTree1() *TreeNode {
	// 创建测试树：[3,9,20,null,null,15,7]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	return root
}

func createTestTree2() *TreeNode {
	// 创建测试树：[1,2,3,4,5]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	return root
}

func createTestTree3() *TreeNode {
	// 创建测试树：[1,2,3,4,null,null,5]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	return root
}

func main() {
	// 测试用例
	testCases := []struct {
		root *TreeNode
		desc string
	}{
		{createTestTree1(), "测试树1: [3,9,20,null,null,15,7]"},
		{createTestTree2(), "测试树2: [1,2,3,4,5]"},
		{createTestTree3(), "测试树3: [1,2,3,4,null,null,5]"},
		{nil, "空树"},
	}

	fmt.Println("=== 解法一：BFS + 双端队列（推荐）===")
	for i, tc := range testCases {
		result := zigzagLevelOrder(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("结果: %v\n\n", result)
	}

	fmt.Println("=== 解法二：BFS + 反转数组 ===")
	for i, tc := range testCases {
		result := zigzagLevelOrder2(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("结果: %v\n\n", result)
	}

	fmt.Println("=== 解法三：BFS + 双端队列（使用切片模拟） ===")
	for i, tc := range testCases {
		result := zigzagLevelOrder3(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("结果: %v\n\n", result)
	}

	fmt.Println("=== 解法四：DFS + 递归 ===")
	for i, tc := range testCases {
		result := zigzagLevelOrder4(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("结果: %v\n\n", result)
	}

	fmt.Println("=== 解法五：BFS + 两个栈 ===")
	for i, tc := range testCases {
		result := zigzagLevelOrder5(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("结果: %v\n\n", result)
	}

	fmt.Println("=== 解法六：BFS + 链表（使用切片模拟） ===")
	for i, tc := range testCases {
		result := zigzagLevelOrder6(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("结果: %v\n\n", result)
	}
}

/*
解题思路：

1. BFS + 双端队列（推荐）：
   - 使用队列进行层次遍历
   - 根据层级决定节点的插入位置
   - 偶数层从左到右，奇数层从右到左
   - 时间复杂度：O(n)，空间复杂度：O(n)

2. BFS + 反转数组：
   - 正常进行层次遍历
   - 对奇数层的数组进行反转
   - 思路简单，但需要额外的反转操作

3. BFS + 双端队列（使用切片模拟）：
   - 偶数层正常添加节点
   - 奇数层在数组开头插入节点
   - 避免了反转操作

4. DFS + 递归：
   - 使用深度优先搜索
   - 记录当前层级
   - 根据层级决定插入位置
   - 空间复杂度：O(h)，h为树的高度

5. BFS + 两个栈：
   - 使用两个栈交替处理
   - 根据层级决定子节点的添加顺序
   - 避免了数组操作

6. BFS + 链表（使用切片模拟）：
   - 预先计算每个节点的最终位置
   - 直接放入正确的位置
   - 避免了数组操作

时间复杂度分析：
- 所有解法：O(n)，每个节点最多访问一次

空间复杂度分析：
- BFS解法：O(n)，需要存储所有节点
- DFS解法：O(h)，h为树的高度

关键点：
1. 理解锯齿形遍历的定义：相邻层交替方向
2. 掌握BFS和DFS的层次遍历方法
3. 注意边界条件：空树、单节点树
4. 理解不同解法的优缺点
5. 掌握数组操作：反转、插入、索引计算

优化技巧：
1. 使用双端队列可以避免反转操作
2. 预先计算索引可以提高效率
3. 使用两个栈可以简化逻辑
4. DFS在空间复杂度上有优势
*/
