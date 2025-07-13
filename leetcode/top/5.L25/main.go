package main

/**
25. K 个一组翻转链表
困难
相关标签
premium lock icon
相关企业

给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。



示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
示例 2：



输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]


提示：
链表中的节点数目为 n
1 <= k <= n <= 5000
0 <= Node.val <= 1000


进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
**/

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy

	for head != nil {
		tail := prev
		//找到当前组的尾节点
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}

		//下一组头结点
		next := tail.Next
		//翻转当前链表
		head, tail = reverse(head, tail)
		//翻转后从新赋值
		prev.Next = head
		tail.Next = next
		prev = tail 
		head = next
		
	}
	return dummy.Next

}

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
