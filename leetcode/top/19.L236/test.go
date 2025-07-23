package main

func test(root, p, q *TreeNode) *TreeNode {

	if root == nil || root == p || root == q {
		return root
	}

	left := test(root.Left, p, q)
	right := test(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right

}
