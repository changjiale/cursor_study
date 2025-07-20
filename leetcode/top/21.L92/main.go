package main

func main() {

}

/*
题目：反转链表 II
难度：中等
标签：链表

题目描述：
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回反转后的链表。

要求：
- 时间复杂度：O(n)
- 空间复杂度：O(1)

示例：
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

输入：head = [5], left = 1, right = 1
输出：[5]
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
解法一：头插法（推荐）
核心思想：使用头插法，将反转区域的节点逐个插入到反转区域的头部

算法步骤：
1. 创建虚拟头节点，避免边界情况处理
2. 找到反转区域的前一个节点pre
3. 使用头插法：将cur后面的节点逐个插入到pre后面
4. 返回虚拟头节点的下一个节点

时间复杂度：O(n)
空间复杂度：O(1)
*/
func reverseBetween(head *ListNode, left, right int) *ListNode {
	// 1. 创建虚拟头节点，避免处理头节点被反转的边界情况
	// 这是链表问题中的常用技巧，可以统一所有情况的处理逻辑
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode

	// 2. 找到反转区域的前一个节点
	// 需要走left-1步，因为pre初始指向虚拟头节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 3. cur指向反转区域的第一个节点
	cur := pre.Next

	// 4. 头插法反转：将cur后面的节点逐个插入到pre后面
	// 需要执行right-left次，因为要反转right-left+1个节点
	for i := 0; i < right-left; i++ {
		next := cur.Next     // 保存下一个要处理的节点
		cur.Next = next.Next // cur指向next的下一个节点（跳过next）
		next.Next = pre.Next // next指向pre的下一个节点（插入到头部）
		pre.Next = next      // pre指向next（更新头部）
	}

	return dummyNode.Next
}

/*
头插法反转过程详解：

假设链表：1 -> 2 -> 3 -> 4 -> 5，反转位置2到4

初始状态：
dummy -> 1 -> 2 -> 3 -> 4 -> 5
         pre  cur

第一次循环：
1. next = cur.Next = 3
2. cur.Next = next.Next = 4  (2 -> 4)
3. next.Next = pre.Next = 2  (3 -> 2)
4. pre.Next = next = 3       (1 -> 3)

结果：dummy -> 1 -> 3 -> 2 -> 4 -> 5
              pre       cur

第二次循环：
1. next = cur.Next = 4
2. cur.Next = next.Next = 5  (2 -> 5)
3. next.Next = pre.Next = 3  (4 -> 3)
4. pre.Next = next = 4       (1 -> 4)

结果：dummy -> 1 -> 4 -> 3 -> 2 -> 5
              pre            cur

最终：1 -> 4 -> 3 -> 2 -> 5
*/

/*
解法二：切断重连法
核心思想：先切断反转区域，反转后再重新连接

算法步骤：
1. 创建虚拟头节点
2. 找到反转区域的边界节点
3. 切断反转区域
4. 反转子链表
5. 重新连接

时间复杂度：O(n)
空间复杂度：O(1)
*/
func reverseBetween2(head *ListNode, left, right int) *ListNode {
	// 1. 创建虚拟头节点，避免头节点变化的复杂处理
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head

	pre := dummyNode
	// 2. 找到反转区域的前一个节点
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	// 3. 找到反转区域的最后一个节点
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	// 4. 切断反转区域
	leftNode := pre.Next   // 反转区域的第一个节点
	curr := rightNode.Next // 反转区域后的第一个节点

	// 切断链接
	pre.Next = nil
	rightNode.Next = nil

	// 5. 反转子链表
	reverseLinkedList(leftNode)

	// 6. 重新连接
	pre.Next = rightNode // 连接反转后的尾部
	leftNode.Next = curr // 连接反转区域后的部分

	return dummyNode.Next
}

/*
反转整个链表的辅助函数
使用迭代法反转链表
*/
func reverseLinkedList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // 保存下一个节点
		cur.Next = pre   // 反转当前节点的指向
		pre = cur        // 更新pre
		cur = next       // 移动到下一个节点
	}
}

/*
两种解法的对比：

1. 头插法（解法一）：
   - 优点：代码简洁，一次遍历完成，空间效率高
   - 缺点：指针操作相对复杂，需要理解头插法的原理
   - 适用：面试中推荐使用，展示对链表的深入理解

2. 切断重连法（解法二）：
   - 优点：逻辑清晰，易于理解，分步骤处理
   - 缺点：需要多次遍历，代码较长
   - 适用：学习阶段使用，便于理解反转过程

关键技巧总结：
1. 虚拟头节点：避免边界情况处理
2. 头插法：高效的反转技巧
3. 指针操作：理解指针的指向关系
4. 边界处理：考虑left=1和right=链表长度的情况

时间复杂度分析：
- 两种解法都是O(n)，每个节点最多访问一次

空间复杂度分析：
- 两种解法都是O(1)，只使用常数额外空间

面试要点：
1. 能够解释头插法的原理
2. 理解虚拟头节点的作用
3. 掌握指针操作的技巧
4. 考虑边界情况的处理
*/
