package main

func test(root *TreeNode) [][]int {

	if root == nil {
		return nil
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
			if level%2 == 0 {
				levelNodes[i] = node.Val
			} else {
				levelNodes[levelSize-i-1] = node.Val
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		result = append(result, levelNodes)
		level++
	}

	return result

}
