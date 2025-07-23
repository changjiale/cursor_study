package main

func test(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		leveSize := len(queue)
		for i := 0; i < leveSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == leveSize-1 {
				result = append(result, node.Val)
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return result
}
