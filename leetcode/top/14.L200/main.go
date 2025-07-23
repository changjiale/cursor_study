package main

/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

 

示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3
 

提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/
func main() {

}



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