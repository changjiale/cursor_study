package main

type Node1 struct {
	key   int
	value int
	pre   *Node1
	next  *Node1
}
type LRUCache1 struct {
	capacity int
	cache    map[int]*Node1
	head     *Node1
	tail     *Node1
}

func New1(capacity int) *LRUCache1 {

	head := &Node1{
		key:   -1,
		value: -1,
	}
	tail := &Node1{
		key:   -1,
		value: -1,
	}

	head.next = tail
	tail.pre = head
	return &LRUCache1{
		capacity: capacity,
		cache:    make(map[int]*Node1),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache1) Get(key int) int {
	if node, exist := l.cache[key]; exist {
		//移动到头部
		l.move2Head(node)
		return node.value
	} else {
		return -1
	}
}

func (l *LRUCache1) Put(key, value int) {
	//已经存在
	if node, exist := l.cache[key]; exist {
		//移动到头部
		l.move2Head(node)
		return
	}

	//新内容

	//满了
	if l.capacity == len(l.cache) {
		//移除尾部
		l.removeTail()
	}

	newNode := &Node1{
		key:   key,
		value: value,
	}

	l.add2Head(newNode)
	l.cache[key] = newNode

}

func (l *LRUCache1) move2Head(node *Node1) {

	//先删除
	l.removeNode(node)
	//在加入头部
	l.add2Head(node)
}

func (l *LRUCache1) removeTail() {
	preNode := l.tail.pre
	l.removeNode(preNode)
	delete(l.cache, preNode.value)
}

func (l *LRUCache1) removeNode(node *Node1) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (l *LRUCache1) add2Head(node *Node1) {

	node.next = l.head.next
	node.pre = l.head
	l.head.next.pre = node
	l.head.next = node
}
