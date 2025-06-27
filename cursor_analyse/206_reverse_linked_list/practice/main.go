package main

import (
	"fmt"
)

/*
题目：反转链表
难度：简单
标签：链表、递归

题目描述：
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

要求：
1. 反转整个链表
2. 不能改变节点内部的值，需要实际进行节点交换
3. 返回反转后的链表头节点

示例：
输入: head = [1,2,3,4,5]
输出: [5,4,3,2,1]

输入: head = [1,2]
输出: [2,1]

输入: head = []
输出: []

提示：
- 使用迭代或递归方法
- 注意处理边界情况
- 可以使用多个指针来辅助反转
*/

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO: 在这里实现你的算法
func reverseList(head *ListNode) *ListNode {
	// 请实现你的代码
	return nil
}

// 辅助函数：创建链表
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

// 辅助函数：打印链表
func printList(head *ListNode) []int {
	var result []int
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

// 辅助函数：比较两个链表是否相等
func equalList(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func main() {
	// 测试用例
	testCases := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}}, // 普通情况
		{[]int{1, 2}, []int{2, 1}},                   // 两个节点
		{[]int{1}, []int{1}},                         // 单个节点
		{[]int{}, []int{}},                           // 空链表
		{[]int{1, 2, 3}, []int{3, 2, 1}},             // 三个节点
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},       // 四个节点
	}

	for _, tc := range testCases {
		fmt.Printf("\n输入: %v\n", tc.input)
		head := createList(tc.input)
		result := reverseList(head)
		output := printList(result)
		fmt.Printf("输出: %v\n", output)
		fmt.Printf("期望: %v\n", tc.output)

		expected := createList(tc.output)
		if equalList(result, expected) {
			fmt.Println("✓ 通过")
		} else {
			fmt.Println("✗ 失败")
		}
	}
}

/*
预期输出：
输入: [1 2 3 4 5]
输出: [5 4 3 2 1]
期望: [5 4 3 2 1]
✓ 通过

输入: [1 2]
输出: [2 1]
期望: [2 1]
✓ 通过

输入: [1]
输出: [1]
期望: [1]
✓ 通过

输入: []
输出: []
期望: []
✓ 通过

输入: [1 2 3]
输出: [3 2 1]
期望: [3 2 1]
✓ 通过

输入: [1 2 3 4]
输出: [4 3 2 1]
期望: [4 3 2 1]
✓ 通过
*/
