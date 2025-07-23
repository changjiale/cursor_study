package main

func test(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	pre := dummyNode

	//先走left -1
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next
	for i := 0; i < right-left; i++ {
		//   1 		2 		3 		4
		//   pre 	cur
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next

	}

	return dummyNode.Next

}
