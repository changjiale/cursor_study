package main

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

func main() {

}

type ListNode struct {
	Value int
	Next  *ListNode
}

func merge(node1 *ListNode, node2 *ListNode) *ListNode {
	summy := &ListNode{}
	head := summy
	for node1 != nil && node2 != nil {
		if node1.Value <= node2.Value {
			summy.Next = &ListNode{Value: node1.Value}
			node1 = node1.Next
		} else {
			summy.Next = &ListNode{Value: node2.Value}
			node2 = node2.Next
		}
	}
	if node1 != nil {
		summy.Next = node1
	}
	if node2 != nil {
		summy.Next = node2
	}

	return head.Next
}
