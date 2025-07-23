package main

func test(matrix [][]int) []int {
	visted := make([][]bool, 0)

	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		visted[i] = make([]bool, cols)
	}

	var (
		total     = rows * cols
		direction = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		order     = make([]int, total)
		row, col  = 0, 0
		dirIndex  = 0
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][col]
		visted[row][col] = true
		nextRow, nextColumn := row+direction[dirIndex][0], col+direction[dirIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= cols || visted[nextRow][nextColumn] {
			dirIndex = (dirIndex + 1) % 4
		}
		row = row + direction[dirIndex][0]
		col = col + direction[dirIndex][1]
	}

	return order

}
