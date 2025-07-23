package main

type ListNode1 struct {
	Value int
	Next  *ListNode1
}

func mergeListNode(node1 *ListNode1, node2 *ListNode1) *ListNode1 {
	summy := &ListNode1{}
	head := summy
	for node1 != nil && node2 != nil {
		if node1.Value < node2.Value {
			head.Next = &ListNode1{Value: node1.Value}
			node1 = node1.Next
		} else {
			head.Next = &ListNode1{Value: node2.Value}
			node2 = node2.Next
		}
		head = head.Next
	}

	if node1 != nil {
		head.Next = node1
	}
	if node2 != nil {
		head.Next = node2
	}

	return summy.Next
}
