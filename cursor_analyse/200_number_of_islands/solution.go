package main

import "fmt"

/*
题目：岛屿数量
难度：中等
标签：深度优先搜索、广度优先搜索、并查集、数组、矩阵

题目描述：
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

要求：
- 时间复杂度：O(m*n)，其中 m 和 n 分别是网格的行数和列数
- 空间复杂度：O(m*n)，最坏情况下整个网格都是陆地

示例：
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3
*/

// 解法一：深度优先搜索（DFS）
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)，最坏情况下递归栈深度为m*n
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	// DFS函数：将当前岛屿的所有陆地标记为已访问
	var dfs func(row, col int)
	dfs = func(row, col int) {
		// 边界检查
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return
		}
		// 如果是水或已访问的陆地，直接返回
		if grid[row][col] == '0' {
			return
		}

		// 标记当前陆地为已访问（改为'0'）
		grid[row][col] = '0'

		// 递归访问上下左右四个方向
		dfs(row-1, col) // 上
		dfs(row+1, col) // 下
		dfs(row, col-1) // 左
		dfs(row, col+1) // 右
	}

	// 遍历整个网格
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// 如果发现未访问的陆地，说明发现了一个新岛屿
			if grid[row][col] == '1' {
				count++
				dfs(row, col) // 标记整个岛屿
			}
		}
	}

	return count
}

// 解法二：广度优先搜索（BFS）
// 时间复杂度：O(m*n)
// 空间复杂度：O(min(m,n))，队列的最大长度
func numIslands2(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	// 方向数组：上、下、左、右
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// BFS函数：使用队列遍历整个岛屿
	var bfs func(row, col int)
	bfs = func(row, col int) {
		queue := [][]int{{row, col}}
		grid[row][col] = '0' // 标记起始点

		for len(queue) > 0 {
			// 取出队首元素
			current := queue[0]
			queue = queue[1:]
			currRow, currCol := current[0], current[1]

			// 检查四个方向
			for _, dir := range directions {
				newRow := currRow + dir[0]
				newCol := currCol + dir[1]

				// 边界检查
				if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
					continue
				}

				// 如果是未访问的陆地，加入队列并标记
				if grid[newRow][newCol] == '1' {
					queue = append(queue, []int{newRow, newCol})
					grid[newRow][newCol] = '0'
				}
			}
		}
	}

	// 遍历整个网格
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				count++
				bfs(row, col)
			}
		}
	}

	return count
}

// 解法三：并查集（Union-Find）
// 时间复杂度：O(m*n * α(m*n))，其中α是阿克曼函数的反函数
// 空间复杂度：O(m*n)
func numIslands3(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])

	// 初始化并查集
	uf := NewUnionFind(rows * cols)

	// 方向数组
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 统计陆地数量
	landCount := 0

	// 遍历网格
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				landCount++
				current := row*cols + col

				// 检查四个方向
				for _, dir := range directions {
					newRow := row + dir[0]
					newCol := col + dir[1]

					// 边界检查
					if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
						continue
					}

					// 如果相邻位置也是陆地，合并两个集合
					if grid[newRow][newCol] == '1' {
						neighbor := newRow*cols + newCol
						uf.Union(current, neighbor)
					}
				}
			}
		}
	}

	// 岛屿数量 = 陆地数量 - 合并次数
	return landCount - uf.GetUnionCount()
}

// 并查集结构
type UnionFind struct {
	parent     []int
	rank       []int
	unionCount int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)

	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 1
	}

	return &UnionFind{
		parent:     parent,
		rank:       rank,
		unionCount: 0,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		// 按秩合并
		if uf.rank[rootX] < uf.rank[rootY] {
			rootX, rootY = rootY, rootX
		}
		uf.parent[rootY] = rootX
		if uf.rank[rootX] == uf.rank[rootY] {
			uf.rank[rootX]++
		}
		uf.unionCount++
	}
}

func (uf *UnionFind) GetUnionCount() int {
	return uf.unionCount
}

func main() {
	// 测试用例
	testCases := []struct {
		grid   [][]byte
		expect int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
		{
			[][]byte{
				{'1', '1', '1'},
				{'0', '1', '0'},
				{'1', '1', '1'},
			},
			1,
		},
		{
			[][]byte{
				{'1'},
			},
			1,
		},
		{
			[][]byte{
				{'0'},
			},
			0,
		},
		{
			[][]byte{},
			0,
		},
	}

	fmt.Println("=== 解法一：深度优先搜索（DFS）===")
	for i, tc := range testCases {
		// 复制网格，避免修改原数据
		grid := make([][]byte, len(tc.grid))
		for j := range tc.grid {
			grid[j] = make([]byte, len(tc.grid[j]))
			copy(grid[j], tc.grid[j])
		}

		result := numIslands(grid)
		fmt.Printf("测试用例 %d: 结果=%d, 期望=%d, 通过=%t\n",
			i+1, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法二：广度优先搜索（BFS）===")
	for i, tc := range testCases {
		grid := make([][]byte, len(tc.grid))
		for j := range tc.grid {
			grid[j] = make([]byte, len(tc.grid[j]))
			copy(grid[j], tc.grid[j])
		}

		result := numIslands2(grid)
		fmt.Printf("测试用例 %d: 结果=%d, 期望=%d, 通过=%t\n",
			i+1, result, tc.expect, result == tc.expect)
	}

	fmt.Println("\n=== 解法三：并查集（Union-Find）===")
	for i, tc := range testCases {
		grid := make([][]byte, len(tc.grid))
		for j := range tc.grid {
			grid[j] = make([]byte, len(tc.grid[j]))
			copy(grid[j], tc.grid[j])
		}

		result := numIslands3(grid)
		fmt.Printf("测试用例 %d: 结果=%d, 期望=%d, 通过=%t\n",
			i+1, result, tc.expect, result == tc.expect)
	}
}

/*
解题思路：

1. 深度优先搜索（DFS）：
   - 遍历网格，遇到未访问的陆地时，开始DFS
   - DFS会将整个岛屿的所有陆地标记为已访问
   - 每次发现新的未访问陆地，岛屿数量+1
   - 使用递归实现，代码简洁

2. 广度优先搜索（BFS）：
   - 使用队列代替递归栈
   - 每次从队列取出一个位置，检查四个方向
   - 将未访问的相邻陆地加入队列
   - 空间复杂度更优，适合处理大型网格

3. 并查集（Union-Find）：
   - 将每个陆地位置看作一个独立的集合
   - 遍历时，将相邻的陆地合并到同一个集合
   - 最终岛屿数量 = 陆地总数 - 合并次数
   - 适合处理动态连接问题

时间复杂度分析：
- 所有解法：O(m*n)，每个位置最多访问一次

空间复杂度分析：
- DFS：O(m*n)，递归栈深度
- BFS：O(min(m,n))，队列长度
- 并查集：O(m*n)，存储父节点数组

关键点：
1. 岛屿的定义：被水包围的连通陆地
2. 访问标记：避免重复计算
3. 边界检查：防止数组越界
4. 方向数组：简化四个方向的遍历
*/
