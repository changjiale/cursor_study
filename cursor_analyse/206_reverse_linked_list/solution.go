package main

import (
	"fmt"
)

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一：迭代法
// 时间复杂度：O(n)，空间复杂度：O(1)
func reverseList1(head *ListNode) *ListNode {
	// 定义前驱节点和当前节点
	var prev *ListNode
	curr := head

	// 遍历链表
	for curr != nil {
		// 保存下一个节点
		next := curr.Next
		// 反转当前节点的指针
		curr.Next = prev
		// 移动前驱节点和当前节点
		prev = curr
		curr = next
	}
	// 返回新的头节点（原链表的尾节点）
	return prev
}

// 解法二：递归法
// 时间复杂度：O(n)，空间复杂度：O(n)（递归调用栈的深度）
func reverseList2(head *ListNode) *ListNode {
	// 递归终止条件：空节点或只有一个节点
	if head == nil || head.Next == nil {
		return head
	}

	// 递归反转剩余部分
	newHead := reverseList2(head.Next)
	// 反转当前节点
	head.Next.Next = head
	// 将当前节点的Next设为nil，避免形成环
	head.Next = nil

	return newHead
}

// 解法三：极简写法（迭代）
// 时间复杂度：O(n)，空间复杂度：O(1)
func reverseListSimple(head *ListNode) *ListNode {
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
	testCases := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2},
		{1},
		{},
	}

	for _, nums := range testCases {
		fmt.Printf("\n原始链表: ")
		head := createList(nums)
		printList(head)

		fmt.Println("解法一（迭代）结果: ")
		reversed1 := reverseList1(head)
		printList(reversed1)

		fmt.Println("解法二（递归）结果: ")
		head = createList(nums) // 重新创建链表
		reversed2 := reverseList2(head)
		printList(reversed2)

		fmt.Println("解法三（极简）结果: ")
		head = createList(nums) // 重新创建链表
		reversed3 := reverseListSimple(head)
		printList(reversed3)
	}
}
