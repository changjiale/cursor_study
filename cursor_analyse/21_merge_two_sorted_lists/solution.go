package main

import (
	"fmt"
)

/*
题目：合并两个有序链表
难度：简单
标签：链表、递归

题目描述：
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

要求：
1. 合并后的链表应该保持升序
2. 不能修改原链表的结构
3. 返回新链表的头节点
*/

// 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一：迭代法
// 时间复杂度：O(n + m)
// 空间复杂度：O(1)
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建虚拟头节点
	dummy := &ListNode{}
	current := dummy

	// 比较两个链表的节点值，选择较小的节点
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			current.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		current = current.Next
	}

	// 处理剩余的节点
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}

// 解法二：递归法
// 时间复杂度：O(n + m)
// 空间复杂度：O(n + m) - 递归调用栈
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 递归终止条件
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 递归合并
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
}

// 解法三：原地合并（修改原链表）
// 时间复杂度：O(n + m)
// 空间复杂度：O(1)
func mergeTwoLists3(l1 *ListNode, l2 *ListNode) *ListNode {
	// 处理空链表的情况
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 选择较小的头节点作为结果的头
	var head, current *ListNode
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	current = head

	// 合并剩余的节点
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// 处理剩余的节点
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return head
}

// 辅助函数：创建链表
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

// 辅助函数：打印链表
func printList(head *ListNode) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func main() {
	// 测试用例
	testCases := []struct {
		l1 []int
		l2 []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}}, // 普通情况
		{[]int{}, []int{}},               // 两个空链表
		{[]int{}, []int{0}},              // 一个空链表
		{[]int{1, 3, 5}, []int{2, 4, 6}}, // 交替合并
		{[]int{1, 2, 3}, []int{4, 5, 6}}, // 顺序合并
		{[]int{4, 5, 6}, []int{1, 2, 3}}, // 逆序合并
	}

	for _, tc := range testCases {
		fmt.Printf("\n输入: l1 = %v, l2 = %v\n", tc.l1, tc.l2)

		// 解法一
		l1_1 := createList(tc.l1)
		l2_1 := createList(tc.l2)
		result1 := mergeTwoLists1(l1_1, l2_1)
		output1 := printList(result1)
		fmt.Printf("解法一（迭代法）结果: %v\n", output1)

		// 解法二
		l1_2 := createList(tc.l1)
		l2_2 := createList(tc.l2)
		result2 := mergeTwoLists2(l1_2, l2_2)
		output2 := printList(result2)
		fmt.Printf("解法二（递归法）结果: %v\n", output2)

		// 解法三
		l1_3 := createList(tc.l1)
		l2_3 := createList(tc.l2)
		result3 := mergeTwoLists3(l1_3, l2_3)
		output3 := printList(result3)
		fmt.Printf("解法三（原地合并）结果: %v\n", output3)
	}
}

/*
预期输出：
输入: l1 = [1 2 4], l2 = [1 3 4]
解法一（迭代法）结果: [1 1 2 3 4 4]
解法二（递归法）结果: [1 1 2 3 4 4]
解法三（原地合并）结果: [1 1 2 3 4 4]

输入: l1 = [], l2 = []
解法一（迭代法）结果: []
解法二（递归法）结果: []
解法三（原地合并）结果: []

输入: l1 = [], l2 = [0]
解法一（迭代法）结果: [0]
解法二（递归法）结果: [0]
解法三（原地合并）结果: [0]

输入: l1 = [1 3 5], l2 = [2 4 6]
解法一（迭代法）结果: [1 2 3 4 5 6]
解法二（递归法）结果: [1 2 3 4 5 6]
解法三（原地合并）结果: [1 2 3 4 5 6]

输入: l1 = [1 2 3], l2 = [4 5 6]
解法一（迭代法）结果: [1 2 3 4 5 6]
解法二（递归法）结果: [1 2 3 4 5 6]
解法三（原地合并）结果: [1 2 3 4 5 6]

输入: l1 = [4 5 6], l2 = [1 2 3]
解法一（迭代法）结果: [1 2 3 4 5 6]
解法二（递归法）结果: [1 2 3 4 5 6]
解法三（原地合并）结果: [1 2 3 4 5 6]
*/
