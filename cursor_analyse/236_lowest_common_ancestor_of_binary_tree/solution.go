package main

import "fmt"

/*
题目：二叉树的最近公共祖先
难度：中等
标签：树、深度优先搜索、二叉树

题目描述：
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先（LCA）。
最近公共祖先的定义为："对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。"

要求：
- 时间复杂度：O(n)，其中 n 是二叉树中节点的个数
- 空间复杂度：O(h)，其中 h 是树的高度

示例：
输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出：3
解释：节点 5 和节点 1 的最近公共祖先是节点 3。

输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出：5
解释：节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
*/

// 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 解法一：递归（推荐）
// 时间复杂度：O(n)
// 空间复杂度：O(h)，h为树的高度
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果root为空，或者root等于p或q，直接返回root
	if root == nil || root == p || root == q {
		return root
	}

	// 递归查找左子树
	left := lowestCommonAncestor(root.Left, p, q)
	// 递归查找右子树
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果左子树和右子树都找到了结果，说明root就是LCA
	if left != nil && right != nil {
		return root
	}

	// 如果只有左子树找到了，返回左子树的结果
	if left != nil {
		return left
	}

	// 如果只有右子树找到了，返回右子树的结果
	return right
}

// 解法二：递归（带路径记录）
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 找到从root到p和q的路径
	pathP := findPath(root, p)
	pathQ := findPath(root, q)

	// 找到两条路径的最后一个公共节点
	i := 0
	for i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i] {
		i++
	}

	return pathP[i-1]
}

// 找到从root到target的路径
func findPath(root, target *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	if root == target {
		return []*TreeNode{root}
	}

	// 查找左子树
	leftPath := findPath(root.Left, target)
	if leftPath != nil {
		return append([]*TreeNode{root}, leftPath...)
	}

	// 查找右子树
	rightPath := findPath(root.Right, target)
	if rightPath != nil {
		return append([]*TreeNode{root}, rightPath...)
	}

	return nil
}

// 解法三：迭代（使用栈）
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	// 使用后序遍历的迭代方式
	stack := []*TreeNode{}
	var lastVisited *TreeNode
	current := root

	// 记录p和q的父节点映射
	parent := make(map[*TreeNode]*TreeNode)

	// 找到p和q
	foundP, foundQ := false, false

	for current != nil || len(stack) > 0 {
		// 遍历到最左节点
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]

		// 如果右子树已经访问过，或者右子树为空
		if current.Right == nil || current.Right == lastVisited {
			// 记录父节点关系
			if current.Left != nil {
				parent[current.Left] = current
			}
			if current.Right != nil {
				parent[current.Right] = current
			}

			// 检查是否找到了p和q
			if current == p {
				foundP = true
			}
			if current == q {
				foundQ = true
			}

			// 如果都找到了，跳出循环
			if foundP && foundQ {
				break
			}

			lastVisited = current
			stack = stack[:len(stack)-1]
			current = nil
		} else {
			current = current.Right
		}
	}

	// 构建从p到root的路径
	ancestors := make(map[*TreeNode]bool)
	for p != nil {
		ancestors[p] = true
		p = parent[p]
	}

	// 从q开始向上查找，直到找到在ancestors中的节点
	for !ancestors[q] {
		q = parent[q]
	}

	return q
}

// 解法四：递归（统计子树中的目标节点）
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func lowestCommonAncestor4(root, p, q *TreeNode) *TreeNode {
	var result *TreeNode

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 统计当前节点及其子树中目标节点的数量
		count := 0
		if node == p || node == q {
			count++
		}

		// 递归统计左子树
		count += dfs(node.Left)
		// 递归统计右子树
		count += dfs(node.Right)

		// 如果当前节点及其子树包含两个目标节点，且还没有找到LCA
		if count == 2 && result == nil {
			result = node
		}

		return count
	}

	dfs(root)
	return result
}

// 解法五：递归（返回找到的节点数量）
// 时间复杂度：O(n)
// 空间复杂度：O(h)
func lowestCommonAncestor5(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode

	var findNodes func(node *TreeNode) int
	findNodes = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 检查当前节点
		found := 0
		if node == p || node == q {
			found = 1
		}

		// 递归查找左子树
		leftFound := findNodes(node.Left)
		// 递归查找右子树
		rightFound := findNodes(node.Right)

		totalFound := found + leftFound + rightFound

		// 如果找到两个节点且还没有确定LCA
		if totalFound == 2 && lca == nil {
			lca = node
		}

		return totalFound
	}

	findNodes(root)
	return lca
}

// 解法六：BFS（层次遍历，找到所有节点的父节点）
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func lowestCommonAncestor6(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 使用BFS建立父节点映射
	parent := make(map[*TreeNode]*TreeNode)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			parent[node.Left] = node
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			queue = append(queue, node.Right)
		}
	}

	// 构建从p到root的路径
	ancestors := make(map[*TreeNode]bool)
	for p != nil {
		ancestors[p] = true
		p = parent[p]
	}

	// 从q开始向上查找，直到找到在ancestors中的节点
	for !ancestors[q] {
		q = parent[q]
	}

	return q
}

// 创建测试用的二叉树
func createTestTree() *TreeNode {
	// 创建测试树：[3,5,1,6,2,0,8,null,null,7,4]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	return root
}

func main() {
	// 创建测试树
	root := createTestTree()

	// 测试用例
	testCases := []struct {
		p, q *TreeNode
		desc string
	}{
		{root.Left, root.Right, "p=5, q=1, 期望LCA=3"},
		{root.Left, root.Left.Right.Right, "p=5, q=4, 期望LCA=5"},
		{root.Left.Left, root.Left.Right.Right, "p=6, q=4, 期望LCA=5"},
		{root.Right.Left, root.Right.Right, "p=0, q=8, 期望LCA=1"},
		{root.Left.Right.Left, root.Left.Right.Right, "p=7, q=4, 期望LCA=2"},
	}

	fmt.Println("=== 解法一：递归（推荐）===")
	for i, tc := range testCases {
		result := lowestCommonAncestor(root, tc.p, tc.q)
		fmt.Printf("测试用例 %d: %s, 结果=%d\n", i+1, tc.desc, result.Val)
	}

	fmt.Println("\n=== 解法二：递归（带路径记录）===")
	for i, tc := range testCases {
		result := lowestCommonAncestor2(root, tc.p, tc.q)
		fmt.Printf("测试用例 %d: %s, 结果=%d\n", i+1, tc.desc, result.Val)
	}

	fmt.Println("\n=== 解法三：迭代（使用栈）===")
	for i, tc := range testCases {
		result := lowestCommonAncestor3(root, tc.p, tc.q)
		fmt.Printf("测试用例 %d: %s, 结果=%d\n", i+1, tc.desc, result.Val)
	}

	fmt.Println("\n=== 解法四：递归（统计子树中的目标节点）===")
	for i, tc := range testCases {
		result := lowestCommonAncestor4(root, tc.p, tc.q)
		fmt.Printf("测试用例 %d: %s, 结果=%d\n", i+1, tc.desc, result.Val)
	}

	fmt.Println("\n=== 解法五：递归（返回找到的节点数量）===")
	for i, tc := range testCases {
		result := lowestCommonAncestor5(root, tc.p, tc.q)
		fmt.Printf("测试用例 %d: %s, 结果=%d\n", i+1, tc.desc, result.Val)
	}

	fmt.Println("\n=== 解法六：BFS（层次遍历）===")
	for i, tc := range testCases {
		result := lowestCommonAncestor6(root, tc.p, tc.q)
		fmt.Printf("测试用例 %d: %s, 结果=%d\n", i+1, tc.desc, result.Val)
	}
}

/*
解题思路：

1. 递归（推荐）：
   - 如果root为空或等于p或q，直接返回root
   - 递归查找左子树和右子树
   - 如果左右子树都找到了结果，说明root就是LCA
   - 如果只有一边找到了，返回那一边的结果

2. 递归（带路径记录）：
   - 找到从root到p和q的路径
   - 找到两条路径的最后一个公共节点
   - 空间复杂度较高，但思路清晰

3. 迭代（使用栈）：
   - 使用后序遍历的迭代方式
   - 建立父节点映射关系
   - 找到p和q后，构建路径并查找LCA

4. 递归（统计子树中的目标节点）：
   - 统计每个节点及其子树中目标节点的数量
   - 当某个节点及其子树包含两个目标节点时，该节点就是LCA

5. 递归（返回找到的节点数量）：
   - 类似解法四，但使用全局变量记录LCA
   - 代码更简洁

6. BFS（层次遍历）：
   - 使用BFS建立所有节点的父节点映射
   - 构建从p到root的路径
   - 从q开始向上查找LCA

时间复杂度分析：
- 所有解法：O(n)，每个节点最多访问一次

空间复杂度分析：
- 递归解法：O(h)，h为树的高度
- 迭代解法：O(n)，需要存储父节点映射

关键点：
1. 递归是最直观的解法
2. 理解LCA的定义：最近的公共祖先
3. 注意边界条件：节点为空、节点等于p或q
4. 理解不同解法的适用场景
5. 掌握树的遍历方式：DFS和BFS
*/
