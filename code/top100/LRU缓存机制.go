package main

//运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上
//限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
// 进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
// 示例：
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4

// Node 双向链表 中的每一个节点的结构体
type Node struct {
	key, value int
	prev, next *Node
}

// DoubleList 双向链表结构体
type DoubleList struct {
	//虚拟头尾节点
	head, tail *Node
	size       int
}

//在链表尾部添加节点node
func (dl *DoubleList) addLast(node *Node) {
	//新加的node节点前指针 指向 链表的尾部的前指针。
	node.prev = dl.tail.prev
	node.next = dl.tail
	dl.tail.prev.next = node
	dl.tail.prev = node
	dl.size++
}

//删除链表中的节点node 当put一个key存在value不同时，做的是先删除后添加的操作，这样新加的永远在队尾
func (dl *DoubleList) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	dl.size--
}

//删除链表第一个真实元素，并且返回该元素 这就是删除最近未使用的功能
func (dl *DoubleList) removeFirst() *Node {
	if dl.head.next == dl.tail {
		return nil
	}
	first := dl.head.next
	dl.remove(first)
	return first
}

type LRUCache struct {
	//map 用于快速找到精确的key对应的value
	cacheMap map[int]*Node
	//双向链表
	cacheList *DoubleList
	//最大容量
	capacity int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		cacheMap: map[int]*Node{},
		cacheList: &DoubleList{
			head: &Node{key: 0, value: 0}, //虚拟一个头节点
			tail: &Node{key: 0, value: 0}, //虚拟一个尾结点
		},
		capacity: capacity,
	}
	//把双向链表虚拟头尾节点 链接起来
	lru.cacheList.head.next = lru.cacheList.tail
	lru.cacheList.tail.prev = lru.cacheList.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	//判断map中是否存在key
	if node, ok := this.cacheMap[key]; ok {
		//存在的话，将这个节点变成最近使用过，也就是先删除，在放到队尾
		this.cacheList.remove(node)
		this.cacheList.addLast(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	//put复杂很多，如果key存在，那就删除之前的，将现在的放入队尾。如果不存在，需要插入队尾并判断容量
	if node, ok := this.cacheMap[key]; ok {
		this.cacheList.remove(node)
	}

	if this.capacity == this.cacheList.size {
		first := this.cacheList.removeFirst()
		delete(this.cacheMap, first.key)
	}
	newNode := &Node{key: key, value: value}
	this.cacheList.addLast(newNode)
	this.cacheMap[key] = newNode
}
