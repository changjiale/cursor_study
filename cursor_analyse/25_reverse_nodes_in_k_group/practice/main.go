package main

import (
	"fmt"
)

/*
题目：K 个一组翻转链表
难度：困难
标签：链表、递归

题目描述：
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

要求：
1. 每k个节点为一组进行翻转
2. 如果剩余节点不足k个，保持原有顺序
3. 不能改变节点内部的值，需要实际交换节点

示例：
输入: head = [1,2,3,4,5], k = 2
输出: [2,1,4,3,5]

输入: head = [1,2,3,4,5], k = 3
输出: [3,2,1,4,5]

输入: head = [1,2,3,4,5], k = 1
输出: [1,2,3,4,5]

提示：
- 使用迭代或递归方法
- 注意处理边界情况
- 可以使用虚拟头节点简化操作
*/

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO: 在这里实现你的算法
func reverseKGroup(head *ListNode, k int) *ListNode {
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
		k      int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},       // 普通情况
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},       // k=3
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},       // k=1
		{[]int{1, 2, 3, 4, 5}, 5, []int{5, 4, 3, 2, 1}},       // k=链表长度
		{[]int{1, 2, 3, 4, 5}, 6, []int{1, 2, 3, 4, 5}},       // k>链表长度
		{[]int{1}, 1, []int{1}},                               // 单个节点
		{[]int{}, 1, []int{}},                                 // 空链表
		{[]int{1, 2, 3, 4}, 2, []int{2, 1, 4, 3}},             // 偶数长度
		{[]int{1, 2, 3, 4, 5, 6}, 2, []int{2, 1, 4, 3, 6, 5}}, // 完整分组
	}

	for _, tc := range testCases {
		fmt.Printf("\n输入: %v, k = %d\n", tc.input, tc.k)
		head := createList(tc.input)
		result := reverseKGroup(head, tc.k)
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
输入: [1 2 3 4 5], k = 2
输出: [2 1 4 3 5]
期望: [2 1 4 3 5]
✓ 通过

输入: [1 2 3 4 5], k = 3
输出: [3 2 1 4 5]
期望: [3 2 1 4 5]
✓ 通过

输入: [1 2 3 4 5], k = 1
输出: [1 2 3 4 5]
期望: [1 2 3 4 5]
✓ 通过

输入: [1 2 3 4 5], k = 5
输出: [5 4 3 2 1]
期望: [5 4 3 2 1]
✓ 通过

输入: [1 2 3 4 5], k = 6
输出: [1 2 3 4 5]
期望: [1 2 3 4 5]
✓ 通过

输入: [1], k = 1
输出: [1]
期望: [1]
✓ 通过

输入: [], k = 1
输出: []
期望: []
✓ 通过

输入: [1 2 3 4], k = 2
输出: [2 1 4 3]
期望: [2 1 4 3]
✓ 通过

输入: [1 2 3 4 5 6], k = 2
输出: [2 1 4 3 6 5]
期望: [2 1 4 3 6 5]
✓ 通过
*/
