package main

import (
	"fmt"
)

/*
题目：LRU 缓存
难度：中等
标签：设计、哈希表、链表、双向链表

题目描述：
请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。

实现 LRUCache 类：
- LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
- int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1
- void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该逐出最久未使用的关键字。

函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

要求：
1. get和put操作的时间复杂度为O(1)
2. 当缓存满时，删除最久未使用的元素
3. 每次访问元素后，将其标记为最近使用

示例：
输入：
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出：
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释：
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

提示：
- 使用哈希表 + 双向链表
- 哈希表提供O(1)的查找
- 双向链表提供O(1)的插入和删除
*/

// Node 定义双向链表节点
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCache 实现LRU缓存
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

// Constructor 初始化LRU缓存
func Constructor(capacity int) LRUCache {
	// TODO: 在这里实现你的代码
	return LRUCache{}
}

// Get 获取缓存值
func (this *LRUCache) Get(key int) int {
	// TODO: 在这里实现你的代码
	return -1
}

// Put 插入或更新缓存
func (this *LRUCache) Put(key int, value int) {
	// TODO: 在这里实现你的代码
}

func main() {
	// 测试用例1：基本操作
	fmt.Println("测试用例1：基本操作")
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Printf("get(1) = %d (期望: 1)\n", lru.Get(1))
	lru.Put(3, 3)
	fmt.Printf("get(2) = %d (期望: -1)\n", lru.Get(2))
	lru.Put(4, 4)
	fmt.Printf("get(1) = %d (期望: -1)\n", lru.Get(1))
	fmt.Printf("get(3) = %d (期望: 3)\n", lru.Get(3))
	fmt.Printf("get(4) = %d (期望: 4)\n", lru.Get(4))

	// 测试用例2：容量为1
	fmt.Println("\n测试用例2：容量为1")
	lru2 := Constructor(1)
	lru2.Put(1, 1)
	fmt.Printf("get(1) = %d (期望: 1)\n", lru2.Get(1))
	lru2.Put(2, 2)
	fmt.Printf("get(1) = %d (期望: -1)\n", lru2.Get(1))
	fmt.Printf("get(2) = %d (期望: 2)\n", lru2.Get(2))

	// 测试用例3：更新已存在的key
	fmt.Println("\n测试用例3：更新已存在的key")
	lru3 := Constructor(2)
	lru3.Put(1, 1)
	lru3.Put(1, 2)
	fmt.Printf("get(1) = %d (期望: 2)\n", lru3.Get(1))

	// 测试用例4：空缓存
	fmt.Println("\n测试用例4：空缓存")
	lru4 := Constructor(2)
	fmt.Printf("get(1) = %d (期望: -1)\n", lru4.Get(1))

	// 测试用例5：复杂操作序列
	fmt.Println("\n测试用例5：复杂操作序列")
	lru5 := Constructor(3)
	lru5.Put(1, 1)
	lru5.Put(2, 2)
	lru5.Put(3, 3)
	fmt.Printf("get(1) = %d (期望: 1)\n", lru5.Get(1))
	lru5.Put(4, 4)
	fmt.Printf("get(2) = %d (期望: -1)\n", lru5.Get(2))
	fmt.Printf("get(1) = %d (期望: 1)\n", lru5.Get(1))
	fmt.Printf("get(3) = %d (期望: 3)\n", lru5.Get(3))
	fmt.Printf("get(4) = %d (期望: 4)\n", lru5.Get(4))
}

/*
预期输出：
测试用例1：基本操作
get(1) = 1 (期望: 1)
get(2) = -1 (期望: -1)
get(1) = -1 (期望: -1)
get(3) = 3 (期望: 3)
get(4) = 4 (期望: 4)

测试用例2：容量为1
get(1) = 1 (期望: 1)
get(1) = -1 (期望: -1)
get(2) = 2 (期望: 2)

测试用例3：更新已存在的key
get(1) = 2 (期望: 2)

测试用例4：空缓存
get(1) = -1 (期望: -1)

测试用例5：复杂操作序列
get(1) = 1 (期望: 1)
get(2) = -1 (期望: -1)
get(1) = 1 (期望: 1)
get(3) = 3 (期望: 3)
get(4) = 4 (期望: 4)
*/
