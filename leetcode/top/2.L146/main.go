package main

/*
*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
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

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put
*/
func main() {

}

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}
type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func New(capacity int) LRUCache {
	head := &Node{key: -1, value: -1}
	tail := &Node{key: -1, value: -1}

	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		//移动到头部
		l.move2Head(node)
		return node.value
	}
	return -1
}
func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.cache[key]; ok {
		node.value = value
		//移动到头部
		l.move2Head(node)
		return
	}

	//满
	if len(l.cache) == l.capacity {
		//删除尾部
		l.removeTail()
	}

	//新内容 构建插入头部
	newNode := &Node{
		key:   key,
		value: value,
	}
	l.cache[key] = newNode
	l.add2Head(newNode)
}

func (l *LRUCache) move2Head(node *Node) {
	//删除节点
	l.removeNode(node)
	//添加到头部
	l.add2Head(node)
}

func (l *LRUCache) removeTail() {
	preNode := l.tail.pre
	l.removeNode(preNode)
	delete(l.cache, preNode.key)
}
func (l *LRUCache) removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (l *LRUCache) add2Head(node *Node) {
	node.next = l.head.next
	node.pre = l.head
	l.head.next.pre = node
	l.head.next = node
}
