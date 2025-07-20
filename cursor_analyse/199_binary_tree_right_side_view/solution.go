package main

import "fmt"

/*
题目：二叉树的右视图
难度：中等
标签：树、深度优先搜索、广度优先搜索、二叉树

题目描述：
给定一个二叉树的根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

要求：
- 时间复杂度：O(n)，其中 n 是二叉树中节点的个数
- 空间复杂度：O(h)，其中 h 是树的高度

示例：
输入：root = [1,2,3,null,5,null,4]
输出：[1,3,4]

输入：root = [1,null,3]
输出：[1,3]

输入：root = []
输出：[]

解释：
- 右视图是从右侧观察二叉树时能看到的节点
- 每一层只能看到最右边的节点
- 如果某一层没有右子节点，则可以看到左子节点
*/

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
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

// 解法二：DFS（深度优先搜索）
// 核心思想：使用DFS，优先访问右子树，记录每一层第一个访问的节点
// 时间复杂度：O(n)
// 空间复杂度：O(h)，h为树的高度
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	dfs(root, 0, &result)
	return result
}

func dfs(node *TreeNode, depth int, result *[]int) {
	if node == nil {
		return
	}

	// 如果当前深度还没有记录节点，说明这是该层第一个访问的节点（最右边的）
	if depth == len(*result) {
		*result = append(*result, node.Val)
	}

	// 优先访问右子树，确保右视图
	dfs(node.Right, depth+1, result)
	dfs(node.Left, depth+1, result)
}

// 解法三：DFS（先序遍历）
// 核心思想：使用先序遍历，记录每一层最右边的节点
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func rightSideView3(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	preorder(root, 0, &result)
	return result
}

func preorder(node *TreeNode, depth int, result *[]int) {
	if node == nil {
		return
	}

	// 如果当前深度还没有记录节点，记录当前节点
	if depth == len(*result) {
		*result = append(*result, node.Val)
	}

	// 先序遍历：根 -> 左 -> 右
	preorder(node.Left, depth+1, result)
	preorder(node.Right, depth+1, result)
}

// 解法四：BFS（使用两个队列）
// 核心思想：使用两个队列交替存储不同层的节点
// 时间复杂度：O(n)
// 空间复杂度：O(w)
func rightSideView4(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	currentLevel := []*TreeNode{root}

	for len(currentLevel) > 0 {
		nextLevel := []*TreeNode{}

		// 处理当前层的所有节点
		for i, node := range currentLevel {
			// 如果是当前层的最后一个节点，加入结果
			if i == len(currentLevel)-1 {
				result = append(result, node.Val)
			}

			// 添加子节点到下一层
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		currentLevel = nextLevel
	}

	return result
}

// 解法五：DFS（后序遍历）
// 核心思想：使用后序遍历，记录每一层最右边的节点
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func rightSideView5(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	postorder(root, 0, &result)
	return result
}

func postorder(node *TreeNode, depth int, result *[]int) {
	if node == nil {
		return
	}

	// 后序遍历：左 -> 右 -> 根
	postorder(node.Left, depth+1, result)
	postorder(node.Right, depth+1, result)

	// 如果当前深度还没有记录节点，记录当前节点
	if depth == len(*result) {
		*result = append(*result, node.Val)
	}
}

// 解法六：BFS（使用map记录层数）
// 核心思想：使用BFS，用map记录每一层最右边的节点
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func rightSideView6(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	// 使用map记录每一层最右边的节点
	rightmost := make(map[int]int)
	queue := []struct {
		node  *TreeNode
		depth int
	}{{root, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// 更新当前层最右边的节点
		rightmost[current.depth] = current.node.Val

		// 添加子节点到队列
		if current.node.Left != nil {
			queue = append(queue, struct {
				node  *TreeNode
				depth int
			}{current.node.Left, current.depth + 1})
		}
		if current.node.Right != nil {
			queue = append(queue, struct {
				node  *TreeNode
				depth int
			}{current.node.Right, current.depth + 1})
		}
	}

	// 将map转换为有序数组
	var result []int
	for i := 0; i < len(rightmost); i++ {
		result = append(result, rightmost[i])
	}

	return result
}

// 解法七：DFS（中序遍历）
// 核心思想：使用中序遍历，记录每一层最右边的节点
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func rightSideView7(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	inorder(root, 0, &result)
	return result
}

func inorder(node *TreeNode, depth int, result *[]int) {
	if node == nil {
		return
	}

	// 中序遍历：左 -> 根 -> 右
	inorder(node.Left, depth+1, result)

	// 如果当前深度还没有记录节点，记录当前节点
	if depth == len(*result) {
		*result = append(*result, node.Val)
	}

	inorder(node.Right, depth+1, result)
}

// 创建测试用的二叉树
func createTestTree1() *TreeNode {
	// 创建测试树：[1,2,3,null,5,null,4]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}

	return root
}

func createTestTree2() *TreeNode {
	// 创建测试树：[1,null,3]
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	return root
}

func createTestTree3() *TreeNode {
	// 创建测试树：[1,2,3,4]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}

	return root
}

func createTestTree4() *TreeNode {
	// 创建测试树：[1,2,3,null,5,null,4,null,6]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}
	root.Left.Right.Right = &TreeNode{Val: 6}

	return root
}

func createTestTree5() *TreeNode {
	// 创建测试树：[1,2,3,4,5,6,7]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	return root
}

func main() {
	// 测试用例
	testCases := []struct {
		root *TreeNode
		desc string
	}{
		{createTestTree1(), "测试树1: [1,2,3,null,5,null,4]"},
		{createTestTree2(), "测试树2: [1,null,3]"},
		{createTestTree3(), "测试树3: [1,2,3,4]"},
		{createTestTree4(), "测试树4: [1,2,3,null,5,null,4,null,6]"},
		{createTestTree5(), "测试树5: [1,2,3,4,5,6,7]"},
		{nil, "空树"},
	}

	fmt.Println("=== 解法一：BFS（层次遍历）===")
	for i, tc := range testCases {
		result := rightSideView(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("右视图: %v\n\n", result)
	}

	fmt.Println("=== 解法二：DFS（深度优先搜索）===")
	for i, tc := range testCases {
		result := rightSideView2(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("右视图: %v\n\n", result)
	}

	fmt.Println("=== 解法三：DFS（先序遍历）===")
	for i, tc := range testCases {
		result := rightSideView3(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("右视图: %v\n\n", result)
	}

	fmt.Println("=== 解法四：BFS（使用两个队列）===")
	for i, tc := range testCases {
		result := rightSideView4(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("右视图: %v\n\n", result)
	}

	fmt.Println("=== 解法五：DFS（后序遍历）===")
	for i, tc := range testCases {
		result := rightSideView5(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("右视图: %v\n\n", result)
	}

	fmt.Println("=== 解法六：BFS（使用map记录层数）===")
	for i, tc := range testCases {
		result := rightSideView6(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("右视图: %v\n\n", result)
	}

	fmt.Println("=== 解法七：DFS（中序遍历）===")
	for i, tc := range testCases {
		result := rightSideView7(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("右视图: %v\n\n", result)
	}
}

/*
解题思路：

1. BFS（层次遍历）：
   - 核心思想：使用层次遍历，记录每一层的最后一个节点
   - 时间复杂度：O(n)，空间复杂度：O(w)，w为树的最大宽度
   - 优点：逻辑清晰，易于理解
   - 适用：面试中推荐使用

2. DFS（深度优先搜索）：
   - 核心思想：使用DFS，优先访问右子树，记录每一层第一个访问的节点
   - 时间复杂度：O(n)，空间复杂度：O(h)，h为树的高度
   - 优点：空间效率高，代码简洁
   - 适用：树的高度较大时使用

3. DFS（先序遍历）：
   - 核心思想：使用先序遍历，记录每一层最右边的节点
   - 时间复杂度：O(n)，空间复杂度：O(h)
   - 优点：实现简单
   - 缺点：需要额外的逻辑来确保记录最右边的节点

4. BFS（使用两个队列）：
   - 核心思想：使用两个队列交替存储不同层的节点
   - 时间复杂度：O(n)，空间复杂度：O(w)
   - 优点：层次清晰，易于理解
   - 缺点：空间复杂度较高

5. DFS（后序遍历）：
   - 核心思想：使用后序遍历，记录每一层最右边的节点
   - 时间复杂度：O(n)，空间复杂度：O(h)
   - 优点：实现简单
   - 缺点：需要额外的逻辑来确保记录最右边的节点

6. BFS（使用map记录层数）：
   - 核心思想：使用BFS，用map记录每一层最右边的节点
   - 时间复杂度：O(n)，空间复杂度：O(n)
   - 优点：思路清晰
   - 缺点：空间复杂度较高

7. DFS（中序遍历）：
   - 核心思想：使用中序遍历，记录每一层最右边的节点
   - 时间复杂度：O(n)，空间复杂度：O(h)
   - 优点：实现简单
   - 缺点：需要额外的逻辑来确保记录最右边的节点

关键点：
1. 理解右视图的定义：从右侧观察二叉树时能看到的节点
2. 掌握BFS和DFS的层次遍历方法
3. 注意边界情况：空树、单节点树
4. 理解不同解法的优缺点和适用场景

时间复杂度分析：
- 所有解法：O(n)，每个节点最多访问一次

空间复杂度分析：
- BFS解法：O(w)，w为树的最大宽度
- DFS解法：O(h)，h为树的高度

面试要点：
1. 能够解释BFS和DFS的原理
2. 理解为什么BFS能够正确获取右视图
3. 掌握不同遍历方式的特点
4. 考虑边界情况的处理
5. 能够分析不同解法的优缺点

优化技巧：
1. BFS是最直观的解法，推荐在面试中使用
2. DFS在空间复杂度上有优势，适合树的高度较大的情况
3. 注意遍历顺序对结果的影响
4. 理解层次遍历和深度优先搜索的区别
*/
