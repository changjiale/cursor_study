package main

func test(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{
		Next: head,
	}

	slow, fast := dummyNode, dummyNode

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummyNode.Next

}
