package main

import "fmt"

/*
题目：二叉树的中序遍历
难度：简单
标签：栈、树、深度优先搜索、二叉树

题目描述：
给定一个二叉树的根节点 root，返回它的中序遍历。

要求：
- 时间复杂度：O(n)，其中 n 是二叉树中节点的个数
- 空间复杂度：O(h)，其中 h 是树的高度

示例：
输入：root = [1,null,2,3]
输出：[1,3,2]

输入：root = []
输出：[]

输入：root = [1]
输出：[1]

解释：
- 中序遍历的顺序：左子树 -> 根节点 -> 右子树
- 对于每个节点，先遍历其左子树，再访问根节点，最后遍历右子树
*/

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

// 解法二：迭代（使用栈）
// 核心思想：使用栈模拟递归过程，手动维护调用栈
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func inorderTraversal2(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		// 遍历到最左节点，将所有左子节点压入栈
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// 弹出栈顶节点并访问
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// 转向右子树
		current = current.Right
	}

	return result
}

// 解法三：迭代（使用栈，优化版本）
// 核心思想：与解法二相同，但代码更简洁
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func inorderTraversal3(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		// 遍历到最左节点
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// 弹出并访问
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// 处理右子树
		current = current.Right
	}

	return result
}

// 解法四：Morris遍历（线索二叉树）
// 核心思想：利用叶子节点的空指针，在遍历过程中建立线索，避免使用栈
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func inorderTraversal4(root *TreeNode) []int {
	var result []int
	current := root

	for current != nil {
		if current.Left == nil {
			// 没有左子树，访问当前节点并转向右子树
			result = append(result, current.Val)
			current = current.Right
		} else {
			// 找到当前节点左子树的最右节点（前驱节点）
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// 第一次访问，建立线索
				predecessor.Right = current
				current = current.Left
			} else {
				// 第二次访问，删除线索并访问当前节点
				predecessor.Right = nil
				result = append(result, current.Val)
				current = current.Right
			}
		}
	}

	return result
}

// 解法五：迭代（使用颜色标记）
// 核心思想：使用颜色标记节点状态，白色表示未访问，灰色表示已访问
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func inorderTraversal5(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	stack := []struct {
		node  *TreeNode
		color int // 0: 白色(未访问), 1: 灰色(已访问)
	}{{root, 0}}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.node == nil {
			continue
		}

		if top.color == 0 {
			// 白色节点，按照中序遍历的逆序压入栈：右 -> 根 -> 左
			stack = append(stack, struct {
				node  *TreeNode
				color int
			}{top.node.Right, 0})
			stack = append(stack, struct {
				node  *TreeNode
				color int
			}{top.node, 1})
			stack = append(stack, struct {
				node  *TreeNode
				color int
			}{top.node.Left, 0})
		} else {
			// 灰色节点，访问
			result = append(result, top.node.Val)
		}
	}

	return result
}

// 解法六：迭代（使用指针标记）
// 核心思想：使用特殊指针标记已访问的节点
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func inorderTraversal6(root *TreeNode) []int {
	var result []int
	stack := []*TreeNode{}
	current := root
	var lastVisited *TreeNode

	for current != nil || len(stack) > 0 {
		// 遍历到最左节点
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]

		// 如果右子树为空或已经访问过，则访问当前节点
		if current.Right == nil || current.Right == lastVisited {
			result = append(result, current.Val)
			lastVisited = current
			stack = stack[:len(stack)-1]
			current = nil
		} else {
			// 转向右子树
			current = current.Right
		}
	}

	return result
}

// 解法七：递归（使用闭包）
// 核心思想：使用闭包简化递归实现
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func inorderTraversal7(root *TreeNode) []int {
	var result []int

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	return result
}

// 解法八：迭代（使用两个栈）
// 核心思想：使用两个栈分别存储节点和访问状态
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func inorderTraversal8(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	nodeStack := []*TreeNode{root}
	visitStack := []bool{false} // false表示未访问，true表示已访问

	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack)-1]
		visited := visitStack[len(visitStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		visitStack = visitStack[:len(visitStack)-1]

		if visited {
			// 已访问，加入结果
			result = append(result, node.Val)
		} else {
			// 未访问，按照中序遍历的逆序压入栈：右 -> 根 -> 左
			if node.Right != nil {
				nodeStack = append(nodeStack, node.Right)
				visitStack = append(visitStack, false)
			}
			nodeStack = append(nodeStack, node)
			visitStack = append(visitStack, true)
			if node.Left != nil {
				nodeStack = append(nodeStack, node.Left)
				visitStack = append(visitStack, false)
			}
		}
	}

	return result
}

// 创建测试用的二叉树
func createTestTree1() *TreeNode {
	// 创建测试树：[1,null,2,3]
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

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

func createTestTree4() *TreeNode {
	// 创建测试树：[1,2,3,null,4,5,6]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	return root
}

func main() {
	// 测试用例
	testCases := []struct {
		root *TreeNode
		desc string
	}{
		{createTestTree1(), "测试树1: [1,null,2,3]"},
		{createTestTree2(), "测试树2: [1,2,3,4,5]"},
		{createTestTree3(), "测试树3: [1,2,3,4,5,6,7]"},
		{createTestTree4(), "测试树4: [1,2,3,null,4,5,6]"},
		{nil, "空树"},
	}

	fmt.Println("=== 解法一：递归（推荐）===")
	for i, tc := range testCases {
		result := inorderTraversal(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("中序遍历: %v\n\n", result)
	}

	fmt.Println("=== 解法二：迭代（使用栈）===")
	for i, tc := range testCases {
		result := inorderTraversal2(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("中序遍历: %v\n\n", result)
	}

	fmt.Println("=== 解法三：迭代（使用栈，优化版本）===")
	for i, tc := range testCases {
		result := inorderTraversal3(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("中序遍历: %v\n\n", result)
	}

	fmt.Println("=== 解法四：Morris遍历（线索二叉树）===")
	for i, tc := range testCases {
		result := inorderTraversal4(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("中序遍历: %v\n\n", result)
	}

	fmt.Println("=== 解法五：迭代（使用颜色标记）===")
	for i, tc := range testCases {
		result := inorderTraversal5(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("中序遍历: %v\n\n", result)
	}

	fmt.Println("=== 解法六：迭代（使用指针标记）===")
	for i, tc := range testCases {
		result := inorderTraversal6(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("中序遍历: %v\n\n", result)
	}

	fmt.Println("=== 解法七：递归（使用闭包）===")
	for i, tc := range testCases {
		result := inorderTraversal7(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("中序遍历: %v\n\n", result)
	}

	fmt.Println("=== 解法八：迭代（使用两个栈）===")
	for i, tc := range testCases {
		result := inorderTraversal8(tc.root)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("中序遍历: %v\n\n", result)
	}
}

/*
解题思路：

1. 递归（推荐）：
   - 核心思想：递归实现中序遍历，先遍历左子树，再访问根节点，最后遍历右子树
   - 时间复杂度：O(n)，空间复杂度：O(h)
   - 优点：代码简洁，易于理解
   - 适用：面试中推荐使用

2. 迭代（使用栈）：
   - 核心思想：使用栈模拟递归过程，手动维护调用栈
   - 时间复杂度：O(n)，空间复杂度：O(h)
   - 优点：避免递归调用栈的开销
   - 适用：需要控制栈空间时使用

3. Morris遍历（线索二叉树）：
   - 核心思想：利用叶子节点的空指针，在遍历过程中建立线索，避免使用栈
   - 时间复杂度：O(n)，空间复杂度：O(1)
   - 优点：空间复杂度最优
   - 缺点：代码复杂，难以理解

4. 迭代（使用颜色标记）：
   - 核心思想：使用颜色标记节点状态，白色表示未访问，灰色表示已访问
   - 时间复杂度：O(n)，空间复杂度：O(h)
   - 优点：思路清晰，易于理解
   - 适用：学习阶段使用

5. 迭代（使用指针标记）：
   - 核心思想：使用特殊指针标记已访问的节点
   - 时间复杂度：O(n)，空间复杂度：O(h)
   - 优点：实现简单
   - 适用：需要标记访问状态时使用

6. 递归（使用闭包）：
   - 核心思想：使用闭包简化递归实现
   - 时间复杂度：O(n)，空间复杂度：O(h)
   - 优点：代码最简洁
   - 适用：函数式编程风格

7. 迭代（使用两个栈）：
   - 核心思想：使用两个栈分别存储节点和访问状态
   - 时间复杂度：O(n)，空间复杂度：O(h)
   - 优点：状态管理清晰
   - 缺点：空间复杂度较高

关键点：
1. 理解中序遍历的定义：左子树 -> 根节点 -> 右子树
2. 掌握递归和迭代的实现方式
3. 理解栈在迭代中的作用
4. 掌握Morris遍历的原理（高级技巧）

时间复杂度分析：
- 所有解法：O(n)，每个节点最多访问一次

空间复杂度分析：
- 递归解法：O(h)，h为树的高度（递归调用栈的深度）
- 迭代解法：O(h)，h为树的高度（栈的最大深度）
- Morris遍历：O(1)，只使用常数额外空间

面试要点：
1. 能够解释中序遍历的定义和顺序
2. 掌握递归和迭代的实现
3. 理解栈在迭代中的作用
4. 能够分析不同解法的优缺点
5. 了解Morris遍历（加分项）

优化技巧：
1. 递归是最直观的解法，推荐在面试中使用
2. 迭代解法避免递归调用栈的开销
3. Morris遍历在空间复杂度上有优势
4. 理解不同解法的适用场景
*/
