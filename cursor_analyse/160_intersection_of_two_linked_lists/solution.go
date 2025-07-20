package main

import "fmt"

/*
题目：相交链表
难度：简单
标签：链表、双指针

题目描述：
给你两个单链表的头节点 headA 和 headB，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null。

要求：
- 时间复杂度：O(m+n)，其中 m 和 n 是链表 A 和 B 的长度
- 空间复杂度：O(1)

示例：
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。

输入：intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Intersected at '2'

输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
*/

// 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一：双指针法（推荐）
// 核心思想：让两个指针分别遍历两个链表，当到达末尾时交换到另一个链表的头部
// 这样两个指针走过的距离相等，如果有交点，必定会在交点相遇
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// 两个指针分别指向两个链表的头部
	ptrA, ptrB := headA, headB

	// 当两个指针不相等时继续遍历
	for ptrA != ptrB {
		// 如果ptrA到达末尾，则指向headB
		if ptrA == nil {
			ptrA = headB
		} else {
			ptrA = ptrA.Next
		}

		// 如果ptrB到达末尾，则指向headA
		if ptrB == nil {
			ptrB = headA
		} else {
			ptrB = ptrB.Next
		}
	}

	// 返回交点（如果不相交，最终都是nil）
	return ptrA
}

// 解法二：计算长度差
// 核心思想：先计算两个链表的长度差，让长链表的指针先走差值步
// 然后两个指针同时前进，必定会在交点相遇
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// 计算两个链表的长度
	lenA, lenB := getLength(headA), getLength(headB)

	// 让长链表的指针先走差值步
	ptrA, ptrB := headA, headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			ptrA = ptrA.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			ptrB = ptrB.Next
		}
	}

	// 两个指针同时前进，直到相遇或到达末尾
	for ptrA != ptrB {
		ptrA = ptrA.Next
		ptrB = ptrB.Next
	}

	return ptrA
}

// 计算链表长度的辅助函数
func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}

// 解法三：哈希集合
// 核心思想：遍历第一个链表，将所有节点加入哈希集合
// 然后遍历第二个链表，检查节点是否在集合中
// 时间复杂度：O(m+n)
// 空间复杂度：O(m) 或 O(n)
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// 使用map存储第一个链表的所有节点
	visited := make(map[*ListNode]bool)

	// 遍历第一个链表，将所有节点加入集合
	ptr := headA
	for ptr != nil {
		visited[ptr] = true
		ptr = ptr.Next
	}

	// 遍历第二个链表，检查是否有节点在集合中
	ptr = headB
	for ptr != nil {
		if visited[ptr] {
			return ptr
		}
		ptr = ptr.Next
	}

	return nil
}

// 解法四：暴力法
// 核心思想：对第一个链表的每个节点，遍历第二个链表查找相同节点
// 时间复杂度：O(m*n)
// 空间复杂度：O(1)
func getIntersectionNode4(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// 遍历第一个链表的每个节点
	ptrA := headA
	for ptrA != nil {
		// 对每个节点，遍历第二个链表
		ptrB := headB
		for ptrB != nil {
			if ptrA == ptrB {
				return ptrA
			}
			ptrB = ptrB.Next
		}
		ptrA = ptrA.Next
	}

	return nil
}

// 解法五：双指针法（优化版本）
// 核心思想：与解法一相同，但代码更简洁
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func getIntersectionNode5(headA, headB *ListNode) *ListNode {
	ptrA, ptrB := headA, headB

	// 当两个指针不相等时继续遍历
	for ptrA != ptrB {
		// 如果ptrA到达末尾，则指向headB，否则前进
		if ptrA == nil {
			ptrA = headB
		} else {
			ptrA = ptrA.Next
		}

		// 如果ptrB到达末尾，则指向headA，否则前进
		if ptrB == nil {
			ptrB = headA
		} else {
			ptrB = ptrB.Next
		}
	}

	return ptrA
}

// 解法六：栈方法
// 核心思想：将两个链表分别压入栈中，然后同时弹出比较
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n)
func getIntersectionNode6(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// 将两个链表分别压入栈中
	stackA := make([]*ListNode, 0)
	stackB := make([]*ListNode, 0)

	// 压入链表A
	ptr := headA
	for ptr != nil {
		stackA = append(stackA, ptr)
		ptr = ptr.Next
	}

	// 压入链表B
	ptr = headB
	for ptr != nil {
		stackB = append(stackB, ptr)
		ptr = ptr.Next
	}

	// 从栈顶开始比较，找到最后一个相同的节点
	var result *ListNode
	for len(stackA) > 0 && len(stackB) > 0 {
		nodeA := stackA[len(stackA)-1]
		nodeB := stackB[len(stackB)-1]

		if nodeA == nodeB {
			result = nodeA
			stackA = stackA[:len(stackA)-1]
			stackB = stackB[:len(stackB)-1]
		} else {
			break
		}
	}

	return result
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

// 创建相交链表
func createIntersectionList(listA, listB []int, intersection []int) (*ListNode, *ListNode) {
	// 创建相交部分
	intersectionHead := createLinkedList(intersection)

	// 创建链表A
	headA := createLinkedList(listA)
	if headA == nil {
		headA = intersectionHead
	} else {
		current := headA
		for current.Next != nil {
			current = current.Next
		}
		current.Next = intersectionHead
	}

	// 创建链表B
	headB := createLinkedList(listB)
	if headB == nil {
		headB = intersectionHead
	} else {
		current := headB
		for current.Next != nil {
			current = current.Next
		}
		current.Next = intersectionHead
	}

	return headA, headB
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

func main() {
	// 测试用例
	testCases := []struct {
		listA        []int
		listB        []int
		intersection []int
		desc         string
	}{
		{[]int{4, 1}, []int{5, 6, 1}, []int{8, 4, 5}, "相交节点值为8"},
		{[]int{1, 9, 1}, []int{3}, []int{2, 4}, "相交节点值为2"},
		{[]int{2, 6, 4}, []int{1, 5}, []int{}, "不相交"},
		{[]int{1}, []int{2}, []int{3, 4, 5}, "相交节点值为3"},
		{[]int{}, []int{1, 2}, []int{3, 4}, "A为空，相交节点值为3"},
		{[]int{1, 2}, []int{}, []int{3, 4}, "B为空，相交节点值为3"},
	}

	fmt.Println("=== 解法一：双指针法（推荐）===")
	for i, tc := range testCases {
		headA, headB := createIntersectionList(tc.listA, tc.listB, tc.intersection)
		result := getIntersectionNode(headA, headB)

		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		fmt.Printf("链表A: %s\n", printList(headA))
		fmt.Printf("链表B: %s\n", printList(headB))
		if result != nil {
			fmt.Printf("相交节点: %d\n", result.Val)
		} else {
			fmt.Printf("相交节点: null\n")
		}
		fmt.Println()
	}

	fmt.Println("=== 解法二：计算长度差 ===")
	for i, tc := range testCases {
		headA, headB := createIntersectionList(tc.listA, tc.listB, tc.intersection)
		result := getIntersectionNode2(headA, headB)

		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		if result != nil {
			fmt.Printf("相交节点: %d\n", result.Val)
		} else {
			fmt.Printf("相交节点: null\n")
		}
	}

	fmt.Println("\n=== 解法三：哈希集合 ===")
	for i, tc := range testCases {
		headA, headB := createIntersectionList(tc.listA, tc.listB, tc.intersection)
		result := getIntersectionNode3(headA, headB)

		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		if result != nil {
			fmt.Printf("相交节点: %d\n", result.Val)
		} else {
			fmt.Printf("相交节点: null\n")
		}
	}

	fmt.Println("\n=== 解法四：暴力法 ===")
	for i, tc := range testCases {
		headA, headB := createIntersectionList(tc.listA, tc.listB, tc.intersection)
		result := getIntersectionNode4(headA, headB)

		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		if result != nil {
			fmt.Printf("相交节点: %d\n", result.Val)
		} else {
			fmt.Printf("相交节点: null\n")
		}
	}

	fmt.Println("\n=== 解法五：双指针法（优化版本）===")
	for i, tc := range testCases {
		headA, headB := createIntersectionList(tc.listA, tc.listB, tc.intersection)
		result := getIntersectionNode5(headA, headB)

		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		if result != nil {
			fmt.Printf("相交节点: %d\n", result.Val)
		} else {
			fmt.Printf("相交节点: null\n")
		}
	}

	fmt.Println("\n=== 解法六：栈方法 ===")
	for i, tc := range testCases {
		headA, headB := createIntersectionList(tc.listA, tc.listB, tc.intersection)
		result := getIntersectionNode6(headA, headB)

		fmt.Printf("测试用例 %d: %s\n", i+1, tc.desc)
		if result != nil {
			fmt.Printf("相交节点: %d\n", result.Val)
		} else {
			fmt.Printf("相交节点: null\n")
		}
	}
}

/*
解题思路：

1. 双指针法（推荐）：
   - 核心思想：让两个指针分别遍历两个链表，当到达末尾时交换到另一个链表的头部
   - 这样两个指针走过的距离相等，如果有交点，必定会在交点相遇
   - 时间复杂度：O(m+n)，空间复杂度：O(1)
   - 优点：代码简洁，空间效率高

2. 计算长度差：
   - 先计算两个链表的长度差，让长链表的指针先走差值步
   - 然后两个指针同时前进，必定会在交点相遇
   - 时间复杂度：O(m+n)，空间复杂度：O(1)
   - 优点：逻辑清晰，易于理解

3. 哈希集合：
   - 遍历第一个链表，将所有节点加入哈希集合
   - 然后遍历第二个链表，检查节点是否在集合中
   - 时间复杂度：O(m+n)，空间复杂度：O(m) 或 O(n)
   - 优点：思路简单，适用于一般情况

4. 暴力法：
   - 对第一个链表的每个节点，遍历第二个链表查找相同节点
   - 时间复杂度：O(m*n)，空间复杂度：O(1)
   - 优点：实现简单，缺点：效率低

5. 双指针法（优化版本）：
   - 与解法一相同，但代码更简洁
   - 时间复杂度：O(m+n)，空间复杂度：O(1)
   - 优点：代码最简洁

6. 栈方法：
   - 将两个链表分别压入栈中，然后同时弹出比较
   - 时间复杂度：O(m+n)，空间复杂度：O(m+n)
   - 优点：思路清晰，缺点：空间复杂度高

关键点：
1. 理解相交链表的定义：两个链表在某个节点后共享相同的节点
2. 掌握双指针技巧：通过交换遍历路径来消除长度差
3. 注意边界情况：空链表、不相交的情况
4. 理解不同解法的优缺点和适用场景

时间复杂度分析：
- 双指针法、计算长度差、哈希集合、栈方法：O(m+n)
- 暴力法：O(m*n)

空间复杂度分析：
- 双指针法、计算长度差、暴力法：O(1)
- 哈希集合：O(m) 或 O(n)
- 栈方法：O(m+n)

面试要点：
1. 能够解释双指针法的原理
2. 理解为什么双指针法能够找到交点
3. 掌握不同解法的优缺点
4. 考虑边界情况的处理
*/
