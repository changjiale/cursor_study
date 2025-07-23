package main

/*
*
重新练 重点
*/
func reverseK(head *ListNode, k int) *ListNode {

	dummy := &ListNode{0, head}
	pre := dummy

	for head != nil {
		tail := pre
		//找到k个一组

		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}

		//下一组的开始
		next := tail.Next
		head, tail = reverse1(head, tail)

		//翻转后重新赋值
		pre.Next = head
		tail.Next = next
		pre = tail
		head = next

	}
	return dummy.Next

}

func reverse1(head, tail *ListNode) (*ListNode, *ListNode) {

	pre := tail.Next
	cur := head
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return tail, head

}
