package main

import (
	"fmt"
)

/*
图例说明：

1. 原始链表：
   1 -> 2 -> 3 -> 4 -> 5 -> nil
   k = 2

2. 第一次反转后：
   2 -> 1 -> 4 -> 3 -> 5 -> nil
   |    |    |    |    |
  组1  组1  组2  组2  组3

3. 反转过程示意：
   原始：    1 -> 2 -> 3 -> 4 -> 5
   |        |    |    |    |    |
   prev    head  next  next  next  next

   反转后：  2 -> 1 -> 3 -> 4 -> 5
   |        |    |    |    |    |
   prev    head  next  next  next  next

4. 递归过程示意：
   第一次：  2 -> 1 -> 3 -> 4 -> 5
   |        |    |    |    |    |
   newHead  head tail  next next next

   第二次：  2 -> 1 -> 4 -> 3 -> 5
   |        |    |    |    |    |
   newHead  head tail  next next next
*/

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一：迭代法
// 时间复杂度：O(n)，空间复杂度：O(1)
func reverseKGroup1(head *ListNode, k int) *ListNode {
	// 创建虚拟头节点，简化边界处理
	dummy := &ListNode{Next: head}
	// prev指向当前组的头节点的前一个节点
	prev := dummy

	for head != nil {
		// 找到当前组的尾节点
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				// 如果剩余节点不足k个，直接返回
				return dummy.Next
			}
		}

		// 保存下一组的头节点
		next := tail.Next

		// 反转当前组
		head, tail = reverse(head, tail)

		// 将反转后的组连接到链表中
		prev.Next = head
		tail.Next = next

		// 更新prev和head，准备处理下一组
		prev = tail
		head = next
	}

	return dummy.Next
}

// 反转从head到tail的链表
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	curr := head

	for prev != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return tail, head
}

// 解法二：递归法
// 时间复杂度：O(n)，空间复杂度：O(n/k)（递归调用栈的深度）
func reverseKGroup2(head *ListNode, k int) *ListNode {
	// 找到当前组的尾节点
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			// 如果剩余节点不足k个，直接返回
			return head
		}
		tail = tail.Next
	}

	// 反转当前组
	newHead := reverseBetween(head, tail)
	// 递归处理剩余部分
	head.Next = reverseKGroup2(tail, k)

	return newHead
}

// 反转从head到tail（不包含tail）的链表
func reverseBetween(head, tail *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// 解法三：极简写法（迭代）
// 时间复杂度：O(n)，空间复杂度：O(1)
func reverseKGroupSimple(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil {
		// 检查剩余节点是否足够k个
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}

		// 反转当前组
		next := tail.Next
		tail.Next = nil
		prev.Next = reverseList(head)
		head.Next = next

		// 更新指针
		prev = head
		head = next
	}

	return dummy.Next
}

// 反转整个链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

// 创建链表
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

// 打印链表
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	// 测试用例
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3, 4, 5}, 2}, // 预期输出：2->1->4->3->5
		{[]int{1, 2, 3, 4, 5}, 3}, // 预期输出：3->2->1->4->5
		{[]int{1, 2, 3, 4, 5}, 1}, // 预期输出：1->2->3->4->5
		{[]int{1}, 1},             // 预期输出：1
	}

	for _, tc := range testCases {
		fmt.Printf("\n输入: nums = %v, k = %d\n", tc.nums, tc.k)

		// 解法一
		head1 := createList(tc.nums)
		fmt.Println("解法一（迭代）结果: ")
		reversed1 := reverseKGroup1(head1, tc.k)
		printList(reversed1)

		// 解法二
		head2 := createList(tc.nums)
		fmt.Println("解法二（递归）结果: ")
		reversed2 := reverseKGroup2(head2, tc.k)
		printList(reversed2)

		// 解法三
		head3 := createList(tc.nums)
		fmt.Println("解法三（极简）结果: ")
		reversed3 := reverseKGroupSimple(head3, tc.k)
		printList(reversed3)
	}
}
