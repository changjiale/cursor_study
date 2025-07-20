package main

import "fmt"

/*
题目：删除链表的倒数第N个节点
难度：中等
标签：链表、双指针

题目描述：
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

要求：
- 时间复杂度：O(n)，其中 n 是链表的长度
- 空间复杂度：O(1)

示例：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

输入：head = [1], n = 1
输出：[]

输入：head = [1,2], n = 1
输出：[1]

注意：
- 链表中结点的数目为 sz
- 1 <= sz <= 30
- 0 <= Node.val <= 100
- 1 <= n <= sz
*/

// 链表节点定义
type ListNode struct {
	Val  int
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

// 解法二：计算链表长度
// 核心思想：先计算链表长度，然后找到要删除节点的前一个节点
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	// 计算链表长度
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	// 如果要删除的是头节点
	if n == length {
		return head.Next
	}

	// 找到要删除节点的前一个节点
	current = head
	for i := 0; i < length-n-1; i++ {
		current = current.Next
	}

	// 删除节点
	current.Next = current.Next.Next

	return head
}

// 解法三：递归法
// 核心思想：使用递归，在回溯时计数，找到倒数第n个节点
// 时间复杂度：O(n)
// 空间复杂度：O(n)，递归调用栈的深度
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	// 使用全局变量记录当前是倒数第几个节点
	var count int

	// 递归函数
	var dfs func(node *ListNode) *ListNode
	dfs = func(node *ListNode) *ListNode {
		if node == nil {
			return nil
		}

		// 递归处理下一个节点
		node.Next = dfs(node.Next)

		// 回溯时计数
		count++

		// 如果当前节点是要删除的节点
		if count == n {
			return node.Next
		}

		return node
	}

	return dfs(head)
}

// 解法四：栈方法
// 核心思想：将所有节点压入栈中，然后弹出n个节点，第n个就是要删除的节点
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func removeNthFromEnd4(head *ListNode, n int) *ListNode {
	// 创建虚拟头节点
	dummy := &ListNode{Val: 0, Next: head}

	// 将所有节点压入栈中
	stack := make([]*ListNode, 0)
	current := dummy
	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	// 弹出n个节点，第n个就是要删除的节点
	for i := 0; i < n; i++ {
		stack = stack[:len(stack)-1]
	}

	// 获取要删除节点的前一个节点
	prev := stack[len(stack)-1]

	// 删除节点
	prev.Next = prev.Next.Next

	return dummy.Next
}

// 解法五：双指针法（优化版本）
// 核心思想：与解法一相同，但代码更简洁
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeNthFromEnd5(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// 快指针先走n+1步
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// 快慢指针同时前进
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 删除节点
	slow.Next = slow.Next.Next

	return dummy.Next
}

// 解法六：两次遍历法
// 核心思想：第一次遍历计算长度，第二次遍历删除节点
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeNthFromEnd6(head *ListNode, n int) *ListNode {
	// 第一次遍历：计算链表长度
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	// 计算要删除节点的位置（从1开始计数）
	targetPos := length - n + 1

	// 如果要删除的是第一个节点
	if targetPos == 1 {
		return head.Next
	}

	// 第二次遍历：找到要删除节点的前一个节点
	current = head
	for i := 1; i < targetPos-1; i++ {
		current = current.Next
	}

	// 删除节点
	current.Next = current.Next.Next

	return head
}

// 解法七：单指针法
// 核心思想：使用单个指针，先计算长度，再找到目标位置
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeNthFromEnd7(head *ListNode, n int) *ListNode {
	// 计算链表长度
	length := 0
	for current := head; current != nil; current = current.Next {
		length++
	}

	// 如果要删除的是头节点
	if n == length {
		return head.Next
	}

	// 找到要删除节点的前一个节点
	current := head
	for i := 0; i < length-n-1; i++ {
		current = current.Next
	}

	// 删除节点
	current.Next = current.Next.Next

	return head
}

// 创建测试用的链表
func createLinkedList(nums []int) *ListNode {
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

// 打印链表
func printList(head *ListNode) string {
	if head == nil {
		return "[]"
	}

	result := "["
	current := head
	for current != nil {
		result += fmt.Sprintf("%d", current.Val)
		if current.Next != nil {
			result += " -> "
		}
		current = current.Next
	}
	result += "]"
	return result
}

// 复制链表
func copyList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := &ListNode{Val: head.Val}
	current := newHead
	original := head.Next

	for original != nil {
		current.Next = &ListNode{Val: original.Val}
		current = current.Next
		original = original.Next
	}

	return newHead
}

func main() {
	// 测试用例
	testCases := []struct {
		nums []int
		n    int
		desc string
	}{
		{[]int{1, 2, 3, 4, 5}, 2, "删除倒数第2个节点"},
		{[]int{1}, 1, "删除唯一节点"},
		{[]int{1, 2}, 1, "删除倒数第1个节点"},
		{[]int{1, 2, 3}, 3, "删除倒数第3个节点（头节点）"},
		{[]int{1, 2, 3, 4}, 4, "删除倒数第4个节点（头节点）"},
		{[]int{1, 2, 3, 4, 5}, 1, "删除倒数第1个节点"},
		{[]int{1, 2, 3, 4, 5}, 5, "删除倒数第5个节点（头节点）"},
	}

	fmt.Println("=== 解法一：双指针法（推荐）===")
	for i, tc := range testCases {
		head := createLinkedList(tc.nums)
		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("原链表: %s\n", printList(head))
		result := removeNthFromEnd(head, tc.n)
		fmt.Printf("删除后: %s\n\n", printList(result))
	}

	fmt.Println("=== 解法二：计算链表长度 ===")
	for i, tc := range testCases {
		head := createLinkedList(tc.nums)
		result := removeNthFromEnd2(head, tc.n)
		fmt.Printf("测试用例 %d: %s, 结果: %s\n", i+1, tc.desc, printList(result))
	}

	fmt.Println("\n=== 解法三：递归法 ===")
	for i, tc := range testCases {
		head := createLinkedList(tc.nums)
		result := removeNthFromEnd3(head, tc.n)
		fmt.Printf("测试用例 %d: %s, 结果: %s\n", i+1, tc.desc, printList(result))
	}

	fmt.Println("\n=== 解法四：栈方法 ===")
	for i, tc := range testCases {
		head := createLinkedList(tc.nums)
		result := removeNthFromEnd4(head, tc.n)
		fmt.Printf("测试用例 %d: %s, 结果: %s\n", i+1, tc.desc, printList(result))
	}

	fmt.Println("\n=== 解法五：双指针法（优化版本） ===")
	for i, tc := range testCases {
		head := createLinkedList(tc.nums)
		result := removeNthFromEnd5(head, tc.n)
		fmt.Printf("测试用例 %d: %s, 结果: %s\n", i+1, tc.desc, printList(result))
	}

	fmt.Println("\n=== 解法六：两次遍历法 ===")
	for i, tc := range testCases {
		head := createLinkedList(tc.nums)
		result := removeNthFromEnd6(head, tc.n)
		fmt.Printf("测试用例 %d: %s, 结果: %s\n", i+1, tc.desc, printList(result))
	}

	fmt.Println("\n=== 解法七：单指针法 ===")
	for i, tc := range testCases {
		head := createLinkedList(tc.nums)
		result := removeNthFromEnd7(head, tc.n)
		fmt.Printf("测试用例 %d: %s, 结果: %s\n", i+1, tc.desc, printList(result))
	}
}

/*
解题思路：

1. 双指针法（推荐）：
   - 核心思想：使用快慢指针，快指针先走n步，然后快慢指针同时前进
   - 当快指针到达末尾时，慢指针指向要删除节点的前一个节点
   - 时间复杂度：O(n)，空间复杂度：O(1)
   - 优点：一次遍历完成，空间效率高

2. 计算链表长度：
   - 先计算链表长度，然后找到要删除节点的前一个节点
   - 时间复杂度：O(n)，空间复杂度：O(1)
   - 优点：逻辑清晰，易于理解

3. 递归法：
   - 使用递归，在回溯时计数，找到倒数第n个节点
   - 时间复杂度：O(n)，空间复杂度：O(n)
   - 优点：代码简洁，缺点：空间复杂度高

4. 栈方法：
   - 将所有节点压入栈中，然后弹出n个节点
   - 时间复杂度：O(n)，空间复杂度：O(n)
   - 优点：思路清晰，缺点：空间复杂度高

5. 双指针法（优化版本）：
   - 与解法一相同，但代码更简洁
   - 时间复杂度：O(n)，空间复杂度：O(1)
   - 优点：代码最简洁

6. 两次遍历法：
   - 第一次遍历计算长度，第二次遍历删除节点
   - 时间复杂度：O(n)，空间复杂度：O(1)
   - 优点：逻辑清晰，缺点：需要两次遍历

7. 单指针法：
   - 使用单个指针，先计算长度，再找到目标位置
   - 时间复杂度：O(n)，空间复杂度：O(1)
   - 优点：实现简单，缺点：需要两次遍历

关键点：
1. 使用虚拟头节点避免处理头节点被删除的情况
2. 理解双指针法的原理：通过快慢指针的步数差来定位
3. 注意边界情况：删除头节点、删除唯一节点
4. 掌握不同解法的优缺点和适用场景

时间复杂度分析：
- 双指针法、计算长度、递归法、栈方法：O(n)
- 两次遍历法、单指针法：O(n)（虽然两次遍历，但总时间复杂度仍是O(n)）

空间复杂度分析：
- 双指针法、计算长度、两次遍历法、单指针法：O(1)
- 递归法、栈方法：O(n)

面试要点：
1. 能够解释双指针法的原理
2. 理解为什么双指针法能够找到倒数第n个节点
3. 掌握虚拟头节点的使用技巧
4. 考虑边界情况的处理
5. 能够分析不同解法的优缺点

优化技巧：
1. 使用虚拟头节点统一处理逻辑
2. 双指针法是最优解，推荐在面试中使用
3. 注意指针的移动顺序和边界条件
4. 理解快慢指针的步数关系
*/
