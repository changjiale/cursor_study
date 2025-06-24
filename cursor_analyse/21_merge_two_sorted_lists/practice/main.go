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

示例：
输入: l1 = [1,2,4], l2 = [1,3,4]
输出: [1,1,2,3,4,4]

输入: l1 = [], l2 = []
输出: []

输入: l1 = [], l2 = [0]
输出: [0]

提示：
1. 可以使用递归或迭代的方法
2. 注意处理空链表的情况
3. 比较节点值的大小来决定合并顺序
*/

// 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO: 在这里实现你的算法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 请实现你的代码
	return nil
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

		l1 := createList(tc.l1)
		l2 := createList(tc.l2)

		result := mergeTwoLists(l1, l2)
		output := printList(result)

		fmt.Printf("输出: %v\n", output)
	}
}

/*
预期输出：
输入: l1 = [1 2 4], l2 = [1 3 4]
输出: [1 1 2 3 4 4]

输入: l1 = [], l2 = []
输出: []

输入: l1 = [], l2 = [0]
输出: [0]

输入: l1 = [1 3 5], l2 = [2 4 6]
输出: [1 2 3 4 5 6]

输入: l1 = [1 2 3], l2 = [4 5 6]
输出: [1 2 3 4 5 6]

输入: l1 = [4 5 6], l2 = [1 2 3]
输出: [1 2 3 4 5 6]
*/
