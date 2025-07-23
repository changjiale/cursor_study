package main

/*
*给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]

提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

进阶：你能尝试使用一趟扫描实现吗？
*/
func main() {

}

type ListNode struct {
	Val int
	Next *ListNode
}	

// 解法一：双指针法（推荐）
// 核心思想：使用快慢指针，快指针先走n步，然后快慢指针同时前进
// 当快指针到达末尾时，慢指针指向要删除节点的前一个节点
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 创建虚拟头节点，避免处理头节点被删除的情况
	dummy := &ListNode{Val: 0, Next: head}

	// 快慢指针都指向虚拟头节点
	fast, slow := dummy, dummy

	// 快指针先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 快慢指针同时前进，直到快指针到达末尾
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 删除慢指针后面的节点
	slow.Next = slow.Next.Next

	return dummy.Next
}
