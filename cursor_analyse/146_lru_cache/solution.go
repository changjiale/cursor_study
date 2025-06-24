package main

import (
	"fmt"
)

// Node 定义双向链表节点
// key: 缓存的键
// value: 缓存的值
// prev: 前驱节点指针
// next: 后继节点指针
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// LRUCache 实现LRU缓存
// capacity: 缓存容量
// cache: 哈希表，用于O(1)时间查找节点
// head: 双向链表头节点（哨兵节点）
// tail: 双向链表尾节点（哨兵节点）
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

// Constructor 初始化LRU缓存
// 创建头尾哨兵节点，初始化容量和哈希表
func Constructor(capacity int) LRUCache {
	// 创建头尾哨兵节点，简化边界处理
	head := &Node{key: -1, value: -1}
	tail := &Node{key: -1, value: -1}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// Get 获取缓存值
// 1. 在哈希表中查找节点
// 2. 如果找到，将节点移到链表头部（表示最近使用）
// 3. 返回节点值，如果未找到返回-1
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

// Put 插入或更新缓存
// 1. 如果key已存在，更新值并移到头部
// 2. 如果缓存已满，删除最久未使用的（尾部节点）
// 3. 创建新节点并加入缓存
func (this *LRUCache) Put(key int, value int) {
	// 如果key已存在，更新值并移到头部
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}

	// 如果缓存已满，删除最久未使用的（尾部节点）
	if len(this.cache) == this.capacity {
		this.removeTail()
	}

	// 创建新节点并加入缓存
	node := &Node{key: key, value: value}
	this.cache[key] = node
	this.addToHead(node)
}

// moveToHead 将节点移到链表头部
// 1. 从原位置删除节点
// 2. 将节点添加到头部
func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

// removeNode 从链表中删除节点
// 通过修改前后节点的指针实现O(1)时间删除
func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// addToHead 将节点添加到链表头部
// 1. 设置新节点的前后指针
// 2. 更新原头节点和新节点的指针
func (this *LRUCache) addToHead(node *Node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

// removeTail 删除链表尾部节点
// 1. 获取尾部节点
// 2. 从链表中删除
// 3. 从哈希表中删除
func (this *LRUCache) removeTail() {
	node := this.tail.prev
	this.removeNode(node)
	delete(this.cache, node.key)
}

// 极简写法：使用自定义List结构
// 将链表操作封装在单独的List结构体中，使代码更模块化

// List 封装双向链表操作
type List struct {
	head *Node
	tail *Node
}

// NewList 创建新的双向链表
// 初始化头尾哨兵节点
func NewList() *List {
	head := &Node{key: -1, value: -1}
	tail := &Node{key: -1, value: -1}
	head.next = tail
	tail.prev = head
	return &List{head: head, tail: tail}
}

// LRUCacheSimple 使用List封装的LRU缓存
type LRUCacheSimple struct {
	capacity int
	cache    map[int]*Node
	list     *List
}

// ConstructorSimple 初始化简化版LRU缓存
func ConstructorSimple(capacity int) LRUCacheSimple {
	return LRUCacheSimple{
		capacity: capacity,
		cache:    make(map[int]*Node),
		list:     NewList(),
	}
}

// Get 获取缓存值（简化版）
func (this *LRUCacheSimple) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.list.moveToHead(node)
		return node.value
	}
	return -1
}

// Put 插入或更新缓存（简化版）
func (this *LRUCacheSimple) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.list.moveToHead(node)
		return
	}

	if len(this.cache) == this.capacity {
		this.list.removeTail()
		delete(this.cache, this.list.tail.prev.key)
	}

	node := &Node{key: key, value: value}
	this.cache[key] = node
	this.list.addToHead(node)
}

// List的方法实现
func (l *List) moveToHead(node *Node) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *List) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *List) addToHead(node *Node) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}

func (l *List) removeTail() {
	l.removeNode(l.tail.prev)
}

func main() {
	// 测试标准解法
	fmt.Println("标准解法测试：")
	lru := Constructor(2)
	lru.Put(1, 1)           // 缓存是 {1=1}
	lru.Put(2, 2)           // 缓存是 {1=1, 2=2}
	fmt.Println(lru.Get(1)) // 返回 1，缓存是 {2=2, 1=1}
	lru.Put(3, 3)           // 删除 2，缓存是 {1=1, 3=3}
	fmt.Println(lru.Get(2)) // 返回 -1 (未找到)
	lru.Put(4, 4)           // 删除 1，缓存是 {3=3, 4=4}
	fmt.Println(lru.Get(1)) // 返回 -1 (未找到)
	fmt.Println(lru.Get(3)) // 返回 3
	fmt.Println(lru.Get(4)) // 返回 4

	fmt.Println("\n极简写法测试：")
	lruSimple := ConstructorSimple(2)
	lruSimple.Put(1, 1)
	lruSimple.Put(2, 2)
	fmt.Println(lruSimple.Get(1)) // 返回 1
	lruSimple.Put(3, 3)           // 删除 2
	fmt.Println(lruSimple.Get(2)) // 返回 -1
	lruSimple.Put(4, 4)           // 删除 1
	fmt.Println(lruSimple.Get(1)) // 返回 -1
	fmt.Println(lruSimple.Get(3)) // 返回 3
	fmt.Println(lruSimple.Get(4)) // 返回 4
}
